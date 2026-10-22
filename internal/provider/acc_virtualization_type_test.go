//go:build acctest

// Virtual machine types, added by NetBox 4.6.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxVirtualMachineType_basic(t *testing.T) {
	testName := testAccGetTestName("vmtype")
	deps := fmt.Sprintf(`
resource "netbox_platform" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine_type" "test" {
  name                = "%[1]s"
  slug                = "%[2]s"
  default_platform_id = netbox_platform.test.id
  default_vcpus       = 4
  default_memory      = 8192
  description         = "Acceptance test virtual machine type."
  comments            = "Created by acceptance test."
}
data "netbox_virtual_machine_type" "test" {
  name = netbox_virtual_machine_type.test.name
}
data "netbox_virtual_machine_types" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_virtual_machine_type.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "default_vcpus", "4"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_type.test", "default_memory", "8192"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_type.test", "default_platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machine_type.test", "id", "netbox_virtual_machine_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_types.test", "virtual_machine_types.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine_type" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_type.test", "default_platform_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_type.test", "default_vcpus"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_type.test", "default_memory"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_type.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_type.test", "comments"),
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

func init() {
	sweep("netbox_virtual_machine_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationVirtualMachineTypesList(virtualization.NewVirtualizationVirtualMachineTypesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationVirtualMachineTypesDestroy(virtualization.NewVirtualizationVirtualMachineTypesDestroyParams().WithID(id), nil)
			return err
		})
}
