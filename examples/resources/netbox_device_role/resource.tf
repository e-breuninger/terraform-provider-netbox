resource "netbox_device_role" "test" {
  name        = "test-device-role"
  slug        = "test-device-role"
  color_hex   = "112233"
  vm_role     = true
  description = "test-description"
}
