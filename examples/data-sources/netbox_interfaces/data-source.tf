data "netbox_interfaces" "myvm_eth0" {
  name_regex = "eth0"
  filter {
    name  = "name"
    value = "myvm"
  }
}
