package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const customFieldsKey = "custom_fields"

var customFieldsSchema = &schema.Schema{
	Type:     schema.TypeMap,
	Optional: true,
	Elem: &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
}
