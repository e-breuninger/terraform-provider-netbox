resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_route_target" "test_a" {
  name = "65000:1"
}

resource "netbox_route_target" "test_b" {
  name = "65000:2"
}

resource "netbox_vrf" "test" {
  name              = "test-vrf"
  rd                = "65000:1"
  description       = "test-description"
  tenant_id         = netbox_tenant.test.id
  import_target_ids = [netbox_route_target.test_a.id]
  export_target_ids = [netbox_route_target.test_a.id, netbox_route_target.test_b.id]
}
