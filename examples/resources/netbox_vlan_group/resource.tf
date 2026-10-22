resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_vlan_group" "test" {
  name        = "test-vlan-group"
  site_id     = netbox_site.test.id
  description = "test-description"
  vid_ranges = [
    { start = 100, end = 199 },
    { start = 300, end = 399 },
  ]
}
