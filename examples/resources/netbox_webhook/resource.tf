resource "netbox_webhook" "test" {
  name              = "test"
  enabled           = "true"
  trigger_on_create = true
  payload_url       = "https://example.com/webhook"
  content_types     = ["dcim.site"]
  bodytemplate      = "Sample body"
}
