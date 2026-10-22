resource "netbox_fhrp_group" "test" {
  name        = "test-fhrp-group"
  protocol    = "vrrp3"
  group_id    = 10
  auth_type   = "plaintext"
  auth_key    = "test-key"
  description = "test-description"
}
