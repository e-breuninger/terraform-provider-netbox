data "netbox_device_bay_templates" "test" {
  filters = [
    { name = "name", value = "test-device_bay_template" },
  ]
}
