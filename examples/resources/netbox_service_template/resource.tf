resource "netbox_service_template" "test" {
  name        = "test"
  protocol    = "tcp"
  ports       = [8080, 8443]
  description = "my description"
  comments    = "some comments"
}
