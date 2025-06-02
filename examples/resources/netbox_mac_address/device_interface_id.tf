// Assuming a device with the id `123` exists
resource "netbox_device_interface" "this" {
  name      = "eth0"
  device_id = 123
  type      = "1000base-t"
}

resource "netbox_mac_address" "this" {
  mac_address          = "00:1A:2B:3C:4D:5E"
  device_interface_id = netbox_device_interface.this.id
}
