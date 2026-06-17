package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var datasourceNetboxVpnTunnelTerminationRoleOptions = []string{"peer", "hub", "spoke"}

func dataSourceNetboxVpnTunnelTerminations() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVpnTunnelTerminationsRead,
		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/features/vpn-tunnels/):

> NetBox can model private tunnels formed among virtual termination points across your network. Typical tunnel implementations include GRE, IP-in-IP, and IPSec. A tunnel may be terminated to two or more device or virtual machine interfaces. For convenient organization, tunnels may be assigned to user-defined groups.`,
		Schema: map[string]*schema.Schema{
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
			},
			"terminations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tunnel_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: buildValidValueDescription(datasourceNetboxVpnTunnelTerminationRoleOptions),
						},
						"termination_type": {
							Type:         schema.TypeString,
							Computed:     true,
						},
						"termination_id": {
							Type:         schema.TypeInt,
							Computed:     true,
						},
						tagsKey: tagsSchema,
					},
				},
			},
		},
	}
}

func dataSourceNetboxVpnTunnelTerminationsRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := vpn.NewVpnTunnelTerminationsListParams()

	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	// Fetch all pages with pagination
	paginationHelper := NewPaginationHelper(userLimit)
	var allTunnelTerminations []*models.TunnelTermination

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Vpn.VpnTunnelTerminationsList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch Tunnel Terminations at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allTunnelTerminations = append(allTunnelTerminations, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allTunnelTerminations)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	trimmedCount := paginationHelper.TrimToLimit(len(allTunnelTerminations))
	filteredTunnelTerminations := allTunnelTerminations[:trimmedCount]

	if len(filteredTunnelTerminations) == 0 {
		return errors.New("no result")
	}

	var s []map[string]interface{}
	for _, v := range filteredTunnelTerminations {
		var mapping = make(map[string]interface{})

		mapping["id"] = v.ID
		mapping["tunnel_id"] = v.Tunnel.ID
		mapping["role"] = v.Role.Value
		mapping["termination_type"] = v.TerminationType
		mapping["termination_id"] = v.TerminationID
		mapping["tags"] = getTagListFromNestedTagList(v.Tags)

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("terminations", s)
}
