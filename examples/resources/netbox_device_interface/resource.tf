// Assumes a device with ID 123 exists
resource "netbox_device_interface" "test" {
  name      = "testinterface"
  device_id = 123
  type      = "1000base-t"
}
