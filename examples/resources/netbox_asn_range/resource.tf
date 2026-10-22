resource "netbox_rir" "test" {
  name = "test-rir"
}

resource "netbox_asn_range" "test" {
  name        = "test-asn-range"
  slug        = "test-asn-range"
  rir_id      = netbox_rir.test.id
  start       = 64512
  end         = 64521
  description = "test-description"
}
