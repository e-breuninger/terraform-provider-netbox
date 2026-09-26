data "netbox_virtual_machine_types" "test" {
  filters = [
    { name = "name", value = "test-virtual_machine_type" },
  ]
}
