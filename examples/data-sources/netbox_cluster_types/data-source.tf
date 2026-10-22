data "netbox_cluster_types" "test" {
  filters = [
    { name = "name", value = "test-cluster_type" },
  ]
}
