resource "netbox_site" "test" {
  name   = "Site 1"
  status = "active"
}

resource "netbox_location" "test" {
  name    = "Location 1"
  site_id = netbox_site.test.id
}

resource "netbox_power_panel" "test" {
  name        = "Power Panel 1"
  site_id     = netbox_site.test.id
  location_id = netbox_location.test.id
}
