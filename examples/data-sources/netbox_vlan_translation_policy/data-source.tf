// Get VLAN translation policy by name
data "netbox_vlan_translation_policy" "test" {
  name = "test"
}

// Get VLAN translation policy by id
data "netbox_vlan_translation_policy" "test_by_id" {
  id = "1"
}
