data "netbox_ike_proposals" "test" {
  filters = [
    { name = "name", value = "test-ike_proposal" },
  ]
}
