# Get VLAN group by name
data "netbox_vlan_group" "example1" {
  name = "example1"
}

# Get VLAN group by stub
data "netbox_vlan_group" "example2" {
  slug = "example2"
}

# Get VLAN group by name and scope_type/id
data "netbox_vlan_group" "example3" {
  name       = "example"
  scope_type = "dcim.site"
  scope_id   = netbox_site.example.id
}
