// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// circuitProviderAccountResourceModel is the Terraform state/plan model of the "circuit_provider_account" resource.
type circuitProviderAccountResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Account           types.String `tfsdk:"account"`
	Name              types.String `tfsdk:"name"`
	CircuitProviderID types.Int64  `tfsdk:"circuit_provider_id"`
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

// expandCircuitProviderAccount converts the Terraform model of type "circuit_provider_account" into an API request.
func expandCircuitProviderAccount(ctx context.Context, model *circuitProviderAccountResourceModel) (*netboxapi.CircuitProviderAccountRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CircuitProviderAccountRequestDTO{}
	requestDTO.Account = conv.StringPtr(model.Account)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Provider = conv.Int64Ptr(model.CircuitProviderID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenCircuitProviderAccount refreshes the Terraform model of type "circuit_provider_account" from an API response.
func flattenCircuitProviderAccount(ctx context.Context, responseDTO *netboxapi.CircuitProviderAccountResponseDTO, model *circuitProviderAccountResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Account = conv.FromStringPtr(responseDTO.Account)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.CircuitProviderID = conv.FromInt64Ptr(responseDTO.Provider)
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
