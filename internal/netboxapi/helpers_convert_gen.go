// Code generated automatically. DO NOT EDIT.

// Value conversions between NetBox's JSON and the flat request and response DTOs of the generated
// adapters.

package netboxapi

import (
	"encoding/json"
	"regexp"
	"strings"
)

var (
	slugSpecial    = regexp.MustCompile(`[^\w\s-]`)
	slugWhitespace = regexp.MustCompile(`[\s-]+`)
)

// Slugify derives a NetBox slug from a name: special characters are stripped, runs of
// whitespace and dashes collapse to a single dash, and the result is trimmed and lowercased.
func Slugify(name string) string {
	s := slugSpecial.ReplaceAllString(name, "")
	s = slugWhitespace.ReplaceAllString(s, "-")
	return strings.ToLower(strings.Trim(s, "-"))
}

// reencode copies src into dst through JSON. It is how nested go-netbox
// values — typed nested models, slices of them, or the untyped interface{}
// NetBox uses for polymorphic fields such as assigned_object — become the
// generated nested response DTOs, matched by JSON field name.
func reencode(src, dst any) error {
	raw, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, dst)
}

// genericObjects renders one side of a cable for the wire: one
// {object_type, object_id} per id, all of the side's type.
func genericObjects(objectType *string, ids []int64) []map[string]any {
	out := make([]map[string]any, 0, len(ids))
	for _, id := range ids {
		var t any
		if objectType != nil {
			t = *objectType
		}
		out = append(out, map[string]any{"object_type": t, "object_id": id})
	}
	return out
}

// terminationSideOf folds NetBox's generic object list of one cable side back
// into its type and ids (through JSON, as the nested adapters do): the type
// is the first object's, nil when the side is empty.
func terminationSideOf(src any) (objectType *string, ids []int64) {
	var objs []map[string]any
	if err := reencode(src, &objs); err != nil {
		return nil, nil
	}
	ids = []int64{}
	for _, obj := range objs {
		if t, ok := obj["object_type"].(string); ok && objectType == nil {
			objectType = &t
		}
		if id, ok := obj["object_id"].(float64); ok {
			ids = append(ids, int64(id))
		}
	}
	return objectType, ids
}

// choiceValue extracts the value of a go-netbox choice object, whose
// Value field is *T on most models and T on some; the zero value (an
// unset choice) becomes nil.
func choiceValue[T comparable](v any) *T {
	var zero T
	switch x := v.(type) {
	case *T:
		if x == nil || *x == zero {
			return nil
		}
		return x
	case T:
		if x == zero {
			return nil
		}
		return &x
	}
	return nil
}

// jsonText renders an untyped JSON value as JSON text; nil stays nil.
func jsonText(v any) *string {
	if v == nil {
		return nil
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	s := string(raw)
	return &s
}

// parseJSONText decodes JSON text for an untyped field; text that is not
// valid JSON is sent as a JSON string.
func parseJSONText(s string) any {
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return s
	}
	return v
}
