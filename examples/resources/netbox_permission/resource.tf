resource "netbox_group" "test" {
  name = "test-group"
}

resource "netbox_user" "test" {
  username = "test-user"
  password = "test-password"
}

resource "netbox_permission" "test" {
  name         = "test-permission"
  description  = "test-description"
  enabled      = true
  object_types = ["dcim.device", "dcim.site"]
  actions      = ["view", "add"]
  group_ids    = [netbox_group.test.id]
  user_ids     = [netbox_user.test.id]
  constraints  = jsonencode({ status = "active" })
}
