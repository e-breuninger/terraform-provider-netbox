package netbox

import (
	"errors"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxPrefixByDescription() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxPrefixByDescriptionRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxPrefixByDescriptionRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	description := d.Get("description").(string)

	id, prefix, err := dataSourceNetboxPrefixAPICall(api, description)

	if err != nil {
		return err
	}

	d.Set("cidr", prefix)
	d.SetId(id)
	return nil
}

func dataSourceNetboxPrefixAPICall(api *client.NetBoxAPI, description string) (string, string, error) {

	params := ipam.NewIpamPrefixesListParams()
	params.Q = &description

	res, err := api.Ipam.IpamPrefixesList(params, nil)

	if err != nil {
		return "", "", err
	}

	if *res.GetPayload().Count == int64(0) {
		return "", "", errors.New(fmt.Sprintf("No result for %s", description))
	}

	var hit *models.Prefix

	for _, result := range res.GetPayload().Results {
		if result.Description == description {
			hit = result
			if hit != nil {
				return "", "", errors.New(fmt.Sprintf("Multiple matches found for %s, can't continue.", description))
			}
		}
	}

	return "", "", errors.New(fmt.Sprintf("No exact match found for %s", description))
}
