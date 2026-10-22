//go:build acctest

package provider_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccNetboxMACAddress_basic(t *testing.T) {
	testName := testAccGetTestName("mac")
	mac := fmt.Sprintf("02:42:%02x:%02x:%02x:%02x",
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255),
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%[2]s"
  description = "%[1]s"
  comments    = "Created by acceptance test."
}
resource "netbox_mac_address" "upper" {
  # Uppercase (NetBox's normal form): safe to import-verify. The lowercase
  # one above proves the configured casing survives create/read/update.
  mac_address = "%[3]s"
  description = "%[1]s"
}
data "netbox_mac_address" "test" {
  id = netbox_mac_address.test.id
}`, testName, mac, strings.ToUpper(mac)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", mac),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "description", testName),
					resource.TestCheckResourceAttrPair("data.netbox_mac_address.test", "id", "netbox_mac_address.test", "id"),
				),
			},
			{
				// Drop description and comments: both must clear.
				Config: fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
}
resource "netbox_mac_address" "upper" {
  mac_address = "%s"
  description = "%s"
}`, mac, strings.ToUpper(mac), testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_mac_address.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_mac_address.test", "comments"),
				),
			},
			{
				// Import reads NetBox's uppercase form with no prior state to take the casing from, so only the
				// uppercase resource is import-verified.
				ResourceName:      "netbox_mac_address.upper",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccNetboxMACAddress_assignment: the assigned_object hook on MAC addresses (the full matrix is
// covered on netbox_ip_address).
func TestAccNetboxMACAddress_assignment(t *testing.T) {
	testName := testAccGetTestName("mac_assign")
	mac := fmt.Sprintf("02:42:%02X:%02X:%02X:%02X",
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255),
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255))
	deps := fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
}
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address     = "%s"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}
data "netbox_mac_addresses" "by_id" {
  filters = [
    { name = "id", value = netbox_mac_address.test.id },
  ]
}`, mac),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "assigned_object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_mac_address.test", "assigned_object_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_mac_address.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_mac_address.test", "device_interface_id"),
					resource.TestCheckResourceAttr("data.netbox_mac_addresses.by_id", "mac_addresses.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_mac_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// The same assignment written as the explicit pair is a no-op.
				Config: deps + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address          = "%s"
  assigned_object_type = "virtualization.vminterface"
  assigned_object_id   = netbox_virtual_machine_interface.test.id
}`, mac),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.TestCheckResourceAttrPair("netbox_mac_address.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
}`, mac),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_mac_address.test", "assigned_object_type"),
					resource.TestCheckNoResourceAttr("netbox_mac_address.test", "assigned_object_id"),
					resource.TestCheckNoResourceAttr("netbox_mac_address.test", "virtual_machine_interface_id"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_mac_address",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimMacAddressesList(dcim.NewDcimMacAddressesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// MAC addresses have no name; tests put the test name in description.
				// 4.6.9 declares description non-nullable, so it is a value string here.
				items = append(items, sweepItem{result.ID, result.Description})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimMacAddressesDestroy(dcim.NewDcimMacAddressesDestroyParams().WithID(id), nil)
			return err
		})
}

// mac_address is normalize mac: a dash-separated spelling keeps that spelling in state although
// NetBox echoes it colon-separated and uppercased, the next plan is empty, and text that is no MAC
// address fails at plan.
func TestAccNetboxMACAddress_normalizedSpelling(t *testing.T) {
	spelled := fmt.Sprintf("02-42-%02x-%02x-%02x-%02x",
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255),
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255))
	config := fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
}`, spelled)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", spelled),
			},
			{
				// Not the last step: the post-test destroy runs the last step's config.
				Config: `
resource "netbox_mac_address" "test" {
  mac_address = "02:42:00"
}`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("(?i)invalid mac address"),
			},
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}
