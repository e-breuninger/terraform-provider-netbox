data "netbox_data_sources" "test" {
  filters = [
    { name = "name", value = "test-data_source" },
  ]
}
