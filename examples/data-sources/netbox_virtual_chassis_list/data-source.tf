data "netbox_virtual_chassis_list" "test" {
  filters = [
    { name = "name", value = "test-virtual_chassis" },
  ]
}
