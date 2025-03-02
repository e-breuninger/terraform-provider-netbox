package helpers

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
)

func ReadTagsFromAPI(tags []netbox.NestedTag) []int32 {
	var tag_names []int32
	for _, element := range tags {
		tag_names = append(tag_names, element.Id)
	}
	return tag_names
}

func WriteTagsToApi(ctx context.Context, client netbox.APIClient, terraformTags types.List) ([]netbox.NestedTagRequest, diag.Diagnostics) {
	var diags = diag.Diagnostics{}
	var tagList []netbox.NestedTagRequest
	if len(terraformTags.Elements()) > 0 {
		elements := make([]int32, 0, len(terraformTags.Elements()))
		diagsConvert := terraformTags.ElementsAs(ctx, &elements, false)
		if diagsConvert.HasError() {
			return nil, diagsConvert
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
	return nil, nil
}

func ReadCustomFieldsFromAPI(customFields map[string]interface{}) map[string]attr.Value {
	elements := map[string]attr.Value{}
	for k, v := range customFields {
		if v == nil {
			elements[k] = types.StringNull()
		} else {
			elements[k] = types.StringValue(v.(string))
		}

	}
	return elements
}

func ReadCustomFieldsFromTerraform(customFields types.Map) map[string]interface{} {
	result := make(map[string]interface{})
	if !customFields.IsUnknown() {
		for k, v := range customFields.Elements() {
			result[k] = v.(types.String).ValueString()
		}
	}
	return result
}
