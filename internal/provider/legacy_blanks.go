package provider

// State written by provider v5 holds "" for every optional string it did not
// set, and older v5 releases hold [] for unset sets such as tags. A 6.x plan
// over that state updates every object, from "" or [] to null, and the plan
// output shows no more than last_updated becoming unknown: keepPriorString
// keeps a prior "" (it stands for a configured "" that NetBox reads back as
// null), and Read keeps prior empty tags (they stand for a configured []).
// Read clears those blanks in state this provider has not written yet;
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

// clearLegacyBlanks sets the empty top-level strings, sets and lists of model
// (a pointer to a resource model) to null unless prior marks the state as
// written by this provider, and marks the state Read writes.
func clearLegacyBlanks(ctx context.Context, prior privateGetter, next privateSetter, model any, diags *diag.Diagnostics) {
	written, d := prior.GetKey(ctx, writtenStateKey)
	diags.Append(d...)
	if written == nil {
		fields := reflect.ValueOf(model).Elem()
		for i := 0; i < fields.NumField(); i++ {
			if null, ok := legacyBlankNull(ctx, fields.Field(i).Interface()); ok {
				fields.Field(i).Set(reflect.ValueOf(null))
			}
		}
	}
	markStateWritten(ctx, next, diags)
}

// legacyBlankNull returns the null value that replaces v when v is blank.
func legacyBlankNull(ctx context.Context, v any) (any, bool) {
	switch v := v.(type) {
	case types.String:
		if !v.IsNull() && !v.IsUnknown() && v.ValueString() == "" {
			return types.StringNull(), true
		}
	case types.Set:
		if !v.IsNull() && !v.IsUnknown() && len(v.Elements()) == 0 {
			return types.SetNull(v.ElementType(ctx)), true
		}
	case types.List:
		if !v.IsNull() && !v.IsUnknown() && len(v.Elements()) == 0 {
			return types.ListNull(v.ElementType(ctx)), true
		}
	}
	return nil, false
}

// markStateWritten marks the state Create and Update write.
func markStateWritten(ctx context.Context, private privateSetter, diags *diag.Diagnostics) {
	diags.Append(private.SetKey(ctx, writtenStateKey, writtenStateValue)...)
}
