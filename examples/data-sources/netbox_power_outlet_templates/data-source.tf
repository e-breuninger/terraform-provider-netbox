data "netbox_power_outlet_templates" "test" {
  filters = [
    { name = "name", value = "test-power_outlet_template" },
  ]
}
