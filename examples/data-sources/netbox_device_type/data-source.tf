# Get device type by model name
data "netbox_device_type" "ex1" {
  model = "7210 SAS-Sx 10/100GE"
}

# Get device type by slug
data "netbox_device_type" "ex2" {
  slug = "7210-sas-sx-10-100GE"
}

# Get device type by manufacturer and part number information
data "netbox_device_type" "ex3" {
  manufacturer = "Nokia"
  part_number  = "3HE11597AARB01"
}
