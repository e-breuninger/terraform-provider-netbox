package netbox

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const customFieldsKey = "custom_fields"

var customFieldsSchema = &schema.Schema{
	Type:     schema.TypeMap,
	Optional: true,
	Computed: true,
	Default:  nil,
	Elem: &schema.Schema{
		Type:    schema.TypeString,
		Default: nil,
	},
}

var customFieldsSchemaRead = &schema.Schema{
	Type:     schema.TypeMap,
	Computed: true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
}

func customFieldsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	cfg := d.GetRawConfig().GetAttr(customFieldsKey)
	cfm, ok := d.Get(customFieldsKey).(map[string]interface{})

	if cfg.IsNull() || !ok {
		d.SetNew(customFieldsKey, nil)
		return nil
	}

	for k, v := range cfm {
		if v == "" {
			delete(cfm, k)
		}
	}

	d.SetNew(customFieldsKey, cfm)
	return nil
}

func computeCustomFieldsModel(d *schema.ResourceData) interface{} {
	oldcf, newcf := d.GetChange(customFieldsKey)

	oldcfm, _ := oldcf.(map[string]interface{})
	newcfm, _ := newcf.(map[string]interface{})

	for k := range oldcfm {
		if _, ok := newcfm[k]; !ok {
			newcfm[k] = ""
		}
	}

	return newcfm
}

func computeCustomFieldsAttr(cf interface{}) map[string]interface{} {
	cfm, _ := cf.(map[string]interface{})
	newcfm := make(map[string]interface{})

	for k, v := range cfm {
		if vs, _ := v.(string); vs != "" {
			newcfm[k] = v
		}
	}

	return newcfm
}
