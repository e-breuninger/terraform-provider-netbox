resource "netbox_virtual_machine" "myvm" {
  name       = "myvm-1"
  cluster_id = netbox_cluster.vmw_cluster_01.id
}

resource "netbox_interface" "myvm_eth0" {
  name               = "eth0"
  virtual_machine_id = netbox_virtual_machine.myvm.id
}

resource "netbox_mac_address" "myvm_mac" {
  mac_address                  = "00:10:FA:63:38:4A"
  virtual_machine_interface_id = netbox_interface.myvm_eth0.id
}

resource "netbox_virtual_machine_interface_primary_mac_address" "myvm_primary_mac" {
  interface_id   = netbox_interface.myvm_eth0.id
  mac_address_id = netbox_mac_address.myvm_mac.id
}
