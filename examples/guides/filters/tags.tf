# Devices carrying both tags.
data "netbox_devices" "monitored_and_backed_up" {
  filters = [
    { name = "tag", value = "monitored" },
    { name = "tag", value = "backup" },
  ]
}

# Devices carrying at least one of the tags.
data "netbox_devices" "monitored_or_backed_up" {
  filters = [
    { name = "tag__any", value = "monitored" },
    { name = "tag__any", value = "backup" },
  ]
}

# Devices carrying neither tag.
data "netbox_devices" "neither" {
  filters = [
    { name = "tag__n", value = "monitored" },
    { name = "tag__n", value = "backup" },
  ]
}
