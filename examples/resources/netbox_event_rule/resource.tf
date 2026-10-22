resource "netbox_webhook" "test" {
  name        = "test-webhook"
  payload_url = "https://example.com/test"
}

resource "netbox_event_rule" "test" {
  name               = "test-event-rule"
  object_types       = ["dcim.site", "dcim.device"]
  event_types        = ["object_created", "object_updated"]
  enabled            = true
  conditions         = jsonencode({ attr = "status", value = "active" })
  action_type        = "webhook"
  action_object_type = "extras.webhook"
  action_object_id   = netbox_webhook.test.id
  description        = "test-description"
}
