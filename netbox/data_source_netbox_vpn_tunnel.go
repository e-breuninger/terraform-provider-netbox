package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxVpnTunnel() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVpnTunnelRead,
		Description: `:meta:subcategory:VPN Tunnels:`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:	  true,
			},
			"encapsualation": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:	  true,
			},
			"tunnel_group_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:	  true,
			},
			"tunnel_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:	  true,
			},
			"tenant_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:	  true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchemaRead,
		},
	}
}

func dataSourceNetboxVpnTunnelRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := vpn.NewVpnTunnelsListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}
	if status, ok := d.Get("status").(string); ok && status != "" {
		params.Status = &status
	}
	if encapsulation, ok := d.Get("encapsulation").(string); ok && encapsulation != "" {
		params.Encapsulation = &encapsulation
	}
	
	if group_id, ok := d.Get("group_id").(string); ok && group_id != "" {
		params.GroupID = &group_id
	}
	if tunnel_id, ok := d.Get("tunnel_id").(string); ok && tunnel_id != "" {
		params.TunnelID = &tunnel_id
	}
	if tenant_id, ok := d.Get("tenant_id").(string); ok && tenant_id != "" {
		params.TenantID = &tenant_id
	}

	if tag, ok := d.Get("tag").(string); ok && tag != "" {
		params.Tag = []string{tag} //TODO: switch schema to list?
	}

	res, err := api.Vpn.VpnTunnelsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one vpn tunnel returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no vpn tunnel found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("id", result.ID)
	d.Set("name", result.Name)
	d.Set("status", result.Status.Value)
	d.Set("encapsualation", result.Encapsulation.Value)
	d.Set("tunnel_group_id", strconv.FormatInt(result.Group.ID, 10))
	d.Set("tunnel_id", strconv.FormatInt(*result.TunnelID, 10))
	d.Set("tenant_id", strconv.FormatInt(result.Tenant.ID, 10))
	d.Set("description", result.Description)
	d.Set("tags", getTagListFromNestedTagList(result.Tags))
	return nil
}
