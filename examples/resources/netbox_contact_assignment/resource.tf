resource "netbox_contact" "test" {
  name = "test"
}

resource "netbox_contact_role" "test" {
  name = "test"
}

// Assumes that a device with id 123 exists
resource "netbox_contact_assignment" "test" {
  content_type = "dcim.device"
  object_id    = 123
  contact_id   = netbox_contact.test.id
  role_id      = netbox_contact_role.test.id
  priority     = "primary"
}
