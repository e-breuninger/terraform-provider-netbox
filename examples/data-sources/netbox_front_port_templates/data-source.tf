data "netbox_front_port_templates" "test" {
  filters = [
    { name = "name", value = "test-front_port_template" },
  ]
}
