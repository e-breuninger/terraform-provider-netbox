resource "netbox_user" "test" {
  username = "test-user"
  password = "test-password"
}

resource "netbox_token" "test" {
  user_id       = netbox_user.test.id
  description   = "test-description"
  write_enabled = false
  enabled       = true
  expires       = "2030-01-01T00:00:00.000Z"
}
