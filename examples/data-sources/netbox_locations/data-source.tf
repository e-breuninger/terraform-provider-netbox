data "netbox_locations" "test" {
  filters = [
    { name = "name", value = "test-location" },
  ]
}
