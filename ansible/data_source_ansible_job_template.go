package ansible

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAnsibleJobTemplate() *schema.Resource {
    return &schema.Resource {
        Read: dataSourceAnsibleJobTemplateRead,
        Schema: map[string]*schema.Schema {
            "name": &schema.Schema {
                Type:       schema.TypeString,
                Required:   true,
            },
        },
    }
}

func dataSourceAnsibleJobTemplateRead(d *schema.ResourceData, meta interface{}) error {
    return ReadJobTemplate(d.Get("name").(string), dataSourceAnsibleJobTemplate(), d, meta)
}
