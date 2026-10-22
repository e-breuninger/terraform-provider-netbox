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

// configContextProfileResourceModel is the Terraform state/plan model of the "config_context_profile" resource.
type configContextProfileResourceModel struct {
	ID          types.Int64          `tfsdk:"id"`
	Name        types.String         `tfsdk:"name"`
	Schema      jsontypes.Normalized `tfsdk:"schema"`
	Description types.String         `tfsdk:"description"`
	Comments    types.String         `tfsdk:"comments"`
	OwnerID     types.Int64          `tfsdk:"owner_id"`
	Created     types.String         `tfsdk:"created"`
	LastUpdated types.String         `tfsdk:"last_updated"`
	URL         types.String         `tfsdk:"url"`
	Tags        types.Set            `tfsdk:"tags"`
	TagsAll     types.Set            `tfsdk:"tags_all"`
}

// expandConfigContextProfile converts the Terraform model of type "config_context_profile" into an API request.
func expandConfigContextProfile(ctx context.Context, model *configContextProfileResourceModel) (*netboxapi.ConfigContextProfileRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ConfigContextProfileRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Schema = conv.StringPtr(model.Schema.StringValue)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	return requestDTO, diags
}

// flattenConfigContextProfile refreshes the Terraform model of type "config_context_profile" from an API response.
func flattenConfigContextProfile(ctx context.Context, responseDTO *netboxapi.ConfigContextProfileResponseDTO, model *configContextProfileResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Schema = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Schema))
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	return diags
}
