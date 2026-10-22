data "netbox_console_port_templates" "test" {
  filters = [
    { name = "name", value = "test-console_port_template" },
  ]
}
