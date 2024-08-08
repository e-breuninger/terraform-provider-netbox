package netbox

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxAvailableInterfaceFullDependencies(testName string) string {
	return fmt.Sprintf(`
	resource "netbox_manufacturer" "test" {
		name = "%[1]s"
	}

	resource "netbox_device_type" "test" {
		model = "%[1]s"
		manufacturer_id = netbox_manufacturer.test.id
	}

	resource "netbox_device_role" "test" {
		name = "%[1]s"
		color_hex = "123456"
	}

	resource "netbox_site" "test" {
		name = "%[1]s"
		status = "active"
	}

	resource "netbox_device" "test" {
		name = "%[1]s"
		device_type_id = netbox_device_type.test.id
		role_id = netbox_device_role.test.id
		site_id = netbox_site.test.id
	}`, testName)
}

func testAccNetboxAvailableInterfaceMultipleInterfaces(names ...string) string {
	var config strings.Builder

	for _, name := range names {
		_, _ = fmt.Fprintf(&config, `
resource "netbox_device_interface" "%[1]s" {
	name      = "%[1]s"
	device_id = netbox_device.test.id
	type      = "1000base-t"
}`, name)
	}

	return config.String()
}

func TestAccNetboxAvailableInterface_basic(t *testing.T) {
	testSlug := "available_interface_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`%s

resource "netbox_available_interface" "test" {
	device_id = netbox_device.test.id
	prefix    = "tun"
	type      = "1000base-t"
}`, testAccNetboxAvailableInterfaceFullDependencies(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_interface.test", "name", "tun0"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableInterface_gap(t *testing.T) {
	testSlug := "available_interface_gap"
	testName := testAccGetTestName(testSlug)

	initial :=
		testAccNetboxAvailableInterfaceFullDependencies(testName) +
			testAccNetboxAvailableInterfaceMultipleInterfaces("tun0", "tun1", "tun3", "tun10")

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: initial,
			},
			{
				Config: initial + `

resource "netbox_available_interface" "test" {
	device_id = netbox_device.test.id
	prefix    = "tun"
	type      = "1000base-t"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_interface.test", "name", "tun2"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableInterface_last(t *testing.T) {
	testSlug := "available_interface_last"
	testName := testAccGetTestName(testSlug)

	initial :=
		testAccNetboxAvailableInterfaceFullDependencies(testName) +
			testAccNetboxAvailableInterfaceMultipleInterfaces(
				"tun0", "tun1", "tun2", "tun3", "tun4", "tun5",
				"tun6", "tun7", "tun8", "tun9", "tun10", "tun11",
			)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: initial,
			},
			{
				Config: initial + `

resource "netbox_available_interface" "test" {
	device_id = netbox_device.test.id
	prefix    = "tun"
	type      = "1000base-t"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_interface.test", "name", "tun12"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableInterface_prefixChange(t *testing.T) {
	testSlug := "available_interface_prefixChange"
	testName := testAccGetTestName(testSlug)

	initial :=
		testAccNetboxAvailableInterfaceFullDependencies(testName) +
			testAccNetboxAvailableInterfaceMultipleInterfaces("tun0", "tun1")

	var resourceID string

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: initial,
			},
			{
				Config: initial + `

resource "netbox_available_interface" "test" {
	device_id = netbox_device.test.id
	prefix    = "tun"
	type      = "1000base-t"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_interface.test", "name", "tun2"),
					func(s *terraform.State) error {
						// record resource ID
						resourceID = s.RootModule().Resources["netbox_available_interface.test"].Primary.Attributes["id"]
						return nil
					},
				),
			},
			{
				Config: initial + `

resource "netbox_available_interface" "test" {
	device_id = netbox_device.test.id
	prefix    = "dummy"
	type      = "1000base-t"
}`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						newID := s.RootModule().Resources["netbox_available_interface.test"].Primary.Attributes["id"]
						if newID == resourceID {
							return errors.New("resource has not been recreated")
						}
						return nil
					},
					resource.TestCheckResourceAttr("netbox_available_interface.test", "name", "dummy0"),
				),
			},
		},
	})
}
