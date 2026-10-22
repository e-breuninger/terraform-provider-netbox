data "netbox_rack_groups" "test" {
  filters = [
    { name = "name", value = "test-rack_group" },
  ]
}
