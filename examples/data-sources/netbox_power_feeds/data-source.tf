data "netbox_power_feeds" "test" {
  filters = [
    { name = "name", value = "test-power_feed" },
  ]
}
