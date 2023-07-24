resource "netbox_rir" "test" {
  name = "testrir"
}
resource "netbox_aggregate" "test" {
  prefix      = "1.1.1.0/25"
  description = "my description"
  rir_id      = netbox_rir.test.id
}
