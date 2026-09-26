// Code generated automatically. DO NOT EDIT.

package provider

import (
	"encoding/json"
	"reflect"
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
// value: the same number, or the same JSON document.
func sameCustomFieldValue(a, b string) bool {
	if numberA, err := strconv.ParseFloat(a, 64); err == nil {
		numberB, err := strconv.ParseFloat(b, 64)
		return err == nil && numberA == numberB
	}
	var documentA, documentB any
	if json.Unmarshal([]byte(a), &documentA) != nil || json.Unmarshal([]byte(b), &documentB) != nil {
		return false
	}
	return reflect.DeepEqual(documentA, documentB)
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
