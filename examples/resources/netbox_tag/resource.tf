resource "netbox_tag" "test" {
  name        = "test-tag"
  slug        = "test-tag"
  color_hex   = "112233"
  description = "test-description"
}
