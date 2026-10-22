data "netbox_tags" "test" {
  filters = [
    { name = "name", value = "test-tag" },
  ]
}
