---
subcategory: ""
page_title: "Create VM and Assign IP - terraform-provider-netbox"
description: |-
    An example of using multiple resource types to create a VM and assign an IP Address
---

# Create VM and Assign IP

A common use case is to leverage Netbox to keep a record of Virtual Machines and also track IP Addressess.

Given that an a cluster already exists, we can:

- Retrieve the Cluster's ID by Cluster Name
- Create a Virtual Machine
- Create an Interface on the Virtual Machine
- Retrieve an IP Address and associate it with the Interface
- Set that IP Address as the primary IP for the VM

Note that you could easily create a new cluster if necessary, or give a specific ID for each resource - however then you must ensure Netbox doesn't have a conflicting resource with the same ID.

```terraform
data "netbox_cluster" "mycluster" {
    name = "my-cluster"
}

resource "netbox_virtual_machine" "vm" {
    name = "testvm"
    cluster_id = data.netbox_cluster.mycluster.id
}

resource "netbox_interface" "vm-int" {
    virtual_machine_id = netbox_virtual_machine.vm.id
    name = "eth0"
}

resource "netbox_ip_address" "ip" {
    ip_address = "10.0.0.20/24"
    status = "active"
    interface_id = netbox_interface.vm-int.id
}

resource "netbox_primary_ip" "vm_primary" {
    virtual_machine_id = netbox_virtual_machine.vm.id
    ip_address_id = netbox_ip_address.ip.id
}
```