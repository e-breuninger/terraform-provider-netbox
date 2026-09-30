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

resource "netbox_site" "test" {
  name = "test-site"
}

# Each resource allocates the next free ASN of the range; the RIR comes from
# the range.

# site_ids set: the ASN is attached to exactly these sites. Other ASNs can
# attach themselves to the same site in the same way.
resource "netbox_available_asn" "to_site" {
  asn_range_id = netbox_asn_range.test.id
  role_id      = netbox_ipam_role.test.id
  tenant_id    = netbox_tenant.test.id
  site_ids     = [netbox_site.test.id]
  description  = "test-description"
}

# site_ids empty: the ASN is detached from all of its sites. Other ASNs of
# those sites stay attached.
resource "netbox_available_asn" "site_explicit_clear" {
  asn_range_id = netbox_asn_range.test.id
  role_id      = netbox_ipam_role.test.id
  tenant_id    = netbox_tenant.test.id
  site_ids     = []
  description  = "test-description"
}

# site_ids unset: the sites of the ASN are left alone, e.g. when a site
# lists the ASN in asn_ids.
resource "netbox_available_asn" "no_site" {
  asn_range_id = netbox_asn_range.test.id
  role_id      = netbox_ipam_role.test.id
  tenant_id    = netbox_tenant.test.id
  description  = "test-description"
}
