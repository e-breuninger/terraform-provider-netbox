data "netbox_console_server_port_templates" "test" {
  filters = [
    { name = "name", value = "test-console_server_port_template" },
  ]
}
