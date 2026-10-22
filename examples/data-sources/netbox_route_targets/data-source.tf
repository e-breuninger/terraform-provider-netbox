data "netbox_route_targets" "test" {
  filters = [
    { name = "name", value = "test-route_target" },
  ]
}
