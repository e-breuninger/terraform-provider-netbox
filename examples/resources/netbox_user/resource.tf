resource "netbox_group" "test" {
  name = "test-group"
}

resource "netbox_user" "test" {
  username   = "test-user"
  password   = "test-password"
  first_name = "test-first-name"
  last_name  = "test-last-name"
  email      = "test@example.com"
  is_active  = true
  group_ids  = [netbox_group.test.id]
}
