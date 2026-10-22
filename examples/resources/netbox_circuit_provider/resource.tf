resource "netbox_rir" "test" {
  name = "test-rir"
}

resource "netbox_asn" "test" {
  asn    = 65000
  rir_id = netbox_rir.test.id
}

resource "netbox_circuit_provider" "test" {
  name        = "test-circuit-provider"
  slug        = "test-circuit-provider"
  description = "test-description"
  asn_ids     = [netbox_asn.test.id]
}
