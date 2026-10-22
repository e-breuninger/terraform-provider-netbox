// Code generated automatically. DO NOT EDIT.

package provider

import (
	"encoding/json"
	"math"
	"reflect"
	"slices"
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
// value: the same number, the same JSON document, or two JSON lists of the
// same whole numbers in any order (a multiobject field's ids, which NetBox
// answers in the related objects' order).
func sameCustomFieldValue(a, b string) bool {
	if numberA, err := strconv.ParseFloat(a, 64); err == nil {
		numberB, err := strconv.ParseFloat(b, 64)
		return err == nil && numberA == numberB
	}
	var documentA, documentB any
	if json.Unmarshal([]byte(a), &documentA) != nil || json.Unmarshal([]byte(b), &documentB) != nil {
		return false
	}
	if reflect.DeepEqual(documentA, documentB) {
		return true
	}
	idsA, isIDsA := sortedWholeNumbers(documentA)
	idsB, isIDsB := sortedWholeNumbers(documentB)
	return isIDsA && isIDsB && slices.Equal(idsA, idsB)
}

// sortedWholeNumbers is a JSON list of whole numbers, sorted; ok is false for any other document.
func sortedWholeNumbers(document any) (numbers []float64, ok bool) {
	list, isList := document.([]any)
	if !isList {
		return nil, false
	}
	for _, element := range list {
		number, isNumber := element.(float64)
		if !isNumber || number != math.Trunc(number) {
			return nil, false
		}
		numbers = append(numbers, number)
	}
	slices.Sort(numbers)
	return numbers, true
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

// keepPriorString returns prior when next is the same text as NetBox stores
// it: prior with surrounding whitespace trimmed, or null for a prior "". It
// returns next otherwise.
func keepPriorString(prior, next types.String) types.String {
	if prior.IsNull() || prior.IsUnknown() || next.IsUnknown() {
		return next
	}
	if next.IsNull() {
		if prior.ValueString() == "" {
			return prior
		}
		return next
	}
	if strings.TrimSpace(prior.ValueString()) == next.ValueString() {
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
