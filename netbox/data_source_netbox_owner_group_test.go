package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxOwnerGroupDataSource_basic(t *testing.T) {
	testSlug := "ownergrp_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_owner_group" "test" {
  name = "%[1]s"
}

data "netbox_owner_group" "test" {
  depends_on = [netbox_owner_group.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_owner_group.test", "id", "netbox_owner_group.test", "id"),
				),
			},
		},
	})
}
