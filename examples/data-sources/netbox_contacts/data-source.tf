data "netbox_contacts" "test" {
  filters = [
    { name = "name", value = "test-contact" },
  ]
}
