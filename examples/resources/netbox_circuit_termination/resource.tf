resource "netbox_circuit_provider" "test" {
  name = "test-circuit-provider"
}

resource "netbox_circuit_type" "test" {
  name = "test-circuit-type"
}

resource "netbox_site" "test" {
  name = "test-site"
}

resource "netbox_circuit" "test" {
  cid                 = "test-cid"
  circuit_provider_id = netbox_circuit_provider.test.id
  circuit_type_id     = netbox_circuit_type.test.id
}

resource "netbox_circuit_termination" "test" {
  circuit_id          = netbox_circuit.test.id
  term_side           = "A"
  site_id             = netbox_site.test.id
  port_speed_kbps     = 1000000
  upstream_speed_kbps = 500000
  xconnect_id         = "test-xconnect-id"
  pp_info             = "test-pp-info"
  mark_connected      = true
  description         = "test-description"
}
