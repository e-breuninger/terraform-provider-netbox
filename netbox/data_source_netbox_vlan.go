package netbox

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxVlan() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxVlanRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				ExactlyOneOf: []string{"name", "vid"},
				Optional:     true,
			},
			"vid": {
				Type:         schema.TypeInt,
				ExactlyOneOf: []string{"name", "vid"},
				Optional:     true,
			},
		},
	}
}

func dataSourceNetboxVlanRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	params := ipam.NewIpamVlansListParams()

	limit := int64(2) // Limit of 2 is enough
	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}
	if vid, ok := d.Get("vid").(int); ok && vid != 0 {
		params.Vid = strToPtr(strconv.Itoa(vid))
	}

	params.Limit = &limit
	res, err := api.Ipam.IpamVlansList(params, nil)
	if err != nil {
		return err
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	if count := *res.GetPayload().Count; count > int64(1) {
		return fmt.Errorf("expected one VLAN, but got %d", count)
	}

	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("id", result.ID)
	d.Set("name", result.Name)
	d.Set("vid", result.Vid)
	return nil
}
