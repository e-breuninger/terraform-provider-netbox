resource "netbox_config_template" "test" {
  name               = "test-config-template"
  description        = "test-description"
  template_code      = "hostname {{ device.name }}"
  environment_params = jsonencode({ trim_blocks = true })
  mime_type          = "text/x-config"
  file_name          = "test-file"
  file_extension     = "cfg"
  as_attachment      = true
}
