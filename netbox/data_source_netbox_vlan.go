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
				ValidateFunc: validation.IntBetween(1, 4094),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
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
				Optional: true,
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
	if groupID, ok := d.Get("group_id").(int); ok && groupID != 0 {
		params.GroupID = strToPtr(strconv.Itoa(groupID))
	}
	if roleID, ok := d.Get("role").(int); ok && roleID != 0 {
		params.RoleID = strToPtr(strconv.Itoa(roleID))
	}
	if tenantID, ok := d.Get("tenant").(int); ok && tenantID != 0 {
		params.TenantID = strToPtr(strconv.Itoa(tenantID))
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

	if vlan.Group != nil {
		d.Set("group_id", vlan.Group.ID)
	}
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
