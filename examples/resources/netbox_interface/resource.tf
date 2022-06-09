// Assumes Netbox already has a VM whos name matches 'dc-west-myvm-20'
data "netbox_virtual_machine" "myvm" {
  name_regex = "dc-west-myvm-20"
}

resource "netbox_interface" "myvm_eth0" {
  name               = "eth0"
  virtual_machine_id = data.netbox_virtual_machine.myvm.id
}
