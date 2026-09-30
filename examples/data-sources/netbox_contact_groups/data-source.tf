data "netbox_contact_groups" "test" {
  filters = [
    { name = "name", value = "test-contact_group" },
  ]
}
