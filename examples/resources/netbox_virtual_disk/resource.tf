resource "netbox_cluster_type" "test" {
  name = "test-cluster-type"
}

resource "netbox_cluster" "test" {
  name            = "test-cluster"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_virtual_machine" "test" {
  name       = "test-virtual-machine"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_virtual_disk" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "test-virtual-disk"
  size_mb            = 20480
  description        = "test-description"
}
