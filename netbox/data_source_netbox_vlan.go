package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVlan() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVlanRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"vid": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"name", "vid"},
				ValidateFunc: validation.IntBetween(1, 4094),
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "vid"},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxVlanRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := ipam.NewIpamVlansListParams()

	params.Limit = int64ToPtr(2)
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}
	if vid, ok := d.Get("vid").(int); ok && vid != 0 {
		params.Vid = strToPtr(strconv.Itoa(vid))
	}

	res, err := api.Ipam.IpamVlansList(params, nil)
	if err != nil {
		return err
	}
	if count := *res.GetPayload().Count; count != int64(1) {
		return fmt.Errorf("expected one device type, but got %d", count)
	}

	vlan := res.GetPayload().Results[0]

	d.SetId(strconv.FormatInt(vlan.ID, 10))
	d.Set("vid", vlan.Vid)
	d.Set("name", vlan.Name)
	d.Set("status", vlan.Status.Value)
	d.Set("description", vlan.Description)

	if vlan.Role != nil {
		d.Set("role", vlan.Role.ID)
	}
	if vlan.Site != nil {
		d.Set("site", vlan.Site.ID)
	}
	if vlan.Tenant != nil {
		d.Set("tenant", vlan.Tenant.ID)
	}

	return nil
}
