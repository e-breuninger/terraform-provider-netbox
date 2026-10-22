data "netbox_fhrp_groups" "test" {
  filters = [
    { name = "name", value = "test-fhrp_group" },
  ]
}
