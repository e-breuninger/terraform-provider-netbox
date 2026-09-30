resource "netbox_rir" "test" {
  name = "test"
}

resource "netbox_tenant" "test" {
  name = "test"
}

resource "netbox_asn_range" "test" {
  name        = "test"
  slug        = "test"
  rir_id      = netbox_rir.test.id
  start       = 65000
  end         = 65100
  tenant_id   = netbox_tenant.test.id
  description = "my description"
  comments    = "some comments"
}
