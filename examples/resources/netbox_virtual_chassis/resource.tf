resource "netbox_virtual_chassis" "test" {
  name        = "test-virtual-chassis"
  domain      = "test-domain"
  description = "test-description"
}
