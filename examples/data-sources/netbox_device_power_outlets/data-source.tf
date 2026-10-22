data "netbox_device_power_outlets" "test" {
  filters = [
    { name = "name", value = "test-device_power_outlet" },
  ]
}
