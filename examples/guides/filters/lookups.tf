# Devices whose name contains "core" (case-insensitive) and which are not decommissioning.
data "netbox_devices" "core" {
  filters = [
    { name = "name__ic", value = "core" },
    { name = "status__n", value = "decommissioning" },
  ]
}

# VLANs with a VID of 100 or higher that have a name.
data "netbox_vlans" "numbered" {
  filters = [
    { name = "vid__gte", value = "100" },
    { name = "name__empty", value = "false" },
  ]
}
