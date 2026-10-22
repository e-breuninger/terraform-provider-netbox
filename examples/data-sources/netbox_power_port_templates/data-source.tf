data "netbox_power_port_templates" "test" {
  filters = [
    { name = "name", value = "test-power_port_template" },
  ]
}
