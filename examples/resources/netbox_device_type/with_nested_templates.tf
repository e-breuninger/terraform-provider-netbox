# Comprehensive device type: a single resource that exercises every nested
# component-template family. Devices later created from this type will be
# pre-stamped with all the components defined here.
#
# A few cross-block rules to be aware of:
#   - power_outlet_templates.power_port refers to a sibling power_port_templates
#     block by name (not by ID).
#   - front_port_templates.rear_port refers to a sibling rear_port_templates
#     block by name. The provider resolves these names to NetBox IDs at apply
#     time, so creation ordering is handled for you.
#   - device_bay_templates require subdevice_role = "parent" on the parent
#     device_type. NetBox enforces this at the API.
resource "netbox_manufacturer" "switches" {
  name = "ExampleCorp Switches"
}

resource "netbox_device_type" "switch" {
  model           = "EX-9000"
  part_number     = "EX9000-2U"
  manufacturer_id = netbox_manufacturer.switches.id
  u_height        = 2
  is_full_depth   = true
  # Required because we attach device_bay_templates below.
  subdevice_role = "parent"

  power_port_templates {
    name           = "psu0"
    type           = "iec-60320-c14"
    maximum_draw   = 750
    allocated_draw = 500
    description    = "Primary PSU"
  }
  power_port_templates {
    name           = "psu1"
    type           = "iec-60320-c14"
    maximum_draw   = 750
    allocated_draw = 500
    description    = "Redundant PSU"
  }

  power_outlet_templates {
    name       = "out0"
    type       = "iec-60320-c13"
    power_port = "psu0"
    feed_leg   = "A"
  }

  interface_templates {
    name      = "mgmt0"
    type      = "1000base-t"
    mgmt_only = true
  }
  interface_templates {
    name        = "eth0"
    type        = "10gbase-x-sfpp"
    description = "Front-panel uplink"
  }

  console_port_templates {
    name = "console0"
    type = "rj-45"
  }

  console_server_port_templates {
    name = "csp0"
    type = "rj-45"
  }

  rear_port_templates {
    name      = "rp0"
    type      = "8p8c"
    positions = 4
  }
  front_port_templates {
    name               = "fp0"
    type               = "8p8c"
    rear_port          = "rp0"
    rear_port_position = 1
  }

  device_bay_templates {
    name        = "bay0"
    description = "Hot-swappable line card slot"
  }

  module_bay_templates {
    name     = "modbay0"
    position = "1"
  }
}
