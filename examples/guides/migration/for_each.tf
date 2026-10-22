variable "site_ids" {
  type = list(number)
}

# Ids are numbers: a for_each over ids needs string keys.
resource "netbox_location" "lab" {
  for_each = toset([for id in var.site_ids : tostring(id)])

  name    = "lab"
  site_id = tonumber(each.key)
}
