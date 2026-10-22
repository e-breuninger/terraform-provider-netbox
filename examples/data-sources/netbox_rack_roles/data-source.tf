data "netbox_rack_roles" "test" {
  filters = [
    { name = "name", value = "test-rack_role" },
  ]
}
