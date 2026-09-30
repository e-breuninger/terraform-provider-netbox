data "netbox_power_panels" "test" {
  filters = [
    { name = "name", value = "test-power_panel" },
  ]
}
