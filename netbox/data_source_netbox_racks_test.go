package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRacksDataSource_basic(t *testing.T) {

	testRacks := []string{"rack1", "rack2", "rack3"}
	testSlug := "racks_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_rack" "test_rack1" {
  name     = "%[2]s"
  site_id  = netbox_site.test.id
  status   = "active"
  width    = 10
  u_height = 40
}

resource "netbox_rack" "test_rack2" {
  name     = "%[3]s"
  site_id  = netbox_site.test.id
  status   = "active"
  width    = 19
  u_height = 41
}

resource "netbox_rack" "test_rack3" {
  name     = "%[4]s"
  site_id  = netbox_site.test.id
  status   = "reserved"
  width    = 21
  u_height = 42
}

data "netbox_racks" "by_name" {
  depends_on = [netbox_rack.test_rack1, netbox_rack.test_rack2, netbox_rack.test_rack3]
  filter {
    name  = "name"
    value = netbox_rack.test_rack3.name
  }
  filter {
    name = "site_id"
    value = netbox_site.test.id
  }
}

data "netbox_racks" "by_status" {
  depends_on = [netbox_rack.test_rack1, netbox_rack.test_rack2, netbox_rack.test_rack3]
  filter {
    name  = "status"
    value = "active"
  }
  filter {
    name = "site_id"
    value = netbox_site.test.id
  }
}
`, testName, testRacks[0], testRacks[1], testRacks[2]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_racks.by_name", "racks.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_racks.by_status", "racks.#", "2"),
				),
			},
		},
	})
}
