package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxIpRange() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIpRangeRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"contains": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
		},
	}
}

func dataSourceNetboxIpRangeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	contains := d.Get("contains").(string)

	params := ipam.NewIpamIPRangesListParams()
	params.Contains = &contains

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamIPRangesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
