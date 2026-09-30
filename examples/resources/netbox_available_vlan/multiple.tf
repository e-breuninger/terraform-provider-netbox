resource "netbox_vlan_group" "test" {
  name = "test-vlan-group"
  vid_ranges = [
    { start = 100, end = 199 },
  ]
}

# Two allocations from one group: the second waits for the first so NetBox hands out the VIDs in
# order.
resource "netbox_available_vlan" "test_a" {
  name     = "test-vlan-a"
  group_id = netbox_vlan_group.test.id
}

resource "netbox_available_vlan" "test_b" {
  depends_on = [netbox_available_vlan.test_a]
  name       = "test-vlan-b"
  group_id   = netbox_vlan_group.test.id
}
