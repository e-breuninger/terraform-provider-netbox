# inventory_item_templates support a parent/child tree (via the parent name
# field) plus an optional polymorphic FK back to any other component template
# on the same device type (component_type + component_id).
#
# Tree-building rules:
#   - parent is a sibling inventory_item_templates name. Leave it unset for
#     root items.
#   - The provider applies inventory items in topological order, so you can
#     safely declare children alongside their parents in any HCL order.
#
# Polymorphic FK rules:
#   - component_type is a NetBox content-type string such as
#     "dcim.interfacetemplate", "dcim.poweroutlettemplate", etc.
#   - component_id is the NetBox ID of the targeted component template. You
#     typically pull it through with a plan-time reference; here we use a
#     trivial example with a sibling interface template defined inline.
resource "netbox_manufacturer" "servers" {
  name = "ExampleCorp Servers"
}

resource "netbox_device_type" "server" {
  model           = "EX-CHASSIS-1U"
  part_number     = "EX1U"
  manufacturer_id = netbox_manufacturer.servers.id

  interface_templates {
    name = "ipmi0"
    type = "1000base-t"
  }

  inventory_item_templates {
    name        = "chassis"
    description = "Outer chassis"
  }
  inventory_item_templates {
    name   = "psu-bay-a"
    parent = "chassis"
    label  = "PSU bay A"
  }
  inventory_item_templates {
    name        = "psu-fan-a"
    parent      = "psu-bay-a"
    description = "Fan inside PSU bay A"
  }

  # Child inventory item that points at a sibling interface template — useful
  # for tracking that a particular optic or transceiver lives on a specific
  # interface.
  inventory_item_templates {
    name           = "ipmi-nic"
    parent         = "chassis"
    component_type = "dcim.interfacetemplate"
    component_id   = 0 # replace with a real ID, e.g. via a data source lookup
  }
}
