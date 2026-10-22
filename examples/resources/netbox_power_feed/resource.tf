resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_power_panel" "test" {
  name    = "test-power-panel"
  site_id = netbox_site.test.id
}

resource "netbox_rack" "test" {
  name    = "test-rack"
  site_id = netbox_site.test.id
}

resource "netbox_power_feed" "test" {
  name                    = "test-power-feed"
  power_panel_id          = netbox_power_panel.test.id
  rack_id                 = netbox_rack.test.id
  status                  = "active"
  type                    = "primary"
  supply                  = "ac"
  phase                   = "three-phase"
  voltage                 = 230
  amperage                = 32
  max_utilization_percent = 75
  mark_connected          = true
  description             = "test-description"
}
