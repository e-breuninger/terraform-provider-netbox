resource "netbox_circuit_provider" "providera" {
    name        = "ProviderA"
    slug        = "providera"
}

resource "netbox_circuit_provider_network" "networka" {
    name        = "NetworkA"
    provider_id = netbox_circuit_provider.providera.id
}
