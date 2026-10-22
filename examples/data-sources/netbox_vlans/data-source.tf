data "netbox_vlans" "test" {
  filters = [
    { name = "name", value = "test-vlan" },
  ]
}
