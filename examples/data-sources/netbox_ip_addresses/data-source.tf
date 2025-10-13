data "netbox_ip_addresses" "filtered_ip_addresses" {
  filter {
    name  = "tag"
    value = "DMZ"
  }
}
