# Nested template blocks on netbox_device_type can coexist with the standalone
# template resources (netbox_interface_template, netbox_device_bay_template,
# etc.) on the same device_type. The contract is:
#
#   any individual NetBox template object is managed by exactly one of
#     (a) a nested block on its parent netbox_device_type
#     (b) a standalone netbox_*_template resource
#
# The provider enforces this with an "ownership gate": the nested-block
# reconciler only deletes templates whose names appear in the prior state of
# the device_type (the templates it previously managed). Templates created by
# standalone resources are invisible to it and are left alone.
#
# Practical use case: most of a device's components are uniform across the
# fleet (declare them once on the device_type), but a few interfaces vary by
# deployment (declare them as standalone resources, possibly in a per-site
# module) so they can be re-used or reshaped without touching the catalog.
resource "netbox_manufacturer" "appliances" {
  name = "ExampleCorp Appliances"
}

resource "netbox_device_type" "appliance" {
  model           = "EX-FIREWALL-1U"
  part_number     = "EXFW1U"
  manufacturer_id = netbox_manufacturer.appliances.id

  # Uniform across all deployments — managed by the device_type.
  interface_templates {
    name      = "mgmt0"
    type      = "1000base-t"
    mgmt_only = true
  }
  interface_templates {
    name = "wan0"
    type = "10gbase-x-sfpp"
  }
}

# Optional / per-deployment interfaces — managed independently. Adding or
# removing these does not trigger a diff on netbox_device_type.appliance, and
# the device_type's reconciler will not delete them.
resource "netbox_interface_template" "appliance_lan_ports" {
  for_each = toset(["lan0", "lan1", "lan2", "lan3"])

  device_type_id = netbox_device_type.appliance.id
  name           = each.key
  type           = "1000base-t"
}
