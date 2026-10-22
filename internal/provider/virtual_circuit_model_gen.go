// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualCircuitResourceModel is the Terraform state/plan model of the "virtual_circuit" resource.
type virtualCircuitResourceModel struct {
	ID                       types.Int64  `tfsdk:"id"`
	Cid                      types.String `tfsdk:"cid"`
	CircuitProviderNetworkID types.Int64  `tfsdk:"circuit_provider_network_id"`
	VirtualCircuitTypeID     types.Int64  `tfsdk:"virtual_circuit_type_id"`
	CircuitProviderAccountID types.Int64  `tfsdk:"circuit_provider_account_id"`
	Status                   types.String `tfsdk:"status"`
	TenantID                 types.Int64  `tfsdk:"tenant_id"`
	Description              types.String `tfsdk:"description"`
	Comments                 types.String `tfsdk:"comments"`
	OwnerID                  types.Int64  `tfsdk:"owner_id"`
	Created                  types.String `tfsdk:"created"`
	LastUpdated              types.String `tfsdk:"last_updated"`
	URL                      types.String `tfsdk:"url"`
	Tags                     types.Set    `tfsdk:"tags"`
	TagsAll                  types.Set    `tfsdk:"tags_all"`
	CustomFields             types.Map    `tfsdk:"custom_fields"`
}

// expandVirtualCircuit converts the Terraform model of type "virtual_circuit" into an API request.
func expandVirtualCircuit(ctx context.Context, model *virtualCircuitResourceModel) (*netboxapi.VirtualCircuitRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualCircuitRequestDTO{}
	requestDTO.Cid = conv.StringPtr(model.Cid)
	requestDTO.ProviderNetwork = conv.Int64Ptr(model.CircuitProviderNetworkID)
	requestDTO.Type = conv.Int64Ptr(model.VirtualCircuitTypeID)
	requestDTO.ProviderAccount = conv.Int64Ptr(model.CircuitProviderAccountID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualCircuit refreshes the Terraform model of type "virtual_circuit" from an API response.
func flattenVirtualCircuit(ctx context.Context, responseDTO *netboxapi.VirtualCircuitResponseDTO, model *virtualCircuitResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Cid = conv.FromStringPtr(responseDTO.Cid)
	model.CircuitProviderNetworkID = conv.FromInt64Ptr(responseDTO.ProviderNetwork)
	model.VirtualCircuitTypeID = conv.FromInt64Ptr(responseDTO.Type)
	model.CircuitProviderAccountID = conv.FromInt64Ptr(responseDTO.ProviderAccount)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
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
