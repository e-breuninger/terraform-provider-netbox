data "netbox_ip_range" "cust_a_prod" {
  contains = "10.0.0.1/24"
}
