resource "netbox_custom_link" "test" {
  name         = "test-custom-link"
  object_types = ["dcim.device", "dcim.site"]
  link_text    = "Open {{ object.name }}"
  link_url     = "https://example.com/test/{{ object.pk }}"
  enabled      = true
  weight       = 100
  group_name   = "test-group"
  button_class = "blue"
  new_window   = true
}
