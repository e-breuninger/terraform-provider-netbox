resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_contact" "test" {
  name = "test-contact"
}

resource "netbox_contact_role" "test" {
  name = "test-contact-role"
}

resource "netbox_contact_assignment" "test" {
  object_type = "dcim.site"
  object_id   = netbox_site.test.id
  contact_id  = netbox_contact.test.id
  role_id     = netbox_contact_role.test.id
  priority    = "primary"
}
