---
page_title: "netbox_device_type Resource - terraform-provider-netbox"
subcategory: "Data Center Inventory Management (DCIM)"
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/features/device-types/#device-types_1:
  A device type represents a particular make and model of hardware that exists in the real world. Device types define the physical attributes of a device (rack height and depth) and its individual components (console, power, network interfaces, and so on).
---

# netbox_device_type (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/features/device-types/#device-types_1):

> A device type represents a particular make and model of hardware that exists in the real world. Device types define the physical attributes of a device (rack height and depth) and its individual components (console, power, network interfaces, and so on).

## Example Usage

### Minimal device type

A bare device-type catalog entry. Devices created from this type get no auto-instantiated components.

```terraform
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
```

### Device type with nested component templates

A single resource that exercises every nested component-template family. Power outlets and front ports use sibling-by-name references; the provider resolves those to NetBox IDs at apply time.

```terraform
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
```

### Inventory item parent tree with polymorphic component FK

`inventory_item_templates` can form a parent/child tree (via `parent` + sibling name) and optionally point at any other component template on the same device type via `component_type` / `component_id`.

```terraform
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
```

### Coexistence with standalone template resources

Nested blocks on `netbox_device_type` can coexist with the standalone `netbox_interface_template` / `netbox_device_bay_template` / etc. resources on the same device_type. The reconciler only manages templates it previously owned, so standalone-managed templates are left alone.

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `manufacturer_id` (Number) ID of the `netbox_manufacturer` this device type belongs to.
- `model` (String) Marketing name of the model.

### Optional

