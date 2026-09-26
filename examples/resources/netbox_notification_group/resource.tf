resource "netbox_group" "test" {
  name = "test-group"
}

resource "netbox_user" "test" {
  username = "test-user"
  password = "test-password"
}

resource "netbox_notification_group" "test" {
  name        = "test-notification-group"
  description = "test-description"
  group_ids   = [netbox_group.test.id]
  user_ids    = [netbox_user.test.id]
}
