resource "netbox_site" "example" {
  name   = "Example Site"
  status = "active"
}

resource "netbox_manufacturer" "example" {
  name = "Example Manufacturer"
}

resource "netbox_device_type" "example" {
  model           = "Example Type"
  manufacturer_id = netbox_manufacturer.example.id
}

resource "netbox_device_role" "example" {
  name      = "Example Role"
  color_hex = "123456"
}

resource "netbox_device" "example" {
  name           = "Example Device"
  device_type_id = netbox_device_type.example.id
  role_id        = netbox_device_role.example.id
  site_id        = netbox_site.example.id
}

resource "netbox_device_interface" "example" {
  device_id = netbox_device.example.id
  name      = "eth0"
  type      = "virtual"
}

resource "netbox_circuit_provider" "provider_a" {
  name = "Provider A"
  slug = "provider-a"
}

resource "netbox_circuit_provider_network" "network_a" {
  name        = "Network A"
  provider_id = netbox_circuit_provider.provider_a.id
}

resource "netbox_virtual_circuit_type" "evpl" {
  name = "EVPL"
  slug = "evpl"
}

resource "netbox_virtual_circuit" "example" {
  cid                 = "VC-0001"
  provider_network_id = netbox_circuit_provider_network.network_a.id
  type_id             = netbox_virtual_circuit_type.evpl.id
  status              = "active"
}

resource "netbox_virtual_circuit_termination" "example" {
  virtual_circuit_id = netbox_virtual_circuit.example.id
  interface_id        = netbox_device_interface.example.id
  role                = "peer"
}
