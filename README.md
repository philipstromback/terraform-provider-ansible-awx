# terraform-provider-ansible-awx
A provider to run ansible-awx common role

To Build:
```HCL
go build -o terraform-provider-ansible ; cp terraform-provider-ansible ~/.terraform.d/plugins/linux_amd64/

export ANSIBLE_AWX_TOKEN=xxxxxxxxxxxxxxxxxxxxxxx
export ANSIBLE_AWX_URL=http://xxxxxxxxxxx
```
