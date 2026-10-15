package netboxapi

// Locks typedCustomFieldValue for object and multiobject custom fields: the
// value is the related object's id (as configured) or the nested object NetBox
// returns (as held in state when set outside Terraform).

import (
	"reflect"
	"testing"
)

func TestTypedCustomFieldValueObject(t *testing.T) {
	nested := `{"id": 12, "url": "https://netbox/api/dcim/sites/12/", "display": "site"}`
	for _, tc := range []struct {
		fieldType, value string
		want             any
		wantErr          bool
	}{
		{"object", "12", int64(12), false},
		{"object", nested, int64(12), false},
		{"object", "abc", nil, true},
		{"object", "0", nil, true},
		{"multiobject", "[1, 2]", []int64{1, 2}, false},
		{"multiobject", "[" + nested + ", 3]", []int64{12, 3}, false},
		{"multiobject", "[]", []int64{}, false},
		{"multiobject", "12", nil, true},
		{"multiobject", `["x"]`, nil, true},
	} {
		got, err := typedCustomFieldValue("cf", tc.fieldType, tc.value)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s %q: error %v, wantErr %v", tc.fieldType, tc.value, err, tc.wantErr)
			continue
		}
		if !tc.wantErr && !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s %q = %#v, want %#v", tc.fieldType, tc.value, got, tc.want)
		}
	}
}
