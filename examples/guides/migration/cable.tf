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

resource "netbox_device" "a" {
  name           = "switch-a"
  device_type_id = netbox_device_type.switch.id
  role_id        = netbox_device_role.switch.id
  site_id        = netbox_site.lab.id
}

resource "netbox_device" "b" {
  name           = "switch-b"
  device_type_id = netbox_device_type.switch.id
  role_id        = netbox_device_role.switch.id
  site_id        = netbox_site.lab.id
}

resource "netbox_device_interface" "a" {
  device_id = netbox_device.a.id
  name      = "uplink"
  type      = "1000base-t"
}

resource "netbox_device_interface" "b" {
  device_id = netbox_device.b.id
  name      = "uplink"
  type      = "1000base-t"
}

# Was one a_termination and one b_termination block per end, each with object_type and object_id.
resource "netbox_cable" "uplink" {
  a_side = { device_interface_ids = [netbox_device_interface.a.id] }
  b_side = { device_interface_ids = [netbox_device_interface.b.id] }
}
