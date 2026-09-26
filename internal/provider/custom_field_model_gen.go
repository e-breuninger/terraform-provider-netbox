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

// customFieldResourceModel is the Terraform state/plan model of the "custom_field" resource.
type customFieldResourceModel struct {
	ID                types.Int64          `tfsdk:"id"`
	Name              types.String         `tfsdk:"name"`
	ObjectTypes       types.Set            `tfsdk:"object_types"`
	Type              types.String         `tfsdk:"type"`
	RelatedObjectType types.String         `tfsdk:"related_object_type"`
	Label             types.String         `tfsdk:"label"`
	Description       types.String         `tfsdk:"description"`
	GroupName         types.String         `tfsdk:"group_name"`
	Required          types.Bool           `tfsdk:"required"`
	FilterLogic       types.String         `tfsdk:"filter_logic"`
	Weight            types.Int64          `tfsdk:"weight"`
	ValidationMinimum types.Float64        `tfsdk:"validation_minimum"`
	ValidationMaximum types.Float64        `tfsdk:"validation_maximum"`
	ValidationRegex   types.String         `tfsdk:"validation_regex"`
	Default           jsontypes.Normalized `tfsdk:"default"`
	ChoiceSetID       types.Int64          `tfsdk:"choice_set_id"`
	OwnerID           types.Int64          `tfsdk:"owner_id"`
	Created           types.String         `tfsdk:"created"`
	LastUpdated       types.String         `tfsdk:"last_updated"`
	URL               types.String         `tfsdk:"url"`
}

// expandCustomField converts the Terraform model of type "custom_field" into an API request.
func expandCustomField(ctx context.Context, model *customFieldResourceModel) (*netboxapi.CustomFieldRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CustomFieldRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.RelatedObjectType = conv.StringPtr(model.RelatedObjectType)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.GroupName = conv.StringPtr(model.GroupName)
	requestDTO.Required = conv.BoolPtr(model.Required)
	requestDTO.FilterLogic = conv.StringPtr(model.FilterLogic)
	requestDTO.Weight = conv.Int64Ptr(model.Weight)
	requestDTO.ValidationMinimum = conv.Float64Ptr(model.ValidationMinimum)
	requestDTO.ValidationMaximum = conv.Float64Ptr(model.ValidationMaximum)
	requestDTO.ValidationRegex = conv.StringPtr(model.ValidationRegex)
	requestDTO.Default = conv.StringPtr(model.Default.StringValue)
	requestDTO.ChoiceSet = conv.Int64Ptr(model.ChoiceSetID)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	return requestDTO, diags
}

// flattenCustomField refreshes the Terraform model of type "custom_field" from an API response.
func flattenCustomField(ctx context.Context, responseDTO *netboxapi.CustomFieldResponseDTO, model *customFieldResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ObjectTypes = conv.SetFrom(ctx, types.StringType, responseDTO.ObjectTypes, false, &diags)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.RelatedObjectType = conv.FromStringPtr(responseDTO.RelatedObjectType)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.GroupName = conv.FromStringPtr(responseDTO.GroupName)
	model.Required = conv.FromBoolPtr(responseDTO.Required)
	model.FilterLogic = conv.FromStringPtr(responseDTO.FilterLogic)
	model.Weight = conv.FromInt64Ptr(responseDTO.Weight)
	model.ValidationMinimum = conv.FromFloat64Ptr(responseDTO.ValidationMinimum)
	model.ValidationMaximum = conv.FromFloat64Ptr(responseDTO.ValidationMaximum)
	model.ValidationRegex = conv.FromStringPtr(responseDTO.ValidationRegex)
	model.Default = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Default))
	model.ChoiceSetID = conv.FromInt64Ptr(responseDTO.ChoiceSet)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
