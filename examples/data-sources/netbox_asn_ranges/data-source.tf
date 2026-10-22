data "netbox_asn_ranges" "test" {
  filters = [
    { name = "name", value = "test-asn_range" },
  ]
}
