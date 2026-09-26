resource "netbox_contact_group" "test_a" {
  name        = "test-contact-group-a"
  description = "test-description"
}

resource "netbox_contact_group" "test_b" {
  name      = "test-contact-group-b"
  parent_id = netbox_contact_group.test_a.id
}
