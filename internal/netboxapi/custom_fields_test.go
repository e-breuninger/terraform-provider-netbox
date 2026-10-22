package netboxapi

// Locks the object and multiobject custom field conversions: NetBox answers related objects, the
// state holds their ids, and a write sends the ids NetBox accepts.

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
