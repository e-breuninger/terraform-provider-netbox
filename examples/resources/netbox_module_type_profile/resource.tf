resource "netbox_module_type_profile" "test" {
  name        = "PSU"
  description = "Power supply unit"
  schema = jsonencode({
    type = "object"
    properties = {
      wattage = {
        type = "integer"
      }
      efficiency_rating = {
        type = "string"
      }
    }
  })
}

resource "netbox_manufacturer" "test" {
  name = "Cisco"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "PWR-C1-715WAC-P"
  part_number     = "PWR-C1-715WAC-P"
  profile_id      = netbox_module_type_profile.test.id
  attributes = jsonencode({
    wattage           = 715
    efficiency_rating = "80 Plus Platinum"
  })
}
