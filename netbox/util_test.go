package netbox

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"log"
	"testing"
)

func TestJoinStringWithFinalConjunction(t *testing.T) {
	for _, tt := range []struct {
		name     string
		list     []string
		sep      string
		con      string
		expected string
	}{
		{
			name:     "Full",
			list:     []string{"foo", "bar", "baz"},
			sep:      ", ",
			con:      "and",
			expected: "foo, bar and baz",
		},
		{
			name:     "OnlyTwoItems",
			list:     []string{"foo", "bar"},
			sep:      ", ",
			con:      "and",
			expected: "foo and bar",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinStringWithFinalConjunction(tt.list, tt.sep, tt.con)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}

func TestBuildValidValueDescription(t *testing.T) {
	for _, tt := range []struct {
		name     string
		list     []string
		expected string
	}{
		{
			name:     "Full",
			list:     []string{"foo", "bar", "baz"},
			expected: "Valid values are `foo`, `bar` and `baz`",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := buildValidValueDescription(tt.list)
			if actual != tt.expected {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}

func TestJsonSemanticCompareEqual(t *testing.T) {
	a := `{"a": [{ "b": [1, 2, 3]}]}`
	b := `{"a":[{"b":[1,2,3]}]}`

	equal, err := jsonSemanticCompare(a, b)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !equal {
		t.Errorf("expected 'a' and 'b' to be semantically equal\n\na: %s\nb: %s\n", a, b)
	}
}

func TestJsonSemanticCompareUnequal(t *testing.T) {
	a := `{"a": [{ "b": [1, 2, 3]}]}`
	b := `{"a": [{ "b": [1, 2, 4]}]}`

	equal, err := jsonSemanticCompare(a, b)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if equal {
		t.Errorf("expected 'a' and 'b' to be semantically unequal\n\na: %s\nb: %s\n", a, b)
	}
}

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	// newProvider is an example function that returns a provider.Provider
	"netbox": func() (tfprotov6.ProviderServer, error) {
		upgradedSdkServer, err := tf5to6server.UpgradeServer(
			context.Background(),
			Provider().GRPCProvider,
		)

		if err != nil {
			log.Fatal(err)
		}

		providers := []func() tfprotov6.ProviderServer{
			providerserver.NewProtocol6(provider.New()()),
			func() tfprotov6.ProviderServer {
				return upgradedSdkServer
			},
		}

		muxServer, err := tf6muxserver.NewMuxServer(context.Background(), providers...)
		if err != nil {
			log.Fatal(err)
		}
		return muxServer, err
	},
}
