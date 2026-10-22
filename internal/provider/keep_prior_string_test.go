package provider

// Locks keepPriorString (normalize_gen.go): configured text survives NetBox's trimmed or null echo
// of the same text, everything else follows NetBox.

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestKeepPriorString(t *testing.T) {
	for _, testCase := range []struct {
		name        string
		prior, next types.String
		want        types.String
	}{
		{"empty string read back as null", types.StringValue(""), types.StringNull(), types.StringValue("")},
		{"trailing newline trimmed", types.StringValue("text\n"), types.StringValue("text"), types.StringValue("text\n")},
		{"surrounding spaces trimmed", types.StringValue("  text  "), types.StringValue("text"), types.StringValue("  text  ")},
		{"whitespace only read back as null", types.StringValue("  "), types.StringNull(), types.StringNull()},
		{"changed text follows NetBox", types.StringValue("text"), types.StringValue("other"), types.StringValue("other")},
		{"inner whitespace is not NetBox's canonical form", types.StringValue("a  b"), types.StringValue("a b"), types.StringValue("a b")},
		{"cleared text follows NetBox", types.StringValue("text"), types.StringNull(), types.StringNull()},
		{"no prior value (import) takes NetBox's", types.StringNull(), types.StringValue("text"), types.StringValue("text")},
		{"unknown prior (computed) takes NetBox's", types.StringUnknown(), types.StringValue("text"), types.StringValue("text")},
	} {
		if got := keepPriorString(testCase.prior, testCase.next); !got.Equal(testCase.want) {
			t.Errorf("%s: keepPriorString(%v, %v) = %v, want %v", testCase.name, testCase.prior, testCase.next, got, testCase.want)
		}
	}
}
