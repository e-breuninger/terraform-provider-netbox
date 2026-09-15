package netbox

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxAvailablePrefixFullDependencies(testName string, parentPrefix string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_prefix" "parent" {
  prefix = "%[2]s"
  description = "%[1]s"
  status = "container"
  tags = [netbox_tag.test.name]
  lifecycle {
    ignore_changes = [%[3]s]
  }
}
`, testName, parentPrefix, customFieldsKey)
}

func TestAccNetboxAvailablePrefix_basic(t *testing.T) {
	testParentPrefix := "1.1.0.0/24"
	testPrefixLength := 25
	expectedPrefix := "1.1.0.0/25"
	testSlug := "prefix"
	testDesc := "test prefix"
	testName := testAccGetTestName(testSlug)

	parentResourceName := "netbox_prefix.parent"
	resourceName := "netbox_available_prefix.test"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  description = "%s"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = true
  is_pool = true
}`, testPrefixLength, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "prefix", expectedPrefix),
					resource.TestCheckResourceAttr(resourceName, "status", "active"),
					resource.TestCheckResourceAttr(resourceName, "description", testDesc),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", testName),
					resource.TestCheckResourceAttr(resourceName, "mark_utilized", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_pool", "true"),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  description = "%s"
  status = "active"
  tags = [netbox_tag.test.name]
  mark_utilized = false
  is_pool = false
}`, testPrefixLength, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "mark_utilized", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_pool", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					parent, ok := s.RootModule().Resources[parentResourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", parentResourceName)
					}
					resource, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}

					return fmt.Sprintf("%s %s %d", parent.Primary.ID, resource.Primary.ID, testPrefixLength), nil
				},
			},
		},
	})
}

func TestAccNetboxAvailablePrefix_cf(t *testing.T) {
	testParentPrefix := "1.1.0.0/24"
	testPrefixLength := 25
	expectedPrefix := "1.1.0.0/25"
	testSlug := "prefix_cf"
	testName := testAccGetTestName(testSlug)

	parentResourceName := "netbox_prefix.parent"
	resourceName := "netbox_available_prefix.test"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name   = "%s"
  type   = "text"
  weight = 100
  content_types = ["ipam.prefix"]
}

resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  status = "active"

  custom_fields = {
    "${netbox_custom_field.test.name}" = "test-field"
  }
}`, testSlug, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "prefix", expectedPrefix),
					resource.TestCheckResourceAttr(resourceName, "status", "active"),
					resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("custom_fields.%s", testSlug), "test-field"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					parent, ok := s.RootModule().Resources[parentResourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", parentResourceName)
					}
					resource, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("Not found: %s", resourceName)
					}

					return fmt.Sprintf("%s %s %d", parent.Primary.ID, resource.Primary.ID, testPrefixLength), nil
				},
			},
		},
	})
}

func TestAccNetboxAvailablePrefix_multiplePrefixesSerial(t *testing.T) {
	testParentPrefix := "1.1.0.0/24"
	testPrefixLength := 25
	expectedPrefixes := []string{
		"1.1.0.0/25",
		"1.1.0.128/25",
	}
	testSlug := "prefix"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test1" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_available_prefix" "test2" {
  depends_on = [netbox_available_prefix.test1]
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = netbox_available_prefix.test1.prefix_length
  status = "active"
  tags = [netbox_tag.test.name]
}
`, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test1", "prefix", expectedPrefixes[0]),
					resource.TestCheckResourceAttr("netbox_available_prefix.test2", "prefix", expectedPrefixes[1]),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test1" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_available_prefix" "test2" {
  depends_on = [netbox_available_prefix.test1]
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = netbox_available_prefix.test1.prefix_length
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_available_prefix" "test3" {
  depends_on = [netbox_available_prefix.test2]
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = netbox_available_prefix.test1.prefix_length
  status = "active"
  tags = [netbox_tag.test.name]
}
`, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test1", "prefix", expectedPrefixes[0]),
					resource.TestCheckResourceAttr("netbox_available_prefix.test2", "prefix", expectedPrefixes[1]),
				),
				ExpectError: regexp.MustCompile(".*Insufficient resources are available to satisfy the request.*"),
			},
		},
	})
}

func testAccNetboxAvailablePrefixScopeDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
}

resource "netbox_region" "test" {
  name = "%[1]s"
}

