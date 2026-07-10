package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxPowerPortTemplate_basic(t *testing.T) {
	testSlug := "power_port_template"
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

resource "netbox_power_port_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
	type = "iec-60320-c14"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "type", "iec-60320-c14"),
					resource.TestCheckResourceAttrPair("netbox_power_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_power_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPowerPortTemplate_opts(t *testing.T) {
	testSlug := "power_port_template"
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

resource "netbox_power_port_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	device_type_id = netbox_device_type.test.id
	type = "iec-60320-c16"
	maximum_draw = 715
	allocated_draw = 500
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "type", "iec-60320-c16"),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "maximum_draw", "715"),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "allocated_draw", "500"),
					resource.TestCheckResourceAttrPair("netbox_power_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
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

resource "netbox_power_port_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	module_type_id = netbox_module_type.test.id
	type = "iec-60320-c14"
	maximum_draw = 350
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "type", "iec-60320-c14"),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "maximum_draw", "350"),
					resource.TestCheckResourceAttrPair("netbox_power_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_power_port_template", &resource.Sweeper{
		Name:         "netbox_power_port_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimPowerPortTemplatesListParams()
			res, err := api.Dcim.DcimPowerPortTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, tmpl := range res.GetPayload().Results {
				if strings.HasPrefix(*tmpl.Name, testPrefix) {
					deleteParams := dcim.NewDcimPowerPortTemplatesDeleteParams().WithID(tmpl.ID)
					_, err := api.Dcim.DcimPowerPortTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a power port template")
				}
			}
			return nil
		},
	})
}
