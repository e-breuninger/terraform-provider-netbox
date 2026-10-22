package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// portMappings is the companion-hook logic behind a front port's rear_ports set — one
// {position, rear_port_id, rear_port_position} object per front-port position, which is what NetBox
// 4.6 speaks — and the flat rear_port_id/rear_port_position pair layered on top of it for the
// common single-position port. Like objectRef, the two forms derive each other deterministically in
// both directions: the pair stands for the one-entry set at position 1, and a set that is exactly
// that fills the pair in, so import and drift work and a configuration can switch between the
// forms without a diff. It serves device_front_port and front_port_template, whose mapping objects
// have the same three int64 attributes.
type portMappings struct {
	positions *types.Int64         // the front port's own position count, as planned
	set       *types.Set           // rear_ports
	id, pos   *types.Int64         // rear_port_id, rear_port_position
	attrTypes map[string]attr.Type // of one rear_ports element
}

const (
	pmPosition         = "position"
	pmRearPortID       = "rear_port_id"
	pmRearPortPosition = "rear_port_position"
)

// plan derives the planned mappings from the configuration: the pair sets rear_ports to the
// one-entry set at position 1 (rear_port_position defaults to 1); an explicit rear_ports set fills
// the pair in when it is exactly such a one-entry set and nulls it otherwise; neither configured
// plans an empty set, which is what clears the mappings in NetBox (the API rejects null). The pair
// next to rear_ports, rear_port_position without rear_port_id, and the pair on a port with more
// than one position are errors.
func (portMappings portMappings) plan(config portMappings, diags *diag.Diagnostics) {
	explicit := !config.set.IsNull()
	pair := !config.id.IsNull()
	switch {
	case explicit && pair:
		diags.AddAttributeError(path.Root(pmRearPortID), "Conflicting mapping",
			"Set either rear_ports or rear_port_id (with rear_port_position), not both.")
	case !pair && !config.pos.IsNull():
		diags.AddAttributeError(path.Root(pmRearPortPosition), "Incomplete mapping",
			pmRearPortPosition+" requires "+pmRearPortID+".")
	case pair && !portMappings.positions.IsNull() && !portMappings.positions.IsUnknown() && portMappings.positions.ValueInt64() > 1:
		diags.AddAttributeError(path.Root(pmRearPortID), "Conflicting mapping",
			pmRearPortID+" maps a single-position front port; use rear_ports for a port with more than one position.")
	case pair:
		*portMappings.id = *config.id
		*portMappings.pos = types.Int64Value(1)
		if !config.pos.IsNull() {
			*portMappings.pos = *config.pos
		}
		portMappings.toSet()
	case explicit:
		*portMappings.set = *config.set
		portMappings.fromSet()
	default:
		*portMappings.set = portMappings.empty()
		*portMappings.id = types.Int64Null()
		*portMappings.pos = types.Int64Null()
	}
}

// toSet sets rear_ports from a known pair (as planned, and again before Create and Update for
// values that were unknown at plan time); an unknown pair makes the set unknown, a null pair
// leaves the set as planned.
func (portMappings portMappings) toSet() {
	switch {
	case portMappings.id.IsNull():
		return
	case portMappings.id.IsUnknown() || portMappings.pos.IsUnknown():
		*portMappings.set = types.SetUnknown(portMappings.elemType())
		return
	}
	obj := types.ObjectValueMust(portMappings.attrTypes, map[string]attr.Value{
		pmPosition:         types.Int64Value(1),
		pmRearPortID:       *portMappings.id,
		pmRearPortPosition: *portMappings.pos,
	})
	*portMappings.set = types.SetValueMust(portMappings.elemType(), []attr.Value{obj})
}

