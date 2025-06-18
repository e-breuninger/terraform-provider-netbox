package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualDiskDataSource_basic(t *testing.T) {
	testSlug := "virtual_disk_ds_basic"
	testName := testAccGetTestName(testSlug)
	resourceName := "netbox_virtual_disk.test"
	dataSourceName := "data.netbox_virtual_disk.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "tag_a" {
  name = "[%[1]s_a]"
  color_hex = "123456"
}
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}
resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_virtual_disk" "test" {
  name = "%[1]s"
  description = "description"
  size_mb = 30
  virtual_machine_id = netbox_virtual_machine.test.id
  tags = [netbox_tag.tag_a.name]
}
data "netbox_virtual_disk" "test" {
  id = netbox_virtual_disk.test.id
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "size_mb", resourceName, "size_mb"),
					resource.TestCheckResourceAttrPair(dataSourceName, "virtual_machine_id", resourceName, "virtual_machine_id"),
				),
			},
		},
	})
}

func TestAccNetboxVirtualDiskDataSource_filter_and_list(t *testing.T) {
	testSlug := "virtual_disk_ds_filter"
	testName := testAccGetTestName(testSlug)
	dataSourceName := "data.netbox_virtual_disk.filtered"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "tag_a" {
  name = "[%[1]s_a]"
  color_hex = "123456"
}
resource "netbox_tag" "tag_b" {
  name = "[%[1]s_b]"
  color_hex = "654321"
}
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}
resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_virtual_disk" "disk_a" {
  name = "%[1]s_disk_a"
  description = "disk a desc"
  size_mb = 10
  virtual_machine_id = netbox_virtual_machine.test.id
  tags = [netbox_tag.tag_a.name]
}
resource "netbox_virtual_disk" "disk_b" {
  name = "%[1]s_disk_b"
  description = "disk b desc"
  size_mb = 20
  virtual_machine_id = netbox_virtual_machine.test.id
  tags = [netbox_tag.tag_b.name]
}
data "netbox_virtual_disk" "filtered" {
  filter = [
    { name = "name" value = "%[1]s_disk_a" },
    { name = "tag" value = netbox_tag.tag_a.name },
  ]
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.0.name", testName+"_disk_a"),
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.0.description", "disk a desc"),
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.0.size_mb", "10"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "tag_a" {
  name = "[%[1]s_a]"
  color_hex = "123456"
}
resource "netbox_tag" "tag_b" {
  name = "[%[1]s_b]"
  color_hex = "654321"
}
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}
resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_virtual_disk" "disk_a" {
  name = "%[1]s_disk_a"
  description = "disk a desc"
  size_mb = 10
  virtual_machine_id = netbox_virtual_machine.test.id
  tags = [netbox_tag.tag_a.name]
}
resource "netbox_virtual_disk" "disk_b" {
  name = "%[1]s_disk_b"
  description = "disk b desc"
  size_mb = 20
  virtual_machine_id = netbox_virtual_machine.test.id
  tags = [netbox_tag.tag_b.name]
}
data "netbox_virtual_disk" "filtered" {
  filter = [
    { name = "tag" value = netbox_tag.tag_b.name },
  ]
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.0.name", testName+"_disk_b"),
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.0.description", "disk b desc"),
					resource.TestCheckResourceAttr(dataSourceName, "virtual_disks.0.size_mb", "20"),
				),
			},
		},
	})
}
