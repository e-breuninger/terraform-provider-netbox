resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_manufacturer" "test" {
  name = "test-manufacturer"
}

resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "test-device-type"
}

resource "netbox_device_role" "test" {
  name = "test-device-role"
}

resource "netbox_tag" "test" {
  name = "test-tag"
}

# The device is managed elsewhere or its tags are left alone: a netbox_device that declares tags
# strips every tag it does not list, so it must ignore them.
resource "netbox_device" "test" {
  name           = "test-device"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id

  lifecycle {
    ignore_changes = [tags]
  }
}

resource "netbox_device_tag" "test" {
  device_id = netbox_device.test.id
  tag_slug  = netbox_tag.test.slug
}