// fromSet sets the pair from rear_ports (as planned, or as returned by NetBox): a one-entry set at
// position 1 fills it, an unknown set, element or position makes it unknown, anything else nulls
// it. A null set (a response without mappings) becomes the empty set so that state and plan agree.
func (portMappings portMappings) fromSet() {
	if portMappings.set.IsUnknown() {
		*portMappings.id, *portMappings.pos = types.Int64Unknown(), types.Int64Unknown()
		return
	}
	if portMappings.set.IsNull() {
		*portMappings.set = portMappings.empty()
	}
	*portMappings.id, *portMappings.pos = types.Int64Null(), types.Int64Null()
	elems := portMappings.set.Elements()
	if len(elems) != 1 {
		return
	}
	obj, ok := elems[0].(types.Object)
	if !ok || obj.IsNull() {
		return
	}
	if obj.IsUnknown() {
		*portMappings.id, *portMappings.pos = types.Int64Unknown(), types.Int64Unknown()
		return
	}
	attrs := obj.Attributes()
	position, _ := attrs[pmPosition].(types.Int64)
	id, _ := attrs[pmRearPortID].(types.Int64)
	pos, _ := attrs[pmRearPortPosition].(types.Int64)
	switch {
	case position.IsUnknown():
		*portMappings.id, *portMappings.pos = types.Int64Unknown(), types.Int64Unknown()
	case position.ValueInt64() == 1:
		*portMappings.id, *portMappings.pos = id, pos
	}
}

// portMapping is one mapping as NetBox holds it.
type portMapping struct{ position, rearPort, rearPortPosition int64 }

// beforeUpdate decides how the planned mappings reach NetBox on an update, working around a NetBox
// 4.6 defect: the nested mapping serializer's uniqueness check counts the rows already in the
// database, so any update that resends an existing mapping — the identical set included — is
// rejected with "The fields rear_port, rear_port_position must make a unique set". A planned set
// equal to current is therefore left out of the body (NetBox keeps mappings it is not told about):
// rear_ports is nulled, which leaves it out of the request DTO's payload. Otherwise the caller
// clears the mappings first (clear is true when there are any), so that the update only ever
// creates rows; the cost is that unchanged rows are recreated too. An unknown planned set is left
// alone.
func (portMappings portMappings) beforeUpdate(current []portMapping) (clear bool) {
	if portMappings.set.IsUnknown() || portMappings.set.IsNull() {
		return false
	}
	planned := portMappings.mappings()
	if planned == nil {
		return false // an element is unknown: nothing to compare yet
	}
	if equalMappings(planned, current) {
		*portMappings.set = types.SetNull(portMappings.elemType())
		return false
	}
	return len(current) > 0
}

// mappings returns the planned set as NetBox rows, or nil when any value in it is unknown.
func (portMappings portMappings) mappings() []portMapping {
	out := []portMapping{}
	for _, element := range portMappings.set.Elements() {
		obj, ok := element.(types.Object)
		if !ok || obj.IsUnknown() || obj.IsNull() {
			return nil
		}
		attrs := obj.Attributes()
		position, _ := attrs[pmPosition].(types.Int64)
		id, _ := attrs[pmRearPortID].(types.Int64)
		pos, _ := attrs[pmRearPortPosition].(types.Int64)
		if position.IsUnknown() || id.IsUnknown() || pos.IsUnknown() {
			return nil
		}
		out = append(out, portMapping{position.ValueInt64(), id.ValueInt64(), pos.ValueInt64()})
	}
	return out
}

func equalMappings(a, b []portMapping) bool {
	if len(a) != len(b) {
		return false
	}
	seen := map[portMapping]int{}
	for _, mapping := range a {
		seen[mapping]++
	}
	for _, mapping := range b {
		if seen[mapping] == 0 {
			return false
		}
		seen[mapping]--
	}
	return true
}

// clearMappingsBody is the PATCH body that removes every mapping of a port.
func clearMappingsBody() map[string]any { return map[string]any{"rear_ports": []any{}} }

func (portMappings portMappings) elemType() attr.Type {
	return types.ObjectType{AttrTypes: portMappings.attrTypes}
}

func (portMappings portMappings) empty() types.Set {
	return types.SetValueMust(portMappings.elemType(), []attr.Value{})
}
