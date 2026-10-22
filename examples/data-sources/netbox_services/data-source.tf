data "netbox_services" "test" {
  filters = [
    { name = "name", value = "test-service" },
  ]
}
