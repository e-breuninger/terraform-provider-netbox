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

# Annotated device type: every optional metadata field NetBox supports on a
# device_type, with a default platform and a couple of custom_fields entries.
# Custom fields take strings; encode complex values with jsonencode().
resource "netbox_platform" "linux" {
  name = "Linux"
}

resource "netbox_device_type" "annotated" {
  model           = "EX-2000"
  part_number     = "EX2000-A"
  manufacturer_id = netbox_manufacturer.example.id

  airflow                  = "front-to-rear"
  weight                   = 12.5
  weight_unit              = "kg"
  description              = "Aggregation switch, top-of-rack"
  comments                 = "## Notes\nMust be paired with redundant PSU."
  default_platform_id      = netbox_platform.linux.id
  exclude_from_utilization = false

  custom_fields = {
    sku          = "EX2000-A"
    system_specs = jsonencode({ ram_gb = 64, cpu_count = 2 })
  }
}
