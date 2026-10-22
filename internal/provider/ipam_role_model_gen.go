// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ipamRoleResourceModel is the Terraform state/plan model of the "ipam_role" resource.
type ipamRoleResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Slug         types.String `tfsdk:"slug"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	Weight       types.Int64  `tfsdk:"weight"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	PrefixCount  types.Int64  `tfsdk:"prefix_count"`
	VlanCount    types.Int64  `tfsdk:"vlan_count"`
	AsnCount     types.Int64  `tfsdk:"asn_count"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandIpamRole converts the Terraform model of type "ipam_role" into an API request.
func expandIpamRole(ctx context.Context, model *ipamRoleResourceModel) (*netboxapi.IpamRoleRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IpamRoleRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Weight = conv.Int64Ptr(model.Weight)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenIpamRole refreshes the Terraform model of type "ipam_role" from an API response.
func flattenIpamRole(ctx context.Context, responseDTO *netboxapi.IpamRoleResponseDTO, model *ipamRoleResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.Weight = conv.FromInt64Ptr(responseDTO.Weight)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.PrefixCount = conv.FromInt64Ptr(responseDTO.PrefixCount)
	model.VlanCount = conv.FromInt64Ptr(responseDTO.VlanCount)
	model.AsnCount = conv.FromInt64Ptr(responseDTO.AsnCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
