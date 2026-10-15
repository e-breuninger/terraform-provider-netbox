// Code generated automatically. DO NOT EDIT.

package provider

import (
	"encoding/json"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// keepPriorCustomFields returns next with, per key, the prior string kept
// where NetBox answered the same value in its canonical spelling: custom
// field values are typed on the wire, so "1.50" on a decimal field comes
// back as "1.5", "007" on an integer field as "7" and a JSON field compact.
// A text field echoes its text as is, so two differing strings that denote
// one number or one JSON document can only be a canonicalised typed field.
func keepPriorCustomFields(prior, next types.Map) types.Map {
	if prior.IsNull() || prior.IsUnknown() || next.IsNull() || next.IsUnknown() {
		return next
	}
	priorElements, nextElements := prior.Elements(), next.Elements()
	kept := make(map[string]attr.Value, len(nextElements))
	changed := false
	for key, nextValue := range nextElements {
		kept[key] = nextValue
		priorString, priorIsString := priorElements[key].(types.String)
		nextString, nextIsString := nextValue.(types.String)
		if !priorIsString || !nextIsString || priorString.IsNull() || priorString.IsUnknown() || nextString.IsNull() || nextString.IsUnknown() {
			continue
		}
		if priorString.ValueString() != nextString.ValueString() && sameCustomFieldValue(priorString.ValueString(), nextString.ValueString()) {
			kept[key] = priorString
			changed = true
		}
	}
	if !changed {
		return next
	}
	return types.MapValueMust(types.StringType, kept)
}

// sameCustomFieldValue reports whether two custom field strings denote one
// value: the same number, the same JSON document, or, for an object or
// multiobject field, the same related objects (configured as ids, answered by
// NetBox as nested objects, in any order).
func sameCustomFieldValue(a, b string) bool {
	numberA, errA := strconv.ParseFloat(a, 64)
	numberB, errB := strconv.ParseFloat(b, 64)
	if errA == nil && errB == nil {
		return numberA == numberB
	}
	var documentA, documentB any
	if json.Unmarshal([]byte(a), &documentA) != nil || json.Unmarshal([]byte(b), &documentB) != nil {
		return false
	}
	if reflect.DeepEqual(documentA, documentB) {
		return true
	}
	idsA, nestedA, okA := relatedObjectIDs(documentA)
	idsB, nestedB, okB := relatedObjectIDs(documentB)
	return okA && okB && (nestedA || nestedB) && reflect.DeepEqual(idsA, idsB)
}

// relatedObjectIDs reads an object or multiobject custom field value as sorted
// ids: an id, a nested object as NetBox returns it, or a list of either. nested
// reports whether NetBox's nested form (id, url and display) was seen, which
// tells an object field apart from a JSON field that merely holds numbers or
// an "id" key.
func relatedObjectIDs(document any) (ids []float64, nested bool, ok bool) {
	one := func(x any) (float64, bool) {
		switch t := x.(type) {
		case float64:
			return t, true
		case map[string]any:
			id, hasID := t["id"].(float64)
			_, hasURL := t["url"]
			_, hasDisplay := t["display"]
			if hasID && hasURL && hasDisplay {
				nested = true
				return id, true
			}
		}
		return 0, false
	}
	list, isList := document.([]any)
	if !isList {
		list = []any{document}
	}
	for _, x := range list {
		id, isID := one(x)
		if !isID {
			return nil, false, false
		}
		ids = append(ids, id)
	}
	sort.Float64s(ids)
	return ids, nested, true
}

// keepPriorFold returns prior when next differs from it only in letter
// case (the API case-normalized the value, e.g. a MAC address uppercased
// by NetBox), and next otherwise.
func keepPriorFold(prior, next types.String) types.String {
	if prior.IsNull() || prior.IsUnknown() || next.IsNull() || next.IsUnknown() {
		return next
	}
	if strings.EqualFold(prior.ValueString(), next.ValueString()) {
		return prior
	}
	return next
}

// reorderToPrior returns prior when next holds exactly the same elements
// in a different order, and next otherwise.
func reorderToPrior(prior, next types.List) types.List {
	if prior.IsNull() || prior.IsUnknown() || next.IsNull() || next.IsUnknown() {
		return next
	}
	pe, ne := prior.Elements(), next.Elements()
	if len(pe) != len(ne) {
		return next
	}
	used := make([]bool, len(ne))
outer:
	for _, priorElement := range pe {
		for i, nextElement := range ne {
			if !used[i] && priorElement.Equal(nextElement) {
				used[i] = true
				continue outer
			}
		}
		return next
	}
	return prior
}

// keepPriorSide returns prior when next is the same cable side — same
// object_type, same ids in any order — since NetBox returns terminations by
// connector and pk rather than as configured; the companion-derived children
// of prior stay valid in that case. Anything else returns next.
func keepPriorSide(prior, next types.Object) types.Object {
	if prior.IsNull() || prior.IsUnknown() || next.IsNull() || next.IsUnknown() {
		return next
	}
	pa, na := prior.Attributes(), next.Attributes()
	pt, nt := pa["object_type"], na["object_type"]
	if pt == nil || nt == nil || !pt.Equal(nt) {
		return next
	}
	pl, ok1 := pa["ids"].(types.List)
	nl, ok2 := na["ids"].(types.List)
	if !ok1 || !ok2 || pl.IsNull() || pl.IsUnknown() || nl.IsNull() || nl.IsUnknown() {
		return next
	}
	if reorderToPrior(pl, nl).Equal(pl) {
		return prior
	}
	return next
}
