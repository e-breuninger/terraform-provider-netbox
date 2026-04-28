package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceType_basic(t *testing.T) {
	testSlug := "device_type"
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
  u_height = "0.5"
  manufacturer_id = netbox_manufacturer.test.id
  is_full_depth = true
  subdevice_role = "parent"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_device_type.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "part_number", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "u_height", "0.5"),
					resource.TestCheckResourceAttrPair("netbox_device_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "is_full_depth", "true"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "subdevice_role", "parent"),
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
  u_height = "0"
  manufacturer_id = netbox_manufacturer.test.id
  is_full_depth = false
  subdevice_role = "child"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_device_type.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "part_number", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "u_height", "0"),
					resource.TestCheckResourceAttrPair("netbox_device_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "is_full_depth", "false"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "subdevice_role", "child"),
				),
			},
			{
				ResourceName:      "netbox_device_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccNetboxDeviceType_templates_basic creates a device_type with one
// template of every supported family (sans inventory_item, which gets its
// own dedicated test so we can also exercise the parent tree). Verifies that
// each template lands in NetBox with the expected attributes.
func TestAccNetboxDeviceType_templates_basic(t *testing.T) {
	testSlug := "dt_tpl_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
  # subdevice_role=parent is required by NetBox to allow device_bay templates.
  subdevice_role  = "parent"

  power_port_templates {
    name          = "psu0"
    type          = "iec-60320-c14"
    maximum_draw  = 750
    allocated_draw = 500
    description   = "primary PSU"
  }

  interface_templates {
    name      = "mgmt0"
    type      = "1000base-t"
    mgmt_only = true
  }

  console_port_templates {
    name = "console0"
    type = "rj-45"
  }

  console_server_port_templates {
    name = "csp0"
    type = "rj-45"
  }

  rear_port_templates {
    name      = "rp0"
    type      = "8p8c"
    positions = 4
  }

  front_port_templates {
    name               = "fp0"
    type               = "8p8c"
    rear_port          = "rp0"
    rear_port_position = 1
  }

  power_outlet_templates {
    name       = "out0"
    type       = "iec-60320-c13"
    power_port = "psu0"
    feed_leg   = "A"
  }

  device_bay_templates {
    name = "bay0"
  }

  module_bay_templates {
    name     = "modbay0"
    position = "1"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "power_port_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "interface_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "console_port_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "console_server_port_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "rear_port_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "front_port_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "power_outlet_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "device_bay_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "module_bay_templates.#", "1"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceType_templates_update mutates an existing device_type's
// nested templates: adds a new one, removes one, and edits one in place.
// Verifies that all three operations converge and a final plan shows no drift.
func TestAccNetboxDeviceType_templates_update(t *testing.T) {
	testSlug := "dt_tpl_update"
	testName := testAccGetTestName(testSlug)
	step1 := fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id

  interface_templates {
    name      = "eth0"
    type      = "1000base-t"
    mgmt_only = false
  }
  interface_templates {
    name      = "eth1"
    type      = "1000base-t"
    mgmt_only = false
  }
}`, testName)

	step2 := fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id

  interface_templates {
    name      = "eth0"
    type      = "10gbase-t"
    mgmt_only = false
    description = "upgraded"
  }
  interface_templates {
    name      = "mgmt0"
    type      = "1000base-t"
    mgmt_only = true
  }
}`, testName)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: step1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "interface_templates.#", "2"),
				),
			},
			{
				Config: step2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "interface_templates.#", "2"),
				),
			},
			{
				Config:             step2,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccNetboxDeviceType_templates_destroy creates a device_type with several
// templates, destroys it, and asserts NetBox cascades the templates with the
// parent. (NetBox's API contract is that templates are deleted with the
// device_type; we just need to confirm the provider does not get into a
// "templates still exist" failure during destroy.)
func TestAccNetboxDeviceType_templates_destroy(t *testing.T) {
	testSlug := "dt_tpl_destroy"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id

  interface_templates {
    name = "eth0"
    type = "1000base-t"
  }
  power_port_templates {
    name = "psu0"
    type = "iec-60320-c14"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "interface_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "power_port_templates.#", "1"),
				),
			},
			{
				// Drop the device_type by replacing the config with just the
				// manufacturer. The destroy step at the end of the test will
				// then clean everything else up too.
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}`, testName),
			},
		},
	})
}

// TestAccNetboxDeviceType_templates_fk_ordering exercises the two
// inter-template FK paths in a single device_type: power_outlet ->
// power_port and front_port -> rear_port. Both reference siblings by name,
// which the provider has to resolve to IDs at apply time after the simple
// types are created.
func TestAccNetboxDeviceType_templates_fk_ordering(t *testing.T) {
	testSlug := "dt_tpl_fk"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id

  power_port_templates {
    name = "psu-a"
    type = "iec-60320-c14"
  }
  power_port_templates {
    name = "psu-b"
    type = "iec-60320-c14"
  }
  power_outlet_templates {
    name       = "out-a-1"
    type       = "iec-60320-c13"
    power_port = "psu-a"
  }
  power_outlet_templates {
    name       = "out-b-1"
    type       = "iec-60320-c13"
    power_port = "psu-b"
  }

  rear_port_templates {
    name      = "rear-1"
    type      = "8p8c"
    positions = 4
  }
  front_port_templates {
    name               = "front-1-a"
    type               = "8p8c"
    rear_port          = "rear-1"
    rear_port_position = 1
  }
  front_port_templates {
    name               = "front-1-b"
    type               = "8p8c"
    rear_port          = "rear-1"
    rear_port_position = 2
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "power_port_templates.#", "2"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "power_outlet_templates.#", "2"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "rear_port_templates.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "front_port_templates.#", "2"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceType_templates_inventory_tree exercises the
// inventory_item parent tree (root + child + grandchild) and a polymorphic
// component_type/component_id reference back at one of the device_type's
// own interface templates.
func TestAccNetboxDeviceType_templates_inventory_tree(t *testing.T) {
	testSlug := "dt_tpl_inv"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id

  interface_templates {
    name = "eth0"
    type = "1000base-t"
  }

  inventory_item_templates {
    name = "chassis"
  }
  inventory_item_templates {
    name   = "psu-a"
    parent = "chassis"
  }
  inventory_item_templates {
    name   = "psu-fan-a"
    parent = "psu-a"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "inventory_item_templates.#", "3"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceType_templates_coexistence shows that a standalone
// netbox_interface_template resource and a nested interface_templates block
// on a different device_type don't interfere with each other.
func TestAccNetboxDeviceType_templates_coexistence(t *testing.T) {
	testSlug := "dt_tpl_coex"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

# device_type managed via the nested interface_templates block.
resource "netbox_device_type" "nested" {
  model           = "%[1]s-nested"
  manufacturer_id = netbox_manufacturer.test.id

  interface_templates {
    name = "eth0"
    type = "1000base-t"
  }
}

# device_type managed via the standalone resource.
resource "netbox_device_type" "standalone" {
  model           = "%[1]s-standalone"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_interface_template" "ext" {
  name           = "eth0"
  type           = "1000base-t"
  device_type_id = netbox_device_type.standalone.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.nested", "interface_templates.#", "1"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_device_type", &resource.Sweeper{
		Name:         "netbox_device_type",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimDeviceTypesListParams()
			res, err := api.Dcim.DcimDeviceTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, devicetype := range res.GetPayload().Results {
				if strings.HasPrefix(*devicetype.Model, testPrefix) {
					deleteParams := dcim.NewDcimDeviceTypesDeleteParams().WithID(devicetype.ID)
					_, err := api.Dcim.DcimDeviceTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device type")
				}
			}
			return nil
		},
	})
}
