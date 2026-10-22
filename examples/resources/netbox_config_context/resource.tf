resource "netbox_config_context_profile" "test" {
  name = "test-config-context-profile"
}

resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_device_role" "test" {
  name = "test-device-role"
}

resource "netbox_tag" "test" {
  name = "test-tag"
}

resource "netbox_config_context" "test" {
  name            = "test-config-context"
  weight          = 500
  description     = "test-description"
  is_active       = true
  data            = jsonencode({ ntp_servers = ["192.0.2.1", "192.0.2.2"] })
  profile_id      = netbox_config_context_profile.test.id
  site_ids        = [netbox_site.test.id]
  device_role_ids = [netbox_device_role.test.id]
  tag_slugs       = [netbox_tag.test.slug]
}
