package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTokenStateAfterAPIResponse(t *testing.T) {
	for _, tc := range []struct {
		name            string
		version         types.Int64
		apiToken, prior types.String
		want            types.String
	}{
		{"write-only ignores echoed secret", types.Int64Value(1), types.StringValue("echo"), types.StringNull(), types.StringNull()},
		{"write-only ignores prior secret", types.Int64Value(2), types.StringNull(), types.StringValue("prior"), types.StringNull()},
		{"legacy keeps echoed secret", types.Int64Null(), types.StringValue("echo"), types.StringNull(), types.StringValue("echo")},
		{"legacy keeps prior secret", types.Int64Null(), types.StringNull(), types.StringValue("prior"), types.StringValue("prior")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := writeOnlyTokenResourceModel{
				tokenResourceModel: tokenResourceModel{Token: tc.apiToken},
				TokenWoVersion:     tc.version,
			}
			model.restoreToken(tc.prior)
			if !model.Token.Equal(tc.want) {
				t.Fatal("unexpected token retention in Terraform state")
			}
		})
	}
}
