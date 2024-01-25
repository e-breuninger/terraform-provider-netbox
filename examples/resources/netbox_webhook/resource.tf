resource "netbox_webhook" "test" {
  name              = "test"
  payload_url       = "https://example.com/webhook"
  bodytemplate      = "Sample body"
}
