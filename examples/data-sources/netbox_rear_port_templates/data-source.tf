data "netbox_rear_port_templates" "test" {
  filters = [
    { name = "name", value = "test-rear_port_template" },
  ]
}
