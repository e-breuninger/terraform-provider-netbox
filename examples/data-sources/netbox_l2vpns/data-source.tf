data "netbox_l2vpns" "test" {
  filters = [
    { name = "name", value = "test-l2vpn" },
  ]
}
