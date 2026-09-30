data "netbox_custom_fields" "test" {
  filters = [
    { name = "name", value = "test-custom_field" },
  ]
}
