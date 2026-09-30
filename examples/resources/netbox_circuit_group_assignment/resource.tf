resource "netbox_circuit_provider" "providera" {
    name = "ProviderA"
    slug = "providera"
}

resource "netbox_circuit_type" "typea" {
    name = "TypeA"
    slug = "typea"
}

resource "netbox_circuit" "circuita" {
    cid         = "CID12345"
    status      = "active"
    provider_id = netbox_circuit_provider.providera.id
    type_id     = netbox_circuit_type.typea.id
}

resource "netbox_circuit_group" "groupa" {
    name = "GroupA"
    slug = "groupa"
}

resource "netbox_circuit_group_assignment" "assignmenta" {
    group_id    = netbox_circuit_group.groupa.id
    member_type = "circuits.circuit"
    member_id   = netbox_circuit.circuita.id
    priority    = "primary"
}
