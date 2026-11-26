//go:build acctest

package provider_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// testAccNetboxAtLeast reports whether the NetBox under test is major.minor or later.
func testAccNetboxAtLeast(t *testing.T, major, minor int) bool {
	t.Helper()
	client, err := testAccClient()
	if err != nil {
		t.Fatalf("error getting client: %s", err)
	}
	_, version, err := netboxapi.Version(context.Background(), client)
	if err != nil {
		t.Fatalf("error reading the NetBox version: %s", err)
	}
	parts := strings.SplitN(version, ".", 3)
	maj, _ := strconv.Atoi(parts[0])
	minr := 0
	if len(parts) > 1 {
		minr, _ = strconv.Atoi(parts[1])
	}
	return maj > major || (maj == major && minr >= minor)
}

// NetBox 4.7 generates every token's secret and ignores a chosen one. TestAccNetboxToken_basic and
// _v1 choose secrets, so they run up to 4.6; these cover 4.7 and later: a generated token through
// create, update and import, and a chosen secret refused at plan time, before anything is created.
func TestAccNetboxToken_generated(t *testing.T) {
	if !testAccNetboxAtLeast(t, 4, 7) {
		t.Skip("NetBox before 4.7 accepts a chosen secret: TestAccNetboxToken_basic covers it")
	}
	testName := testAccGetTestName("token_gen")
	deps := fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Acceptance-Test-Pw-9271"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_token" "test" {
  user_id       = netbox_user.test.id
  description   = "Acceptance test token."
  write_enabled = false
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_token.test", "user_id", "netbox_user.test", "id"),
					resource.TestCheckResourceAttr("netbox_token.test", "write_enabled", "false"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "key"),
					resource.TestCheckResourceAttrWith("netbox_token.test", "token", func(v string) error {
						if len(v) != 40 {
							return fmt.Errorf("generated token %q is not 40 characters", v)
						}
						return nil
					}),
				),
			},
			{
				// An update leaves the generated secret alone.
				Config: deps + `
resource "netbox_token" "test" {
  user_id       = netbox_user.test.id
  description   = "relabelled"
  write_enabled = false
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_token.test", "description", "relabelled"),
					resource.TestCheckResourceAttrSet("netbox_token.test", "token"),
				),
			},
			{
				ResourceName:            "netbox_token.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token"},
			},
		},
	})
}

func TestAccNetboxToken_chosenSecretRefused(t *testing.T) {
	if !testAccNetboxAtLeast(t, 4, 7) {
		t.Skip("NetBox before 4.7 accepts a chosen secret: TestAccNetboxToken_basic covers it")
	}
	testName := testAccGetTestName("token_chosen")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Acceptance-Test-Pw-9271"
}
resource "netbox_token" "chosen" {
  user_id = netbox_user.test.id
  token   = "0123456789abcdef0123456789abcdef01234567"
}`, testName),
				ExpectError: regexp.MustCompile(`NetBox generates the token's secret`),
			},
		},
	})
}
