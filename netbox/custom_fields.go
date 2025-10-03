package netbox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const customFieldsKey = "custom_fields"

var customFieldsSchema = &schema.Schema{
	Type:     schema.TypeMap,
	Optional: true,
	Default:  nil,
	Elem: &schema.Schema{
		Type:    schema.TypeString,
		Default: nil,
	},
	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
		if old == "" && new == "0" {
			return true // treat empty and "0" as equal? Wait, for maps it's different
		}
		// For maps, old and new are JSON strings
		if old == "{}" && new == "" {
			return true
		}
		if old == "" && new == "{}" {
			return true
		}
		return false
	},
}

func getCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(cfm) == 0 {
		return map[string]interface{}{}
	}
	result := make(map[string]interface{})
	for k, v := range cfm {
		if v != nil {
			if m, ok := v.(map[string]interface{}); ok {
				// Handle object references by extracting ID
				if id, ok := m["id"]; ok {
					result[k] = fmt.Sprintf("%v", id)
				} else {
					result[k] = fmt.Sprintf("%v", v)
				}
			} else {
				result[k] = fmt.Sprintf("%v", v)
			}
		} else {
			result[k] = ""
		}
	}
	return result
}
