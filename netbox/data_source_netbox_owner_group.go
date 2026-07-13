package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxOwnerGroup() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxOwnerGroupRead,
		Description: `:meta:subcategory:Authentication:`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxOwnerGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	params := users.NewUsersOwnerGroupsListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough

	params.Limit = &limit

	res, err := api.Users.UsersOwnerGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one owner group returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no owner group found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	return nil
}
