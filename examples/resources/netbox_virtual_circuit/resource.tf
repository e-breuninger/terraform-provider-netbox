resource "netbox_circuit_provider" "provider_a" {
  name = "Provider A"
  slug = "provider-a"
}

resource "netbox_circuit_provider_network" "network_a" {
  name        = "Network A"
  provider_id = netbox_circuit_provider.provider_a.id
}

resource "netbox_virtual_circuit_type" "evpl" {
  name = "EVPL"
  slug = "evpl"
}

resource "netbox_virtual_circuit" "example" {
  cid                 = "VC-0001"
  provider_network_id = netbox_circuit_provider_network.network_a.id
  type_id             = netbox_virtual_circuit_type.evpl.id
  status              = "active"
  description         = "Example virtual circuit"
}
