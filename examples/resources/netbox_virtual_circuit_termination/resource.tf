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

resource "netbox_device" "test" {
  name           = "test-device"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

# Virtual circuits terminate on virtual interfaces only.
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "test-device-interface"
  type      = "virtual"
}

resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_provider_network" "test" {
  name                = "test-circuit-provider-network"
  circuit_provider_id = netbox_circuit_provider.test.id
}

resource "netbox_circuit_provider_account" "test" {
  account             = "test-account"
  circuit_provider_id = netbox_circuit_provider.test.id
}

resource "netbox_virtual_circuit_type" "test" {
  name = "test-virtual-circuit-type"
}

resource "netbox_virtual_circuit" "test" {
  cid                         = "test-cid"
  circuit_provider_network_id = netbox_circuit_provider_network.test.id
  virtual_circuit_type_id     = netbox_virtual_circuit_type.test.id
}

resource "netbox_virtual_circuit_termination" "test" {
  virtual_circuit_id  = netbox_virtual_circuit.test.id
  device_interface_id = netbox_device_interface.test.id
  role                = "hub"
  description         = "test-description"
}
