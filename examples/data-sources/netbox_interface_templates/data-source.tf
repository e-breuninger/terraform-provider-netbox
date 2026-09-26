data "netbox_interface_templates" "test" {
  filters = [
    { name = "name", value = "test-interface_template" },
  ]
}
