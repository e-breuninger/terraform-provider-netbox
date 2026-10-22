resource "netbox_wireless_lan_group" "test_a" {
  name = "test-wireless-lan-group-parent"
}

resource "netbox_wireless_lan_group" "test" {
  name        = "test-wireless-lan-group"
  slug        = "test-wireless-lan-group"
  parent_id   = netbox_wireless_lan_group.test_a.id
  description = "test-description"
}
