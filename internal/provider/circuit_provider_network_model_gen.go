// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// circuitProviderNetworkResourceModel is the Terraform state/plan model of the "circuit_provider_network" resource.
type circuitProviderNetworkResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	CircuitProviderID types.Int64  `tfsdk:"circuit_provider_id"`
	ServiceID         types.String `tfsdk:"service_id"`
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

// expandCircuitProviderNetwork converts the Terraform model of type "circuit_provider_network" into an API request.
func expandCircuitProviderNetwork(ctx context.Context, model *circuitProviderNetworkResourceModel) (*netboxapi.CircuitProviderNetworkRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CircuitProviderNetworkRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Provider = conv.Int64Ptr(model.CircuitProviderID)
	requestDTO.ServiceID = conv.StringPtr(model.ServiceID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenCircuitProviderNetwork refreshes the Terraform model of type "circuit_provider_network" from an API response.
func flattenCircuitProviderNetwork(ctx context.Context, responseDTO *netboxapi.CircuitProviderNetworkResponseDTO, model *circuitProviderNetworkResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.CircuitProviderID = conv.FromInt64Ptr(responseDTO.Provider)
	model.ServiceID = conv.FromStringPtr(responseDTO.ServiceID)
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
