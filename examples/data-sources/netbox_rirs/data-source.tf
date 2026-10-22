data "netbox_rirs" "test" {
  filters = [
    { name = "name", value = "test-rir" },
  ]
}
