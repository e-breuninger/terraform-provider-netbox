//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxRackReservation_basic(t *testing.T) {
	testName := testAccGetTestName("rackres")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_rack" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// The docker test setup's admin user has id 1.
				Config: deps + fmt.Sprintf(`
resource "netbox_rack_reservation" "test" {
  rack_id     = netbox_rack.test.id
  units       = [1, 2, 3]
  user_id     = 1
  tenant_id   = netbox_tenant.test.id
  description = "%[1]s"
  comments    = "Reserved by acceptance test."
  status      = "pending"
}
data "netbox_rack_reservation" "test" {
  id = netbox_rack_reservation.test.id
}
data "netbox_rack_reservations" "by_id" {
  filters = [
    { name = "id", value = netbox_rack_reservation.test.id },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "units.#", "3"),
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "user_id", "1"),
					resource.TestCheckResourceAttrPair("netbox_rack_reservation.test", "rack_id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "description", testName),
					resource.TestCheckResourceAttrPair("data.netbox_rack_reservation.test", "id", "netbox_rack_reservation.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rack_reservations.by_id", "rack_reservations.#", "1"),
				),
			},
			{
				// Shrink the reservation, drop tenant and comments.
				Config: deps + fmt.Sprintf(`
resource "netbox_rack_reservation" "test" {
  rack_id     = netbox_rack.test.id
  units       = [2]
  user_id     = 1
  description = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "units.#", "1"),
					resource.TestCheckNoResourceAttr("netbox_rack_reservation.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_rack_reservation.test", "comments"),
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "status", "pending"),
				),
			},
			{
				ResourceName:      "netbox_rack_reservation.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_rack_reservation",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimRackReservationsList(dcim.NewDcimRackReservationsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// Reservations have no name; tests put the test name in description.
				items = append(items, sweepItem{result.ID, deref(result.Description)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimRackReservationsDestroy(dcim.NewDcimRackReservationsDestroyParams().WithID(id), nil)
			return err
		})
}
