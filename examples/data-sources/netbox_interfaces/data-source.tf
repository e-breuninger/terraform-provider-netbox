data "netbox_interface" "myvm_eth0" {
  name_regex = "eth0"
  filter {
    name  = "name"
    value = "myvm"
  }
}
