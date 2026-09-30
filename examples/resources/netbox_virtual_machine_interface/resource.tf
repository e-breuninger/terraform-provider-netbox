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

resource "netbox_vlan" "test_a" {
  name = "test-vlan-a"
  vid  = 100
}

resource "netbox_vlan" "test_b" {
  name = "test-vlan-b"
  vid  = 101
}

resource "netbox_vrf" "test" {
  name = "test-vrf"
}

resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "test-virtual-machine-interface"
  mtu                = 1500
  mode               = "tagged"
  untagged_vlan_id   = netbox_vlan.test_a.id
  tagged_vlan_ids    = [netbox_vlan.test_b.id]
  vrf_id             = netbox_vrf.test.id
  description        = "test-description"
}
