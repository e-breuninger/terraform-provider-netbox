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

		Description: `:meta:subcategory:Authentication:From the official NetBox documentation:

> An owner represents a user and/or user group(s) which can be assigned responsibility for NetBox objects.`,

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
	data := models.WritableOwner{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Description = description
	data.Group = getOptionalInt(d, "group_id")
	data.UserGroups = toInt64List(d.Get("user_group_ids"))
	data.Users = toInt64List(d.Get("user_ids"))

	params := users.NewUsersOwnersCreateParams().WithData(&data)
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

	owner := res.GetPayload()
	if owner.Name != nil {
		d.Set("name", owner.Name)
	}
	d.Set("description", owner.Description)

	if owner.Group != nil {
		d.Set("group_id", owner.Group.ID)
	} else {
		d.Set("group_id", nil)
	}

	d.Set("user_group_ids", getIDsFromNestedGroup(owner.UserGroups))
	d.Set("user_ids", getIDsFromNestedUser(owner.Users))

	return nil
}

func resourceNetboxOwnerUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableOwner{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Description = description
	data.Group = getOptionalInt(d, "group_id")
	data.UserGroups = toInt64List(d.Get("user_group_ids"))
	data.Users = toInt64List(d.Get("user_ids"))

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
	d.SetId("")
	return nil
}
