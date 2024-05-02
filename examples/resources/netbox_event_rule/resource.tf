resource "netbox_webhook" "test" {
  name        = "my-webhook"
  payload_url = "https://example.com/webhook"
}

resource "netbox_event_rule" "test" {
  name              = "my-event-rule"
  content_types     = ["dcim.site", "virtualization.cluster"]
  action_type       = "webhook"
  action_object_id  = netbox_webhook.test.id
  trigger_on_create = true
}
