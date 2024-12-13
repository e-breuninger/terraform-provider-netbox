package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxServiceFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
}

`, testName)
}

func TestAccNetboxService_basic(t *testing.T) {
	testSlug := "svc_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxServiceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_service" "test" {
  name = "%s"
  virtual_machine_id = netbox_virtual_machine.test.id
  ports = [666]
  protocol = "tcp"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_service.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.#", "1"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.0", "666"),
					resource.TestCheckResourceAttr("netbox_service.test", "protocol", "tcp"),
				),
			},
			{
				ResourceName:      "netbox_service.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxService_customFields(t *testing.T) {
	testSlug := "svc_custom_fields"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxServiceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name          = "custom_field"
  type          = "text"
  content_types = ["ipam.service"]
}
resource "netbox_service" "test_customfield" {
  name = "%s"
  virtual_machine_id = netbox_virtual_machine.test.id
  ports = [333]
  protocol = "tcp"
  custom_fields = {"${netbox_custom_field.test.name}" = "testtext"}
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test_customfield", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_service.test_customfield", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("netbox_service.test_customfield", "ports.#", "1"),
					resource.TestCheckResourceAttr("netbox_service.test_customfield", "ports.0", "333"),
					resource.TestCheckResourceAttr("netbox_service.test_customfield", "protocol", "tcp"),
					resource.TestCheckResourceAttr("netbox_service.test_customfield", "custom_fields.custom_field", "testtext"),
				),
			},
			{
				ResourceName:      "netbox_service.test_customfield",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckServiceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each service
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_service" {
			continue
		}

		// Retrieve our service by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := ipam.NewIpamServicesReadParams().WithID(stateID)
		_, err := conn.Ipam.IpamServicesRead(params, nil)

		if err == nil {
			return fmt.Errorf("service (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*ipam.IpamServicesReadDefault); ok {
				errorcode := errresp.Code()
				if errorcode == 404 {
					return nil
				}
			}
			return err
		}
	}
	return nil
}

func TestAccNetboxService_withDescriptionDeviceID(t *testing.T) {
	testSlug := "svc_with_desc_tags_device"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_service" "test" {
  name = "%s"
  device_id = netbox_device.test_device.id
  ports = [666]
  protocol = "tcp"
  description = "Test service description"
}
  resource "netbox_site" "test_site" {
  name = "%[1]s_site"
  slug = "%[1]s_site"
}

resource "netbox_device_role" "test_role" {
  name = "%[1]s_role"
  slug = "%[1]s_role"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test_manufacturer" {
  name = "%[1]s_manufacturer"
}

resource "netbox_device_type" "test_type" {
  model = "%[1]s_type"
  manufacturer_id = netbox_manufacturer.test_manufacturer.id
}

resource "netbox_device" "test_device" {
  name = "%[1]s_device"
  role_id = netbox_device_role.test_role.id
  device_type_id = netbox_device_type.test_type.id
  site_id = netbox_site.test_site.id
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_service.test", "device_id", "netbox_device.test_device", "id"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.#", "1"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.0", "666"),
					resource.TestCheckResourceAttr("netbox_service.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("netbox_service.test", "description", "Test service description"),
				),
			},
			{
				ResourceName:      "netbox_service.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxService_withDescriptionTagsVirtualMachine(t *testing.T) {
	testSlug := "svc_with_desc_tags_device"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxServiceFullDependencies(testName) + fmt.Sprintf(
					`
					resource "netbox_tag" "tag1" {
						name = "tag1"
						slug = "tag1"
					}
					resource "netbox_tag" "tag2" {
						name = "tag2"
						slug = "tag2"
					}
					resource "netbox_service" "test" {
						name = "%s"
						virtual_machine_id = netbox_virtual_machine.test.id
						ports = [666]
						protocol = "tcp"
						description = "Test service description"
						tags = [netbox_tag.tag1.name, netbox_tag.tag2.name]
					}
				`,
					testName,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_service.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("netbox_service.test", "tags.0", "tag1"),
					resource.TestCheckResourceAttr("netbox_service.test", "tags.1", "tag2"),
					resource.TestCheckResourceAttrPair("netbox_service.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.#", "1"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.0", "666"),
					resource.TestCheckResourceAttr("netbox_service.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("netbox_service.test", "description", "Test service description"),
				),
			},
			{
				ResourceName:      "netbox_service.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_service", &resource.Sweeper{
		Name:         "netbox_service",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamServicesListParams()
			res, err := api.Ipam.IpamServicesList(params, nil)
			if err != nil {
				return err
			}
			for _, intrface := range res.GetPayload().Results {
				if strings.HasPrefix(*intrface.Name, testPrefix) {
					deleteParams := ipam.NewIpamServicesDeleteParams().WithID(intrface.ID)
					_, err := api.Ipam.IpamServicesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an interface")
				}
			}
			return nil
		},
	})
}
