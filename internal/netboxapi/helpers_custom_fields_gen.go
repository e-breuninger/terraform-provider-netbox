// Code generated automatically. DO NOT EDIT.

// Handling of the custom_fields attribute every resource carries: converting its string values
// to and from the types NetBox uses for each kind of custom field.

package netboxapi

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
)

// TypedCustomFields converts the string custom field values in body to the
// types their NetBox definitions require (integer, decimal, boolean,
// multiselect, json, object, multiobject), looked up live from the custom
// field definitions. String-typed fields and unknown names pass through
// unchanged. nil values (keys cleared by ClearRemovedCustomFields) stay null.
func TypedCustomFields(ctx context.Context, c *Client, body map[string]any) error {
	var fields map[string]any
	switch cf := body["custom_fields"].(type) {
	case map[string]string:
		fields = make(map[string]any, len(cf))
		for name, value := range cf {
			fields[name] = value
		}
	case map[string]any:
		fields = cf
	default:
		return nil
	}
	body["custom_fields"] = fields
	var stringNames []string
	for name, value := range fields {
		if _, ok := value.(string); ok {
			stringNames = append(stringNames, name)
		}
	}
	if len(stringNames) == 0 {
		return nil
	}
	// The definitions of the string-valued fields, page by page: NetBox caps a page at its
	// MAX_PAGE_SIZE.
	params := extras.NewExtrasCustomFieldsListParams().WithName(stringNames)
	defs := map[string]string{}
	offset := int64(0)
	for {
		params.SetOffset(&offset)
		res, err := c.Extras.ExtrasCustomFieldsListContext(ctx, params, nil)
		if err != nil {
			return fmt.Errorf("listing custom field definitions: %w", err)
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, definition := range res.Payload.Results {
			if definition != nil && definition.Name != nil && definition.Type != nil && definition.Type.Value != "" {
				defs[*definition.Name] = definition.Type.Value
			}
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}
	for name, value := range fields {
		s, ok := value.(string)
		if !ok {
			continue
		}
		t, ok := defs[name]
		if !ok {
			continue
		}
		tv, err := typedCustomFieldValue(name, t, s)
		if err != nil {
			return err
		}
		fields[name] = tv
	}
	return nil
}

// typedCustomFieldValue converts one string value to the JSON type NetBox
// expects for the field type. The conversions are exact inverses of
// customFieldValues so values round-trip without drift.
func typedCustomFieldValue(name, fieldType, v string) (any, error) {
	switch fieldType {
	case "integer":
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not an integer", name, v)
		}
		return n, nil
	case "decimal":
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not a decimal", name, v)
		}
		return f, nil
	case "boolean":
		switch v {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
		return nil, fmt.Errorf("custom field %q: %q is not a boolean (use \"true\" or \"false\")", name, v)
	case "multiselect":
		var xs []any
		if err := json.Unmarshal([]byte(v), &xs); err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not a JSON array", name, v)
		}
		return xs, nil
	case "json":
		var x any
		if err := json.Unmarshal([]byte(v), &x); err != nil {
			return v, nil // bare text is sent as a JSON string
		}
		return x, nil
	case "object":
		var x any
		if err := json.Unmarshal([]byte(v), &x); err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not an object id", name, v)
		}
		id, ok := relatedObjectID(x)
		if !ok {
			return nil, fmt.Errorf("custom field %q: %q is not an object id", name, v)
		}
		return id, nil
	case "multiobject":
		var xs []any
		if err := json.Unmarshal([]byte(v), &xs); err != nil {
			return nil, fmt.Errorf("custom field %q: %q is not a JSON array of object ids", name, v)
		}
		ids := make([]int64, 0, len(xs))
		for _, x := range xs {
			id, ok := relatedObjectID(x)
			if !ok {
				return nil, fmt.Errorf("custom field %q: %q is not a JSON array of object ids", name, v)
			}
			ids = append(ids, id)
		}
		return ids, nil
	default:
		return v, nil
	}
}

// relatedObjectID is the id an object custom field value denotes: a positive whole number, or a
// related object as NetBox returns it (id, url and display), which an older state may still hold.
// Numbers are float64 (encoding/json) or json.Number (go-netbox responses).
func relatedObjectID(x any) (int64, bool) {
	switch value := x.(type) {
	case float64:
		if value > 0 && value == float64(int64(value)) {
			return int64(value), true
		}
	case json.Number:
		if id, err := value.Int64(); err == nil && id > 0 {
			return id, true
		}
	case map[string]any:
		if isRelatedObject(value) {
			return relatedObjectID(value["id"])
		}
	}
	return 0, false
}

// isRelatedObject reports whether a custom field value is a related object as NetBox returns it
// for object and multiobject fields: a map with an id, url and display.
func isRelatedObject(value map[string]any) bool {
	_, hasURL := value["url"]
	_, hasDisplay := value["display"]
	if !hasURL || !hasDisplay {
		return false
	}
	_, hasID := relatedObjectID(value["id"])
	return hasID
}

// relatedObjectIDs is the sorted ids of a multiobject custom field value as NetBox returns it (a
// non-empty list of related objects); ok is false for any other value.
func relatedObjectIDs(value []any) (ids []int64, ok bool) {
	if len(value) == 0 {
		return nil, false
	}
	for _, element := range value {
		object, isObject := element.(map[string]any)
		if !isObject || !isRelatedObject(object) {
			return nil, false
		}
		id, _ := relatedObjectID(object)
		ids = append(ids, id)
	}
	slices.Sort(ids)
	return ids, true
}

// customFieldValues converts NetBox's custom_fields object (arbitrary JSON
// values, unset fields null) to string values: strings as-is, numbers in
// their shortest form, booleans as true/false, a related object (object
// field) as its id, a list of them (multiobject field) as a JSON list of the
// sorted ids, anything else JSON-encoded.
// Null entries are dropped, so an object without values is an empty map
// (custom_fields = {} round-trips); only a missing object yields nil.
func customFieldValues(v any) map[string]string {
	m, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	out := map[string]string{}
	for name, value := range m {
		switch x := value.(type) {
		case nil:
		case string:
			out[name] = x
		case float64:
			out[name] = strconv.FormatFloat(x, 'f', -1, 64)
		case bool:
			out[name] = strconv.FormatBool(x)
		default:
			if object, isObject := x.(map[string]any); isObject && isRelatedObject(object) {
				id, _ := relatedObjectID(object)
				out[name] = strconv.FormatInt(id, 10)
				continue
			}
			if list, isList := x.([]any); isList {
				if ids, ok := relatedObjectIDs(list); ok {
					value = ids
				}
			}
			raw, err := json.Marshal(value)
			if err != nil {
				continue
			}
			out[name] = string(raw)
		}
	}
	return out
}

// ClearRemovedCustomFields adds a null entry to body's custom_fields for
// every key of prior (the previous state) the plan no longer sets, so
// NetBox clears it instead of keeping the old value.
func ClearRemovedCustomFields(body map[string]any, prior map[string]string) {
	if len(prior) == 0 {
		return
	}
	cf, _ := body["custom_fields"].(map[string]string)
	merged := map[string]any{}
	for name, value := range cf {
		merged[name] = value
	}
	for name := range prior {
		if _, ok := cf[name]; !ok {
			merged[name] = nil
		}
	}
	body["custom_fields"] = merged
}
