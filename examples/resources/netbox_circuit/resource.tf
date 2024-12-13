resource "netbox_tenant" "test" {
  name = "test"
}

resource "netbox_circuit_provider" "test" {
  name = "test"
}

resource "netbox_circuit_type" "test" {
  name = "test"
}

resource "netbox_circuit" "test" {
  cid         = "test"
  status      = "active"
  provider_id = netbox_circuit_provider.test.id
  type_id     = netbox_circuit_type.test.id
}
