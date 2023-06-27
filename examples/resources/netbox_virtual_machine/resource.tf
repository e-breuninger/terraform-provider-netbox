// Assumes vmw-cluster-01 exists in Netbox
data "netbox_cluster" "vmw_cluster_01" {
  name = "vmw-cluster-01"
}

resource "netbox_virtual_machine" "base_vm" {
  cluster_id = data.netbox_cluster.vmw_cluster_01.id
  name       = "myvm-1"
}
// Assumes vmw-cluster-01 exists in Netbox
data "netbox_cluster" "vmw_cluster_01" {
  name = "vmw-cluster-01"
}

resource "netbox_virtual_machine" "basic_vm" {
  cluster_id   = data.netbox_cluster.vmw_cluster_01.id
  name         = "myvm-2"
  disk_size_gb = 40
  memory_mb    = 4092
  vcpus        = "2"
}
// Assumes vmw-cluster-01 exists as a cluster in Netbox
data "netbox_cluster" "vmw_cluster_01" {
  name = "vmw-cluster-01"
}

// Assumes customer-a exists as a tenant in Netbox
data "netbox_tenant" "customer_a" {
  name = "Customer A"
}

resource "netbox_virtual_machine" "full_vm" {
  cluster_id   = data.netbox_cluster.vmw_cluster_01.id
  name         = "myvm-3"
  disk_size_gb = 40
  memory_mb    = 4092
  vcpus        = "2"
  role_id      = 31 // This corresponds to the Netbox ID for a given role
  tenant_id    = data.netbox_tenant.customer_a.id
  local_context_data = jsonencode({
    "setting_a" = "Some Setting"
    "setting_b" = 42
  })
}
