data "netbox_webhooks" "test" {
  filters = [
    { name = "name", value = "test-webhook" },
  ]
}
