data "netbox_cable_bundles" "test" {
  filters = [
    { name = "name", value = "test-cable_bundle" },
  ]
}
