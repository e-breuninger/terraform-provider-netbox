package netbox

import (
	"context"
	"fmt"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testPrefix = "test"

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"netbox": testAccProvider,
	}
}

func testAccGetTestName(testSlug string) string {
	randomString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	return strings.Join([]string{testPrefix, testSlug, randomString}, "-")
}

func testAccGetTestToken() string {
	randomToken := acctest.RandStringFromCharSet(40, "0123456789")
	return randomToken
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NETBOX_SERVER_URL"); v == "" {
		t.Fatal("NETBOX_SERVER_URL must be set for acceptance tests.")
	}
	if v := os.Getenv("NETBOX_API_TOKEN"); v == "" {
		t.Fatal("NETBOX_API_TOKEN must be set for acceptance tests.")
	}
}

func testProviderConfig(platform string) string {
	return fmt.Sprintf(`
	resource "netbox_platform" "testplatform" {
    name = "%s"
	}`, platform)
}

func providerInvalidConfigure() schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		config := &Config{}
		config.ServerURL = "https://fake.netbox.server"
		config.APIToken = "1234567890"

		netboxClient, clientError := config.Client()
		if clientError != nil {
			return nil, diag.FromErr(clientError)
		}

		return netboxClient, diags
	}
}

func TestAccNetboxProviderConfigure_failure(t *testing.T) {
	var testAccInvalidProviders map[string]*schema.Provider

	testAccInvalidProvider := Provider()
	testAccInvalidProvider.ConfigureContextFunc = providerInvalidConfigure()
	testAccInvalidProviders = map[string]*schema.Provider{
		"netbox": testAccInvalidProvider,
	}

	resource.Test(t, resource.TestCase{
		Providers: testAccInvalidProviders,
		Steps: []resource.TestStep{
			{
				Config:      testProviderConfig(acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)),
				ExpectError: regexp.MustCompile("Post \"https://fake.netbox.server/api/dcim/platforms/\": dial tcp: lookup fake.netbox.server.*: no such host"),
			},
		},
	})
}

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
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
