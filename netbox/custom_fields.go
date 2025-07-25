package netbox

import (
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

func flattenMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range m {
		switch v := value.(type) {
		case map[string]interface{}:
			// Recursively flatten the nested map
			nestedResult := flattenMap(v)
			for nestedKey, nestedValue := range nestedResult {
				result[key+"__"+nestedKey] = nestedValue
			}
		case []interface{}:
			// Flatten the array by adding each element with an index to the result
			for i, element := range v {
				if elementMap, ok := element.(map[string]interface{}); ok {
					nestedResult := flattenMap(elementMap)
					for nestedKey, nestedValue := range nestedResult {
						result[key+"__"+fmt.Sprintf("%d", i)+"__"+nestedKey] = nestedValue
					}
				} else {
					result[key+"__"+fmt.Sprintf("%d", i)] = element
				}
			}
		default:
			result[key] = value
		}
	}

	return result
}

func getCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok || len(cfm) == 0 {
		return nil
	}
	flattenMap(cfm)
}
