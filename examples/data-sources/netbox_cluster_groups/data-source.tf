data "netbox_cluster_groups" "test" {
  filters = [
    { name = "name", value = "test-cluster_group" },
  ]
}
