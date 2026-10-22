# Devices whose custom field "tier" matches "gold": a substring match if the field's filter logic
# is loose, an exact match if it is exact.
data "netbox_devices" "gold" {
  filters = [
    { name = "cf_tier", value = "gold" },
  ]
}
