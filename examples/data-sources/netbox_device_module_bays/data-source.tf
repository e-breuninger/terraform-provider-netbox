data "netbox_device_module_bays" "test" {
  filters = [
    { name = "name", value = "test-device_module_bay" },
  ]
}
