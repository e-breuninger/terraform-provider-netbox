package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxLocation() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxLocationRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxLocationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimLocationsListParams()

	params.Limit = int64ToPtr(2)
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.SetName(&name)
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.SetSlug(&slug)
	}
	if id, ok := d.Get("id").(string); ok && id != "0" {
		params.SetID(&id)
	}
	if site, ok := d.Get("site_id").(int); ok && site != 0 {
		siteID := fmt.Sprintf("%v", site)
		params.SetSiteID(&siteID)
	}
	res, err := api.Dcim.DcimLocationsList(params, nil)

	if err != nil {
		return err
	}
	if count := *res.GetPayload().Count; count != 1 {
		return fmt.Errorf("expected one site, but got %d", count)
	}

	location := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(location.ID, 10))
	d.Set("description", location.Description)
	d.Set("name", location.Name)
	d.Set("site_id", location.Site.ID)
	d.Set("slug", location.Slug)

	if location.Status != nil {
		d.Set("status", location.Status.Value)
	}
	if location.Tenant != nil {
		d.Set("tenant_id", location.Tenant.ID)
	}

	return nil
}
