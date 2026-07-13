package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxOwner() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxOwnerCreate,
		Read:   resourceNetboxOwnerRead,
		Update: resourceNetboxOwnerUpdate,
		Delete: resourceNetboxOwnerDelete,

		Description: `:meta:subcategory:Authentication:From the NetBox documentation:

> An Owner represents a set of users and/or groups responsible for administering NetBox resources. Owner assignments are useful for indicating which parties are responsible for the administration of a particular object.
>
> Most objects within NetBox can be assigned an owner, although this is not required.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"user_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"user_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxOwnerCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	groupID := int64(d.Get("group_id").(int))
	userGroupIDs := toInt64List(d.Get("user_group_ids"))
	userIDs := toInt64List(d.Get("user_ids"))

	data := &models.WritableOwner{}
	data.Name = &name
	data.Description = description
	data.UserGroups = userGroupIDs
	data.Users = userIDs

	if groupID != 0 {
		data.Group = &groupID
	}

	params := users.NewUsersOwnersCreateParams().WithData(data)

	res, err := api.Users.UsersOwnersCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxOwnerRead(d, m)
}

func resourceNetboxOwnerRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersOwnersReadParams().WithID(id)

	res, err := api.Users.UsersOwnersRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersOwnersReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("description", res.GetPayload().Description)
	if res.GetPayload().Group != nil {
		d.Set("group_id", res.GetPayload().Group.ID)
	} else {
		d.Set("group_id", nil)
	}
	d.Set("user_group_ids", getIDsFromNestedGroup(res.GetPayload().UserGroups))
	d.Set("user_ids", getIDsFromNestedUser(res.GetPayload().Users))

	return nil
}

func resourceNetboxOwnerUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableOwner{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	groupID := int64(d.Get("group_id").(int))
	userGroupIDs := toInt64List(d.Get("user_group_ids"))
	userIDs := toInt64List(d.Get("user_ids"))

	data.Name = &name
	data.Description = description
	data.UserGroups = userGroupIDs
	data.Users = userIDs

	if groupID != 0 {
		data.Group = &groupID
	}

	params := users.NewUsersOwnersPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Users.UsersOwnersPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxOwnerRead(d, m)
}

func resourceNetboxOwnerDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersOwnersDeleteParams().WithID(id)

	_, err := api.Users.UsersOwnersDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersOwnersDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}

func getIDsFromNestedUser(nestedUsers []*models.NestedUser) []int64 {
	var userIDs []int64
	for _, user := range nestedUsers {
		userIDs = append(userIDs, user.ID)
	}
	return userIDs
}
