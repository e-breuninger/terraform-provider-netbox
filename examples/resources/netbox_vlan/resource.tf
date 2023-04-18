resource "netbox_vlan" "example1" {
  name = "VLAN 1"
  vid  = 1777
  tags = []
}

# Assume netbox_tenant, netbox_site, and netbox_tag resources exist
resource "netbox_vlan" "example2" {
  name        = "VLAN 2"
  vid         = 1778
  status      = "reserved"
  description = "Reserved example VLAN"
  tenant_id   = netbox_tenant.ex.id
  site_id     = netbox_site.ex.id
  group_id    = netbox_vlan_group.ex.id
  tags        = [netbox_tag.ex.name]
}
