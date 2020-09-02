package netbox

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
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

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NETBOX_SERVER_URL"); v == "" {
		t.Fatal("NETBOX_SERVER must be set for acceptance tests.")
	}
	if v := os.Getenv("NETBOX_API_TOKEN"); v == "" {
		t.Fatal("NETBOX_API_TOKEN must be set for acceptance tests.")
	}
}
