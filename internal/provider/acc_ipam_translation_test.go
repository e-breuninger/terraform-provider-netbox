//go:build acctest

// The ipam objects NetBox 4.6 added: ASN ranges and VLAN translation policies with their rules.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxAsnRange_basic(t *testing.T) {
	testName := testAccGetTestName("asnrange")
	deps := fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_asn_range" "test" {
  name        = "%[1]s"
  rir_id      = netbox_rir.test.id
  start       = 64512
  end         = 64520
  tenant_id   = netbox_tenant.test.id
  description = "Acceptance test ASN range."
  comments    = "Created by acceptance test."
}
data "netbox_asn_range" "test" {
  name = netbox_asn_range.test.name
}
data "netbox_asn_ranges" "test" {
  filters = [
    { name = "rir_id", value = netbox_rir.test.id },
  ]
  depends_on = [netbox_asn_range.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_asn_range.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "start", "64512"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "end", "64520"),
					resource.TestCheckResourceAttrPair("data.netbox_asn_range.test", "id", "netbox_asn_range.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_asn_ranges.test", "asn_ranges.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_asn_range" "test" {
  name   = "%[1]s"
  rir_id = netbox_rir.test.id
  start  = 64512
  end    = 64520
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_asn_range.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_asn_range.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_asn_range.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_asn_range.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVlanTranslationPolicy_basic(t *testing.T) {
	testName := testAccGetTestName("vlanxlate")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vlan_translation_policy" "test" {
  name        = "%[1]s"
  description = "Acceptance test VLAN translation policy."
  comments    = "Created by acceptance test."
}
resource "netbox_vlan_translation_rule" "test" {
  vlan_translation_policy_id = netbox_vlan_translation_policy.test.id
  local_vid                  = 100
  remote_vid                 = 200
  description                = "Acceptance test VLAN translation rule."
}
data "netbox_vlan_translation_policy" "test" {
  name = netbox_vlan_translation_policy.test.name
}
data "netbox_vlan_translation_rules" "test" {
  filters = [
    { name = "policy_id", value = netbox_vlan_translation_policy.test.id },
  ]
  depends_on                 = [netbox_vlan_translation_rule.test]
}
data "netbox_vlan_translation_policies" "test" {
  depends_on = [netbox_vlan_translation_policy.test]
  filters = [
    { name = "name", value = "%[1]s" },
  ]
}
data "netbox_vlan_translation_rule" "test" {
  id = netbox_vlan_translation_rule.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_translation_policy.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan_translation_rule.test", "local_vid", "100"),
					resource.TestCheckResourceAttr("netbox_vlan_translation_rule.test", "remote_vid", "200"),
					resource.TestCheckResourceAttrPair("netbox_vlan_translation_rule.test", "vlan_translation_policy_id", "netbox_vlan_translation_policy.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_translation_policy.test", "id", "netbox_vlan_translation_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_rules.test", "vlan_translation_rules.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_policies.test", "vlan_translation_policies.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_translation_rule.test", "id", "netbox_vlan_translation_rule.test", "id"),
				),
			},
			{
				// Shrink: both descriptions and the comments clear.
				Config: fmt.Sprintf(`
resource "netbox_vlan_translation_policy" "test" {
  name = "%[1]s"
}
resource "netbox_vlan_translation_rule" "test" {
  vlan_translation_policy_id = netbox_vlan_translation_policy.test.id
  local_vid                  = 100
  remote_vid                 = 200
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_vlan_translation_policy.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_vlan_translation_policy.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_vlan_translation_rule.test", "description"),
				),
			},
			{
				ResourceName:      "netbox_vlan_translation_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_vlan_translation_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_asn_range",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamAsnRangesList(ipam.NewIpamAsnRangesListParams(), nil)
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
			_, err := client.Ipam.IpamAsnRangesDestroy(ipam.NewIpamAsnRangesDestroyParams().WithID(id), nil)
			return err
		})
	// Translation rules are deleted with their policy, so only the policy needs a sweeper.
	sweep("netbox_vlan_translation_policy",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamVlanTranslationPoliciesList(ipam.NewIpamVlanTranslationPoliciesListParams(), nil)
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
			_, err := client.Ipam.IpamVlanTranslationPoliciesDestroy(ipam.NewIpamVlanTranslationPoliciesDestroyParams().WithID(id), nil)
			return err
		})
}
