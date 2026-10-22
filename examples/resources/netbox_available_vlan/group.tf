resource "netbox_vlan_group" "test" {
  name = "test-vlan-group"
  vid_ranges = [
    { start = 100, end = 199 },
  ]
}

# Allocates the next free VID of the group.
resource "netbox_available_vlan" "test" {
  name        = "test-vlan"
  group_id    = netbox_vlan_group.test.id
  status      = "reserved"
  description = "test-description"
}
