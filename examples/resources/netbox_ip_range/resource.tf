resource "netbox_ip_range" "cust_a_prod" {
  start_address = "10.0.0.1/24"
  end_address   = "10.0.0.50/24"
  tags          = ["customer-a", "prod"]
}
