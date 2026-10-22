package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// cableSide is the companion-hook logic behind one side of a cable, the a_side/b_side object:
// object_type plus ids (NetBox's wire list of {object_type, object_id} folded by the backend's
// termination_side kind, one type per side by NetBox's own rule) and one alias id list per content
// type layered on top (device_interface_ids, front_port_ids, ...). Like objectRef, the two forms
// derive each other deterministically in both directions, order included: with a cable profile the
// position in the list is the connector.
type cableSide struct {
	name      string // attribute name, for diagnostics
	obj       *types.Object
	attrTypes map[string]attr.Type
}

const (
	csObjectType = "object_type"
	csIDs        = "ids"
)

// cableSideAliases pairs the alias children with their content types, in the order NetBox's
// COMPATIBLE_TERMINATION_TYPES lists them.
var cableSideAliases = []struct{ name, contentType string }{
	{"device_interface_ids", "dcim.interface"},
	{"front_port_ids", "dcim.frontport"},
	{"rear_port_ids", "dcim.rearport"},
	{"console_port_ids", "dcim.consoleport"},
	{"console_server_port_ids", "dcim.consoleserverport"},
	{"power_port_ids", "dcim.powerport"},
	{"power_outlet_ids", "dcim.poweroutlet"},
	{"power_feed_ids", "dcim.powerfeed"},
	{"circuit_termination_ids", "circuits.circuittermination"},
}

func cableSideAliasNames() string {
	s := ""
	for i, alias := range cableSideAliases {
		if i > 0 {
			s += ", "
		}
		s += alias.name
	}
	return s
}

// plan derives the planned side from the configured one: a configured alias list sets object_type
// and ids and nulls the other aliases; configured object_type and ids fill in the alias of that
// type. More than one alias, an alias next to object_type/ids, half of that pair, or nothing at all
// (a cable needs terminations on both sides) are errors. An unknown side is left alone.
func (cableSide cableSide) plan(config cableSide, diags *diag.Diagnostics) {
	if config.obj.IsUnknown() {
		return
	}
	cfg := map[string]attr.Value{}
	if !config.obj.IsNull() {
		cfg = config.obj.Attributes()
	}
	isSet := func(k string) bool { v, ok := cfg[k]; return ok && !v.IsNull() }
	var set []string
	for _, alias := range cableSideAliases {
		if isSet(alias.name) {
			set = append(set, alias.name)
		}
	}
	explicitType, explicitIDs := isSet(csObjectType), isSet(csIDs)
	p := path.Root(cableSide.name)
	switch {
	case len(set) > 1:
		diags.AddAttributeError(p.AtName(set[0]), "Conflicting terminations",
			"Set only one of "+cableSideAliasNames()+".")
	case len(set) == 1 && (explicitType || explicitIDs):
		diags.AddAttributeError(p.AtName(set[0]), "Conflicting terminations",
			"Set either object_type and ids or one of "+cableSideAliasNames()+", not both.")
	case explicitType != explicitIDs:
		diags.AddAttributeError(p.AtName(csObjectType), "Incomplete terminations",
			"object_type and ids must be set together.")
	case len(set) == 1:
		ids, _ := cfg[set[0]].(types.List)
		objectType := types.StringValue("")
		for _, alias := range cableSideAliases {
			if alias.name == set[0] {
				objectType = types.StringValue(alias.contentType)
			}
		}
		cableSide.set(objectType, ids)
	case explicitType:
		objectType, _ := cfg[csObjectType].(types.String)
		ids, _ := cfg[csIDs].(types.List)
		cableSide.set(objectType, ids)
	default:
		diags.AddAttributeError(p, "Missing terminations",
			"Set object_type and ids, or one of "+cableSideAliasNames()+"; a cable needs terminations on both sides.")
	}
}

// fromWire derives the aliases from object_type and ids as NetBox returned them.
func (cableSide cableSide) fromWire() {
	if cableSide.obj.IsNull() || cableSide.obj.IsUnknown() {
		return
	}
	attrs := cableSide.obj.Attributes()
	objectType, _ := attrs[csObjectType].(types.String)
	ids, _ := attrs[csIDs].(types.List)
	cableSide.set(objectType, ids)
}

// set writes the side: object_type and ids as given, the alias of that type holding the same ids,
// the other aliases null. An unknown object_type makes every alias unknown.
func (cableSide cableSide) set(objectType types.String, ids types.List) {
	vals := map[string]attr.Value{csObjectType: objectType, csIDs: ids}
	for _, alias := range cableSideAliases {
		switch {
		case objectType.IsUnknown():
			vals[alias.name] = types.ListUnknown(types.Int64Type)
		case !objectType.IsNull() && objectType.ValueString() == alias.contentType:
			vals[alias.name] = ids
		default:
			vals[alias.name] = types.ListNull(types.Int64Type)
		}
	}
	*cableSide.obj = types.ObjectValueMust(cableSide.attrTypes, vals)
}
