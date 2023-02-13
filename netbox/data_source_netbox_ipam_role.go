package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxIPAMRole() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIPAMRoleRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
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

func dataSourceNetboxIPAMRoleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	params := ipam.NewIpamRolesListParams()
	params.Name = &name

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Ipam.IpamRolesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("slug", result.Slug)
	d.Set("name", result.Name)
	if result.Weight != nil {
		d.Set("weight", result.Weight)
	}
	if result.Description != "" {
		d.Set("description", result.Description)

	}
	return nil
}
