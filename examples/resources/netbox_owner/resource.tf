resource "netbox_owner_group" "test" {
  name = "test-owner-group"
}

resource "netbox_group" "test" {
  name = "test-group"
}

resource "netbox_user" "test" {
  username = "test-user"
  password = "test-password"
}

resource "netbox_owner" "test" {
  name           = "test-owner"
  owner_group_id = netbox_owner_group.test.id
  user_ids       = [netbox_user.test.id]
  user_group_ids = [netbox_group.test.id]
  description    = "test-description"
}
