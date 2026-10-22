// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// wirelessLanGroupResourceModel is the Terraform state/plan model of the "wireless_lan_group" resource.
type wirelessLanGroupResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Slug             types.String `tfsdk:"slug"`
	ParentID         types.Int64  `tfsdk:"parent_id"`
	Description      types.String `tfsdk:"description"`
	Comments         types.String `tfsdk:"comments"`
	OwnerID          types.Int64  `tfsdk:"owner_id"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
	WirelessLanCount types.Int64  `tfsdk:"wireless_lan_count"`
	Tags             types.Set    `tfsdk:"tags"`
	TagsAll          types.Set    `tfsdk:"tags_all"`
	CustomFields     types.Map    `tfsdk:"custom_fields"`
}

// expandWirelessLanGroup converts the Terraform model of type "wireless_lan_group" into an API request.
func expandWirelessLanGroup(ctx context.Context, model *wirelessLanGroupResourceModel) (*netboxapi.WirelessLanGroupRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.WirelessLanGroupRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Parent = conv.Int64Ptr(model.ParentID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenWirelessLanGroup refreshes the Terraform model of type "wireless_lan_group" from an API response.
func flattenWirelessLanGroup(ctx context.Context, responseDTO *netboxapi.WirelessLanGroupResponseDTO, model *wirelessLanGroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.ParentID = conv.FromInt64Ptr(responseDTO.Parent)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.WirelessLanCount = conv.FromInt64Ptr(responseDTO.WirelesslanCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
