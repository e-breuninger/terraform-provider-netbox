data "netbox_ipsec_profiles" "test" {
  filters = [
    { name = "name", value = "test-ipsec_profile" },
  ]
}
