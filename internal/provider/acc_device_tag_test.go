//go:build acctest

// The netbox_device_tag companion: it owns the presence of a single tag on a device without
// owning the device, so the checks read the device straight from the API — the netbox_device
// resource in the config carries lifecycle { ignore_changes = [tags] } and its state says
// nothing about the tags (without it, its updates would strip every companion-attached tag).
package provider_test

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkDeviceTagSlugs compares the device's tag slugs in NetBox against want, order-insensitive.
func checkDeviceTagSlugs(deviceName string, want ...string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		var id int64
		if err := stateID(deviceName, &id)(state); err != nil {
			return err
		}
		client, err := testAccClient()
		if err != nil {
			return err
		}
		res, err := client.Dcim.DcimDevicesRetrieve(dcim.NewDcimDevicesRetrieveParams().WithID(id), nil)
		if err != nil {
			return err
		}
		got := []string{}
		for _, tag := range res.Payload.Tags {
			if tag.Slug != nil {
				got = append(got, *tag.Slug)
			}
		}
		sort.Strings(got)
		wantSorted := append([]string{}, want...)
		sort.Strings(wantSorted)
		if strings.Join(got, ",") != strings.Join(wantSorted, ",") {
			return fmt.Errorf("device %d has tags [%s], want [%s]",
				id, strings.Join(got, ", "), strings.Join(wantSorted, ", "))
		}
		return nil
	}
}

// TestAccNetboxDeviceTag_basic: two tags attach concurrently (exercising the per-device patch
// lock), one detaches without touching the other, the last detach leaves the device untagged, and
// every step's plan converges thanks to the device's ignore_changes = [tags].
// netbox_device_tag has no shrink step: it is a pure link resource with no optional attributes.
func TestAccNetboxDeviceTag_basic(t *testing.T) {
	testName := testAccGetTestName("devtag")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id

  lifecycle {
    ignore_changes = [tags]
  }
}
resource "netbox_tag" "a" {
  name = "%[1]s-a"
}
resource "netbox_tag" "b" {
  name = "%[1]s-b"
}`, testName)
	tagA := `
resource "netbox_device_tag" "a" {
  device_id = netbox_device.test.id
  tag_slug  = netbox_tag.a.slug
}`
	tagB := `
resource "netbox_device_tag" "b" {
  device_id = netbox_device.test.id
  tag_slug  = netbox_tag.b.slug
}`
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + tagA + tagB + `
data "netbox_device" "test" {
  id         = netbox_device.test.id
  depends_on = [netbox_device_tag.a, netbox_device_tag.b]
}
data "netbox_devices" "by_tags" {
  filters = [
    { name = "tag", value = netbox_tag.a.slug },
    { name = "tag", value = netbox_tag.b.slug },
  ]
  depends_on = [netbox_device_tag.a, netbox_device_tag.b]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_tag.a", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_tag.a", "tag_slug", "netbox_tag.a", "slug"),
					resource.TestCheckTypeSetElemAttrPair("data.netbox_device.test", "tags.*", "netbox_tag.a", "slug"),
					resource.TestCheckTypeSetElemAttrPair("data.netbox_device.test", "tags.*", "netbox_tag.b", "slug"),
					// Repeated tag filters are ANDed by NetBox: both tags are this test's alone,
					// so the two-tag lookup finds exactly our device.
					resource.TestCheckResourceAttr("data.netbox_devices.by_tags", "devices.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.by_tags", "devices.0.id", "netbox_device.test", "id"),
					checkDeviceTagSlugs("netbox_device.test", getSlug(testName+"-a"), getSlug(testName+"-b")),
				),
			},
			{
				ResourceName:      "netbox_device_tag.a",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Detaching one tag leaves the other alone.
				Config: deps + tagA,
				Check:  checkDeviceTagSlugs("netbox_device.test", getSlug(testName+"-a")),
			},
			{
				// With tag b detached (previous step — a data source in that step could be read
				// before the detach applies), the ANDed two-tag filter no longer matches; an OR
				// would still find the device by its remaining tag.
				Config: deps + tagA + `
data "netbox_devices" "by_tags" {
  filters = [
    { name = "tag", value = netbox_tag.a.slug },
    { name = "tag", value = netbox_tag.b.slug },
  ]
  depends_on = [netbox_device_tag.a]
}`,
				Check: resource.ComposeTestCheckFunc(
					checkDeviceTagSlugs("netbox_device.test", getSlug(testName+"-a")),
					resource.TestCheckResourceAttr("data.netbox_devices.by_tags", "devices.#", "0"),
				),
			},
			{
				Config: deps,
				Check:  checkDeviceTagSlugs("netbox_device.test"),
			},
		},
	})
}

// TestAccNetboxDeviceTag_missingTag: the tag must already exist.
func TestAccNetboxDeviceTag_missingTag(t *testing.T) {
	testName := testAccGetTestName("devtag-no-tag")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_tag" "test" {
  device_id = netbox_device.test.id
  tag_slug  = "%s-does-not-exist"
}`, testName),
				ExpectError: regexp.MustCompile("no tag with slug"),
			},
		},
	})
}
