package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVpnTunnelGroup() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVpnTunnelGroupRead,
		Description: `:meta:subcategory:VPN Tunnels:`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": tagsSchemaRead,
		},
	}
}

func dataSourceNetboxVpnTunnelGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := vpn.NewVpnTunnelGroupsListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}
	if slug, ok := d.Get("slug").(string); ok && slug != "" {
		params.Slug = &slug
	}

	if tag, ok := d.Get("tag").(string); ok && tag != "" {
		params.Tag = []string{tag} //TODO: switch schema to list?
	}

	res, err := api.Vpn.VpnTunnelGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one tunnel group returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no tunnel group found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("id", result.ID)
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	d.Set("description", result.Description)
	d.Set("tags", getTagListFromNestedTagList(result.Tags))
	return nil
}
