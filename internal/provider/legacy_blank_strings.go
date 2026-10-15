package provider

// State written by provider v5 holds "" for every optional string it did not
// set. keepPriorString keeps a prior "" (it stands for a configured "" that
// NetBox reads back as null), so after an upgrade from v5 every object would
// plan an update from "" to null that the plan output does not even show.
// Read clears those empty strings in state this provider has not written yet;
// Create, Update and Read mark the state they write in its private data.

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const writtenStateKey = "state_written"

var writtenStateValue = []byte(`true`)

type privateGetter interface {
	GetKey(ctx context.Context, key string) ([]byte, diag.Diagnostics)
}

type privateSetter interface {
	SetKey(ctx context.Context, key string, value []byte) diag.Diagnostics
}

// clearLegacyBlankStrings sets the empty top-level strings of model (a
// pointer to a resource model) to null unless prior marks the state as
// written by this provider, and marks the state Read writes.
func clearLegacyBlankStrings(ctx context.Context, prior privateGetter, next privateSetter, model any, diags *diag.Diagnostics) {
	written, d := prior.GetKey(ctx, writtenStateKey)
	diags.Append(d...)
	if written == nil {
		fields := reflect.ValueOf(model).Elem()
		for i := 0; i < fields.NumField(); i++ {
			s, ok := fields.Field(i).Interface().(types.String)
			if ok && !s.IsNull() && !s.IsUnknown() && s.ValueString() == "" {
				fields.Field(i).Set(reflect.ValueOf(types.StringNull()))
			}
		}
	}
	markStateWritten(ctx, next, diags)
}

// markStateWritten marks the state Create and Update write.
func markStateWritten(ctx context.Context, private privateSetter, diags *diag.Diagnostics) {
	diags.Append(private.SetKey(ctx, writtenStateKey, writtenStateValue)...)
}
