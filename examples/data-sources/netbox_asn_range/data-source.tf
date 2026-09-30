// Get ASN range by name
data "netbox_asn_range" "test" {
  name = "test"
}

// Get ASN range by id
data "netbox_asn_range" "test_by_id" {
  id = "1"
}
