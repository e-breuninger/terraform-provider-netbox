data "netbox_virtual_machine_interfaces" "test" {
  filters = [
    { name = "name", value = "test-virtual_machine_interface" },
  ]
}
