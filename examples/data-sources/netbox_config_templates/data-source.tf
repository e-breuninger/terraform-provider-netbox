data "netbox_config_templates" "test" {
  filters = [
    { name = "name", value = "test-config_template" },
  ]
}
