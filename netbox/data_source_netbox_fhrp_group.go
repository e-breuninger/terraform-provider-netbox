package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxFhrpGroup() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxFhrpGroupRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
			},
			"group_id": {
				Type:         schema.TypeInt,
				Optional:     false,
				Required: true,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchemaRead,
		},
	}
}

func dataSourceNetboxFhrpGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := ipam.NewIpamFhrpGroupsListParams()

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	if group_id, ok := d.Get("group_id").(int); ok && group_id != 0 {
		params.GroupID = strToPtr(strconv.FormatInt(int64(group_id), 10))
	}

	if protocol, ok := d.Get("protocol").(string); ok && protocol != "" {
		params.Protocol = &protocol
	}

	res, err := api.Ipam.IpamFhrpGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one group returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no group found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.Set("id", result.ID)
	d.Set("protocol", result.Protocol)
	d.Set("group_id", result.GroupID)
	d.Set("auth_type", result.AuthType)
	d.Set("auth_key", result.AuthKey)
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set("tags", getTagListFromNestedTagList(result.Tags))
	d.SetId(strconv.FormatInt(result.ID, 10))
	return nil
}