- `console_port_templates` (Block Set) Console port templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/consoleporttemplate/). (see [below for nested schema](#nestedblock--console_port_templates))
- `console_server_port_templates` (Block Set) Console server port templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/consoleserverporttemplate/). (see [below for nested schema](#nestedblock--console_server_port_templates))
- `device_bay_templates` (Block Set) Device bay templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/devicebaytemplate/). (see [below for nested schema](#nestedblock--device_bay_templates))
- `front_port_templates` (Block Set) Front port templates instantiated on every device of this type. Each must reference a sibling `rear_port_templates` block by name. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/frontporttemplate/). (see [below for nested schema](#nestedblock--front_port_templates))
- `interface_templates` (Block Set) Network interface templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/interfacetemplate/). (see [below for nested schema](#nestedblock--interface_templates))
- `inventory_item_templates` (Block Set) Inventory item templates instantiated on every device of this type. Supports a parent tree via the `parent` field and an optional polymorphic FK via `component_type`/`component_id`. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/inventoryitemtemplate/). (see [below for nested schema](#nestedblock--inventory_item_templates))
- `is_full_depth` (Boolean) Whether the device occupies the full rack depth.
- `module_bay_templates` (Block Set) Module bay templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/modulebaytemplate/). (see [below for nested schema](#nestedblock--module_bay_templates))
- `part_number` (String) Manufacturer part number / SKU.
- `power_outlet_templates` (Block Set) Power outlet templates instantiated on every device of this type. May reference a sibling `power_port_templates` block by name. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/poweroutlettemplate/). (see [below for nested schema](#nestedblock--power_outlet_templates))
- `power_port_templates` (Block Set) Power port templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/powerporttemplate/). (see [below for nested schema](#nestedblock--power_port_templates))
- `rear_port_templates` (Block Set) Rear port templates instantiated on every device of this type. Front ports reference these by name. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/rearporttemplate/). (see [below for nested schema](#nestedblock--rear_port_templates))
- `slug` (String) URL-safe identifier for the device type. Defaults to a slugified `model` if not given.
- `subdevice_role` (String) For chassis-style devices: `parent` for the chassis, `child` for the modules. Leave unset for a single-piece device.
- `tags` (Set of String)
- `u_height` (Number) Rack height in U. Defaults to `1.0`. Defaults to `1.0`.

### Read-Only

- `id` (String) The ID of this resource.
- `tags_all` (Set of String)

<a id="nestedblock--console_port_templates"></a>
### Nested Schema for `console_port_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `type` (String) Console port connector type, e.g. `de-9`, `rj-45`, `usb-c`. See the NetBox docs for the full enumeration.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--console_server_port_templates"></a>
### Nested Schema for `console_server_port_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `type` (String) Console server port connector type. See the NetBox docs for the full enumeration.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--device_bay_templates"></a>
### Nested Schema for `device_bay_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--front_port_templates"></a>
### Nested Schema for `front_port_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.
- `rear_port` (String) Name of the sibling `rear_port_templates` block this front port is mapped to. Resolved to the corresponding template ID at apply time.
- `type` (String) Front port connector type, e.g. `8p8c`, `lc`, `mpo`. See the NetBox docs for the full enumeration.

Optional:

- `color` (String) Hex color code (without leading `#`) used for the port in the UI.
- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `rear_port_position` (Number) Which numbered position on the rear port this front port maps to. Defaults to `1`. Defaults to `1`.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--interface_templates"></a>
### Nested Schema for `interface_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.
- `type` (String) Interface type, e.g. `1000base-t`, `25gbase-x-sfp28`. See the NetBox docs for the full enumeration.

Optional:

- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `mgmt_only` (Boolean) If true, this interface is for out-of-band management only.
- `poe_mode` (String) PoE mode (`pd`, `pse`).
- `poe_type` (String) PoE type, e.g. `type1-ieee802.3af`.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--inventory_item_templates"></a>
### Nested Schema for `inventory_item_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `component_id` (Number) Polymorphic FK target ID. Use the `id` computed attribute of another nested template to wire this up.
- `component_type` (String) Polymorphic FK type, e.g. `dcim.interfacetemplate`, `dcim.consoleporttemplate`. Pair with `component_id` to attach this inventory item to another component on the same device_type.
- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `manufacturer_id` (Number) Optional manufacturer ID for this inventory item.
- `parent` (String) Name of the sibling `inventory_item_templates` block that should be the parent of this item. Forms a tree; the root has no parent.
- `part_id` (String) Manufacturer part number / SKU for this inventory item.
- `role_id` (Number) Optional inventory item role ID.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--module_bay_templates"></a>
### Nested Schema for `module_bay_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `position` (String) Position designator inside the chassis, used by NetBox when {module} substitution is performed on child component template names.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--power_outlet_templates"></a>
### Nested Schema for `power_outlet_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `description` (String) Free-form description shown in the NetBox UI.
- `feed_leg` (String) Power feed leg this outlet is connected to. Valid values are `A`, `B`, `C`.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `power_port` (String) Name of the sibling `power_port_templates` block this outlet is downstream of. Resolved to the corresponding template ID at apply time.
- `type` (String) Power outlet connector type, e.g. `iec-60320-c13`. See the NetBox docs for the full enumeration.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--power_port_templates"></a>
### Nested Schema for `power_port_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.

Optional:

- `allocated_draw` (Number) Allocated power draw in watts.
- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `maximum_draw` (Number) Maximum power draw in watts.
- `type` (String) Power port connector type, e.g. `iec-60320-c14`. See the NetBox docs for the full enumeration.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


<a id="nestedblock--rear_port_templates"></a>
### Nested Schema for `rear_port_templates`

Required:

- `name` (String) Name of the template. Must be unique within the parent device_type and is used as the identity key for the nested set.
- `type` (String) Rear port connector type, e.g. `8p8c`, `lc`, `mpo`. See the NetBox docs for the full enumeration.

Optional:

- `color` (String) Hex color code (without leading `#`) used for the port in the UI.
- `description` (String) Free-form description shown in the NetBox UI.
- `label` (String) Optional physical label, e.g. text printed on the chassis next to the port.
- `positions` (Number) Number of front positions this rear port can be split into. Defaults to 1 if not set. Defaults to `1`.

Read-Only:

- `id` (Number) NetBox-assigned ID of the template, populated after Create.


