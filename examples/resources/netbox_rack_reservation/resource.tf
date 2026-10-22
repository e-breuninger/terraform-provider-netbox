resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_rack" "test" {
  name    = "test-rack"
  site_id = netbox_site.test.id
}

resource "netbox_user" "test" {
  username = "test-user"
  password = "test-password"
}

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_rack_reservation" "test" {
  rack_id     = netbox_rack.test.id
  units       = [1, 2, 3]
  user_id     = netbox_user.test.id
  tenant_id   = netbox_tenant.test.id
  description = "test-description"
}
