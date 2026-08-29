package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxModuleTypeProfile_basic(t *testing.T) {
	testSlug := "module_type_profile"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_module_type_profile" "test" {
	name = "%[1]s"
	schema = jsonencode({
		type = "object"
		properties = {
			wattage = {
				type = "integer"
			}
		}
	})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "name", testName),
					resource.TestCheckResourceAttrSet("netbox_module_type_profile.test", "schema"),
				),
			},
			{
				ResourceName:      "netbox_module_type_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModuleTypeProfile_opts(t *testing.T) {
	testSlug := "module_type_profile"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_module_type_profile" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	comments = "%[1]s comments"
	schema = jsonencode({
		type = "object"
		properties = {
			cores = {
				type = "integer"
			}
		}
	})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "comments", fmt.Sprintf("%s comments", testName)),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_module_type_profile" "test" {
	name = "%[1]s"
	description = "%[1]s new description"
	comments = "%[1]s new comments"
	schema = jsonencode({
		type = "object"
		properties = {
			cores = {
				type = "integer"
			}
			speed_mhz = {
				type = "integer"
			}
		}
	})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "description", fmt.Sprintf("%s new description", testName)),
					resource.TestCheckResourceAttr("netbox_module_type_profile.test", "comments", fmt.Sprintf("%s new comments", testName)),
				),
			},
		},
	})
}

func TestAccNetboxModuleType_profileAndAttributes(t *testing.T) {
	testSlug := "module_type_w_profile"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_module_type_profile" "test" {
	name = "%[1]s"
	schema = jsonencode({
		type = "object"
		properties = {
			wattage = {
				type = "integer"
			}
		}
	})
}

resource "netbox_module_type" "test" {
	manufacturer_id = netbox_manufacturer.test.id
	model           = "%[1]s"
	profile_id      = netbox_module_type_profile.test.id
	attributes = jsonencode({
		wattage = 715
	})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_module_type.test", "profile_id", "netbox_module_type_profile.test", "id"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "attributes", `{"wattage":715}`),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_module_type_profile" "test" {
	name = "%[1]s"
	schema = jsonencode({
		type = "object"
		properties = {
			wattage = {
				type = "integer"
			}
		}
	})
}

resource "netbox_module_type" "test" {
	manufacturer_id = netbox_manufacturer.test.id
	model           = "%[1]s"
	profile_id      = netbox_module_type_profile.test.id
	attributes = jsonencode({
		wattage = 350
	})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "attributes", `{"wattage":350}`),
				),
			},
			{
				ResourceName:      "netbox_module_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModuleType_profileWithoutAttributes(t *testing.T) {
	testSlug := "module_type_profile_only"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_module_type_profile" "test" {
	name = "%[1]s"
	schema = jsonencode({
		type = "object"
		properties = {
			wattage = {
				type = "integer"
			}
		}
	})
}

resource "netbox_module_type" "test" {
	manufacturer_id = netbox_manufacturer.test.id
	model           = "%[1]s"
	profile_id      = netbox_module_type_profile.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_module_type.test", "profile_id", "netbox_module_type_profile.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_module_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_module_type_profile", &resource.Sweeper{
		Name:         "netbox_module_type_profile",
		Dependencies: []string{"netbox_module_type"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimModuleTypeProfilesListParams()
			res, err := api.Dcim.DcimModuleTypeProfilesList(params, nil)
			if err != nil {
				return err
			}
			for _, profile := range res.GetPayload().Results {
				if strings.HasPrefix(*profile.Name, testPrefix) {
					deleteParams := dcim.NewDcimModuleTypeProfilesDeleteParams().WithID(profile.ID)
					_, err := api.Dcim.DcimModuleTypeProfilesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a module type profile")
				}
			}
			return nil
		},
	})
}
