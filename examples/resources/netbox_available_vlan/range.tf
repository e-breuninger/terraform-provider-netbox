resource "netbox_site" "testSite" {
  name = "test site"
  slug = "test-site"
}

resource "netbox_vlan_group" "testGroup" {
  name     = "Group One"
  slug     = "group-one"
  scope_id = netbox_site.testSite.id
  scope_type = "dcim.site"
  description = "First VLAN group"
  vid_ranges  = [[1,2], [7,17]]
}

resource "netbox_available_vlan" "testVlan" {
  name        = "test-vlan"
  status      = "active"
  description = "Virtual network for testing purposes"
  group_id    = netbox_vlan_group.testGroup.id
  site_id     = netbox_vlan_group.testGroup.scope_id
}

resource "netbox_available_vlan" "testVlan2" {
  name        = "test-vlan2"
  status      = "active"
  description = "Virtual network for testing purposes"
  group_id    = netbox_vlan_group.testGroup.id
  site_id     = netbox_vlan_group.testGroup.scope_id
}

resource "netbox_available_vlan" "testVlan3" {
  name        = "test-vlan3"
  status      = "active"
  description = "Virtual network for testing purposes"
  group_id    = netbox_vlan_group.testGroup.id
  site_id     = netbox_vlan_group.testGroup.scope_id
}
