resource "netbox_fhrp_group" "test" {
  protocol    = "vrrp"
  group_id    = 1234
  auth_type   = "md5"
  auth_key    = "SuperSecretKey"
  name        = "test-fhrp-group"
  description = "This is a test group"
}
