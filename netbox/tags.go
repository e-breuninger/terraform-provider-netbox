package netbox

import (
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const tagsKey = "tags"

var tagsSchema = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Optional: true,
	Set:      schema.HashString,
}

var tagsSchemaRead = &schema.Schema{
	Type: schema.TypeSet,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
	Computed: true,
	Set:      schema.HashString,
}

func getNestedTagListFromResourceDataSet(client *client.NetBoxAPI, d interface{}) ([]*models.NestedTag, diag.Diagnostics) {
	var diags diag.Diagnostics

	tagList := d.(*schema.Set).List()
	tags := []*models.NestedTag{}
	for _, tag := range tagList {
		tagString := tag.(string)
		params := extras.NewExtrasTagsListParams()
		params.Name = &tagString
		limit := int64(2) // We search for a unique tag. Having two hits suffices to know its not unique.
		params.Limit = &limit
		res, err := client.Extras.ExtrasTagsList(params, nil)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Error retrieving tag %s from netbox", tag.(string)),
				Detail:   fmt.Sprintf("API Error trying to retrieve tag %s from netbox", tag.(string)),
			})
			return tags, diags
		}
		payload := res.GetPayload()
		switch *payload.Count {
		case int64(0):
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Error retrieving tag %s from netbox", tag.(string)),
				Detail:   fmt.Sprintf("Could not locate referenced tag %s in netbox", tag.(string)),
			})
			return tags, diags
		case int64(1):
			tags = append(tags, &models.NestedTag{
				Name: payload.Results[0].Name,
				Slug: payload.Results[0].Slug,
			})
		default:
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("Error retrieving tag %s from netbox", tag.(string)),
				Detail:   fmt.Sprintf("Could not map tag %s to unique tag in netbox", tag.(string)),
			})
		}
	}

	return tags, diags
}

func getTagListFromNestedTagList(nestedTags []*models.NestedTag) []string {
	tags := []string{}
	for _, nestedTag := range nestedTags {
		tags = append(tags, *nestedTag.Name)
	}
	return tags
}
