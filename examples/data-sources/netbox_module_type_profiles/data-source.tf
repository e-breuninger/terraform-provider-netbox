data "netbox_module_type_profiles" "test" {
  filters = [
    { name = "name", value = "test-module_type_profile" },
  ]
}
