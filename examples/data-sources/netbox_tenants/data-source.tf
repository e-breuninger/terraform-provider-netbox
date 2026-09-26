data "netbox_tenants" "test" {
  filters = [
    { name = "name", value = "test-tenant" },
  ]
}
