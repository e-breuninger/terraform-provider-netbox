data "netbox_module_bay_templates" "test" {
  filters = [
    { name = "name", value = "test-module_bay_template" },
  ]
}
