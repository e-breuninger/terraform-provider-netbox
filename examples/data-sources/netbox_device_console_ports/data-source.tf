data "netbox_device_console_ports" "test" {
  filters = [
    { name = "name", value = "test-device_console_port" },
  ]
}
