data "netbox_fhrp_group" "test" {
  protocol    = "vrrp"
  group_id    = 1234
}