resource "netbox_webhook" "test" {
  name               = "test-webhook"
  payload_url        = "https://example.com/test"
  http_method        = "POST"
  http_content_type  = "application/json"
  additional_headers = "X-Test: test"
  body_template      = "{{ event }}"
  secret             = "test-secret"
  ssl_verification   = true
  description        = "test-description"
}
