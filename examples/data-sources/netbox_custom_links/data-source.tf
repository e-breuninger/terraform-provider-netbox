data "netbox_custom_links" "test" {
  filters = [
    { name = "name", value = "test-custom_link" },
  ]
}
