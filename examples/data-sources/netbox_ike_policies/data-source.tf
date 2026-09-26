data "netbox_ike_policies" "test" {
  filters = [
    { name = "name", value = "test-ike_policy" },
  ]
}
