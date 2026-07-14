package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVlansSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_vlan" "test_1" {
  name = "VLAN1234"
  vid  = 1234
  tags = [netbox_tag.test.name]
}

resource "netbox_vlan" "test_2" {
  name = "VLAN1235"
  vid  = 1235
  tags = [netbox_tag.test.name]
}

resource "netbox_vlan" "test_3" {
  name = "VLAN1236"
  vid  = 1236
  tags = [netbox_tag.test.name]
}`, testName)
}

func testAccNetboxVlansByVid() string {
	return `
data "netbox_vlans" "test" {
  filter {
	name  = "vid"
	value = "1234"
  }

  filter {
	name  = "tag"
	value = netbox_tag.test.slug
  }
}`
}

func testAccNetboxVlansByVidN() string {
	return `
data "netbox_vlans" "test" {
  filter {
	name = "vid__n"
	value = "1234"
  }

  filter {
	name  = "tag"
	value = netbox_tag.test.slug
  }
}`
}

func testAccNetboxVlansByVidRange() string {
	return `
data "netbox_vlans" "test" {
  filter {
	name = "vid__gte"
	value = "1235"
  }

  filter {
	name = "vid__lte"
	value = "1236"
  }

  filter {
	name  = "tag"
	value = netbox_tag.test.slug
  }
}`
}

func TestAccNetboxVlansDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("vlans_ds_basic")
	setUp := testAccNetboxVlansSetUp(testName)
	// The VLANs are tagged with a unique per-test tag and every query filters on
	// that tag, so concurrently-created VLANs from other tests no longer affect
	// the vid__n / range counts.
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test_1", "vid", "1234"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVid(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.id", "netbox_vlan.test_1", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVid(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.vid", "netbox_vlan.test_1", "vid"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVidN(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.vid", "netbox_vlan.test_2", "vid"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVidRange(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.vid", "netbox_vlan.test_2", "vid"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.1.vid", "netbox_vlan.test_3", "vid"),
				),
			},
		},
	})
}

func TestAccNetboxVlansDataSource_customFields(t *testing.T) {
	testName := testAccGetTestName("vlans_ds_custom_fields")
	testField := fmt.Sprintf("vlans_ds_cf_%s", acctest.RandStringFromCharSet(10, "abcdefghijklmnopqrstuvwxyz"))
	testVidMatch := acctest.RandIntRange(1000, 2000)
	testVidSkip := acctest.RandIntRange(2001, 4000)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlansWithCustomFields(testName, testField, testVidMatch, testVidSkip),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.by_vid", "vlans.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.by_vid", "vlans.0.id", "netbox_vlan.test_match", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlans.by_vid", fmt.Sprintf("vlans.0.custom_fields.%s", testField), "match"),
					resource.TestCheckResourceAttr("data.netbox_vlans.by_custom_fields", "vlans.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.by_custom_fields", "vlans.0.id", "netbox_vlan.test_match", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlans.by_custom_fields", fmt.Sprintf("vlans.0.custom_fields.%s", testField), "match"),
				),
			},
		},
	})
}

func testAccNetboxVlansWithCustomFields(testName string, testField string, testVidMatch int, testVidSkip int) string {
	return fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name          = "%[2]s"
	type          = "text"
	content_types = ["ipam.vlan"]
}

resource "netbox_vlan" "test_match" {
	name = "%[1]s-match"
	vid  = %[3]d
	tags = []

	custom_fields = {
		(netbox_custom_field.test.name) = "match"
	}
}

resource "netbox_vlan" "test_skip" {
	name = "%[1]s-skip"
	vid  = %[4]d
	tags = []

	custom_fields = {
		(netbox_custom_field.test.name) = "skip"
	}
}

data "netbox_vlans" "by_vid" {
	depends_on = [netbox_vlan.test_match, netbox_vlan.test_skip]

	filter {
		name  = "vid"
		value = netbox_vlan.test_match.vid
	}
}

data "netbox_vlans" "by_custom_fields" {
	depends_on = [netbox_vlan.test_match, netbox_vlan.test_skip]

	custom_fields = {
		(netbox_custom_field.test.name) = "match"
	}
}
`, testName, testField, testVidMatch, testVidSkip)
}
