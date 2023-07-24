resource "netbox_site" "test" {
  name   = "test"
  status = "active"
}

resource "netbox_rack" "test" {
  name     = "test"
  site_id  = netbox_site.test.id
  status   = "reserved"
  width    = 19
  u_height = 48
}
