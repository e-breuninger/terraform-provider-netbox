package netbox

import (
	"fmt"
	"regexp"
	"testing"

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
