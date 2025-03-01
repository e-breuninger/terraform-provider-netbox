package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/netbox-community/go-netbox/v4"
)

func readTags(tags []netbox.NestedTag) []int32 {
	var tag_names []int32
	for _, element := range tags {
		tag_names = append(tag_names, element.Id)
	}
	return tag_names
}

func readCustomFieldsFromAPI(customFields map[string]interface{}) (basetypes.ObjectValue, diag.Diagnostics) {

	elementTypes := map[string]attr.Type{}
	elements := map[string]attr.Value{}
	for k, v := range customFields {
		switch value := v.(type) {
		case int64:
			elementTypes[k] = types.Int64Type
			elements[k] = types.Int64Value(value)
		case string:
			elementTypes[k] = types.StringType
			elements[k] = types.StringValue(value)
		case bool:
			elementTypes[k] = types.BoolType
			elements[k] = types.BoolValue(value)
		case float64:
			elementTypes[k] = types.Float64Type
			elements[k] = types.Float64Value(value)
		case nil:
			elementTypes[k] = types.StringType
			elements[k] = types.StringNull()
		default:
			var diags diag.Diagnostics
			diags.AddError(fmt.Sprintf("Unknown type %T", v), "")
			return basetypes.NewObjectUnknown(elementTypes), diags
		}

	}
	return types.ObjectValue(elementTypes, elements)
}

func readCustomFieldsFromTerraform(customFields types.Dynamic) map[string]interface{} {
	//TODO: API Call to see if the type is OK
	if !customFields.IsUnknown() {
		var result map[string]interface{}
		obj := customFields.UnderlyingValue().(types.Object)
		for k, v := range obj.Attributes() {
			switch value := v.(type) {
			case types.Int64:
				result[k] = value.ValueInt64()
			case types.String:
				result[k] = value.ValueString()
			case types.Bool:
				result[k] = value.ValueBool()
			case types.Number:
				result[k] = value.ValueBigFloat()
			}
		}
		return result
	}
	return nil
}

func testClient(client *netbox.APIClient) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	//Validate that the APIClient exist.
	if client == nil {
		diags.AddError(
			"Create: Unconfigured API Client",
			"Expected configured API Client. Please report this issue to the provider developers.",
		)
		return nil
	}
	return diags
}

func writeTagsToApi(ctx context.Context, client netbox.APIClient, terraformTags types.List) ([]netbox.NestedTagRequest, diag.Diagnostics) {
	var diags = diag.Diagnostics{}
	var tagList []netbox.NestedTagRequest

	elements := make([]int32, 0, len(terraformTags.Elements()))
	diagsConvert := terraformTags.ElementsAs(ctx, &elements, false)
	if diagsConvert.HasError() {
		diags.Append(diagsConvert...)
		return nil, diags
	}
	paginatedTags, _, err := client.ExtrasAPI.ExtrasTagsList(ctx).Id(elements).Execute()
	if err != nil {
		diags.AddError("Unable to retrieve tags.",
			err.Error())
		return nil, diags
	}
	for _, element := range paginatedTags.Results {
		tagList = append(tagList, *netbox.NewNestedTagRequest(element.Name, element.Slug))
	}
	return tagList, diags
}
