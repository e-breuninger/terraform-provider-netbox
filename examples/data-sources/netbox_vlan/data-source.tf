# Get VLAN by name
data "netbox_vlan" "vlan1" {
  name = "vlan-1"
}

# Get VLAN by VID and IPAM role ID
data "netbox_vlan" "vlan2" {
  vid  = 1234
  role = netbox_ipam_role.example.id
}

# Get VLAN by name and tenant ID
data "netbox_vlan" "vlan3" {
  name   = "vlan-3"
  tenant = netbox_tenant.example.id
}
