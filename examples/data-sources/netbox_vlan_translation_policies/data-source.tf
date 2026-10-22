data "netbox_vlan_translation_policies" "test" {
  filters = [
    { name = "name", value = "test-vlan_translation_policy" },
  ]
}
