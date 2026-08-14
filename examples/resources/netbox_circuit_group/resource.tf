resource "netbox_tenant" "tenanta" {
    name = "TenantA"
}

resource "netbox_circuit_group" "groupa" {
    name      = "GroupA"
    slug      = "groupa"
    tenant_id = netbox_tenant.tenanta.id
}
