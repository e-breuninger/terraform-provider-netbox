package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type fakePrivate map[string][]byte

func (p fakePrivate) GetKey(_ context.Context, key string) ([]byte, diag.Diagnostics) {
	return p[key], nil
}

func (p fakePrivate) SetKey(_ context.Context, key string, value []byte) diag.Diagnostics {
	p[key] = value
	return nil
}

type blankStringsModel struct {
	ID          types.Int64
	Name        types.String
	Description types.String
	Comments    types.String
	Label       types.String
	Tags        types.Set
	VlanIDs     types.List
	TagsAll     types.Set
}

func TestClearLegacyBlanks(t *testing.T) {
	ctx := context.Background()
	fresh := func() *blankStringsModel {
		return &blankStringsModel{
			ID:          types.Int64Value(7),
			Name:        types.StringValue("x"),
			Description: types.StringValue(""),
			Comments:    types.StringNull(),
			Label:       types.StringUnknown(),
			Tags:        types.SetValueMust(types.StringType, []attr.Value{}),
			VlanIDs:     types.ListValueMust(types.Int64Type, []attr.Value{}),
			TagsAll:     types.SetValueMust(types.StringType, []attr.Value{types.StringValue("a")}),
		}
	}

	// State v5 wrote: no marker, the blanks are cleared and the state marked.
	var diags diag.Diagnostics
	private := fakePrivate{}
	model := fresh()
	clearLegacyBlanks(ctx, private, private, model, &diags)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if !model.Description.IsNull() || model.Name.ValueString() != "x" || !model.Comments.IsNull() || !model.Label.IsUnknown() || model.ID.ValueInt64() != 7 {
		t.Fatalf("legacy state: got %+v", model)
	}
	if !model.Tags.IsNull() || !model.Tags.ElementType(ctx).Equal(types.StringType) || !model.VlanIDs.IsNull() || len(model.TagsAll.Elements()) != 1 {
		t.Fatalf("legacy state collections: got %+v", model)
	}
	if private[writtenStateKey] == nil {
		t.Fatal("legacy state: not marked as written")
	}

	// State this provider wrote: a blank is a configured "" and stays.
	model = fresh()
	clearLegacyBlanks(ctx, private, private, model, &diags)
	if model.Description.IsNull() || model.Description.ValueString() != "" || model.Tags.IsNull() {
		t.Fatalf("written state: got %+v", model)
	}
}
