resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_type" "test" {
  name = "test-circuit-type"
}

resource "netbox_circuit" "test" {
  cid                 = "test-cid"
  circuit_provider_id = netbox_circuit_provider.test.id
  circuit_type_id     = netbox_circuit_type.test.id
}

resource "netbox_circuit_group" "test" {
  name = "test-circuit-group"
}

resource "netbox_circuit_group_assignment" "test" {
  circuit_group_id = netbox_circuit_group.test.id
  member_type      = "circuits.circuit"
  member_id        = netbox_circuit.test.id
  priority         = "primary"
}
