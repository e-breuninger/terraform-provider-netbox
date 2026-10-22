data "netbox_event_rules" "test" {
  filters = [
    { name = "name", value = "test-event_rule" },
  ]
}
