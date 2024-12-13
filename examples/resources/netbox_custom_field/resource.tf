resource "netbox_custom_field" "test" {
  name             = "test"
  type             = "text"
  content_types    = ["virtualization.vminterface"]
  weight           = 100
  validation_regex = "^.*$"
}
