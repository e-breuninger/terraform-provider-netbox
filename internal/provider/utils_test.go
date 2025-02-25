package provider

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/netbox"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"log"
	"os"
	"strings"
	"testing"
)

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NETBOX_API_TOKEN"); v == "" {
		t.Fatal("Environment variable NETBOX_API_TOKEN must be set!")
	}
	if v := os.Getenv("NETBOX_SERVER_URL"); v == "" {
		t.Fatal("Environment variable NETBOX_SERVER_URL must be set!")
	}
}

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	// newProvider is an example function that returns a provider.Provider
	"netbox": func() (tfprotov6.ProviderServer, error) {
		upgradedSdkServer, err := tf5to6server.UpgradeServer(
			context.Background(),
			netbox.Provider().GRPCProvider,
		)

		if err != nil {
			log.Fatal(err)
		}

		providers := []func() tfprotov6.ProviderServer{
			providerserver.NewProtocol6(New()()),
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

var testPrefix = "test"

func testAccGetTestName(testSlug string) string {
	randomString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	return strings.Join([]string{testPrefix, testSlug, randomString}, "-")
}
