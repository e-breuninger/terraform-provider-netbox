data "netbox_permissions" "test" {
  filters = [
    { name = "name", value = "test-permission" },
  ]
}
