data "netbox_device_power_ports" "test" {
  filters = [
    { name = "name", value = "test-device_power_port" },
  ]
}
