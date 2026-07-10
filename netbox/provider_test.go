package netbox

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	semver "github.com/Masterminds/semver/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
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

// testAccNetboxVersionAtLeast returns true if the NETBOX_VERSION env var is >= the given constraint.
// If NETBOX_VERSION is unset or unparseable, it returns false.
func testAccNetboxVersionAtLeast(constraint string) bool {
	raw := strings.TrimPrefix(os.Getenv("NETBOX_VERSION"), "v")
	v, err := semver.NewVersion(raw)
	if err != nil {
		return false
	}
	c, err := semver.NewConstraint(">= " + constraint)
	if err != nil {
		return false
	}
	return c.Check(v)
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

// TestProviderConfigure_nonJSONStatusResponse verifies that when the Netbox
// version check (GET /api/status/) gets a non-JSON response - for example
// from a misconfigured server_url or a proxy/ingress returning an HTML error
// page - providerConfigure returns a clear, actionable diagnostic instead of
// the raw go-openapi "is not supported by the TextConsumer" error. See
// https://github.com/e-breuninger/terraform-provider-netbox/issues/263,
// https://github.com/e-breuninger/terraform-provider-netbox/issues/606, and
// others reporting this same unhelpful error message.
func TestProviderConfigure_nonJSONStatusResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<html><body>not json</body></html>"))
	}))
	defer ts.Close()

	raw := map[string]interface{}{
		"server_url": ts.URL,
		"api_token":  "0123456789abcdef0123456789abcdef01234567",
	}
	d := schema.TestResourceDataRaw(t, Provider().Schema, raw)

	_, diags := providerConfigure(context.Background(), d)

	if assert.True(t, diags.HasError(), "expected providerConfigure to return an error") {
		found := false
		for _, diagnostic := range diags {
			if diagnostic.Severity != diag.Error {
				continue
			}
			found = true
			assert.Equal(t, "Failed to determine Netbox version", diagnostic.Summary)
			assert.Contains(t, diagnostic.Detail, "skip_version_check")
			assert.Contains(t, diagnostic.Detail, ts.URL+"/api/status/")
			assert.True(t, strings.Contains(diagnostic.Detail, "Original error:"),
				"expected the original go-openapi error to still be included for debugging, prefixed by an explanation")
		}
		assert.True(t, found, "expected an error diagnostic")
	}
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
