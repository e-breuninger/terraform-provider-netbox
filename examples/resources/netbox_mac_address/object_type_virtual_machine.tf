// Assuming a virtual machine with the id `123` exists
resource "netbox_interface" "this" {
  name               = "eth0"
  virtual_machine_id = 123
}

resource "netbox_mac_address" "this" {
  mac_address   = "00:1A:2B:3C:4D:5E"
  interface_id = netbox_interface.this.id
  object_type  = "virtualization.vminterface"
}
