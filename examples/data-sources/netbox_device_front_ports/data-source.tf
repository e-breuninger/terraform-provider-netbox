data "netbox_device_front_ports" "test" {
  filters = [
    { name = "name", value = "test-device_front_port" },
  ]
}
