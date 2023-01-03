---
page_title: "netbox_available_ip_address Resource - terraform-provider-netbox"
subcategory: "IP Address Management (IPAM)"
description: |-
  Per the docs https://netbox.readthedocs.io/en/stable/models/ipam/ipaddress/:
  An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
  Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchya.
  Each IP address can also be assigned an operational status and a functional role. Statuses are hard-coded in NetBox and include the following:
  * Active
  * Reserved
  * Deprecated
  * DHCP
  * SLAAC (IPv6 Stateless Address Autoconfiguration)
  This resource will retrieve the next available IP address from a given prefix or IP range (specified by ID)
---

# netbox_available_ip_address (Resource)

Per [the docs](https://netbox.readthedocs.io/en/stable/models/ipam/ipaddress/):

> An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
> Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchya.
>
> Each IP address can also be assigned an operational status and a functional role. Statuses are hard-coded in NetBox and include the following:
> * Active
> * Reserved
> * Deprecated
> * DHCP
> * SLAAC (IPv6 Stateless Address Autoconfiguration)

This resource will retrieve the next available IP address from a given prefix or IP range (specified by ID)

## Example Usage
### Creating an IP in a prefix
```terraform
data "netbox_prefix" "test" {
  cidr = "10.0.0.0/24"
}

resource "netbox_available_ip_address" "test" {
  prefix_id = data.netbox_prefix.test.id
}
```

### Creating an IP in an IP range
```terraform
data "netbox_ip_range" "test" {
  start_address = "10.0.0.1/24"
  end_address   = "10.0.0.50/24"
}

resource "netbox_available_ip_address" "test" {
  ip_range_id = data.netbox_ip_range.test.id
}
```

### Marking an IP active and assigning to interface
```terraform
// Assumes Netbox already has a VM whos name matches 'dc-west-myvm-20'
data "netbox_virtual_machine" "myvm" {
  name_regex = "dc-west-myvm-20"
}

data "netbox_prefix" "test" {
  cidr = "10.0.0.0/24"
}

resource "netbox_interface" "myvm-eth0" {
  name               = "eth0"
  virtual_machine_id = data.netbox_virtual_machine.myvm.id
}

resource "netbox_available_ip_address" "myvm-ip" {
  prefix_id    = data.netbox_prefix.test.id
  status       = "active"
  interface_id = netbox_interface.myvm-eth0.id
}
```

## Schema

### Required

- Either **prefix_id** or **ip_range_id** (String)

### Optional

- **description** (String)
- **dns_name** (String)
- **interface_id** (Number)
- **status** (String) Defaults to "active".  Choose from "active", "reserved", "deprecated", "dhcp", or "slaac"
- **tags** (Set of String)
- **tenant_id** (Number)
- **vrf_id** (Number)

### Read-Only

- **id** (String) The ID of this resource.
- **ip_address** (String)
