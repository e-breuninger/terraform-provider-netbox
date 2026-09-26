// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// customFieldChoiceSetResourceModel is the Terraform state/plan model of the "custom_field_choice_set" resource.
type customFieldChoiceSetResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	BaseChoices         types.String `tfsdk:"base_choices"`
	ExtraChoices        types.List   `tfsdk:"extra_choices"`
	OrderAlphabetically types.Bool   `tfsdk:"order_alphabetically"`
	ChoicesCount        types.Int64  `tfsdk:"choices_count"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
}

// expandCustomFieldChoiceSet converts the Terraform model of type "custom_field_choice_set" into an API request.
func expandCustomFieldChoiceSet(ctx context.Context, model *customFieldChoiceSetResourceModel) (*netboxapi.CustomFieldChoiceSetRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CustomFieldChoiceSetRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.BaseChoices = conv.StringPtr(model.BaseChoices)
	if !model.ExtraChoices.IsNull() && !model.ExtraChoices.IsUnknown() {
		var items []customFieldChoiceSetExtraChoicesModel
		diags.Append(model.ExtraChoices.ElementsAs(ctx, &items, false)...)
		requestDTO.ExtraChoices = make([]*netboxapi.CustomFieldChoiceSetRequestDTOExtraChoices, 0, len(items))
		for i := range items {
			v, d := expandCustomFieldChoiceSetExtraChoices(ctx, &items[i])
			diags.Append(d...)
			requestDTO.ExtraChoices = append(requestDTO.ExtraChoices, v)
		}
	}
	requestDTO.OrderAlphabetically = conv.BoolPtr(model.OrderAlphabetically)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	return requestDTO, diags
}

// flattenCustomFieldChoiceSet refreshes the Terraform model of type "custom_field_choice_set" from an API response.
func flattenCustomFieldChoiceSet(ctx context.Context, responseDTO *netboxapi.CustomFieldChoiceSetResponseDTO, model *customFieldChoiceSetResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.BaseChoices = conv.FromStringPtr(responseDTO.BaseChoices)
	if responseDTO.ExtraChoices == nil || (len(responseDTO.ExtraChoices) == 0 && model.ExtraChoices.IsNull() || model.ExtraChoices.IsUnknown()) {
		model.ExtraChoices = types.ListNull(types.ObjectType{AttrTypes: customFieldChoiceSetExtraChoicesAttrTypes()})
	} else {
		items := make([]customFieldChoiceSetExtraChoicesModel, 0, len(responseDTO.ExtraChoices))
		for _, element := range responseDTO.ExtraChoices {
			var o customFieldChoiceSetExtraChoicesModel
			diags.Append(flattenCustomFieldChoiceSetExtraChoices(ctx, element, &o)...)
			items = append(items, o)
		}
		v, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: customFieldChoiceSetExtraChoicesAttrTypes()}, items)
		diags.Append(d...)
		model.ExtraChoices = v
	}
	model.OrderAlphabetically = conv.FromBoolPtr(responseDTO.OrderAlphabetically)
	model.ChoicesCount = conv.FromInt64Ptr(responseDTO.ChoicesCount)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}

// customFieldChoiceSetExtraChoicesModel is the model of the extra_choices object.
type customFieldChoiceSetExtraChoicesModel struct {
	Value types.String `tfsdk:"value"`
	Label types.String `tfsdk:"label"`
}

// customFieldChoiceSetExtraChoicesAttrTypes is the attribute type map of customFieldChoiceSetExtraChoicesModel.
func customFieldChoiceSetExtraChoicesAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"value": types.StringType,
		"label": types.StringType,
	}
}

func expandCustomFieldChoiceSetExtraChoices(ctx context.Context, model *customFieldChoiceSetExtraChoicesModel) (*netboxapi.CustomFieldChoiceSetRequestDTOExtraChoices, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CustomFieldChoiceSetRequestDTOExtraChoices{}
	requestDTO.Value = conv.StringPtr(model.Value)
	requestDTO.Label = conv.StringPtr(model.Label)
	return requestDTO, diags
}

func flattenCustomFieldChoiceSetExtraChoices(ctx context.Context, responseDTO *netboxapi.CustomFieldChoiceSetResponseDTOExtraChoices, model *customFieldChoiceSetExtraChoicesModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.CustomFieldChoiceSetResponseDTOExtraChoices{}
	}
	model.Value = conv.FromStringPtr(responseDTO.Value)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	return diags
}
