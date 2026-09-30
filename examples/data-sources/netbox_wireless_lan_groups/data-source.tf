data "netbox_wireless_lan_groups" "test" {
  filters = [
    { name = "name", value = "test-wireless_lan_group" },
  ]
}
