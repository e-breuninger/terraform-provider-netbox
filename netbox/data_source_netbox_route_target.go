package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxRouteTarget() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxRouteTargetRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(1, 21),
				Required:     true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			tagsKey: tagsSchema,
		},
	}
}

func dataSourceNetboxRouteTargetRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	params := ipam.NewIpamRouteTargetsListParams()
	params.Name = &name

	limit := int64(2)
	params.Limit = &limit

	res, err := api.Ipam.IpamRouteTargetsList(params, nil)

	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	}
	if result.Description != "" {
		d.Set("description", result.Description)
	}
	if result.Tags != nil {
		d.Set(tagsKey, result.Tags)

	}

	return nil
}
