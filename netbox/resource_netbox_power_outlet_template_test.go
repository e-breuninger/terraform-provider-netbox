package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxPowerOutletTemplate_basic(t *testing.T) {
	testSlug := "power_outlet_template"
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

resource "netbox_power_outlet_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
	type = "iec-60320-c13"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "type", "iec-60320-c13"),
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_power_outlet_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPowerOutletTemplate_opts(t *testing.T) {
	testSlug := "power_outlet_template"
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
	name = "%[1]s inlet"
	device_type_id = netbox_device_type.test.id
	type = "iec-60320-c20"
}

resource "netbox_power_outlet_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	device_type_id = netbox_device_type.test.id
	type = "iec-60320-c13"
	power_port_id = netbox_power_port_template.test.id
	feed_leg = "A"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "type", "iec-60320-c13"),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "feed_leg", "A"),
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "power_port_id", "netbox_power_port_template.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "device_type_id", "netbox_device_type.test", "id"),
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

resource "netbox_power_outlet_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	module_type_id = netbox_module_type.test.id
	type = "iec-60320-c13"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "type", "iec-60320-c13"),
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "module_type_id", "netbox_module_type.test", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_power_outlet_template", &resource.Sweeper{
		Name:         "netbox_power_outlet_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimPowerOutletTemplatesListParams()
			res, err := api.Dcim.DcimPowerOutletTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, tmpl := range res.GetPayload().Results {
				if strings.HasPrefix(*tmpl.Name, testPrefix) {
					deleteParams := dcim.NewDcimPowerOutletTemplatesDeleteParams().WithID(tmpl.ID)
					_, err := api.Dcim.DcimPowerOutletTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a power outlet template")
				}
			}
			return nil
		},
	})
}
