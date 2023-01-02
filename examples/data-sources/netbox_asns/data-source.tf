data "netbox_asns" "asns" {
  filter {
    name = "asn__gte"
    value = "1000"
  }
  filter {
    name = "asn__lte"
    value = "2000"
  }
}
