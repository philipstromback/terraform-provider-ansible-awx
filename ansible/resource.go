package ansible

import (
    "encoding/json"
    "fmt"
    "strconv"
    "log"
    "time"
    "github.com/hashicorp/terraform/helper/schema"
)


type jobTemplate struct {
    Results []struct {
        Id int
        Name string
    }
}

type job struct {
    Job int
    Status string
}

type jobStatus struct {
    Id  int
    Status string
}

type delJob struct{}

func ReadJobTemplate(name string, resource *schema.Resource, d *schema.ResourceData, m interface{}) error {
    client := m.(*API)

    uri := "/api/v2/job_templates/"
    o, err := client.makeRequest("GET", uri, nil)

    var data jobTemplate
    json.Unmarshal(o, &data)

    if len(data.Results) == 0 {
        err := fmt.Errorf("No job templates for this user....")
        return err
    }

    for _, record := range data.Results {
        if (record.Name == name) {
            d.SetId(strconv.Itoa(record.Id))
            break
        }
    }

    if err != nil {
        d.SetId("")
        return err
    }

    return nil
}

func CreateLaunchJobTemplate(name string, job_template string, service string, resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

    time.Sleep(10 * time.Second)

    client := m.(*API)

    uri := "/api/v2/job_templates/" +  job_template + "/launch/"

    var par params
    par.Limit = d.Get("name").(string)
    par.ExtraVars.Service = d.Get("service").(string)
    props := par

    r, err := client.makeRequest("POST", uri, props)

    var data job
    json.Unmarshal(r, &data)

    uri_job := "/api/v2/jobs/" + strconv.Itoa(data.Job) + "/"


    sum := 0
    for i := 1; i < 50; i++ {
        sum += i
        res, _ := client.makeRequest("GET", uri_job, nil)

        var job_data jobStatus
        json.Unmarshal(res, &job_data)

        if (job_data.Status == "successful") {
            log.Printf("Job is complete")
            d.SetId(strconv.Itoa(data.Job))
            break
        } else if (job_data.Status == "failed") {
            err := fmt.Errorf("Job %v has failed", data.Job)
            d.SetId("")
            return (err)
            break

        } else {
            log.Printf("Job is running, waiting...")
            time.Sleep(10 * time.Second)
        }
    }

    if err != nil {
        d.SetId("")
        return err
    }

    return nil
}

func ReadLaunchJobTemplate(name string, resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

    client := m.(*API)

    uri := "/api/v2/jobs/" + d.Id() + "/"

    r, err := client.makeRequest("GET", uri, nil)

    var data jobStatus
    json.Unmarshal(r, &data)

    d.SetId(strconv.Itoa(data.Id))

    if err != nil {
        d.SetId("")
        return err
    }

    return nil
}

func UpdateLaunchJobTemplate(name string, job_template string, service string, source *schema.Resource, d *schema.ResourceData, m interface{}) error  {

    client := m.(*API)

    uri := "/api/v2/job_templates/" +  job_template + "/launch/"

    var par params
    par.Limit = d.Get("name").(string)
    par.ExtraVars.Service = d.Get("service").(string)
    props := par

    r, err := client.makeRequest("POST", uri, props)

    var data job
    json.Unmarshal(r, &data)

    uri_job := "/api/v2/jobs/" + strconv.Itoa(data.Job) + "/"


    sum := 0
    for i := 1; i < 50; i++ {
        sum += i
        res, _ := client.makeRequest("GET", uri_job, nil)

        var job_data jobStatus
        json.Unmarshal(res, &job_data)

        if (job_data.Status == "successful") {
            log.Printf("Job is complete")
            d.SetId(strconv.Itoa(data.Job))
            break
        } else if (job_data.Status == "failed") {
            err := fmt.Errorf("Job %v has failed", data.Job)
            d.SetId("")
            return (err)
            break

        } else {
            log.Printf("Job is running, waiting...")
            time.Sleep(10 * time.Second)
        }
    }

    if err != nil {
        d.SetId("")
        return err
    }

    return nil
}


func DeleteLaunchJobTemplate(job string, d *schema.ResourceData, m interface{}) error {

    client := m.(*API)

    uri := "/api/v2/jobs/" + d.Id() + "/"

    r, err := client.makeRequest("DELETE", uri, nil)

    var data delJob
    json.Unmarshal(r, &data)

    if err != nil {
        d.SetId("")
        return err
    }

    d.SetId("")

    return nil
}
