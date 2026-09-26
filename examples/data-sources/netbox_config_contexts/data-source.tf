data "netbox_config_contexts" "test" {
  filters = [
    { name = "name", value = "test-config_context" },
  ]
}
