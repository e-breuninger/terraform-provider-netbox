data "netbox_clusters" "test" {
  filters = [
    { name = "name", value = "test-cluster" },
  ]
}
