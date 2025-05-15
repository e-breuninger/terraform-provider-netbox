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
  vid_ranges  = [[1,20]]
}

resource "netbox_available_vlan" "testVlan" {
  name        = "test-vlan"
  status      = "active"
  description = "Virtual network for testing purposes"
  group_id    = netbox_vlan_group.testGroup.id
  site_id     = netbox_vlan_group.testGroup.scope_id
}

resource "netbox_prefix" "testPrefix" {
  prefix      = "192.168.40.0/24"
  status      = "active"
  description = "test prefix"
  site_id     = netbox_site.testSite.id
  vlan_id     = netbox_available_vlan.testVlan.id
}
