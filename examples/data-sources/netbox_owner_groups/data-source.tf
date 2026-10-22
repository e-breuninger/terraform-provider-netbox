data "netbox_owner_groups" "test" {
  filters = [
    { name = "name", value = "test-owner_group" },
  ]
}
