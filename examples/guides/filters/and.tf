data "netbox_site" "berlin" {
  name = "Berlin"
}

# Entries with different names combine with a logical AND: active devices at the site.
data "netbox_devices" "active_in_berlin" {
  filters = [
    { name = "site_id", value = data.netbox_site.berlin.id },
    { name = "status", value = "active" },
  ]
}
