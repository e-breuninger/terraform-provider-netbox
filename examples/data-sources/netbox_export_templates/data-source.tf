data "netbox_export_templates" "test" {
  filters = [
    { name = "name", value = "test-export_template" },
  ]
}
