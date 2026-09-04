---
page_title: "netbox_ip_address Resource - terraform-provider-netbox"
subcategory: "IP Address Management (IPAM)"
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses:
  An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
  Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchy.
  Either set ip_address directly, or set prefix_id/ip_range_id to have Netbox allocate the next free address from that prefix or range (mirrors netbox_available_ip_address).
---

# netbox_ip_address (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses):

> An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
>
> Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchy.

Either set `ip_address` directly, or set `prefix_id`/`ip_range_id` to have Netbox allocate the next free address from that prefix or range (mirrors `netbox_available_ip_address`).

## Example Usage

### Creating an IP address that is assigned to a virtual machine interface

Starting with provider version 3.5.0, you can use the `virtual_machine_interface_id` attribute to assign an IP address to a virtual machine interface.
You can also use the `interface_id` and `object_type` attributes instead.

With `virtual_machine_interface_id`:
```terraform
// Assuming a virtual machine with the id `123` exists
resource "netbox_interface" "this" {
  name               = "eth0"
  virtual_machine_id = 123
}

resource "netbox_ip_address" "this" {
  ip_address                   = "10.0.0.60/24"
  status                       = "active"
  virtual_machine_interface_id = netbox_interface.this.id
}
```

With `object_type` and `interface_id`:
```terraform
// Assuming a virtual machine with the id `123` exists
resource "netbox_interface" "this" {
  name               = "eth0"
  virtual_machine_id = 123
}

resource "netbox_ip_address" "this" {
  ip_address   = "10.0.0.60/24"
  status       = "active"
  interface_id = netbox_interface.this.id
  object_type  = "virtualization.vminterface"
}
```

### Creating an IP address that is assigned to a device interface

Starting with provider version 3.5.0, you can use the `device_interface_id` attribute to assign an IP address to a device interface.
You can also use the `interface_id` and `object_type` attributes instead.

With `device_interface_id`:
```terraform
// Assuming a device with the id `123` exists
resource "netbox_device_interface" "this" {
  name      = "eth0"
  device_id = 123
  type      = "1000base-t"
}

resource "netbox_ip_address" "this" {
  ip_address          = "10.0.0.60/24"
  status              = "active"
  device_interface_id = netbox_device_interface.this.id
}
```

With `object_type` and `interface_id`:
```terraform
// Assuming a device with the id `123` exists
resource "netbox_device_interface" "this" {
  name      = "eth0"
  device_id = 123
  type      = "1000base-t"
}

resource "netbox_ip_address" "this" {
  ip_address   = "10.0.0.60/24"
  status       = "active"
  interface_id = netbox_device_interface.this.id
  object_type  = "dcim.interface"
}
```

### Creating an IP address that is not assigned to anything

You can create an IP address that is not assigned to anything by omitting the attributes mentioned above.

```terraform
resource "netbox_ip_address" "this" {
  ip_address = "10.0.0.50/24"
  status     = "reserved"
}
```

### Allocating the next available address from a prefix

Omit `ip_address` and set `prefix_id` to have Netbox pick the next free address from that prefix.

```terraform
resource "netbox_prefix" "test" {
  prefix = "10.0.0.0/24"
}

resource "netbox_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  status    = "active"
}
```

### Allocating the next available address from an IP range

Same idea, from an IP range instead of a prefix.

```terraform
resource "netbox_ip_range" "test" {
  start_address = "10.0.0.1/24"
  end_address   = "10.0.0.50/24"
}

resource "netbox_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
  status      = "active"
}
```

## Migrating from `netbox_available_ip_address`

`netbox_available_ip_address` is superseded by the `prefix_id` / `ip_range_id`
attributes on this resource. There is no automatic migration: `terraform state
mv` refuses to move state between different resource types, and `moved` blocks
require the provider to implement `MoveResourceState`, which this provider does
not. Migrate manually, one resource at a time:

```shell
terraform state rm netbox_available_ip_address.example
```

Then change the resource block's type from `netbox_available_ip_address` to
`netbox_ip_address` in your configuration, and import it back:

```shell
terraform import netbox_ip_address.example <id>
```

The underlying Netbox object and its ID are unchanged, so the import is a no-op
against the live API and the next plan is empty.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `status` (String) Valid values are `active`, `reserved`, `deprecated`, `dhcp` and `slaac`.

### Optional

- `custom_fields` (Map of String)
- `description` (String)
- `device_interface_id` (Number) Conflicts with `interface_id` and `virtual_machine_interface_id`.
- `dns_name` (String)
- `interface_id` (Number) Required when `object_type` is set.
- `ip_address` (String) Exactly one of `ip_address`, `prefix_id` or `ip_range_id` must be given.
- `ip_range_id` (Number) Exactly one of `ip_address`, `prefix_id` or `ip_range_id` must be given.
- `nat_inside_address_id` (Number)
- `object_type` (String) Valid values are `virtualization.vminterface`, `dcim.interface` and `ipam.fhrpgroup`. Required when `interface_id` is set.
- `prefix_id` (Number) Exactly one of `ip_address`, `prefix_id` or `ip_range_id` must be given.
- `role` (String) Valid values are `loopback`, `secondary`, `anycast`, `vip`, `vrrp`, `hsrp`, `glbp` and `carp`.
- `tags` (Set of String)
- `tenant_id` (Number)
- `virtual_machine_interface_id` (Number) Conflicts with `interface_id` and `device_interface_id`.
- `vrf_id` (Number)

### Read-Only

- `id` (String) The ID of this resource.
- `nat_outside_addresses` (List of Object) (see [below for nested schema](#nestedatt--nat_outside_addresses))
- `tags_all` (Set of String)

<a id="nestedatt--nat_outside_addresses"></a>
### Nested Schema for `nat_outside_addresses`

Read-Only:

- `address_family` (Number)
- `id` (Number)
- `ip_address` (String)


