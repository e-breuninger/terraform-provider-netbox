package netbox

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		return &providerState{NetBoxAPI: netboxClient}, diags
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

func TestAccNetboxProviderDefaultTags(t *testing.T) {
	defaultTag := fmt.Sprintf("managed-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"netbox": func() (*schema.Provider, error) {
				p := Provider()
				p.ConfigureContextFunc = func(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
					rd.Set("default_tags", []string{defaultTag})
					return providerConfigure(ctx, rd)
				}
				return p, nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "netbox_tag" "managed" {
						name = "%s"
					}

					resource "netbox_site" "testsite" {
						name = "%s"

						depends_on = [
							netbox_tag.managed
						]
					}
					`, defaultTag, acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.testsite", "tags_all.#", "1"),
					resource.TestCheckResourceAttr("netbox_site.testsite", "tags_all.0", defaultTag),
				),
			},
		},
	})
}
