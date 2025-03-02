package models

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
)

type RegionTerraformModel struct {
	CustomFields types.Map    `tfsdk:"custom_fields"`
	Description  types.String `tfsdk:"description"`
	Id           types.Int32  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Parent       types.Int32  `tfsdk:"parent"`
	Slug         types.String `tfsdk:"slug"`
	Tags         types.List   `tfsdk:"tags"`
}

func (r *RegionTerraformModel) ReadAPI(ctx context.Context, region *netbox.Region) diag.Diagnostics {
	r.Id = types.Int32Value(region.Id)
	r.Name = types.StringValue(region.Name)
	r.Slug = types.StringValue(region.Slug)
	r.Description = types.StringPointerValue(region.Description)

	if region.Parent.Get() != nil {
		r.Parent = types.Int32Value(region.Parent.Get().Id)
	} else {
		r.Parent = types.Int32Null()
	}

	customFieldsFromAPI, diagData := types.MapValueFrom(ctx, types.StringType, helpers.ReadCustomFieldsFromAPI(region.CustomFields))
	if diagData.HasError() {
		return diagData
	}

	//Let's only add custom fields that we know
	if r.CustomFields.IsUnknown() || r.CustomFields.IsNull() {
		r.CustomFields = customFieldsFromAPI
	} else {
		for k, _ := range r.CustomFields.Elements() {
			if val, ok := customFieldsFromAPI.Elements()[k]; ok {
				r.CustomFields.Elements()[k] = val
			}
		}
	}

	tags := helpers.ReadTagsFromAPI(region.Tags)
	tagsdata, diagData := types.ListValueFrom(ctx, types.Int32Type, tags)
	if diagData.HasError() {
		return diagData
	}
	r.Tags = tagsdata
	return nil
}
