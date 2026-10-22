data "netbox_platforms" "test" {
  filters = [
    { name = "name", value = "test-platform" },
  ]
}
