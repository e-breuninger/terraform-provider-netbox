resource "netbox_owner_group" "noc" {
  name = "NOC"
}

resource "netbox_group" "noc" {
  name = "noc-engineers"
}

resource "netbox_user" "jdoe" {
  username = "jdoe"
  password = "changeme"
}

resource "netbox_owner" "noc" {
  name           = "Network Operations Center"
  group_id       = netbox_owner_group.noc.id
  user_group_ids = [netbox_group.noc.id]
  user_ids       = [netbox_user.jdoe.id]
}
