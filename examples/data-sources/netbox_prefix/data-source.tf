//Retrieve resource by cidr
resource "netbox_prefix" "cidr" {
    cidr = "10.0.0.0/16"
}

//Retrieve resource by description
resource "netbox_prefix" "description" {
    description = "prod-eu-west-1a"
}