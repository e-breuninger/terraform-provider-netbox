resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
}

resource "netbox_circuit_type" "test" {
  name = "%[1]s"
}

resource "netbox_circuit" "test" {
  cid         = "%[1]s"
  status      = "active"
  provider_id = netbox_circuit_provider.test.id
  type_id     = netbox_circuit_type.test.id
}

resource "netbox_circuit_termination" "test" {
  circuit_id     = netbox_circuit.test.id
  term_side      = "A"
  site_id        = netbox_site.test.id
  port_speed     = 100000
  upstream_speed = 50000
}
