resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
  subdevice_role = "parent"
}

resource "netbox_device_type" "test_installed" {
  model = "%[1]s_installed"
  manufacturer_id = netbox_manufacturer.test.id
  u_height = 0
  subdevice_role = "child"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device" "test_installed" {
  name = "%[1]s_installed"
  device_type_id = netbox_device_type.test_installed.id
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  label = "%[1]s_label"
  description = "%[1]s_description"
  installed_device_id = netbox_device.test_installed.id
}