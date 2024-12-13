resource "netbox_site" "test" {
  name = "test"
}

resource "netbox_tenant" "test" {
  name = "test"
}

resource "netbox_location" "test" {
  name        = "test"
  description = "my description"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
}
