data "netbox_site_groups" "test" {
  filters = [
    { name = "name", value = "test-site_group" },
  ]
}