resource "netbox_site_group" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxAvailablePrefix_scopes(t *testing.T) {
	testParentPrefix := "16.1.0.0/24"
	testPrefixLength := 25
	testSlug := "prefix-scopes"
	testName := testAccGetTestName(testSlug)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + testAccNetboxAvailablePrefixScopeDependencies(testName) + fmt.Sprintf(`
resource "netbox_available_prefix" "with_site_id" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %[2]d
  status = "active"
  site_id = netbox_site.test.id
}`, testSlug, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_prefix.with_site_id", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + testAccNetboxAvailablePrefixScopeDependencies(testName) + fmt.Sprintf(`
resource "netbox_available_prefix" "with_location_id" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %[2]d
  status = "active"
  location_id = netbox_location.test.id
}`, testSlug, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_prefix.with_location_id", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + testAccNetboxAvailablePrefixScopeDependencies(testName) + fmt.Sprintf(`
resource "netbox_available_prefix" "with_region_id" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %[2]d
  status = "active"
  region_id = netbox_region.test.id
}`, testSlug, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_prefix.with_region_id", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + testAccNetboxAvailablePrefixScopeDependencies(testName) + fmt.Sprintf(`
resource "netbox_available_prefix" "with_site_group_id" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %[2]d
  status = "active"
  site_group_id = netbox_site_group.test.id
}`, testSlug, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_prefix.with_site_group_id", "site_group_id", "netbox_site_group.test", "id"),
				),
			},
		},
	})
}

// TestAccNetboxAvailablePrefix_parallel allocates several prefixes out of one
// parent without any depends_on between them, so terraform creates them in
// parallel and they all race for the same pool. Unlike
// TestAccNetboxAvailablePrefix_multiplePrefixesSerial, which chains the
// resources deliberately, this is the case that used to return conflicts from
// Netbox and hand out overlapping prefixes.
//
// The prefixes themselves are allocated in whatever order the requests land,
// so the test checks that every prefix is distinct rather than checking for
// specific values.
func TestAccNetboxAvailablePrefix_parallel(t *testing.T) {
	testParentPrefix := "17.1.0.0/24"
	testPrefixLength := 28
	testSlug := "prefix-parallel"
	testName := testAccGetTestName(testSlug)

	// A /24 split into /28s leaves plenty of room for these.
	const count = 10

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_available_prefix" "test" {
  count = %d
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length = %d
  status = "active"
  tags = [netbox_tag.test.name]
}`, count, testPrefixLength),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPrefixesAreDistinct("netbox_available_prefix.test", count),
				),
			},
		},
	})
}

// testAccCheckPrefixesAreDistinct fails if any two of the given resources ended
// up with the same prefix, which is what an unserialized allocation produces.
func testAccCheckPrefixesAreDistinct(resourcePrefix string, count int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		seen := make(map[string]string, count)
		for i := 0; i < count; i++ {
			name := fmt.Sprintf("%s.%d", resourcePrefix, i)
			rs, ok := s.RootModule().Resources[name]
			if !ok {
				return fmt.Errorf("resource %s not found in state", name)
			}

			prefix := rs.Primary.Attributes["prefix"]
			if prefix == "" {
				return fmt.Errorf("resource %s has no prefix set", name)
			}
			if previous, duplicate := seen[prefix]; duplicate {
				return fmt.Errorf("%s and %s were both allocated %s", previous, name, prefix)
			}
			seen[prefix] = name
		}
		return nil
	}
}

