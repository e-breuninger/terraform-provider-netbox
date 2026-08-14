// Get VLAN translation rule by id
data "netbox_vlan_translation_rule" "test_by_id" {
  id = "1"
}

// Get VLAN translation rule by policy and local VLAN ID
data "netbox_vlan_translation_rule" "test" {
  policy_id = netbox_vlan_translation_policy.test.id
  local_vid = 100
}
