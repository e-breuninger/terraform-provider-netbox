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
				AtLeastOneOf: []string{"id", "name"},
				Optional:     true,
				Type:         schema.TypeString,
			},
			"id": {
				AtLeastOneOf: []string{"id", "name"},
				Computed:     true,
				Optional:     true,
				Type:         schema.TypeString,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"group_id": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"user_group_ids": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"user_ids": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceNetboxOwnerRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := users.NewUsersOwnersListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Users.UsersOwnersList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
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
	} else {
		d.Set("group_id", nil)
	}

	d.Set("user_group_ids", getIDsFromNestedGroup(result.UserGroups))
	d.Set("user_ids", getIDsFromNestedUser(result.Users))

	return nil
}
