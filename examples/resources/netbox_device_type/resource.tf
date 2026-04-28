# Minimal device type: just the catalog entry, no nested component templates.
# Devices created from this type will have no auto-instantiated components.
resource "netbox_manufacturer" "example" {
  name = "ExampleCorp"
}

resource "netbox_device_type" "minimal" {
  model           = "EX-1000"
  part_number     = "EX1000-A"
  manufacturer_id = netbox_manufacturer.example.id
}
