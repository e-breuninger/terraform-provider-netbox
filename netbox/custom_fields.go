package netbox

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
	return cfm
}

// customFieldsSchemaFunc is a function that returns the schema for all custom
// fields.
func customFieldsSchemaFunc() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Description: "A JSON string that defines the custom fields as defined under the `custom_fields` key in the object's api." +
			"This is best managed with the `jsonencode()` & `jsondecode()` functions.",
		ValidateFunc: validation.StringIsJSON,
	}
}

// handleCustomFieldUpdate is a function that takes in the old and new values returned
// from a terraform d.GetChange() function and returns a map with the values to be sent
// in the custom_fields field on an update to the netbox api.
// This function handles setting the custom_field fields to null when needed. It does
// this by comparing the custom fields that were previously in terraform state, and
// compares them with then new state. It then sets any fields that were previously
// set to nil, if the new state does not include them.
func handleCustomFieldUpdate(old, new interface{}) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	var newData map[string]interface{}

	if new.(string) != "" {
		err := json.Unmarshal([]byte(new.(string)), &newData)
		if err != nil {
			return nil, fmt.Errorf("err1: %w", err)
		}
		for k, v := range newData {
			ret[k] = v
		}
	}
	var oldData map[string]interface{}
	if old.(string) != "" {
		err := json.Unmarshal([]byte(old.(string)), &oldData)
		if err != nil {
			return nil, fmt.Errorf("err2: %w", err)
		}
		for k := range oldData {
			if val, ok := ret[k]; !ok {
				ret[k] = nil
			} else {
				ret[k] = val
			}
		}
	}
	return ret, nil
}

// handleCustomFieldRead is a function that take an input of the interface
// from the CustomField struct field, and returns a string with the value
// to set for terraform and an error.
// This function checks the number of keys in the map, and if the number of
// nil fields equal the number of fields, it returns an empty string. Since
// this means the custom fields are not managed. Otherwise, it will return
// the result of unmarshalling the field to a json string.
// This allows the custom_field field to be set to empty and have an empty plan
// even though the api still has the custom fields with null values.
func handleCustomFieldRead(cf interface{}) (string, error) {
	cfMap, ok := cf.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("cannot cast %v to map[string]interface{}", cf)
	}
	numNull := 0
	for k := range cfMap {
		if cfMap[k] == nil {
			numNull += 1
		}
	}
	if len(cfMap) == 0 || numNull == len(cfMap) {
		return "", nil
	}
	b, err := json.Marshal(cf)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
