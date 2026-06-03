# Filter by name
data "netbox_virtual_disk" "disk_by_name" {
  filter {
    name  = "name"
    value = "disk1"
  }
}

# Filter by virtual machine ID
data "netbox_virtual_disk" "disk_by_vm" {
  filter {
    name  = "virtual_machine_id"
    value = netbox_virtual_machine.example.id
  }
}

# Filter by virtual machine ID, ordered by ID (creation order)
data "netbox_virtual_disk" "disk_by_vm_ordered" {
  filter {
    name  = "virtual_machine_id"
    value = netbox_virtual_machine.example.id
  }
  ordering = ["id"]
}

# Filter by tag
data "netbox_virtual_disk" "disk_by_tag" {
  filter {
    name  = "tag"
    value = "production"
  }
}

# Multiple filters
data "netbox_virtual_disk" "disk_filtered" {
  filter {
    name  = "name"
    value = "disk1"
  }
  filter {
    name  = "tag"
    value = "production"
  }
}

# Filter with name regex
data "netbox_virtual_disk" "disk_regex" {
  name_regex = "^disk[0-9]+"
  limit      = 10
}
