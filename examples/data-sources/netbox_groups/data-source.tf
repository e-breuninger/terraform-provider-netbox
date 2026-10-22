data "netbox_groups" "test" {
  filters = [
    { name = "name", value = "test-group" },
  ]
}
