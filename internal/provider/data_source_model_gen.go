// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// dataSourceResourceModel is the Terraform state/plan model of the "data_source" resource.
type dataSourceResourceModel struct {
	ID           types.Int64          `tfsdk:"id"`
	Name         types.String         `tfsdk:"name"`
	Type         types.String         `tfsdk:"type"`
	SourceURL    types.String         `tfsdk:"source_url"`
	Enabled      types.Bool           `tfsdk:"enabled"`
	IgnoreRules  types.String         `tfsdk:"ignore_rules"`
	Parameters   jsontypes.Normalized `tfsdk:"parameters"`
	SyncInterval types.Int64          `tfsdk:"sync_interval"`
	Description  types.String         `tfsdk:"description"`
	Comments     types.String         `tfsdk:"comments"`
	Status       types.String         `tfsdk:"status"`
	LastSynced   types.String         `tfsdk:"last_synced"`
	OwnerID      types.Int64          `tfsdk:"owner_id"`
	Created      types.String         `tfsdk:"created"`
	LastUpdated  types.String         `tfsdk:"last_updated"`
	URL          types.String         `tfsdk:"url"`
	FileCount    types.Int64          `tfsdk:"file_count"`
	CustomFields types.Map            `tfsdk:"custom_fields"`
}

// expandDataSource converts the Terraform model of type "data_source" into an API request.
func expandDataSource(ctx context.Context, model *dataSourceResourceModel) (*netboxapi.DataSourceRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DataSourceRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.SourceURL = conv.StringPtr(model.SourceURL)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.IgnoreRules = conv.StringPtr(model.IgnoreRules)
	requestDTO.Parameters = conv.StringPtr(model.Parameters.StringValue)
	requestDTO.SyncInterval = conv.Int64Ptr(model.SyncInterval)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDataSource refreshes the Terraform model of type "data_source" from an API response.
func flattenDataSource(ctx context.Context, responseDTO *netboxapi.DataSourceResponseDTO, model *dataSourceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.SourceURL = conv.FromStringPtr(responseDTO.SourceURL)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.IgnoreRules = conv.FromStringPtr(responseDTO.IgnoreRules)
	model.Parameters = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Parameters))
	model.SyncInterval = conv.FromInt64Ptr(responseDTO.SyncInterval)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.LastSynced = conv.FromStringPtr(responseDTO.LastSynced)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.FileCount = conv.FromInt64Ptr(responseDTO.FileCount)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
