package provider

// Locks sameCustomFieldValue for object and multiobject custom fields: ids as
// configured match the nested objects NetBox answers, in any order; JSON fields
// that merely hold numbers or an "id" key are not mistaken for them.

import "testing"

func TestSameCustomFieldValueObjects(t *testing.T) {
	site := func(id string) string {
		return `{"id": ` + id + `, "url": "https://netbox/api/dcim/sites/` + id + `/", "display": "s` + id + `"}`
	}
	for _, tc := range []struct {
		a, b string
		want bool
	}{
		{"12", site("12"), true},
		{"13", site("12"), false},
		{"[2, 1]", "[" + site("1") + ", " + site("2") + "]", true},
		{"[1]", "[" + site("1") + ", " + site("2") + "]", false},
		{`{"id": 5, "name": "x"}`, `{"id": 5, "name": "y"}`, false}, // JSON field, not an object field
		{"5", "[5]", false},
		{"[1, 2]", "[2, 1]", false}, // JSON lists keep their order
	} {
		if got := sameCustomFieldValue(tc.a, tc.b); got != tc.want {
			t.Errorf("sameCustomFieldValue(%s, %s) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
