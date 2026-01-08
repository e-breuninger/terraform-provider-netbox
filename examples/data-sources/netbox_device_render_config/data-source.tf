# Get the rendered configuration for a device
data "netbox_device_render_config" "server_config" {
  device_id = 60
}

# Use the rendered configuration
output "rendered_config" {
  value = data.netbox_device_render_config.server_config.content
}

output "template_used" {
  value = data.netbox_device_render_config.server_config.config_template_name
}

# Example: Write the config to a file using local_file resource
# resource "local_file" "kickstart" {
#   content  = data.netbox_device_render_config.server_config.content
#   filename = "${path.module}/kickstart.cfg"
# }

