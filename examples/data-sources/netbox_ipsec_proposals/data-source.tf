data "netbox_ipsec_proposals" "test" {
  filters = [
    { name = "name", value = "test-ipsec_proposal" },
  ]
}
