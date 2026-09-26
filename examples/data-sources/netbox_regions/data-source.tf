data "netbox_regions" "test" {
  filters = [
    { name = "name", value = "test-region" },
  ]
}
