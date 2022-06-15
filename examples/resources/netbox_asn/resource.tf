resource "netbox_rir" "test" {
  name = "testrir"
}

resource "netbox_asn" "test" {
  asn    = 1337
  rir_id = netbox_rir.test.id
}
