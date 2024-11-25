resource "netbox_user" "test" {
  username = "johndoe"
  password = "Abcdefghijkl1"
}

resource "netbox_token" "test_basic" {
  user_id       = netbox_user.test.id
  key           = "0123456789012345678901234567890123456789"
  allowed_ips   = ["2.4.8.16/32"]
  write_enabled = false
}
