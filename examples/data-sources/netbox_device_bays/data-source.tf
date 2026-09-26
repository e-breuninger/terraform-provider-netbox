data "netbox_device_bays" "test" {
  filters = [
    { name = "name", value = "test-device_bay" },
  ]
}
