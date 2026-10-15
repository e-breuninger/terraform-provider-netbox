package netbox

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestFlattenCustomFields(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: nil,
		},
		{
			name: "simple string value",
			input: map[string]interface{}{
				"field1": "value1",
			},
			expected: map[string]interface{}{
				"field1": "value1",
			},
		},
		{
			name: "numeric values",
			input: map[string]interface{}{
				"int_field":   42,
				"int64_field": int64(100),
				"float_field": 3.14,
			},
			expected: map[string]interface{}{
				"int_field":   "42",
				"int64_field": "100",
				"float_field": "3.14",
			},
		},
		{
			name: "boolean value",
			input: map[string]interface{}{
				"bool_field": true,
			},
			expected: map[string]interface{}{
				"bool_field": "true",
			},
		},
		{
			name: "null value",
			input: map[string]interface{}{
				"null_field": nil,
			},
			expected: map[string]interface{}{
				"null_field": "",
			},
		},
		{
			name: "complex nested object (IP address reference)",
			input: map[string]interface{}{
				"gateway": map[string]interface{}{
					"id":      9,
					"address": "10.21.10.254/24",
					"display": "10.21.10.254/24",
					"family": map[string]interface{}{
						"value": 4,
						"label": "IPv4",
					},
					"url": "https://netbox.example.com/api/ipam/ip-addresses/9/",
				},
			},
			expected: map[string]interface{}{
				"gateway": `{"address":"10.21.10.254/24","display":"10.21.10.254/24","family":{"label":"IPv4","value":4},"id":9,"url":"https://netbox.example.com/api/ipam/ip-addresses/9/"}`,
			},
		},
		{
			name: "array value",
			input: map[string]interface{}{
				"tags": []string{"tag1", "tag2", "tag3"},
			},
			expected: map[string]interface{}{
				"tags": `["tag1","tag2","tag3"]`,
			},
		},
		{
			name: "mixed types",
			input: map[string]interface{}{
				"text_field":   "simple text",
				"number_field": 123,
				"bool_field":   false,
				"null_field":   nil,
				"object_field": map[string]interface{}{
					"nested": "value",
				},
			},
			expected: map[string]interface{}{
				"text_field":   "simple text",
				"number_field": "123",
				"bool_field":   "false",
				"null_field":   "",
				"object_field": `{"nested":"value"}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenCustomFields(tt.input)

			if tt.expected == nil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if result == nil {
				t.Errorf("expected non-nil result, got nil")
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d fields, got %d", len(tt.expected), len(result))
			}

			for key, expectedValue := range tt.expected {
				actualValue, ok := result[key]
				if !ok {
					t.Errorf("expected field %q not found in result", key)
					continue
				}

				// For JSON strings, compare the parsed objects to avoid formatting differences
				expectedStr, expectedIsStr := expectedValue.(string)
				actualStr, actualIsStr := actualValue.(string)

				if expectedIsStr && actualIsStr {
					// Try to parse as JSON
					var expectedJSON, actualJSON interface{}
					expectedIsJSON := json.Unmarshal([]byte(expectedStr), &expectedJSON) == nil
					actualIsJSON := json.Unmarshal([]byte(actualStr), &actualJSON) == nil

					if expectedIsJSON && actualIsJSON {
						// Compare as JSON objects
						expectedJSONStr, _ := json.Marshal(expectedJSON)
						actualJSONStr, _ := json.Marshal(actualJSON)
						if string(expectedJSONStr) != string(actualJSONStr) {
							t.Errorf("field %q: expected JSON %q, got %q", key, expectedStr, actualStr)
						}
						continue
					}
				}

				// Compare as strings
				if actualValue != expectedValue {
					t.Errorf("field %q: expected %q, got %q", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestFlattenCustomFields_ComplexRealWorldExample(t *testing.T) {
	// Simulate a real NetBox custom field response with an IP address object reference
	input := map[string]interface{}{
		"gateway": map[string]interface{}{
			"id":          9,
			"url":         "https://netbox.zonda.systems/api/ipam/ip-addresses/9/",
			"display":     "10.21.10.254/24",
			"address":     "10.21.10.254/24",
			"description": "",
			"family": map[string]interface{}{
				"value": float64(4),
				"label": "IPv4",
			},
		},
		"vlan_purpose": "management",
		"monitoring":   true,
		"priority":     10,
	}

	result := flattenCustomFields(input)

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	// Check that gateway is a JSON string
	gateway, ok := result["gateway"].(string)
	if !ok {
		t.Errorf("expected gateway to be a string, got %T", result["gateway"])
	}

	// Verify we can parse the gateway JSON
	var gatewayObj map[string]interface{}
	if err := json.Unmarshal([]byte(gateway), &gatewayObj); err != nil {
		t.Errorf("failed to parse gateway JSON: %v", err)
	}

	// Verify the gateway object has expected fields
	if gatewayObj["address"] != "10.21.10.254/24" {
		t.Errorf("expected address 10.21.10.254/24, got %v", gatewayObj["address"])
	}

	// Check simple fields
	if result["vlan_purpose"] != "management" {
		t.Errorf("expected vlan_purpose=management, got %v", result["vlan_purpose"])
	}

	if result["monitoring"] != "true" {
		t.Errorf("expected monitoring=true, got %v", result["monitoring"])
	}

	if result["priority"] != "10" {
		t.Errorf("expected priority=10, got %v", result["priority"])
	}
}

func TestGetCustomFields(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: nil,
		},
		{
			name: "valid map",
			input: map[string]interface{}{
				"field1": "value1",
				"field2": 123,
			},
			expected: map[string]interface{}{
				"field1": "value1",
				"field2": "123",
			},
		},
		{
			name: "non-string values are rendered as strings",
			input: map[string]interface{}{
				"bool_field":   false,
				"int_field":    int64(2),
				"float_field":  3.14,
				"object_field": map[string]interface{}{"name": "vlan1"},
				"array_field":  []interface{}{"a", "b"},
			},
			expected: map[string]interface{}{
				"bool_field":   "false",
				"int_field":    "2",
				"float_field":  "3.14",
				"object_field": `{"name":"vlan1"}`,
				"array_field":  `["a","b"]`,
			},
		},
		{
			name:     "invalid type",
			input:    "not a map",
			expected: nil,
		},
		{
			name: "map with nil values excluded",
			input: map[string]interface{}{
				"set_field":   "value1",
				"unset_field": nil,
			},
			expected: map[string]interface{}{
				"set_field": "value1",
			},
		},
		{
			name: "map with only nil values returns nil",
			input: map[string]interface{}{
				"unset1": nil,
				"unset2": nil,
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getCustomFields(tt.input)

			if tt.expected == nil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if result == nil {
				t.Errorf("expected non-nil result, got nil")
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d fields, got %d", len(tt.expected), len(result))
			}

			for key, expectedValue := range tt.expected {
				actualValue, ok := result[key]
				if !ok {
					t.Errorf("expected field %q not found in result", key)
					continue
				}

				if actualValue != expectedValue {
					t.Errorf("field %q: expected %v, got %v", key, expectedValue, actualValue)
				}
			}
		})
	}
}

// TestCustomFieldsWriteToStateIsComplete guards the read path against a partial
// write to state.
//
// The custom_fields elements are decoded strictly on the way into state, so a
// single value that is not a string makes the SDK give up part way through the
// map and leave the keys it had already written behind. Map iteration order is
// randomised, so the surviving subset differs on every read and the resource
// reports drift on fields nothing has touched.
func TestCustomFieldsWriteToStateIsComplete(t *testing.T) {
	// An object with a custom field of every type, as Netbox returns it: an unset
	// text field comes back as "", an unset field of any other type as null, and
	// the boolean, integer and object fields carry a real boolean, a real number
	// and a nested object.
	apiResponse := func() map[string]interface{} {
		return map[string]interface{}{
			"text_field":        "value1",
			"empty_text_field":  "",
			"unset_field":       nil,
			"true_field":        true,
			"false_field":       false,
			"integer_field":     float64(2),
			"decimal_field":     3.14,
			"selection_field":   "choice1",
			"object_field":      map[string]interface{}{"name": "object1"},
			"multiselect_field": []interface{}{"choice1", "choice2"},
		}
	}

	set := map[string]interface{}{
		"text_field":        "value1",
		"empty_text_field":  "",
		"true_field":        "true",
		"false_field":       "false",
		"integer_field":     "2",
		"decimal_field":     "3.14",
		"selection_field":   "choice1",
		"object_field":      `{"name":"object1"}`,
		"multiselect_field": `["choice1","choice2"]`,
	}

	// flattenCustomFields keeps an unset field as an empty string, getCustomFields
	// drops it.
	flattened := map[string]interface{}{"unset_field": ""}
	for k, v := range set {
		flattened[k] = v
	}

	cases := []struct {
		name string
		fn   func(interface{}) map[string]interface{}
		want map[string]interface{}
	}{
		{"getCustomFields", getCustomFields, set},
		{"flattenCustomFields", flattenCustomFields, flattened},
	}

	r := &schema.Resource{Schema: map[string]*schema.Schema{
		customFieldsKey: customFieldsSchema,
	}}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// Repeat: a writer that aborts mid-map drops a different, randomly
			// chosen set of keys each time, and can drop none at all by luck.
			for i := 0; i < 50; i++ {
				d := r.TestResourceData()
				if err := d.Set(customFieldsKey, tt.fn(apiResponse())); err != nil {
					t.Fatalf("writing custom fields to state failed: %s", err)
				}

				got := d.Get(customFieldsKey)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("state holds %v, want %v", got, tt.want)
				}
			}
		})
	}
}
