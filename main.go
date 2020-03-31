package main

import (
    "github.com/philipstromback/terraform-provider-ansible-awx/ansible"
    "github.com/hashicorp/terraform/plugin"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: ansible.Provider})
}
