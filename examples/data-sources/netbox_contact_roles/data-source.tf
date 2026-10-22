data "netbox_contact_roles" "test" {
  filters = [
    { name = "name", value = "test-contact_role" },
  ]
}
