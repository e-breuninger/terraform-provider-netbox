// Assumes Netbox already has a VM whos name matches 'dc-west-myvm-20'
data "netbox_virtual_machine" "myvm" {
  name_regex = "dc-west-myvm-20"
}

data "netbox_prefix" "test" {
  cidr = "10.0.0.0/24"
}

resource "netbox_interface" "myvm-eth0" {
  name               = "eth0"
  virtual_machine_id = data.netbox_virtual_machine.myvm.id
}

resource "netbox_available_ip_address" "myvm-ip" {
  prefix_id    = data.netbox_prefix.test.id
  status       = "active"
  interface_id = netbox_interface.myvm-eth0.id
}
