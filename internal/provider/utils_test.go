package provider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
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

var testPrefix = "test"

func testAccGetTestName(testSlug string) string {
	randomString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	return strings.Join([]string{testPrefix, testSlug, randomString}, "-")
}

func testAccGetTestCustomFieldName(testSlug string) string {
	randomString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	return strings.Join([]string{testPrefix, testSlug, randomString}, "_")
}
