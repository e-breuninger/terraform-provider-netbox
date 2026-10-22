data "netbox_virtual_device_contexts" "test" {
  filters = [
    { name = "name", value = "test-virtual_device_context" },
  ]
}
