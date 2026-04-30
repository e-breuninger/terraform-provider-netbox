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
}

func getCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok || len(cfm) == 0 {
		return nil
	}

	// Filter out keys that NetBox returns as "unset". NetBox includes every
	// custom_field registered for the content_type in the response, regardless
	// of whether this particular object set a value. The exact null shape
	// depends on CF type and NetBox version: text comes back as "" on 4.4,
	// other types as JSON null. Treat empty string + nil as "not set" so they
	// don't pollute Terraform state with ghost keys for users who never opted
	// in to that CF on this resource.
	result := make(map[string]interface{})
	for key, value := range cfm {
		if value == nil {
			continue
		}
		if s, isStr := value.(string); isStr && s == "" {
			continue
		}
		result[key] = value
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

// customFieldsForUpdate returns the value to assign to the CustomFields field
// of a writable model in an Update path.
//
// NetBox treats custom_fields as a sparse JSON dict: keys absent from the PATCH
// payload are left untouched, and keys with a literal null value are cleared.
// Combined with the writable models' `json:"custom_fields,omitempty"` tag, that
// means simply assigning the user's current map (which is empty when they have
// removed all CF entries from their HCL) is not enough — the empty map is
// dropped by omitempty and NetBox keeps the stale values.
//
// This helper walks the schema diff and produces a map that:
//   - sets every key still present in the user's config to its new value
//   - explicitly nulls every key that used to be in state but is gone from config,
//     so NetBox actually clears them
//
// When neither old state nor new config has any custom fields, it returns nil
// so omitempty correctly omits the field entirely.
func customFieldsForUpdate(d *schema.ResourceData) interface{} {
	oldRaw, newRaw := d.GetChange(customFieldsKey)
	oldMap, _ := oldRaw.(map[string]interface{})
	newMap, _ := newRaw.(map[string]interface{})

	if len(oldMap) == 0 && len(newMap) == 0 {
		return nil
	}

	result := make(map[string]interface{}, len(oldMap)+len(newMap))
	for k := range oldMap {
		result[k] = nil
	}
	for k, v := range newMap {
		result[k] = v
	}
	return result
}

// flattenCustomFields converts custom fields to a map where all values are strings.
// Complex nested objects (like IP address references) are converted to JSON strings.
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
		default:
			// For complex types (maps, arrays, objects), convert to JSON string
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
