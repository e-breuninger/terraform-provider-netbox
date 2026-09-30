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

resource "netbox_tenant" "test" {
  name = "test-tenant"
}

resource "netbox_platform" "test" {
  name = "test-platform"
}

resource "netbox_location" "test" {
  name    = "test-location"
  site_id = netbox_site.test.id
}

resource "netbox_rack" "test" {
  name        = "test-rack"
  site_id     = netbox_site.test.id
  location_id = netbox_location.test.id
}

resource "netbox_device" "test" {
  name           = "test-device"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
  location_id    = netbox_location.test.id
  rack_id        = netbox_rack.test.id
  rack_position  = 10
  rack_face      = "front"
  status         = "active"
  airflow        = "front-to-rear"
  tenant_id      = netbox_tenant.test.id
  platform_id    = netbox_platform.test.id
  serial         = "test-serial"
  asset_tag      = "test-asset-tag"
  description    = "test-description"

  local_context_data = jsonencode({ ntp = ["192.0.2.1"] })
}
