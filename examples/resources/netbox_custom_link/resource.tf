resource "netbox_custom_link" "device_docs" {
  name          = "device-docs"
  content_types = ["dcim.device"]
  link_text     = "View documentation"
  link_url      = "https://wiki.example.com/devices/{{ object.name }}"
  enabled       = true
  weight        = 100
  group_name    = "documentation"
  button_class  = "blue"
  new_window    = true
}
