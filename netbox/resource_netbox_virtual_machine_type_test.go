package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualMachineType_basic(t *testing.T) {
	testSlug := "vm_type_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_virtual_machine_type" "test" {
  name              = "%[1]s"
  slug              = "%[2]s"
  description       = "This is my virtual machine type"
  comments          = "Some comments"
  default_vcpus     = 2
  default_memory_mb = 4096
  tags              = [netbox_tag.test.name]
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "description", "This is my virtual machine type"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "comments", "Some comments"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "default_vcpus", "2"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "default_memory_mb", "4096"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "tags.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_virtual_machine_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualMachineType_defaultSlug(t *testing.T) {
	testSlug := "vm_type_default_slug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_virtual_machine_type" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "slug", testName),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_virtual_machine_type", &resource.Sweeper{
		Name:         "netbox_virtual_machine_type",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := virtualization.NewVirtualizationVirtualMachineTypesListParams()
			res, err := api.Virtualization.VirtualizationVirtualMachineTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, vmType := range res.GetPayload().Results {
				if strings.HasPrefix(*vmType.Name, testPrefix) {
					deleteParams := virtualization.NewVirtualizationVirtualMachineTypesDeleteParams().WithID(vmType.ID)
					_, err := api.Virtualization.VirtualizationVirtualMachineTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a virtual machine type")
				}
			}
			return nil
		},
	})
}
