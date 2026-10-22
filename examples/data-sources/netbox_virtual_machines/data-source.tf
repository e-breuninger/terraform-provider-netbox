data "netbox_virtual_machines" "test" {
  filters = [
    { name = "name", value = "test-virtual_machine" },
  ]
}
