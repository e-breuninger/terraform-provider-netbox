package netboxapi

// Locks the object, multiobject, select and multiselect custom field conversions: NetBox answers
// related objects and {value, label} choices, the state holds their ids and values, and a write
// sends what NetBox accepts.

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTypedCustomFieldValueObject(t *testing.T) {
	nested := `{"id": 12, "url": "https://netbox.example.com/api/dcim/sites/12/", "display": "site"}`
	for _, testCase := range []struct {
		fieldType, value string
		want             any
	}{
		{"object", "12", int64(12)},
		{"object", nested, int64(12)}, // an older state's related object
		{"multiobject", "[3, 1]", []int64{3, 1}},
		{"multiobject", "[" + nested + ", 3]", []int64{12, 3}},
		{"multiobject", "[]", []int64{}},
	} {
		got, err := typedCustomFieldValue("cf", testCase.fieldType, testCase.value)
		if err != nil {
			t.Errorf("%s %s: %v", testCase.fieldType, testCase.value, err)
			continue
		}
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("%s %s = %#v, want %#v", testCase.fieldType, testCase.value, got, testCase.want)
		}
	}
	for _, testCase := range []struct{ fieldType, value string }{
		{"object", "abc"},
		{"object", "0"},
		{"object", "1.5"},
		{"object", `{"id": 12}`},
		{"multiobject", "12"},
		{"multiobject", `["x"]`},
	} {
		if _, err := typedCustomFieldValue("cf", testCase.fieldType, testCase.value); err == nil {
			t.Errorf("%s %s: want an error", testCase.fieldType, testCase.value)
		}
	}
}

func TestCustomFieldValuesRelatedObjects(t *testing.T) {
	site := func(id float64) map[string]any {
		return map[string]any{"id": id, "url": "https://netbox.example.com/api/dcim/sites/", "display": "site", "name": "site"}
	}
	got := customFieldValues(map[string]any{
		"object":      site(12),
		"multiobject": []any{site(13), site(2)},
		"decoded":     map[string]any{"id": json.Number("7"), "url": "u", "display": "d"},
		"json_object": map[string]any{"id": float64(5), "name": "x"},
		"json_list":   []any{float64(3), float64(1)},
		"empty_list":  []any{},
		"text":        "12",
		"unset":       nil,
	})
	want := map[string]string{
		"object":      "12",                  // the related object's id
		"decoded":     "7",                   // a json.Number id, as go-netbox decodes responses
		"multiobject": "[2,13]",              // the related objects' ids, sorted
		"json_object": `{"id":5,"name":"x"}`, // a JSON field that merely has an id key
		"json_list":   "[3,1]",               // a JSON field keeps its order
		"empty_list":  "[]",
		"text":        "12",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("customFieldValues = %v, want %v", got, want)
	}
}

func TestCustomFieldValuesSelections(t *testing.T) {
	choice := func(value, label string) map[string]any { return map[string]any{"value": value, "label": label} }
	got := customFieldValues(map[string]any{
		"select":       choice("a", "Alpha"),
		"multiselect":  []any{choice("b", "Beta"), choice("a", "Alpha")},
		"json_choice":  map[string]any{"value": "a", "label": "Alpha", "extra": true},
		"json_numeric": map[string]any{"value": float64(1), "label": "One"},
		"empty_list":   []any{},
	})
	want := map[string]string{
		"select":       "a",                                          // the choice value, as the configuration sets it
		"multiselect":  `["b","a"]`,                                  // the choice values, in NetBox's order
		"json_choice":  `{"extra":true,"label":"Alpha","value":"a"}`, // a JSON field with more keys
		"json_numeric": `{"label":"One","value":1}`,                  // a JSON field whose value is not a string
		"empty_list":   "[]",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("customFieldValues = %v, want %v", got, want)
	}
}

func TestTypedCustomFieldValueSelections(t *testing.T) {
	for _, testCase := range []struct {
		fieldType, value string
		want             any
	}{
		{"select", "a", "a"},
		{"select", `{"value": "a", "label": "Alpha"}`, "a"}, // an older state's choice
		{"multiselect", `["b", "a"]`, []any{"b", "a"}},
		{"multiselect", `[{"value": "b", "label": "Beta"}, {"value": "a", "label": "Alpha"}]`, []string{"b", "a"}},
		{"multiselect", "[]", []any{}},
	} {
		got, err := typedCustomFieldValue("cf", testCase.fieldType, testCase.value)
		if err != nil {
			t.Errorf("%s %s: %v", testCase.fieldType, testCase.value, err)
			continue
		}
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("%s %s = %#v, want %#v", testCase.fieldType, testCase.value, got, testCase.want)
		}
	}
}
