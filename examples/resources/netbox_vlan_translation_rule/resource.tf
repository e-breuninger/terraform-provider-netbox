resource "netbox_vlan_translation_policy" "test" {
  name = "test"
}

resource "netbox_vlan_translation_rule" "test" {
  policy_id   = netbox_vlan_translation_policy.test.id
  local_vid   = 100
  remote_vid  = 200
  description = "my description"
}
