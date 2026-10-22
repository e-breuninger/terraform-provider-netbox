data "netbox_service_templates" "test" {
  filters = [
    { name = "name", value = "test-service_template" },
  ]
}
