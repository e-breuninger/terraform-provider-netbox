resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_location" "test" {
  name    = "test-location"
  site_id = netbox_site.test.id
}

resource "netbox_power_panel" "test" {
  name        = "test-power-panel"
  site_id     = netbox_site.test.id
  location_id = netbox_location.test.id
  description = "test-description"
}
