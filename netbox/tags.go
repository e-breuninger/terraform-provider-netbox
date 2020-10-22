package netbox

import (
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
				Severity: diag.Warning,
				Summary: "Error retrieving tag",
				Detail: fmt.Sprintf("Error trying to retrieve tag %s from netbox", tag.(string)),
			})
		} else {
			payload := res.GetPayload()
			one := int64(1) // oh. my. god.
			if (payload.Count == &one) {
				tags = append(tags, &models.NestedTag{
					Name: payload.Results[0].Name,
					Slug: payload.Results[0].Slug,
				})
			}
		}
		//tags = append(tags, *models.NestedTag{tag})
	}
	return tags, diags
}
