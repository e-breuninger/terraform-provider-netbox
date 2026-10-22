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

// eventRuleResourceModel is the Terraform state/plan model of the "event_rule" resource.
type eventRuleResourceModel struct {
	ID               types.Int64          `tfsdk:"id"`
	Name             types.String         `tfsdk:"name"`
	ObjectTypes      types.Set            `tfsdk:"object_types"`
	EventTypes       types.Set            `tfsdk:"event_types"`
	Enabled          types.Bool           `tfsdk:"enabled"`
	Conditions       jsontypes.Normalized `tfsdk:"conditions"`
	ActionType       types.String         `tfsdk:"action_type"`
	ActionObjectType types.String         `tfsdk:"action_object_type"`
	ActionObjectID   types.Int64          `tfsdk:"action_object_id"`
	Description      types.String         `tfsdk:"description"`
	OwnerID          types.Int64          `tfsdk:"owner_id"`
	Created          types.String         `tfsdk:"created"`
	LastUpdated      types.String         `tfsdk:"last_updated"`
	URL              types.String         `tfsdk:"url"`
	Tags             types.Set            `tfsdk:"tags"`
	TagsAll          types.Set            `tfsdk:"tags_all"`
	CustomFields     types.Map            `tfsdk:"custom_fields"`
}

// expandEventRule converts the Terraform model of type "event_rule" into an API request.
func expandEventRule(ctx context.Context, model *eventRuleResourceModel) (*netboxapi.EventRuleRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.EventRuleRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	requestDTO.EventTypes = conv.SetTo[string](ctx, model.EventTypes, &diags)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.Conditions = conv.StringPtr(model.Conditions.StringValue)
	requestDTO.ActionType = conv.StringPtr(model.ActionType)
	requestDTO.ActionObjectType = conv.StringPtr(model.ActionObjectType)
	requestDTO.ActionObjectID = conv.Int64Ptr(model.ActionObjectID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenEventRule refreshes the Terraform model of type "event_rule" from an API response.
func flattenEventRule(ctx context.Context, responseDTO *netboxapi.EventRuleResponseDTO, model *eventRuleResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ObjectTypes = conv.SetFrom(ctx, types.StringType, responseDTO.ObjectTypes, false, &diags)
	model.EventTypes = conv.SetFrom(ctx, types.StringType, responseDTO.EventTypes, false, &diags)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.Conditions = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Conditions))
	model.ActionType = conv.FromStringPtr(responseDTO.ActionType)
	model.ActionObjectType = conv.FromStringPtr(responseDTO.ActionObjectType)
	model.ActionObjectID = conv.FromInt64Ptr(responseDTO.ActionObjectID)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
