// Assumes vmw-cluster-01 exists as a cluster in Netbox
data "netbox_cluster" "vmw_cluster_01" {
  name = "vmw-cluster-01"
}

data "netbox_virtual_machines" "base_vm" {
  name_regex = "myvm-1"
  filter {
    name  = "cluster_id"
    value = data.netbox_cluster.vmw_cluster_01.id
  }
}

// Custom fields are matched exactly and are combined with the other filters
data "netbox_virtual_machines" "prod_backup_vms" {
  custom_fields = {
    environment = "production"
    backup_tier = "gold"
  }
}
