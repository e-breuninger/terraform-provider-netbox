resource "netbox_site" "test" {
  name   = "test"
  status = "active"
}

resource "netbox_rack" "test" {
  name     = "test"
  site_id  = netbox_site.test.id
  status   = "active"
  width    = 10
  u_height = 40
}
resource "netbox_rack_reservation" "test" {
  rack_id     = netbox_rack.test.id
  units       = [1, 2, 3, 4, 5]
  user_id     = 1
  description = "my description"
}
