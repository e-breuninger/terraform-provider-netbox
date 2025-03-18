package models

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
)

type TagTerraformModel struct {
	Description types.String `tfsdk:"description"`
	Id          types.Int32  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ObjectTypes types.List   `tfsdk:"object_types"`
	Slug        types.String `tfsdk:"slug"`
	ColorHex    types.String `tfsdk:"color_hex"`
}

func (data *TagTerraformModel) ReadAPI(ctx context.Context, tag *netbox.Tag) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	data.Id = types.Int32Value(tag.Id)
	data.Name = types.StringValue(tag.Name)
	data.Slug = types.StringValue(tag.Slug)
	data.ColorHex = types.StringPointerValue(tag.Color)
	data.Description = types.StringPointerValue(tag.Description)
	listObjectTypes, err := types.ListValueFrom(ctx, types.StringType, tag.ObjectTypes)
	if err != nil {
		return err
	}
	data.ObjectTypes = listObjectTypes

	return diags
}
