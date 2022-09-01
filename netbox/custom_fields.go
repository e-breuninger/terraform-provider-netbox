package netbox

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const customFieldsKey = "custom_field"

// Boolean string for custom field
const CustomFieldBoolean = "boolean"

var customFieldsSchema = &schema.Schema{
	Type:     schema.TypeMap,
	Optional: true,
	Default:  nil,
	Elem: &schema.Schema{
		Type:    schema.TypeString,
		Default: nil,
	},
}

var customFieldSchema = &schema.Schema{
	Type:     schema.TypeSet,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the existing custom field.",
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{"text", "integer", "boolean",
					"date", "url", "selection", "multiple"}, false),
				Description: "Type of the existing custom field (text, integer, boolean, url, selection, multiple).",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value of the existing custom field.",
			},
		},
	},
	Description: "Existing custom fields to associate to this prefix (ipam module).",
}

func getCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok || len(cfm) == 0 {
		return nil
	}
	return cfm
}

func convertArrayInterfaceString(arrayInterface []interface{}) string {
	var arrayString []string

	for _, item := range arrayInterface {
		switch v := item.(type) {
		case string:
			arrayString = append(arrayString, v)
		case int:
			strV := strconv.FormatInt(int64(v), 10)
			arrayString = append(arrayString, strV)
		}

	}

	sort.Strings(arrayString)
	result := strings.Join(arrayString, ",")

	return result
}

// Pick the custom fields in the state file and update values with data from API
func updateCustomFieldsFromAPI(stateCustomFields, customFields interface{}) []map[string]string {
	var tfCms []map[string]string

	switch t := customFields.(type) {
	case map[string]interface{}:
		for _, stateCustomField := range stateCustomFields.([]interface{}) {
			for key, value := range t {
				if stateCustomField.(map[string]interface{})["name"].(string) == key {
					var strValue string

					cm := map[string]string{}
					cm["name"] = key
					cm["type"] = stateCustomField.(map[string]interface{})["type"].(string)

					if value != nil {
						switch v := value.(type) {
						case []interface{}:
							strValue = convertArrayInterfaceString(v)
						default:
							strValue = fmt.Sprintf("%v", v)
						}

						if strValue == "1" && cm["type"] == CustomFieldBoolean {
							strValue = "true"
						} else if strValue == "0" && cm["type"] == CustomFieldBoolean {
							strValue = "false"
						}

						cm["value"] = strValue
					} else {
						cm["value"] = ""
					}

					tfCms = append(tfCms, cm)
				}
			}
		}
	}

	return tfCms
}

// Convert custom field regarding his type
func convertCustomFieldsFromTerraformToAPI(stateCustomFields []interface{}, customFields []interface{}) map[string]interface{} {
	toReturn := make(map[string]interface{})

	for _, stateCf := range stateCustomFields {
		stateCustomField := stateCf.(map[string]interface{})
		toReturn[stateCustomField["name"].(string)] = nil
	}

	for _, cf := range customFields {
		customField := cf.(map[string]interface{})

		cfName := customField["name"].(string)
		cfType := customField["type"].(string)
		cfValue := customField["value"].(string)

		if len(cfValue) > 0 {
			if cfType == "integer" {
				cfValueInt, _ := strconv.Atoi(cfValue)
				toReturn[cfName] = cfValueInt
			} else if cfType == CustomFieldBoolean {
				if cfValue == "true" {
					toReturn[cfName] = true
				} else if cfValue == "false" {
					toReturn[cfName] = false
				}
			} else if cfType == "multiple" {
				cfValueArray := strings.Split(cfValue, ",")
				sort.Strings(cfValueArray)
				toReturn[cfName] = cfValueArray
			} else {
				toReturn[cfName] = cfValue
			}
		} else {
			toReturn[cfName] = nil
		}
	}

	return toReturn
}
