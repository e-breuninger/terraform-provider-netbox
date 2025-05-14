package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxVrf() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVrfRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataSourceNetboxVrfRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	params := ipam.NewIpamVrfsListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if tenantID, ok := d.Get("tenant_id").(int); ok && tenantID != 0 {
		// Note that tenant_id is a string pointer in the netbox filter, but we use a number in the provider
		params.TenantID = strToPtr(strconv.Itoa(tenantID))
	}

	res, err := api.Ipam.IpamVrfsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one vrf returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no vrf found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}
	return nil
}
