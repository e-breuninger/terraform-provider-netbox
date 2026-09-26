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

resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "test-interface"
}

resource "netbox_mac_address" "test" {
  mac_address                  = "00:11:22:33:44:55"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}
