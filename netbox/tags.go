package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func getTagListFromResourceDataSet(d interface{}) []string {
	tagList := d.(*schema.Set).List()
	tags := []string{}
	for _, tag := range tagList {
		tags = append(tags, tag.(string))
	}
	return tags
}
