resource "netbox_rir" "test" {
  name = "test-rir"
}

resource "netbox_ipam_role" "test" {
  name = "test-ipam-role"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_asn" "test" {
  asn         = 64512
  rir_id      = netbox_rir.test.id
  role_id     = netbox_ipam_role.test.id
  tenant_id   = netbox_tenant.test.id
  description = "test-description"
}
