resource "netbox_site" "example1" {
  name      = "Example site 1"
  asn       = 1337
  facility  = "Data center"
  latitude  = "-45.4085"
  longitude = "30.1496"
  status    = "staging"
  timezone  = "Africa/Johannesburg"
}
