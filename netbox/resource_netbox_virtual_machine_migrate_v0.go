package netbox

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxVirtualMachineResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_size_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
			"primary_ipv4": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}
func resourceNetboxVirtualMachineStateUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {

	v, ok := rawState["vcpus"]
	if !ok {
		return rawState, nil
	}

	s, ok := v.(string)
	if !ok {
		// since the provider was already released without this state migration, we have to accept that this field already contains non-string content
		return rawState, nil
	}

	log.Printf("[DEBUG] vcpus before migration: %#v", rawState["vcpus"])

	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		rawState["vcpus"] = f
	} else {
		rawState["vcpus"] = float64(0)
		log.Printf("[DEBUG] Schema upgrade: vcpus has been migrated to %g", f)
	}

	log.Printf("[DEBUG] vcpus after migration: %#v", rawState["vcpus"])
	return rawState, nil
}
