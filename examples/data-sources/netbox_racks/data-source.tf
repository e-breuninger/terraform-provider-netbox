data "netbox_racks" "test" {
  filters = [
    { name = "name", value = "test-rack" },
  ]
}
