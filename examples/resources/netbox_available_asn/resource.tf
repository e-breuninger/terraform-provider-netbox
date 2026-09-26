resource "netbox_rir" "test" {
  name = "test-rir"
}

resource "netbox_asn_range" "test" {
  name   = "test-asn-range"
  slug   = "test-asn-range"
  rir_id = netbox_rir.test.id
  start  = 64512
  end    = 64521
}

resource "netbox_ipam_role" "test" {
  name = "test-ipam-role"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

# Allocates the next free ASN of the range; the RIR comes from the range.
resource "netbox_available_asn" "test" {
  asn_range_id = netbox_asn_range.test.id
  role_id      = netbox_ipam_role.test.id
  tenant_id    = netbox_tenant.test.id
  description  = "test-description"
}
