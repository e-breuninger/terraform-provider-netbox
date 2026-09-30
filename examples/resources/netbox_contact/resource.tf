resource "netbox_contact_group" "test" {
  name = "test-contact-group"
}

resource "netbox_contact" "test" {
  name        = "test-contact"
  title       = "test-title"
  phone       = "+1 555 0100"
  email       = "test-contact@example.com"
  address     = "test-address"
  link        = "https://example.com/test-contact"
  description = "test-description"
  group_ids   = [netbox_contact_group.test.id]
}
