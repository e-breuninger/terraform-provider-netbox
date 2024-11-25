// Assuming a device with the id `123` exists
resource "netbox_device_interface" "this" {
  name      = "eth0"
  device_id = 123
  type      = "1000base-t"
}

resource "netbox_ip_address" "this" {
  ip_address = "10.0.0.60/24"
  status     = "active"
}

resource "netbox_ip_address_assignment" "this" {
  ip_address_id = netbox_ip_address.this.id
  interface_id  = netbox_device_interface.this.id
  object_type   = "dcim.interface"
}
