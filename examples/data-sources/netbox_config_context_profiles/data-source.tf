data "netbox_config_context_profiles" "test" {
  filters = [
    { name = "name", value = "test-config_context_profile" },
  ]
}
