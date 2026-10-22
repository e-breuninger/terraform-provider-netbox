data "netbox_devices" "test" {
  filters = [
    { name = "name", value = "test-device" },
  ]
}
