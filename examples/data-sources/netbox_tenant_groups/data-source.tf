data "netbox_tenant_groups" "test" {
  filters = [
    { name = "name", value = "test-tenant_group" },
  ]
}
