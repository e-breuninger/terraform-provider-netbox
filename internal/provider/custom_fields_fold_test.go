package provider

// Locks keepPriorCustomFields (normalize_gen.go): the configured spelling of a typed custom field
// survives NetBox's canonical echo, everything else follows NetBox.

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestKeepPriorCustomFields(t *testing.T) {
	stringMap := func(kv map[string]string) types.Map {
		elements := map[string]attr.Value{}
		for key, value := range kv {
			elements[key] = types.StringValue(value)
		}
		return types.MapValueMust(types.StringType, elements)
	}
	prior := stringMap(map[string]string{
		"ratio":   "1.50",
		"tier":    "007",
		"tags":    "{\"a\": 1, \"b\": [1, 2]}",
		"note":    "1.50",
		"removed": "x",
	})
	next := stringMap(map[string]string{
		"ratio": "1.5",
		"tier":  "7",
		"tags":  "{\"a\":1,\"b\":[1,2]}",
		"note":  "1.50 (edited)",
		"added": "new",
	})
	got := keepPriorCustomFields(prior, next).Elements()
	for key, want := range map[string]string{
		"ratio": "1.50",                      // decimal: same number, configured spelling kept
		"tier":  "007",                       // integer: same number, configured spelling kept
		"tags":  "{\"a\": 1, \"b\": [1, 2]}", // JSON: same document, configured spelling kept
		"note":  "1.50 (edited)",             // text: a different value follows NetBox
		"added": "new",                       // a key NetBox added appears
	} {
		if value, has := got[key]; !has || value.(types.String).ValueString() != want {
			t.Errorf("%s = %v, want %q", key, got[key], want)
		}
	}
	if _, has := got["removed"]; has {
		t.Error("a key NetBox dropped must not survive")
	}
	if !keepPriorCustomFields(types.MapNull(types.StringType), next).Equal(next) {
		t.Error("without a prior value (data sources, import) NetBox's map is taken as is")
	}
	if !keepPriorCustomFields(prior, prior).Equal(prior) {
		t.Error("an unchanged map must come back unchanged")
	}
}
