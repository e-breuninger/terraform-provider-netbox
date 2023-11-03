data "netbox_tags" "all_tags" {
}

data "netbox_tags" "ansible_tags" {
  filter {
    name = "name__isw"
    value = "ansible_"
  }
}

data "netbox_tags" "not_ansible_tags" {
  filter {
    name = "name__nisw"
    value = "ansible_"
  }
}
