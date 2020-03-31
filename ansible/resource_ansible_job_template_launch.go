package ansible

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceAnsibleJobTemplateLaunch() *schema.Resource {
    return &schema.Resource {
        Create: resourceAnsibleJobTemplateLaunchCreate,
        Read: resourceAnsibleJobTemplateLaunchRead,
        Update: resourceAnsibleJobTemplateLaunchUpdate,
        Delete: resourceAnsibleJobTemplateLaunchDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema {
            "job":  &schema.Schema {
                Type:   schema.TypeString,
                Computed: true,
            },
            "name":  &schema.Schema {
                Type:     schema.TypeString,
                Required: true,
            },
            "job_template": &schema.Schema {
                Type:     schema.TypeString,
                Required: true,
            },
            "service": &schema.Schema {
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceAnsibleJobTemplateLaunchCreate(d *schema.ResourceData, meta interface{}) error {
    return CreateLaunchJobTemplate(d.Get("name").(string),d.Get("job_template").(string),d.Get("service").(string),resourceAnsibleJobTemplateLaunch(), d, meta)
}

func resourceAnsibleJobTemplateLaunchRead(d *schema.ResourceData, meta interface{}) error {
    return ReadLaunchJobTemplate(d.Get("name").(string),resourceAnsibleJobTemplateLaunch(), d, meta)
}

func resourceAnsibleJobTemplateLaunchUpdate(d *schema.ResourceData, meta interface{}) error {
    return UpdateLaunchJobTemplate(d.Get("name").(string),d.Get("job_template").(string),d.Get("service").(string),resourceAnsibleJobTemplateLaunch(), d, meta)
}

func resourceAnsibleJobTemplateLaunchDelete(d *schema.ResourceData, meta interface{}) error {
    return DeleteLaunchJobTemplate(d.Get("job").(string), d, meta)
}

