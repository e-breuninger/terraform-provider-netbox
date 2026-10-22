data "netbox_device_interfaces" "test" {
  filters = [
    { name = "name", value = "test-device_interface" },
  ]
}
