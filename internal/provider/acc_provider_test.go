//go:build acctest

// Package provider_test holds the acceptance tests of the generated NetBox provider. The files live
// in acctest/ of terraform-codegen-netbox and are copied next to the generated code by "mage
// acctest" (they are *_test.go companions, so regeneration never touches them). They are adapted
// from the pre-6.0 provider's suite; where the generated provider's schema
// deliberately differs (null instead of "" for unset strings, tags by slug), the tests follow the
// generated provider.
package provider_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider"
)

// testPrefix marks objects created by the suite; sweepers delete every object whose name starts
// with it.
const testPrefix = "test"

// testAccProtoV6ProviderFactories serves the generated provider to the test harness over protocol
// 6.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"netbox": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func testAccGetTestName(testSlug string) string {
	randomString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	return strings.Join([]string{testPrefix, testSlug, randomString}, "-")
}

func testAccPreCheck(t *testing.T) {
	t.Helper()
	if v := os.Getenv("NETBOX_SERVER_URL"); v == "" {
		t.Fatal("NETBOX_SERVER_URL must be set for acceptance tests.")
	}
	if v := os.Getenv("NETBOX_API_TOKEN"); v == "" {
		t.Fatal("NETBOX_API_TOKEN must be set for acceptance tests.")
	}
}

// getSlug is the slug the provider derives from a name.
func getSlug(name string) string { return netboxapi.Slugify(name) }

// testAccClient returns a go-netbox client for out-of-band checks and sweepers.
func testAccClient() (*client.NetBoxAPI, error) {
	return netboxapi.New(os.Getenv("NETBOX_SERVER_URL"), os.Getenv("NETBOX_API_TOKEN"), netboxapi.Options{})
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

type sweepItem struct {
	id   int64
	name string
}

// sweep registers a sweeper that deletes every object with the test prefix via the given list and
// delete callbacks.
func sweep(name string, list func(client *client.NetBoxAPI) ([]sweepItem, error), del func(client *client.NetBoxAPI, id int64) error) {
	resource.AddTestSweepers(name, &resource.Sweeper{
		Name: name,
		F: func(region string) error {
			client, err := testAccClient()
			if err != nil {
				return fmt.Errorf("error getting client: %w", err)
			}
			items, err := list(client)
			if err != nil {
				return err
			}
			for _, item := range items {
				if !strings.HasPrefix(item.name, testPrefix) {
					continue
				}
				if err := del(client, item.id); err != nil {
					return err
				}
				log.Printf("[DEBUG] swept %s %d (%s)", name, item.id, item.name)
			}
			return nil
		},
	})
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
