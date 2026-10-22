# Repeating a name combines its values with a logical OR: devices in either status.
data "netbox_devices" "planned_or_staged" {
  filters = [
    { name = "status", value = "planned" },
    { name = "status", value = "staged" },
  ]
}
