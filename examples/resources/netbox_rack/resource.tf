resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_location" "test" {
  name    = "test-location"
  site_id = netbox_site.test.id
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_rack_role" "test" {
  name = "test-rack-role"
}

resource "netbox_rack" "test" {
  name        = "test-rack"
  site_id     = netbox_site.test.id
  location_id = netbox_location.test.id
  tenant_id   = netbox_tenant.test.id
  role_id     = netbox_rack_role.test.id
  status      = "active"
  form_factor = "4-post-cabinet"
  width       = 19
  u_height    = 42
  desc_units  = false
  serial      = "test-serial"
  asset_tag   = "test-asset-tag"
  facility_id = "test-facility-id"
  outer_width = 600
  outer_depth = 1200
  outer_unit  = "mm"
  description = "test-description"
}
