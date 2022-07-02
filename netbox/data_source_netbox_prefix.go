package netbox

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxPrefixRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cidr": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceNetboxPrefixRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	cidr := d.Get("cidr").(string)
	description := d.Get("description").(string)

	if cidr == "" && description == "" {
		return errors.New("Either a prefix or a description should be given.")
	} else if cidr == "" && description != "" {
		err := dataSourceNetBoxPrefixReadByDesc(api, d, description)
		if err != nil {
			return err
		}
	} else if cidr != "" {
		err := dataSourceNetBoxPrefixReadByCidr(api, d, cidr)
		if err != nil {
			return err
		}
	}

	return nil
}

func dataSourceNetBoxPrefixReadByCidr(api *client.NetBoxAPI, d *schema.ResourceData, cidr string) error {
	params := ipam.NewIpamPrefixesListParams()
	params.Prefix = &cidr

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("More than one result. Specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("No result")
	}

	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.Set("description", result.Description)
	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}

func dataSourceNetBoxPrefixReadByDesc(api *client.NetBoxAPI, d *schema.ResourceData, description string) error {
	params := ipam.NewIpamPrefixesListParams().WithDefaults()
	params.Q = &description

	res, err := api.Ipam.IpamPrefixesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count == int64(0) {
		return errors.New(fmt.Sprintf("No result for %s", description))
	}

	var hit *models.Prefix

	for _, result := range res.GetPayload().Results {
		if result.Description == description {
			if hit != nil {
				return errors.New(fmt.Sprintf("Multiple matches found for %s, can't continue.", description))
			}

			hit = result
		}
	}

	if hit == nil {
		return errors.New(fmt.Sprintf("No exact match found for %s", description))
	}

	d.Set("id", hit.ID)
	d.Set("cidr", hit.Prefix)
	d.SetId(strconv.FormatInt(hit.ID, 10))
	return nil
}
