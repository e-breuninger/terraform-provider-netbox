resource "netbox_webhook" "test" {
  name        = "my-webhook"
  payload_url = "https://example.com/webhook"
}

resource "netbox_event_rule" "webhook" {
  name             = "my-event-rule"
  content_types    = ["dcim.site", "virtualization.cluster"]
  action_type      = "webhook"
  action_object_id = netbox_webhook.test.id
  event_types = [
    "object_created",
    "object_updated",
    "object_deleted",
    "job_started",
    "job_completed",
    "job_failed",
    "job_errored"
  ]
}

resource "netbox_event_rule" "script" {
  name             = "my-script-event-rule"
  content_types    = ["dcim.site"]
  action_type      = "script"
  action_object_id = 42 # existing NetBox Script ID
  event_types      = ["object_created"]
}
