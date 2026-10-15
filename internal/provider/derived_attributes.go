package provider

// Attributes NetBox fills in itself when they are not set. As plain Optional
// attributes they failed every apply that left them unset with "Provider
// produced inconsistent result after apply". They are Optional+Computed with
// one of two plan modifiers, depending on when NetBox derives them:
//
//   - derivedFrom: NetBox derives the value from another attribute on every
//     save (a device's location from its rack, a virtual machine's site from
//     its cluster, a rack's dimensions from its rack type). The prior value is
//     kept while that attribute is unchanged, left unknown when it changes so
//     NetBox derives it again, and cleared when it is unset, as before.
//   - derivedOnCreate: NetBox derives the value only when the object is created
//     (a device's platform and airflow from its device type, a virtual
//     machine's platform, vCPUs and memory from its type). It is left unknown on
//     create; afterwards an unset value is cleared, as before, so an inherited
//     value that should stay must be set explicitly.

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// derivedPlan is what a derived attribute's plan modifier decided.
type derivedPlan int

const (
	derivedUnknown derivedPlan = iota // leave the plan unknown: NetBox decides
	derivedKeep                       // plan the prior state value
	derivedNull                       // plan null: the attribute is cleared
)

func derivedFrom(source string) derivedFromModifier {
	return derivedFromModifier{source: path.Root(source), name: source}
}

type derivedFromModifier struct {
	source path.Path
	name   string
	// keepWhenSourceUnset keeps the prior value when the source is unset
	// instead of clearing it, for attributes NetBox requires and defaults.
	keepWhenSourceUnset bool
}

func (m derivedFromModifier) keepingWhenSourceUnset() derivedFromModifier {
	m.keepWhenSourceUnset = true
	return m
}

func (m derivedFromModifier) Description(context.Context) string {
	return "NetBox derives this value from " + m.name + " when it is not set."
}

func (m derivedFromModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m derivedFromModifier) decide(ctx context.Context, configNull, stateNull bool, plan tfsdk.Plan, state tfsdk.State) derivedPlan {
	if !configNull {
		return derivedUnknown // the plan already holds the configured value
	}
	var planned types.Int64
	if plan.GetAttribute(ctx, m.source, &planned).HasError() || planned.IsUnknown() {
		return derivedUnknown
	}
	if planned.IsNull() {
		if m.keepWhenSourceUnset {
			if stateNull {
				return derivedUnknown
			}
			return derivedKeep
		}
		return derivedNull
	}
	if state.Raw.IsNull() || stateNull {
		return derivedUnknown
	}
	var prior types.Int64
	if state.GetAttribute(ctx, m.source, &prior).HasError() || !planned.Equal(prior) {
		return derivedUnknown
	}
	return derivedKeep
}

func derivedOnCreate() derivedOnCreateModifier {
	return derivedOnCreateModifier{}
}

type derivedOnCreateModifier struct{}

func (m derivedOnCreateModifier) Description(context.Context) string {
	return "NetBox derives this value when the object is created and it is not set."
}

func (m derivedOnCreateModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m derivedOnCreateModifier) decide(ctx context.Context, configNull, stateNull bool, plan tfsdk.Plan, state tfsdk.State) derivedPlan {
	switch {
	case !configNull:
		return derivedUnknown
	case state.Raw.IsNull():
		return derivedUnknown
	default:
		return derivedNull
	}
}

type derivedDecider interface {
	decide(ctx context.Context, configNull, stateNull bool, plan tfsdk.Plan, state tfsdk.State) derivedPlan
}

// The typed plan modifiers share one decision.

type derivedInt64 struct{ derivedDecider }
type derivedString struct{ derivedDecider }
type derivedFloat64 struct{ derivedDecider }
type derivedBool struct{ derivedDecider }

func (m derivedInt64) Description(ctx context.Context) string {
	return describeDerived(ctx, m.derivedDecider)
}
func (m derivedInt64) MarkdownDescription(ctx context.Context) string { return m.Description(ctx) }
func (m derivedInt64) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	switch m.decide(ctx, req.ConfigValue.IsNull(), req.StateValue.IsNull(), req.Plan, req.State) {
	case derivedKeep:
		resp.PlanValue = req.StateValue
	case derivedNull:
		resp.PlanValue = types.Int64Null()
	}
}

func (m derivedString) Description(ctx context.Context) string {
	return describeDerived(ctx, m.derivedDecider)
}
func (m derivedString) MarkdownDescription(ctx context.Context) string { return m.Description(ctx) }
func (m derivedString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	switch m.decide(ctx, req.ConfigValue.IsNull(), req.StateValue.IsNull(), req.Plan, req.State) {
	case derivedKeep:
		resp.PlanValue = req.StateValue
	case derivedNull:
		resp.PlanValue = types.StringNull()
	}
}

func (m derivedFloat64) Description(ctx context.Context) string {
	return describeDerived(ctx, m.derivedDecider)
}
func (m derivedFloat64) MarkdownDescription(ctx context.Context) string { return m.Description(ctx) }
func (m derivedFloat64) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	switch m.decide(ctx, req.ConfigValue.IsNull(), req.StateValue.IsNull(), req.Plan, req.State) {
	case derivedKeep:
		resp.PlanValue = req.StateValue
	case derivedNull:
		resp.PlanValue = types.Float64Null()
	}
}

func (m derivedBool) Description(ctx context.Context) string {
	return describeDerived(ctx, m.derivedDecider)
}
func (m derivedBool) MarkdownDescription(ctx context.Context) string { return m.Description(ctx) }
func (m derivedBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	switch m.decide(ctx, req.ConfigValue.IsNull(), req.StateValue.IsNull(), req.Plan, req.State) {
	case derivedKeep:
		resp.PlanValue = req.StateValue
	case derivedNull:
		resp.PlanValue = types.BoolNull()
	}
}

func describeDerived(ctx context.Context, d derivedDecider) string {
	if described, ok := d.(interface{ Description(context.Context) string }); ok {
		return described.Description(ctx)
	}
	return ""
}

var (
	_ planmodifier.Int64   = derivedInt64{}
	_ planmodifier.String  = derivedString{}
	_ planmodifier.Float64 = derivedFloat64{}
	_ planmodifier.Bool    = derivedBool{}
)
