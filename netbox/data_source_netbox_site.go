package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxSite() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxSiteRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asn": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxSiteRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := dcim.NewDcimSitesListParams()

	params.Limit = int64ToPtr(2)
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.SetName(&name)
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.SetSlug(&slug)
	}

	res, err := api.Dcim.DcimSitesList(params, nil)
	if err != nil {
		return err
	}
	if count := *res.GetPayload().Count; count != 1 {
		return fmt.Errorf("expected one site, but got %d", count)
	}

	site := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(site.ID, 10))
	d.Set("asn", site.Asn)
	d.Set("comments", site.Comments)
	d.Set("description", site.Description)
	d.Set("name", site.Name)
	d.Set("site_id", site.ID)
	d.Set("slug", site.Slug)
	d.Set("time_zone", site.TimeZone)

	if site.Group != nil {
		d.Set("group_id", site.Group.ID)
	}
	if site.Region != nil {
		d.Set("region_id", site.Region.ID)
	}
	if site.Status != nil {
		d.Set("status", site.Status.Value)
	}
	if site.Tenant != nil {
		d.Set("tenant_id", site.Tenant.ID)
	}

	return nil
}
