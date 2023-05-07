package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxContactGroupDataSource_basic(t *testing.T) {

	testSlug := "cntctgrp_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact_group" "test" {
  name = "%[1]s"
}

data "netbox_contact_group" "test" {
  name = netbox_contact_group.test.name
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_contact_group.test", "id", "netbox_contact_group.test", "id"),
				),
			},
		},
	})
}
