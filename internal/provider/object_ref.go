package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// objectRef is the companion-hook logic behind a NetBox polymorphic object reference pair
// (assigned_object_type + assigned_object_id, or parent_object_type + parent_object_id) and the
// per-content-type alias attributes layered on top of it (device_interface_id,
// virtual_machine_interface_id, device_id, ...). It ports the convenience of
// the pre-6.0 netbox_ip_address without its state dependence: aliases and pair derive each other
// deterministically in both directions, so import and drift work and no attribute needs to be
// ignored.
//
// Every attribute involved is optional+computed; the aliases have kind derived on both sides (never
// sent, never read) and the pair is sent as null when unset (clear_when_unset).
type objectRef struct {
	typeName, idName string // attribute names, for diagnostics
	typ              *types.String
	id               *types.Int64
	aliases          []objectRefAlias
	// required rejects an empty assignment: NetBox refuses the object without one (a service must
	// have a parent), so leaving every alias and the pair unset is a plan-time error rather than a
	// cleared assignment.
	required bool
}

// objectRefAlias is one alias attribute standing for one content type of the pair.
type objectRefAlias struct {
	name        string // attribute name
	contentType string // NetBox content type, e.g. dcim.interface
	id          *types.Int64
}

func (objectRef objectRef) aliasNames() string {
	s := ""
	for i, alias := range objectRef.aliases {
		if i > 0 {
			s += ", "
		}
		s += alias.name
	}
	return s
}

// plan derives the planned assignment from the configuration: a configured alias sets the pair
// (type = its content type, id = the alias, unknown stays unknown) and nulls the other aliases; a
// configured pair fills in the alias whose content type matches (an unknown type makes every alias
// unknown); nothing configured nulls all of them, or errors when the assignment is required. More
// than one alias, an alias next to the pair, or half a pair are errors.
func (objectRef objectRef) plan(config objectRef, diags *diag.Diagnostics) {
	var set []objectRefAlias
	for i := range config.aliases {
		if !config.aliases[i].id.IsNull() {
			set = append(set, config.aliases[i])
		}
	}
	explicit := !config.typ.IsNull() || !config.id.IsNull()
	switch {
	case len(set) > 1:
		diags.AddAttributeError(path.Root(set[0].name), "Conflicting assignment",
			"Set only one of "+objectRef.aliasNames()+".")
	case len(set) == 1 && explicit:
		diags.AddAttributeError(path.Root(objectRef.typeName), "Conflicting assignment",
			"Set either "+objectRef.typeName+" and "+objectRef.idName+" or one of "+objectRef.aliasNames()+", not both.")
	case explicit && (config.typ.IsNull() || config.id.IsNull()):
		diags.AddAttributeError(path.Root(objectRef.typeName), "Incomplete assignment",
			objectRef.typeName+" and "+objectRef.idName+" must be set together.")
	case len(set) == 1:
		*objectRef.typ = types.StringValue(set[0].contentType)
		*objectRef.id = *set[0].id
		for i := range objectRef.aliases {
			if objectRef.aliases[i].name == set[0].name {
				*objectRef.aliases[i].id = *set[0].id
			} else {
				*objectRef.aliases[i].id = types.Int64Null()
			}
		}
	case explicit:
		*objectRef.typ = *config.typ
		*objectRef.id = *config.id
		objectRef.fromPair()
	case objectRef.required:
		diags.AddError("Missing assignment",
			"Set one of "+objectRef.aliasNames()+", or "+objectRef.typeName+" and "+objectRef.idName+" together.")
	default:
		*objectRef.typ = types.StringNull()
		*objectRef.id = types.Int64Null()
		for i := range objectRef.aliases {
			*objectRef.aliases[i].id = types.Int64Null()
		}
	}
}

// toPair sets the pair from a known alias (before Create and Update, for values that were unknown
// at plan time); a null or unknown alias leaves the pair as planned.
func (objectRef objectRef) toPair() {
	for _, alias := range objectRef.aliases {
		if !alias.id.IsNull() && !alias.id.IsUnknown() {
			*objectRef.typ = types.StringValue(alias.contentType)
			*objectRef.id = *alias.id
			return
		}
	}
}

// fromPair sets the aliases from the pair (as planned, or as returned by NetBox): the alias whose
// content type matches gets the id, the others are null; an unknown type makes every alias unknown.
func (objectRef objectRef) fromPair() {
	for i := range objectRef.aliases {
		switch {
		case objectRef.typ.IsUnknown():
			*objectRef.aliases[i].id = types.Int64Unknown()
		case !objectRef.typ.IsNull() && objectRef.typ.ValueString() == objectRef.aliases[i].contentType:
			*objectRef.aliases[i].id = *objectRef.id
		default:
			*objectRef.aliases[i].id = types.Int64Null()
		}
	}
}

// changed reports whether the pair differs between two models.
func (objectRef objectRef) changed(other objectRef) bool {
	return !objectRef.typ.Equal(*other.typ) || !objectRef.id.Equal(*other.id)
}
