package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxFrontPortTemplate_basic(t *testing.T) {
	testSlug := "front_port_template"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_rear_port_template" "test" {
	name = "%[1]s rear"
	device_type_id = netbox_device_type.test.id
	type = "8p8c"
}

resource "netbox_front_port_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
	type = "8p8c"
	rear_port_id = netbox_rear_port_template.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "type", "8p8c"),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_port_position", "1"),
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "rear_port_id", "netbox_rear_port_template.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_front_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxFrontPortTemplate_opts(t *testing.T) {
	testSlug := "front_port_template"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_rear_port_template" "test" {
	name = "%[1]s rear"
	device_type_id = netbox_device_type.test.id
	type = "mpo"
	positions = 12
}

resource "netbox_front_port_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	device_type_id = netbox_device_type.test.id
	type = "lc-upc"
	rear_port_id = netbox_rear_port_template.test.id
	rear_port_position = 3
	color_hex = "f44336"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "type", "lc-upc"),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_port_position", "3"),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "color_hex", "f44336"),
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "rear_port_id", "netbox_rear_port_template.test", "id"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_rear_port_template" "test" {
	name = "%[1]s rear"
	device_type_id = netbox_device_type.test.id
	type = "mpo"
	positions = 12
}

resource "netbox_front_port_template" "test" {
	name = "%[1]s"
	description = "%[1]s new description"
	label = "%[1]s label"
	device_type_id = netbox_device_type.test.id
	type = "lc-upc"
	rear_port_id = netbox_rear_port_template.test.id
	rear_port_position = 7
	color_hex = "f44336"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "description", fmt.Sprintf("%s new description", testName)),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_port_position", "7"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_module_type" "test" {
	manufacturer_id = netbox_manufacturer.test.id
	model           = "%[1]s"
}

resource "netbox_rear_port_template" "test" {
	name = "%[1]s rear"
	module_type_id = netbox_module_type.test.id
	type = "mpo"
	positions = 12
}

resource "netbox_front_port_template" "test" {
	name = "%[1]s"
	module_type_id = netbox_module_type.test.id
	type = "lc-upc"
	rear_port_id = netbox_rear_port_template.test.id
	rear_port_position = 5
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_port_position", "5"),
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "rear_port_id", "netbox_rear_port_template.test", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_front_port_template", &resource.Sweeper{
		Name:         "netbox_front_port_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimFrontPortTemplatesListParams()
			res, err := api.Dcim.DcimFrontPortTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, tmpl := range res.GetPayload().Results {
				if strings.HasPrefix(*tmpl.Name, testPrefix) {
					deleteParams := dcim.NewDcimFrontPortTemplatesDeleteParams().WithID(tmpl.ID)
					_, err := api.Dcim.DcimFrontPortTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a front port template")
				}
			}
			return nil
		},
	})
}
