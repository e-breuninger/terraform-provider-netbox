data "netbox_sites" "test" {
  filters = [
    { name = "name", value = "test-site" },
  ]
}
