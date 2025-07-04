resource "netbox_tenant" "example" {
  name = "example_tenant"
}

resource "netbox_site" "example" {
  name = "example_site"
  status = "active"
}

resource "netbox_manufacturer" "example" {
  name = "example_manufacturer"
}

resource "netbox_device_type" "example" {
  model = "example_device_type"
  manufacturer_id = netbox_manufacturer.example.id
  subdevice_role = "parent"
}

resource "netbox_device_type" "example_installed" {
  model = "example_device_type_installed"
  manufacturer_id = netbox_manufacturer.example.id
  u_height = 0
  subdevice_role = "child"
}

resource "netbox_device_role" "example" {
  name = "example_role"
  color_hex = "123456"
}

resource "netbox_device" "example" {
  name = "example_device"
  device_type_id = netbox_device_type.example.id
  tenant_id = netbox_tenant.example.id
  role_id = netbox_device_role.example.id
  site_id = netbox_site.example.id
}

resource "netbox_device" "example_installed" {
  name = "example_device_installed"
  device_type_id = netbox_device_type.example_installed.id
  tenant_id = netbox_tenant.example.id
  role_id = netbox_device_role.example.id
  site_id = netbox_site.example.id
}

resource "netbox_device_bay" "example" {
  device_id = netbox_device.example.id
  name = "example_device_bay"
  label = "example_label"
  description = "example_description"
  installed_device_id = netbox_device.example_installed.id
}
