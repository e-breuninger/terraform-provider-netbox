package netbox

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
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
	if !ok || len(cfm) == 0 {
		return nil
	}
	return cfm
}

// flattenCustomFields converts custom fields to a map where all values are strings.
// Object references (maps with "id" field) are converted to just their ID string.
// Other complex types are converted to JSON strings.
func flattenCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok || len(cfm) == 0 {
		return nil
	}

	result := make(map[string]interface{})
	for key, value := range cfm {
		if value == nil {
			result[key] = ""
			continue
		}

		// Check if the value is a simple type (string, number, bool)
		switch v := value.(type) {
		case string:
			result[key] = v
		case float64, int, int64, bool:
			result[key] = fmt.Sprintf("%v", v)
		case map[string]interface{}:
			// Check if this is an object reference with an ID
			if id, hasID := v["id"]; hasID {
				// Extract just the ID for object references
				result[key] = fmt.Sprintf("%v", id)
			} else {
				// For other complex objects without ID, convert to JSON
				if jsonBytes, err := json.Marshal(value); err == nil {
					result[key] = string(jsonBytes)
				} else {
					result[key] = fmt.Sprintf("%v", value)
				}
			}
		default:
			// For other complex types (arrays, etc.), convert to JSON string
			if jsonBytes, err := json.Marshal(value); err == nil {
				result[key] = string(jsonBytes)
			} else {
				// Fallback to string representation
				result[key] = fmt.Sprintf("%v", value)
			}
		}
	}

	return result
}

type CustomFieldParams struct {
	params runtime.ClientRequestWriter
	cfm    map[string]interface{}
}

func (o *CustomFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := o.params.WriteToRequest(r, reg); err != nil {
		return err
	}

	for k, v := range o.cfm {
		if vs, ok := v.(string); ok {
			if err := r.SetQueryParam(fmt.Sprintf("cf_%s", url.QueryEscape(k)), vs); err != nil {
				return err
			}
		}
	}

	return nil
}

func WithCustomFieldParamsOption(cfm map[string]interface{}) func(*runtime.ClientOperation) {
	if cfm == nil {
		cfm = make(map[string]interface{})
	}

	return func(co *runtime.ClientOperation) {
		co.Params = &CustomFieldParams{
			params: co.Params,
			cfm:    cfm,
		}
	}
}
