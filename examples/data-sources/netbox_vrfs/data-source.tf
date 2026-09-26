data "netbox_vrfs" "test" {
  filters = [
    { name = "name", value = "test-vrf" },
  ]
}
