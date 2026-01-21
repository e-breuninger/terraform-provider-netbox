// Assuming a virtual machine with the id `123` exists
resource "netbox_interface" "this" {
  name               = "eth0"
  virtual_machine_id = 123
}

resource "netbox_ip_address" "this" {
  ip_address = "10.0.0.60/24"
  status     = "active"
}


resource "netbox_ip_address_assignment" "this" {
  ip_address_id                = netbox_ip_address.this.id
  virtual_machine_interface_id = netbox_interface.this.id
}
