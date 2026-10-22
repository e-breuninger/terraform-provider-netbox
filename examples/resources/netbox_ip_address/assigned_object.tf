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
  name               = "eth0"
}

# The generic pair; device_interface_id and virtual_machine_interface_id are aliases of it.
resource "netbox_ip_address" "test" {
  ip_address           = "192.0.2.30/24"
  status               = "active"
  assigned_object_type = "virtualization.vminterface"
  assigned_object_id   = netbox_virtual_machine_interface.test.id
}
