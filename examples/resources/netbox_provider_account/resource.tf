resource "netbox_circuit_provider" "providera" {
    name        = "ProviderA"
    slug        = "providera"
}

resource "netbox_provider_account" "accounta" {
    account     = "12345"
    name        = "AccountA"
    provider_id = netbox_circuit_provider.providera.id
}
