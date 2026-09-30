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
	ID                  types.Int64          `tfsdk:"id"`
	Name                types.String         `tfsdk:"name"`
	ObjectTypes         types.Set            `tfsdk:"object_types"`
	Type                types.String         `tfsdk:"type"`
	RelatedObjectType   types.String         `tfsdk:"related_object_type"`
	RelatedObjectFilter jsontypes.Normalized `tfsdk:"related_object_filter"`
	Label               types.String         `tfsdk:"label"`
	Description         types.String         `tfsdk:"description"`
	Comments            types.String         `tfsdk:"comments"`
	GroupName           types.String         `tfsdk:"group_name"`
	Required            types.Bool           `tfsdk:"required"`
	Unique              types.Bool           `tfsdk:"unique"`
	FilterLogic         types.String         `tfsdk:"filter_logic"`
	Weight              types.Int64          `tfsdk:"weight"`
	SearchWeight        types.Int64          `tfsdk:"search_weight"`
	UiVisible           types.String         `tfsdk:"ui_visible"`
	UiEditable          types.String         `tfsdk:"ui_editable"`
	IsCloneable         types.Bool           `tfsdk:"is_cloneable"`
	ValidationMinimum   types.Float64        `tfsdk:"validation_minimum"`
	ValidationMaximum   types.Float64        `tfsdk:"validation_maximum"`
	ValidationRegex     types.String         `tfsdk:"validation_regex"`
	ValidationSchema    jsontypes.Normalized `tfsdk:"validation_schema"`
	Default             jsontypes.Normalized `tfsdk:"default"`
	ChoiceSetID         types.Int64          `tfsdk:"choice_set_id"`
	OwnerID             types.Int64          `tfsdk:"owner_id"`
	Created             types.String         `tfsdk:"created"`
	LastUpdated         types.String         `tfsdk:"last_updated"`
	URL                 types.String         `tfsdk:"url"`
}

// expandCustomField converts the Terraform model of type "custom_field" into an API request.
func expandCustomField(ctx context.Context, model *customFieldResourceModel) (*netboxapi.CustomFieldRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CustomFieldRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.RelatedObjectType = conv.StringPtr(model.RelatedObjectType)
	requestDTO.RelatedObjectFilter = conv.StringPtr(model.RelatedObjectFilter.StringValue)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.GroupName = conv.StringPtr(model.GroupName)
	requestDTO.Required = conv.BoolPtr(model.Required)
	requestDTO.Unique = conv.BoolPtr(model.Unique)
	requestDTO.FilterLogic = conv.StringPtr(model.FilterLogic)
	requestDTO.Weight = conv.Int64Ptr(model.Weight)
	requestDTO.SearchWeight = conv.Int64Ptr(model.SearchWeight)
	requestDTO.UIVisible = conv.StringPtr(model.UiVisible)
	requestDTO.UIEditable = conv.StringPtr(model.UiEditable)
	requestDTO.IsCloneable = conv.BoolPtr(model.IsCloneable)
	requestDTO.ValidationMinimum = conv.Float64Ptr(model.ValidationMinimum)
	requestDTO.ValidationMaximum = conv.Float64Ptr(model.ValidationMaximum)
	requestDTO.ValidationRegex = conv.StringPtr(model.ValidationRegex)
	requestDTO.ValidationSchema = conv.StringPtr(model.ValidationSchema.StringValue)
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
	model.RelatedObjectFilter = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.RelatedObjectFilter))
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.GroupName = conv.FromStringPtr(responseDTO.GroupName)
	model.Required = conv.FromBoolPtr(responseDTO.Required)
	model.Unique = conv.FromBoolPtr(responseDTO.Unique)
	model.FilterLogic = conv.FromStringPtr(responseDTO.FilterLogic)
	model.Weight = conv.FromInt64Ptr(responseDTO.Weight)
	model.SearchWeight = conv.FromInt64Ptr(responseDTO.SearchWeight)
	model.UiVisible = conv.FromStringPtr(responseDTO.UIVisible)
	model.UiEditable = conv.FromStringPtr(responseDTO.UIEditable)
	model.IsCloneable = conv.FromBoolPtr(responseDTO.IsCloneable)
	model.ValidationMinimum = conv.FromFloat64Ptr(responseDTO.ValidationMinimum)
	model.ValidationMaximum = conv.FromFloat64Ptr(responseDTO.ValidationMaximum)
	model.ValidationRegex = conv.FromStringPtr(responseDTO.ValidationRegex)
	model.ValidationSchema = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.ValidationSchema))
	model.Default = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Default))
	model.ChoiceSetID = conv.FromInt64Ptr(responseDTO.ChoiceSet)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
