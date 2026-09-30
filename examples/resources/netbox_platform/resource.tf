resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_platform" "test" {
  name            = "test-platform"
  slug            = "test-platform"
  manufacturer_id = netbox_manufacturer.test.id
  description     = "test-description"
}
