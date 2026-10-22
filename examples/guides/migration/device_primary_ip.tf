resource "netbox_site" "lab" {
  name = "lab"
}

resource "netbox_manufacturer" "acme" {
  name = "acme"
}

resource "netbox_device_type" "switch" {
  manufacturer_id = netbox_manufacturer.acme.id
  model           = "switch"
}

resource "netbox_device_role" "switch" {
  name = "switch"
}

resource "netbox_device" "switch" {
  name           = "switch"
  device_type_id = netbox_device_type.switch.id
  role_id        = netbox_device_role.switch.id
  site_id        = netbox_site.lab.id
}

resource "netbox_device_interface" "mgmt" {
  device_id = netbox_device.switch.id
  name      = "mgmt"
  type      = "1000base-t"
}

# Was interface_id plus object_type = "dcim.interface".
resource "netbox_ip_address" "switch" {
  ip_address          = "192.0.2.20/24"
  device_interface_id = netbox_device_interface.mgmt.id
}

# Was netbox_device_primary_ip.
resource "netbox_primary_ip" "switch" {
  device_id     = netbox_device.switch.id
  ip_address_id = netbox_ip_address.switch.id
}
