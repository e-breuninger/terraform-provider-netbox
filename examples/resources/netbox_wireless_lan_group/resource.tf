resource "netbox_wireless_lan_group" "parent" {
  name = "campus"
}

resource "netbox_wireless_lan_group" "child" {
  name      = "building-a"
  parent_id = netbox_wireless_lan_group.parent.id
}
