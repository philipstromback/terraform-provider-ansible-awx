package ansible

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAnsibleJobTemplateLaunch() *schema.Resource {
    return &schema.Resource {
        Read: dataSourceAnsibleJobTemplateLaunchRead,
        Schema: map[string]*schema.Schema {
            "job": &schema.Schema {
                Type: schema.TypeString,
                Required: true,
            },
        },
    }
}

func dataSourceAnsibleJobTemplateLaunchRead(d *schema.ResourceData, meta interface{}) error {
    return ReadLaunchJobTemplate(d.Get("job").(string), dataSourceAnsibleJobTemplateLaunch(), d, meta)
}
