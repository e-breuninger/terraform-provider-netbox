package provider

// Locks keepPriorString (normalize_gen.go): an optional string keeps its
// configured spelling where NetBox stored the same text canonically (empty
// string read back as null, surrounding whitespace trimmed); anything else
// follows NetBox.

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestKeepPriorString(t *testing.T) {
	for _, tc := range []struct {
		name        string
		prior, next types.String
		want        types.String
	}{
		{"empty string read back as null", types.StringValue(""), types.StringNull(), types.StringValue("")},
		{"trailing newline trimmed", types.StringValue("note\n"), types.StringValue("note"), types.StringValue("note\n")},
		{"surrounding spaces trimmed", types.StringValue("  padded  "), types.StringValue("padded"), types.StringValue("  padded  ")},
		{"same value", types.StringValue("x"), types.StringValue("x"), types.StringValue("x")},
		{"changed outside Terraform", types.StringValue("x"), types.StringValue("y"), types.StringValue("y")},
		{"cleared outside Terraform", types.StringValue("x"), types.StringNull(), types.StringNull()},
		{"unset in config", types.StringNull(), types.StringValue("y"), types.StringValue("y")},
		{"unknown prior (create)", types.StringUnknown(), types.StringNull(), types.StringNull()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := keepPriorString(tc.prior, tc.next); !got.Equal(tc.want) {
				t.Errorf("keepPriorString(%s, %s) = %s, want %s", tc.prior, tc.next, got, tc.want)
			}
		})
	}
}
