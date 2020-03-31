package ansible

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

var descriptions map[string]string

func init() {
    descriptions = map[string]string{
        "token":               "Token.",
        "url":                 "The Ansible-awx url",
    }
}

//Provider function
func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "url": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("ANSIBLE_AWX_URL", "http://localhost:80"),
                Description: "The Ansible AWX URL",
            },
            "token": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("ANSIBLE_AWX_TOKEN", nil),
                Description: "The Ansible AWX token auth",
            },
        },
        DataSourcesMap: map[string]*schema.Resource{
            "ansible_job_template":          dataSourceAnsibleJobTemplate(),
            "ansible_job_template_launch":   dataSourceAnsibleJobTemplateLaunch(),
        },
        ResourcesMap: map[string]*schema.Resource{
            "ansible_job_template_launch":  resourceAnsibleJobTemplateLaunch(),
        },

        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    config := Config{
        URL:                 d.Get("url").(string),
        Token:               d.Get("token").(string),
    }
    return config.Client(config)
}
