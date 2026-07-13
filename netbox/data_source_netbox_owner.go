package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxOwner() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxOwnerRead,
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
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_group_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"user_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceNetboxOwnerRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	params := users.NewUsersOwnersListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough

	params.Limit = &limit

	res, err := api.Users.UsersOwnersList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one owner returned, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no owner found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("description", result.Description)
	if result.Group != nil {
		d.Set("group_id", result.Group.ID)
	}
	d.Set("user_group_ids", getIDsFromNestedGroup(result.UserGroups))
	d.Set("user_ids", getIDsFromNestedUser(result.Users))
	return nil
}
