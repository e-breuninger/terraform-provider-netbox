// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// circuitResourceModel is the Terraform state/plan model of the "circuit" resource.
type circuitResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Cid               types.String `tfsdk:"cid"`
	CircuitProviderID types.Int64  `tfsdk:"circuit_provider_id"`
	CircuitTypeID     types.Int64  `tfsdk:"circuit_type_id"`
	Status            types.String `tfsdk:"status"`
	TenantID          types.Int64  `tfsdk:"tenant_id"`
	InstallDate       types.String `tfsdk:"install_date"`
	TerminationDate   types.String `tfsdk:"termination_date"`
	CommitRateKbps    types.Int64  `tfsdk:"commit_rate_kbps"`
	Description       types.String `tfsdk:"description"`
	Comments          types.String `tfsdk:"comments"`
	OwnerID           types.Int64  `tfsdk:"owner_id"`
	Created           types.String `tfsdk:"created"`
	LastUpdated       types.String `tfsdk:"last_updated"`
	URL               types.String `tfsdk:"url"`
	Tags              types.Set    `tfsdk:"tags"`
	TagsAll           types.Set    `tfsdk:"tags_all"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
}

// expandCircuit converts the Terraform model of type "circuit" into an API request.
func expandCircuit(ctx context.Context, model *circuitResourceModel) (*netboxapi.CircuitRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CircuitRequestDTO{}
	requestDTO.Cid = conv.StringPtr(model.Cid)
	requestDTO.Provider = conv.Int64Ptr(model.CircuitProviderID)
	requestDTO.Type = conv.Int64Ptr(model.CircuitTypeID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.InstallDate = conv.StringPtr(model.InstallDate)
	requestDTO.TerminationDate = conv.StringPtr(model.TerminationDate)
	requestDTO.CommitRate = conv.Int64Ptr(model.CommitRateKbps)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenCircuit refreshes the Terraform model of type "circuit" from an API response.
func flattenCircuit(ctx context.Context, responseDTO *netboxapi.CircuitResponseDTO, model *circuitResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Cid = conv.FromStringPtr(responseDTO.Cid)
	model.CircuitProviderID = conv.FromInt64Ptr(responseDTO.Provider)
	model.CircuitTypeID = conv.FromInt64Ptr(responseDTO.Type)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.InstallDate = conv.FromStringPtr(responseDTO.InstallDate)
	model.TerminationDate = conv.FromStringPtr(responseDTO.TerminationDate)
	model.CommitRateKbps = conv.FromInt64Ptr(responseDTO.CommitRate)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
