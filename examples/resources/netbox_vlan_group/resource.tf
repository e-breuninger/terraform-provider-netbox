#Basic VLAN Group example
resource "netbox_vlan_group" "example1" {
  name    = "example1"
  slug    = "example1"
}

#Full VLAN Group example
resource "netbox_vlan_group" "example2" {
  name        = "Second Example"
  slug        = "example2"
  scope_type  = "dcim.site"
  scope_id    = netbox_site.example.id
  description = "Second Example VLAN Group"
  tags        = [netbox_tag.example.id]
}
