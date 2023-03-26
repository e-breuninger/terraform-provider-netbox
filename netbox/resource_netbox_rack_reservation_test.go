package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxRackReservationFullDependencies(testName string) string {
	return fmt.Sprintf(`
	resource "netbox_site" "test" {
		name = "%[1]s"
		status = "active"
	}

	resource "netbox_tenant" "test" {
		name = "%[1]s"
	}

	resource "netbox_rack" "test" {
		name     = "%[1]s"
		site_id  = netbox_site.test.id
		status   = "active"
		width    = 10
		u_height = 40
	}`, testName)
}

func TestAccNetboxRackReservation_basic(t *testing.T) {

	testSlug := "rack_reservation_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxRackReservationFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_rack_reservation" "test" {
  rack_id = netbox_rack.test.id
	units = [1,2,3,4,5]
	user_id = 1
	description = "%[1]sdescription"
	tenant_id = netbox_tenant.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_rack_reservation.test", "rack_id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_rack_reservation.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "units.#", "5"),
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "user_id", "1"),
					resource.TestCheckResourceAttr("netbox_rack_reservation.test", "description", testName+"description"),
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
	resource.AddTestSweepers("netbox_rack_reservation", &resource.Sweeper{
		Name:         "netbox_rack_reservation",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimRackReservationsListParams()
			res, err := api.Dcim.DcimRackReservationsList(params, nil)
			if err != nil {
				return err
			}
			for _, rack_res := range res.GetPayload().Results {
				if strings.HasPrefix(*rack_res.Description, testPrefix) {
					deleteParams := dcim.NewDcimRackReservationsDeleteParams().WithID(rack_res.ID)
					_, err := api.Dcim.DcimRackReservationsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a rack_reservation")
				}
			}
			return nil
		},
	})
}
