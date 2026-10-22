resource "netbox_service_template" "test" {
  name        = "test-service-template"
  protocol    = "udp"
  ports       = [53]
  description = "test-description"
}
