data "netbox_device_rear_ports" "test" {
  filters = [
    { name = "name", value = "test-device_rear_port" },
  ]
}
