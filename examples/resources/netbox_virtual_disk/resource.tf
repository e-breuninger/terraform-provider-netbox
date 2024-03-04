// Assumes vmw-cluster-01 exists in Netbox
data "netbox_cluster" "vmw_cluster_01" {
  name = "vmw-cluster-01"
}

resource "netbox_virtual_machine" "base_vm" {
  cluster_id = data.netbox_cluster.vmw_cluster_01.id
  name       = "myvm-1"
}

resource "netbox_virtual_disk" "example" {
  name               = "disk-01"
  description        = "Main disk"
  size               = 50
  virtual_machine_id = netbox_virtual_machine.base_vm.id
}
