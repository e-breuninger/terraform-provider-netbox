data "netbox_device_roles" "test" {
  filters = [
    { name = "name", value = "test-device_role" },
  ]
}
