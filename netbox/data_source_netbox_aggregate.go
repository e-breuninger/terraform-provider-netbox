package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxAggregate() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxAggregateRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"rir_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxAggregateRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamAggregatesListParams()

	prefix := d.Get("prefix").(string)
	params.Prefix = &prefix

	if rirID, ok := d.GetOk("rir_id"); ok {
		rirIDStr := strconv.Itoa(rirID.(int))
		params.RirID = &rirIDStr
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamAggregatesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one aggregate returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no aggregate found matching filter")
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("prefix", result.Prefix)
	if result.Rir != nil {
		d.Set("rir_id", result.Rir.ID)
	}
	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	}
	d.Set("description", result.Description)
	return nil
}
