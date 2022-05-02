---
subcategory: ""
page_title: "Create VM and Assign IP - terraform-provider-netbox"
description: |-
    An example of using multiple resource types to create a VM and assign an IP Address
---

# Create VM and Assign IP

A common use case is to leverage Netbox to keep a record of Virtual Machines and also track IP Addressess.

Given that a cluster already exists, we can:

- Retrieve the Cluster's ID by Cluster Name
- Create a Virtual Machine
- Create an Interface on the Virtual Machine
- Retrieve an IP Address and associate it with the Interface
- Set that IP Address as the primary IP for the VM

Note that you could easily create a new cluster if necessary, or give a specific ID for each resource - however, then you must ensure Netbox doesn't have a conflicting resource with the same ID.

```terraform
data "netbox_cluster" "mycluster" {
    name = "my-cluster"
}

resource "netbox_virtual_machine" "myvm" {
    name       = "testvm"
    cluster_id = data.netbox_cluster.mycluster.id
}

resource "netbox_interface" "myint" {
    name               = "eth0"
    virtual_machine_id = netbox_virtual_machine.myvm.id
}

resource "netbox_ip_address" "myip" {
    ip_address   = "10.0.0.20/24"
    status       = "active"
    interface_id = netbox_interface.myint.id
}

resource "netbox_primary_ip" "myprimary" {
    virtual_machine_id = netbox_virtual_machine.myvm.id
    ip_address_id      = netbox_ip_address.myip.id
}
```
