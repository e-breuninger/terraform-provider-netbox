resource "netbox_vlan_translation_policy" "test" {
  name = "test-vlan-translation-policy"
}

resource "netbox_vlan_translation_rule" "test" {
  vlan_translation_policy_id = netbox_vlan_translation_policy.test.id
  local_vid                  = 100
  remote_vid                 = 200
  description                = "test-description"
}
