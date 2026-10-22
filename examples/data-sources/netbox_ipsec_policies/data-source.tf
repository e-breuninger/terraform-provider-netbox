data "netbox_ipsec_policies" "test" {
  filters = [
    { name = "name", value = "test-ipsec_policy" },
  ]
}
