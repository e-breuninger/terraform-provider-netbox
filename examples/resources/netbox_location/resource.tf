resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_location" "test" {
  name        = "test-location"
  slug        = "test-location"
  site_id     = netbox_site.test.id
  status      = "active"
  facility    = "test-facility"
  tenant_id   = netbox_tenant.test.id
  description = "test-description"
}

resource "netbox_location" "test_child" {
  name      = "test-location-child"
  site_id   = netbox_site.test.id
  parent_id = netbox_location.test.id
}
