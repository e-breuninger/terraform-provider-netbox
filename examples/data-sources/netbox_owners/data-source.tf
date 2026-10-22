data "netbox_owners" "test" {
  filters = [
    { name = "name", value = "test-owner" },
  ]
}
