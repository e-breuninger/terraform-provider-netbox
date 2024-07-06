package netbox

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetboxJournal(t *testing.T) {
	cl, err := sharedClientForRegion("test")
	if err != nil {
		t.Fatal(err)
	}
	testClient := cl.(*client.NetBoxAPI)

	netboxProvider := *testAccProvider
	origConfigure := netboxProvider.ConfigureContextFunc
	netboxProvider.ConfigureContextFunc = func(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
		rd.Set("journal_entry", "Test journal entry to be written")
		return origConfigure(ctx, rd)
	}
	providers := map[string]*schema.Provider{
		"netbox": &netboxProvider,
	}

	resource.ParallelTest(t, resource.TestCase{
		Providers: providers,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: `
resource "netbox_site" "test" {
	name = "Test site for journal"
}`,
				Check: func(s *terraform.State) error {
					site, ok := s.RootModule().Resources["netbox_site.test"]
					if !ok {
						return errors.New("site resource not found in state")
					}

					p := extras.NewExtrasJournalEntriesListParams()
					p.AssignedObjectID = &site.Primary.ID
					res, err := testClient.Extras.ExtrasJournalEntriesList(p, nil)
					if err != nil {
						return fmt.Errorf("failed to get journal from API: %w", err)
					}
					entries := res.GetPayload().Results
					if len(entries) != 1 {
						return fmt.Errorf("invalid number of journal entries: %d", len(entries))
					}
					entry := entries[0]
					if *entry.AssignedObjectType != "dcim.site" {
						return fmt.Errorf("invalid object type on entry: %s", *entry.AssignedObjectType)
					}
					if *entry.Kind.Value != models.JournalEntryKindValueSuccess {
						return fmt.Errorf("invalid kind: %v", *entry.Kind)
					}
					if *entry.Comments != "Test journal entry to be written" {
						return fmt.Errorf("unexpected comment: %s", *entry.Comments)
					}
					return nil
				},
			},
		},
	})
}
