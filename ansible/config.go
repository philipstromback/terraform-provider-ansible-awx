package ansible

import (
    "fmt"
    "log"
)

type Config struct {
        URL                 string
        Token               string
}

func (c *Config) Client(prop Config) (*API, error) {

    client, err := New(prop)

    if err != nil {
        return nil, fmt.Errorf("Error creating new ansible client: %s", err)
    }

    log.Printf("[INFO] Ansible AWX Client configured")

    return client, nil
}
