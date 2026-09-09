resource "netbox_user" "test" {
  username = "johndoe"
  password = "Abcdefghijkl1"
}

resource "netbox_token" "test_basic" {
  user_id       = netbox_user.test.id
  token         = "0123456789012345678901234567890123456789"
  write_enabled = false
  expires       = "2036-01-02T15:04:05.000Z"
}
