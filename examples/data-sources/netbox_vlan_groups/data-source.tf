data "netbox_vlan_groups" "test" {
  filters = [
    { name = "name", value = "test-vlan_group" },
  ]
}
