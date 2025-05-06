resource "netbox_site" "testSite" {
  name = "test site"
  slug = "test-site"
}

resource "netbox_vlan_group" "group1" {
  name     = "Group One"
  slug     = "group-one"
  scope_id = netbox_site.testSite.id
  scope_type = "dcim.site"
  description = "First VLAN group"
  vid_ranges  = [[1,2], [7,17]]
}

resource "netbox_available_vlan" "vlan1" {
  name        = "vlan1"
  status      = "active"
  description = "Virtual network for team 1"
  group_id    = netbox_vlan_group.group1.id
  site_id     = netbox_vlan_group.group1.scope_id
}

resource "netbox_available_vlan" "vlan2" {
  name        = "vlan2"
  status      = "active"
  description = "Virtual network for team 2"
  group_id    = netbox_vlan_group.group1.id
  site_id     = netbox_vlan_group.group1.scope_id
}

resource "netbox_available_vlan" "vlan3" {
  name        = "vlan3"
  status      = "active"
  description = "Virtual network for team 3"
  group_id    = netbox_vlan_group.group1.id
  site_id     = netbox_vlan_group.group1.scope_id
}
