resource "netbox_config_context_profile" "test" {
  name        = "test-config-context-profile"
  description = "test-description"
  schema = jsonencode({
    properties = {
      ntp_servers = { type = "array" }
    }
  })
}
