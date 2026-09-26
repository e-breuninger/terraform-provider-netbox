//go:build acctest

// The VPN policy objects NetBox 4.6 added to go-netbox: IKE policies and the IPSec proposal,
// policy and profile family.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxIkePolicy_basic(t *testing.T) {
	testName := testAccGetTestName("ikepolicy")
	deps := fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                  = "%[1]s"
  authentication_method = "preshared-keys"
  encryption_algorithm  = "aes-256-cbc"
  group                 = 14
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_ike_policy" "test" {
  name             = "%[1]s"
  version          = 1
  mode             = "main"
  ike_proposal_ids = [netbox_ike_proposal.test.id]
  preshared_key    = "acceptance-test-psk"
  description      = "Acceptance test IKE policy."
  comments         = "Created by acceptance test."
}
data "netbox_ike_policy" "test" {
  name = netbox_ike_policy.test.name
}
data "netbox_ike_policies" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_ike_policy.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "version", "1"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "mode", "main"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "ike_proposal_ids.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ike_policy.test", "id", "netbox_ike_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ike_policies.test", "ike_policies.#", "1"),
				),
			},
			{
				// Shrink: everything optional clears, except mode, which NetBox refuses to blank
				// (it stays). The version stays 1 for the same reason: NetBox rejects a mode on
				// IKEv2, and the mode cannot be removed.
				Config: deps + fmt.Sprintf(`
resource "netbox_ike_policy" "test" {
  name             = "%[1]s"
  version          = 1
  ike_proposal_ids = [netbox_ike_proposal.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "version", "1"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "mode", "main"),
					resource.TestCheckNoResourceAttr("netbox_ike_policy.test", "preshared_key"),
					resource.TestCheckNoResourceAttr("netbox_ike_policy.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ike_policy.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_ike_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIpsecProposal_basic(t *testing.T) {
	testName := testAccGetTestName("ipsecprop")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ipsec_proposal" "test" {
  name                     = "%[1]s"
  encryption_algorithm     = "aes-256-cbc"
  authentication_algorithm = "hmac-sha256"
  sa_lifetime_seconds      = 3600
  sa_lifetime_data         = 102400
  description              = "Acceptance test IPSec proposal."
  comments                 = "Created by acceptance test."
}
data "netbox_ipsec_proposal" "test" {
  name = netbox_ipsec_proposal.test.name
}
data "netbox_ipsec_proposals" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_ipsec_proposal.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "encryption_algorithm", "aes-256-cbc"),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "sa_lifetime_seconds", "3600"),
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_proposal.test", "id", "netbox_ipsec_proposal.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposals.test", "ipsec_proposals.#", "1"),
				),
			},
			{
				// Shrink: the lifetimes and the text clear; authentication_algorithm cannot be blanked,
				// so it keeps its value.
				Config: fmt.Sprintf(`
resource "netbox_ipsec_proposal" "test" {
  name                 = "%[1]s"
  encryption_algorithm = "aes-256-cbc"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_proposal.test", "sa_lifetime_seconds"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_proposal.test", "sa_lifetime_data"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_proposal.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_proposal.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_ipsec_proposal.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIpsecPolicy_basic(t *testing.T) {
	testName := testAccGetTestName("ipsecpol")
	deps := fmt.Sprintf(`
resource "netbox_ipsec_proposal" "test" {
  name                 = "%[1]s"
  encryption_algorithm = "aes-256-cbc"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_ipsec_policy" "test" {
  name               = "%[1]s"
  ipsec_proposal_ids = [netbox_ipsec_proposal.test.id]
  pfs_group          = 14
  description        = "Acceptance test IPSec policy."
  comments           = "Created by acceptance test."
}
data "netbox_ipsec_policy" "test" {
  name = netbox_ipsec_policy.test.name
}
data "netbox_ipsec_policies" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_ipsec_policy.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "pfs_group", "14"),
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "ipsec_proposal_ids.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_policy.test", "id", "netbox_ipsec_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_policies.test", "ipsec_policies.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_ipsec_policy" "test" {
  name               = "%[1]s"
  ipsec_proposal_ids = [netbox_ipsec_proposal.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					// pfs_group cannot be blanked, so it keeps its value.
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "pfs_group", "14"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_policy.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_policy.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_ipsec_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIpsecProfile_basic(t *testing.T) {
	testName := testAccGetTestName("ipsecprof")
	deps := fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                  = "%[1]s"
  authentication_method = "preshared-keys"
  encryption_algorithm  = "aes-256-cbc"
  group                 = 14
}
resource "netbox_ike_policy" "test" {
  name             = "%[1]s"
  version          = 2
  ike_proposal_ids = [netbox_ike_proposal.test.id]
}
resource "netbox_ipsec_proposal" "test" {
  name                 = "%[1]s"
  encryption_algorithm = "aes-256-cbc"
}
resource "netbox_ipsec_policy" "test" {
  name               = "%[1]s"
  ipsec_proposal_ids = [netbox_ipsec_proposal.test.id]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_ipsec_profile" "test" {
  name            = "%[1]s"
  mode            = "esp"
  ike_policy_id   = netbox_ike_policy.test.id
  ipsec_policy_id = netbox_ipsec_policy.test.id
  description     = "Acceptance test IPSec profile."
  comments        = "Created by acceptance test."
}
data "netbox_ipsec_profile" "test" {
  name = netbox_ipsec_profile.test.name
}
data "netbox_ipsec_profiles" "test" {
  filters = [
    { name = "ike_policy_id", value = netbox_ike_policy.test.id },
  ]
  depends_on    = [netbox_ipsec_profile.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "mode", "esp"),
					resource.TestCheckResourceAttrPair("netbox_ipsec_profile.test", "ike_policy_id", "netbox_ike_policy.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ipsec_profile.test", "ipsec_policy_id", "netbox_ipsec_policy.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_profile.test", "id", "netbox_ipsec_profile.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_profiles.test", "ipsec_profiles.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_ipsec_profile" "test" {
  name            = "%[1]s"
  mode            = "ah"
  ike_policy_id   = netbox_ike_policy.test.id
  ipsec_policy_id = netbox_ipsec_policy.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "mode", "ah"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_profile.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ipsec_profile.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_ipsec_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_ipsec_profile",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnIpsecProfilesList(vpn.NewVpnIpsecProfilesListParams(), nil)
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
			_, err := client.Vpn.VpnIpsecProfilesDestroy(vpn.NewVpnIpsecProfilesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_ipsec_policy",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnIpsecPoliciesList(vpn.NewVpnIpsecPoliciesListParams(), nil)
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
			_, err := client.Vpn.VpnIpsecPoliciesDestroy(vpn.NewVpnIpsecPoliciesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_ipsec_proposal",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnIpsecProposalsList(vpn.NewVpnIpsecProposalsListParams(), nil)
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
			_, err := client.Vpn.VpnIpsecProposalsDestroy(vpn.NewVpnIpsecProposalsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_ike_policy",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnIkePoliciesList(vpn.NewVpnIkePoliciesListParams(), nil)
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
			_, err := client.Vpn.VpnIkePoliciesDestroy(vpn.NewVpnIkePoliciesDestroyParams().WithID(id), nil)
			return err
		})
}
