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

resource "netbox_power_feed" "test" {
  power_panel_id          = netbox_power_panel.test.id
  name                    = "Power Feed 1"
  status                  = "active"
  type                    = "primary"
  supply                  = "ac"
  phase                   = "single-phase"
  voltage                 = 250
  amperage                = 100
  max_percent_utilization = 80
}
