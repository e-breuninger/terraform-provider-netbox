package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxSite() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxSiteRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"facility": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asn_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
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
	api := m.(*providerState)
	params := dcim.NewDcimSitesListParams()

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
	if facility, ok := d.Get("facility").(string); ok && facility != "" {
		params.SetFacility(&facility)
	}

	res, err := api.Dcim.DcimSitesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one site returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no site found matching filter")
	}

	site := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(site.ID, 10))
	d.Set("asn_ids", getIDsFromNestedASNList(site.Asns))
	d.Set("comments", site.Comments)
	d.Set("description", site.Description)
	d.Set("name", site.Name)
	d.Set("site_id", site.ID)
	d.Set("slug", site.Slug)
	d.Set("time_zone", site.TimeZone)
	d.Set("facility", site.Facility)

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
