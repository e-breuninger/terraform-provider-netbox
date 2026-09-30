resource "netbox_group" "test" {
  name = "network-admins"
}

resource "netbox_user" "test" {
  username = "jdoe"
  password = "Abcdefghijkl1"
  active   = true
}

resource "netbox_owner_group" "test" {
  name = "Infrastructure Team"
}

resource "netbox_owner" "test" {
  name           = "Network Operations"
  description    = "Owner for network infrastructure objects"
  group_id       = netbox_owner_group.test.id
  user_group_ids = [netbox_group.test.id]
  user_ids       = [netbox_user.test.id]
}
