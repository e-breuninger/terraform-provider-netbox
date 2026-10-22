data "netbox_ipam_roles" "test" {
  filters = [
    { name = "name", value = "test-ipam_role" },
  ]
}
