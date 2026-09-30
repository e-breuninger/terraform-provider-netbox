resource "netbox_region" "test" {
  name = "test-region"
}

resource "netbox_site_group" "test" {
  name = "test-site-group"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_rir" "test" {
  name = "test-rir"
}

resource "netbox_asn" "test" {
  asn    = 65000
  rir_id = netbox_rir.test.id
}

resource "netbox_tag" "test" {
  name = "test-tag"
}

# asn_ids set: the site has exactly these ASNs. An ASN that attaches itself
# to this site with site_ids is removed again, so list every ASN here.
resource "netbox_site" "with_asn" {
  name             = "test-site"
  slug             = "test-site"
  status           = "active"
  facility         = "test-facility"
  physical_address = "test-physical-address"
  shipping_address = "test-shipping-address"
  timezone         = "Europe/Berlin"
  latitude         = 52.5
  longitude        = 13.4
  region_id        = netbox_region.test.id
  group_id         = netbox_site_group.test.id
  tenant_id        = netbox_tenant.test.id
  asn_ids          = [netbox_asn.test.id]
  tags             = [netbox_tag.test.slug]
  description      = "test-description"
}

# asn_ids empty: all ASNs are removed from the site.
resource "netbox_site" "asn_explicit_clear" {
  name    = "test-site-without-asns"
  asn_ids = []
}

# asn_ids unset: the ASNs of the site are left alone, e.g. when ASNs attach
# themselves with site_ids.
resource "netbox_site" "no_asn" {
  name = "test-site-asns-unmanaged"
}
