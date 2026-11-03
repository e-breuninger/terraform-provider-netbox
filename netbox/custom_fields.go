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

// customFieldValueToString renders a custom field value as returned by the
// Netbox API as a string.
//
// The custom_fields schema is a map of strings, and the SDK's state writer
// decodes each element strictly (mapstructure.Decode, not WeakDecode). Handing
// it anything but a string aborts the write part way through the map, and since
// Go randomises map iteration order the keys that made it into state differ on
// every read. The result is a resource that reports drift on custom fields it
// never touched, with a different set of them each plan.
//
// Netbox returns a boolean for a boolean field, a number for an integer or
// decimal field and an object for an object reference, so every value has to be
// rendered here before it reaches state.
func customFieldValueToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case float64, int, int64, bool:
		return fmt.Sprintf("%v", v)
	default:
		// For complex types (maps, arrays, objects), convert to JSON string
		if jsonBytes, err := json.Marshal(value); err == nil {
			return string(jsonBytes)
		}
		// Fallback to string representation
		return fmt.Sprintf("%v", value)
	}
}

func getCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok || len(cfm) == 0 {
		return nil
	}

	result := make(map[string]interface{})
	for key, value := range cfm {
		// Netbox reports an unset field as null. Storing it as "" would show up
		// as a permanent diff for configurations that do not declare the field.
		if value != nil {
			result[key] = customFieldValueToString(value)
		}
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

// flattenCustomFields converts custom fields to a map where all values are strings.
// Complex nested objects (like IP address references) are converted to JSON strings.
// Unlike getCustomFields, an unset field is kept as an empty string.
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

		result[key] = customFieldValueToString(value)
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
