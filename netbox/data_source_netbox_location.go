package netbox

import (
	"errors"
	"fmt"
	"strconv"

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
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxLocationRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
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
	if parent, ok := d.Get("parent_id").(int); ok && parent != 0 {
		parentID := fmt.Sprintf("%v", parent)
		params.SetParentID(&parentID)
	}
	res, err := api.Dcim.DcimLocationsList(params, nil)

	if err != nil {
		return err
	}
	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one location returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no location found matching filter")
	}

	location := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(location.ID, 10))
	d.Set("description", location.Description)
	d.Set("name", location.Name)
	d.Set("site_id", location.Site.ID)
	d.Set("slug", location.Slug)

	if location.Parent != nil {
		d.Set("parent_id", location.Parent.ID)
	}
	if location.Status != nil {
		d.Set("status", location.Status.Value)
	}
	if location.Tenant != nil {
		d.Set("tenant_id", location.Tenant.ID)
	}

	return nil
}
