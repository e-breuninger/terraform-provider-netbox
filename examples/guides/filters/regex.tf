# Up to 50 active devices named like "<site>-core-<number>". NetBox matches the name, so the
# limit applies to the matching devices.
data "netbox_devices" "core" {
  filters = [
    { name = "status", value = "active" },
    { name = "name__regex", value = "^[a-z]+-core-[0-9]+$" },
  ]
  limit = 50
}

output "core_device_names" {
  value = data.netbox_devices.core.devices[*].name
}
