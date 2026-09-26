// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vrfResourceModel is the Terraform state/plan model of the "vrf" resource.
type vrfResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Rd             types.String `tfsdk:"rd"`
	Description    types.String `tfsdk:"description"`
	EnforceUnique  types.Bool   `tfsdk:"enforce_unique"`
	TenantID       types.Int64  `tfsdk:"tenant_id"`
	ImportTargets  types.Set    `tfsdk:"import_targets"`
	ExportTargets  types.Set    `tfsdk:"export_targets"`
	OwnerID        types.Int64  `tfsdk:"owner_id"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	IPAddressCount types.Int64  `tfsdk:"ip_address_count"`
	PrefixCount    types.Int64  `tfsdk:"prefix_count"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
	CustomFields   types.Map    `tfsdk:"custom_fields"`
}

// expandVrf converts the Terraform model of type "vrf" into an API request.
func expandVrf(ctx context.Context, model *vrfResourceModel) (*netboxapi.VrfRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VrfRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Rd = conv.StringPtr(model.Rd)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.EnforceUnique = conv.BoolPtr(model.EnforceUnique)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.ImportTargets = conv.SetTo[int64](ctx, model.ImportTargets, &diags)
	requestDTO.ExportTargets = conv.SetTo[int64](ctx, model.ExportTargets, &diags)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVrf refreshes the Terraform model of type "vrf" from an API response.
func flattenVrf(ctx context.Context, responseDTO *netboxapi.VrfResponseDTO, model *vrfResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Rd = conv.FromStringPtr(responseDTO.Rd)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.EnforceUnique = conv.FromBoolPtr(responseDTO.EnforceUnique)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.ImportTargets = conv.SetFrom(ctx, types.Int64Type, responseDTO.ImportTargets, model.ImportTargets.IsNull() || model.ImportTargets.IsUnknown(), &diags)
	model.ExportTargets = conv.SetFrom(ctx, types.Int64Type, responseDTO.ExportTargets, model.ExportTargets.IsNull() || model.ExportTargets.IsUnknown(), &diags)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.IPAddressCount = conv.FromInt64Ptr(responseDTO.IpaddressCount)
	model.PrefixCount = conv.FromInt64Ptr(responseDTO.PrefixCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
