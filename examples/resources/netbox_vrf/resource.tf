resource "netbox_vrf" "cust_a_prod" {
  name = "cust-a-prod"
  tags = ["customer-a", "prod"]
}
