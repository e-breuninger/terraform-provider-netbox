// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vlanResourceModel is the Terraform state/plan model of the "vlan" resource.
type vlanResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Vid          types.Int64  `tfsdk:"vid"`
	Status       types.String `tfsdk:"status"`
	SiteID       types.Int64  `tfsdk:"site_id"`
	GroupID      types.Int64  `tfsdk:"group_id"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	RoleID       types.Int64  `tfsdk:"role_id"`
	QinqRole     types.String `tfsdk:"qinq_role"`
	QinqSvlanID  types.Int64  `tfsdk:"qinq_svlan_id"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	PrefixCount  types.Int64  `tfsdk:"prefix_count"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandVlan converts the Terraform model of type "vlan" into an API request.
func expandVlan(ctx context.Context, model *vlanResourceModel) (*netboxapi.VlanRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VlanRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Vid = conv.Int64Ptr(model.Vid)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Site = conv.Int64Ptr(model.SiteID)
	requestDTO.Group = conv.Int64Ptr(model.GroupID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.QinqRole = conv.StringPtr(model.QinqRole)
	requestDTO.QinqSvlan = conv.Int64Ptr(model.QinqSvlanID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVlan refreshes the Terraform model of type "vlan" from an API response.
func flattenVlan(ctx context.Context, responseDTO *netboxapi.VlanResponseDTO, model *vlanResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Vid = conv.FromInt64Ptr(responseDTO.Vid)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.SiteID = conv.FromInt64Ptr(responseDTO.Site)
	model.GroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.QinqRole = conv.FromStringPtr(responseDTO.QinqRole)
	model.QinqSvlanID = conv.FromInt64Ptr(responseDTO.QinqSvlan)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.PrefixCount = conv.FromInt64Ptr(responseDTO.PrefixCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