// TestAccNetboxAvailablePrefix_retriesExternalConflict proves retryAllocation
// recovers from a conflict that comes from outside this provider process,
// e.g. a second terraform apply or somebody freeing space in the Netbox UI
// while our request is queued.
//
// It bypasses lockAllocation on purpose and calls the Netbox client directly:
// the parent prefix has room for exactly 3 free /30s (a 4th is held by an
// "occupier" prefix), and 4 goroutines race for it. One of them necessarily
// loses and gets a genuine 409 from Netbox. Shortly after, the occupier is
// deleted, freeing a slot. If retryAllocation is working, the loser's next
// attempt succeeds and all 4 requesters end up with distinct prefixes.
func TestAccNetboxAvailablePrefix_retriesExternalConflict(t *testing.T) {
	testParentPrefix := "60.60.60.0/28" // room for four /30s
	testOccupierPrefix := "60.60.60.0/30"
	testPrefixLength := 30
	testSlug := "prefix-retry-external"
	testName := testAccGetTestName(testSlug)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAvailablePrefixFullDependencies(testName, testParentPrefix) + fmt.Sprintf(`
resource "netbox_prefix" "occupier" {
  prefix      = "%s"
  description = "%s-occupier"
  status      = "active"
  tags        = [netbox_tag.test.name]
  depends_on  = [netbox_prefix.parent]
}`, testOccupierPrefix, testName),
				Check: testAccCheckRetriesExternalPrefixConflict("netbox_prefix.parent", "netbox_prefix.occupier", testPrefixLength),
				// The check deliberately deletes netbox_prefix.occupier out from
				// under Terraform to free the slot mid-retry, so the post-apply
				// refresh sees drift.
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckRetriesExternalPrefixConflict(parentResource, occupierResource string, prefixLength int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*providerState)

		parentRS, ok := s.RootModule().Resources[parentResource]
		if !ok {
			return fmt.Errorf("resource %s not found in state", parentResource)
		}
		parentID, err := strconv.ParseInt(parentRS.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		occupierRS, ok := s.RootModule().Resources[occupierResource]
		if !ok {
			return fmt.Errorf("resource %s not found in state", occupierResource)
		}
		occupierID, err := strconv.ParseInt(occupierRS.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		type allocation struct {
			id     int64
			prefix string
		}

		const requesters = 4
		results := make(chan allocation, requesters)
		errs := make(chan error, requesters)

		var wg sync.WaitGroup
		for i := 0; i < requesters; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				length := int64(prefixLength)
				params := ipam.NewIpamPrefixesAvailablePrefixesCreateParams().
					WithID(parentID).
					WithData(&models.PrefixLength{PrefixLength: &length})

				var got allocation
				err := retryAllocation(context.Background(), func() error {
					res, err := conn.Ipam.IpamPrefixesAvailablePrefixesCreate(params, nil)
					if err != nil {
						return err
					}
					payload := res.GetPayload()
					if payload == nil || payload.Prefix == nil {
						return fmt.Errorf("no available prefix in parent prefix %d", parentID)
					}
					got = allocation{id: payload.ID, prefix: *payload.Prefix}
					return nil
				})
				if err != nil {
					errs <- err
					return
				}
				results <- got
			}()
		}

		// Let the requesters race for the 3 truly free slots so the loser hits its
		// first genuine 409, then free the slot the occupier was holding. This
		// lands comfortably inside the loser's backoff window (250ms-750ms before
		// the first retry, growing from there), well before allocationRetryTimeout.
		time.Sleep(700 * time.Millisecond)
		if _, err := conn.Ipam.IpamPrefixesDelete(ipam.NewIpamPrefixesDeleteParams().WithID(occupierID), nil); err != nil {
			return fmt.Errorf("freeing occupier slot: %w", err)
		}

		wg.Wait()
		close(results)
		close(errs)

		for err := range errs {
			return fmt.Errorf("allocation did not recover from external conflict: %w", err)
		}

		seen := make(map[string]bool, requesters)
		var allocated []allocation
		for got := range results {
			if seen[got.prefix] {
				return fmt.Errorf("prefix %s allocated twice", got.prefix)
			}
			seen[got.prefix] = true
			allocated = append(allocated, got)
		}
		if len(allocated) != requesters {
			return fmt.Errorf("expected %d allocations, got %d", requesters, len(allocated))
		}

		// These prefixes were created directly through the client, so Terraform
		// never learns about them and won't clean them up on its own.
		for _, got := range allocated {
			if _, err := conn.Ipam.IpamPrefixesDelete(ipam.NewIpamPrefixesDeleteParams().WithID(got.id), nil); err != nil {
				return fmt.Errorf("cleaning up allocated prefix %d: %w", got.id, err)
			}
		}

		return nil
	}
}

func init() {
	resource.AddTestSweepers("netbox_available_prefix", &resource.Sweeper{
		Name:         "netbox_available_prefix",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamPrefixesListParams()
			res, err := api.Ipam.IpamPrefixesList(params, nil)
			if err != nil {
				return err
			}
			for _, prefix := range res.GetPayload().Results {
				if len(prefix.Tags) > 0 && (prefix.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
					deleteParams := ipam.NewIpamPrefixesDeleteParams().WithID(prefix.ID)
					_, err := api.Ipam.IpamPrefixesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a prefix")
				}
			}
			return nil
		},
	})
}
