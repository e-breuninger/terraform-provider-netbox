package netbox

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVirtualDiskResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"size_mb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtual_machine_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
	}
}

func resourceNetboxVirtualDiskStateUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	v, ok := rawState["size_gb"].(float64)
	if !ok {
		log.Printf("[DEBUG] disk size before migration isnt float64: %#v\n but a %T", rawState["size_gb"], rawState["size_gb"])
		rawState["size_mb"] = float64(0)
		delete(rawState, "size_gb")
		return rawState, nil
	}

	log.Printf("[DEBUG] disk size in GB before migration: %#v\n", rawState["size_gb"])

	// set new disk size
	rawState["size_mb"] = v * 1000
	log.Printf("[DEBUG] disk size in MB after migration: %#v\n", rawState["size_mb"])

	delete(rawState, "size_gb")

	return rawState, nil
}
