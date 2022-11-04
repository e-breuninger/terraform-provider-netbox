data "netbox_asn" "asn_1" {
  asn = "1111"
  tag = "tag-1"
}

data "netbox_asn" "asn_2" {
  tag    = "tag-1"
  tag__n = "tag-2"
}
