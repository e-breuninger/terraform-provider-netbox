# Was filter blocks, name_regex and the vms result list.
data "netbox_virtual_machines" "web" {
  filters = [
    { name = "status", value = "active" },
    { name = "name__regex", value = "^web-[0-9]+$" },
  ]
}

output "web_ids" {
  # Was vm_id on each item.
  value = data.netbox_virtual_machines.web.virtual_machines[*].id
}
