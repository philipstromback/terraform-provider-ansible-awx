package ansible

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/pkg/errors"
)

type params struct {
	Limit string `json:"limit"`
	ExtraVars struct {
		Service string `json:"service"`
	} `json:"extra_vars"`
}

type API struct {
	Auth                string
	BaseURL             string
	Insecure            bool
	httpClient          *http.Client
}

func New(config Config) (*API, error) {

	api := &API{
		Auth:                config.Token,
		BaseURL:             config.URL,
	}

	// Setup HTTP Client based on insecure
	api.httpClient = cleanhttp.DefaultClient()
	api.httpClient.Timeout = time.Second * 20
	if api.Insecure {
		transport := cleanhttp.DefaultTransport()
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		api.httpClient.Transport = transport
	}

	return api, nil
}

func (api *API) makeRequest(method, uri string, params interface{}) ([]byte, error) {

	// Replace nil with a JSON object if needed
	var reqBody io.Reader
	if params != nil {
		json, err := json.Marshal(params)
		if err != nil {
			return nil, errors.Wrap(err, "error marshalling params to JSON")
		}
		reqBody = bytes.NewReader(json)
	} else {
		reqBody = nil
	}

	log.Printf("[DEBUG] Body: %q ", reqBody)

	resp, err := api.request(method, uri, reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "could not read response body")
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, 204:
		break
	case http.StatusUnauthorized:
		return nil, errors.Errorf("HTTP status %d: invalid credentials", resp.StatusCode)
	case http.StatusForbidden:
		return nil, errors.Errorf("HTTP status %d: insufficient permissions", resp.StatusCode)
	case http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout,
		522, 523, 524:
		return nil, errors.Errorf("HTTP status %d: service failure", resp.StatusCode)
	default:
		var s string
		if body != nil {
			s = string(body)
		}
		return nil, errors.Errorf("HTTP status %d: content %q", resp.StatusCode, s)
	}

	return body, nil
}

func (api *API) request(method, uri string, reqBody io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, api.BaseURL+uri, reqBody)
	req.Header.Add("Authorization", "Bearer "+api.Auth)

	if err != nil {
		return nil, errors.Wrap(err, "HTTP request creation failed")
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := api.httpClient.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "HTTP request failed")
	}

	return resp, nil
}
