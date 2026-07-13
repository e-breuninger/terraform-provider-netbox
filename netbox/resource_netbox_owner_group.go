package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxOwnerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxOwnerGroupCreate,
		Read:   resourceNetboxOwnerGroupRead,
		Update: resourceNetboxOwnerGroupUpdate,
		Delete: resourceNetboxOwnerGroupDelete,

		Description: `:meta:subcategory:Authentication:From the NetBox documentation:

> Groups are used to correlate and organize owners. The assignment of an owner to a group has no bearing on the relationship of owned objects to their owners; groups exist purely as an organizational convenience for administrators.`,

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
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxOwnerGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.OwnerGroup{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Description = description

	params := users.NewUsersOwnerGroupsCreateParams().WithData(&data)
	res, err := api.Users.UsersOwnerGroupsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxOwnerGroupRead(d, m)
}

func resourceNetboxOwnerGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersOwnerGroupsReadParams().WithID(id)

	res, err := api.Users.UsersOwnerGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersOwnerGroupsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if res.GetPayload().Name != nil {
		d.Set("name", res.GetPayload().Name)
	}
	d.Set("description", res.GetPayload().Description)

	return nil
}

func resourceNetboxOwnerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.OwnerGroup{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Description = description

	params := users.NewUsersOwnerGroupsPartialUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersOwnerGroupsPartialUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxOwnerGroupRead(d, m)
}

func resourceNetboxOwnerGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersOwnerGroupsDeleteParams().WithID(id)
	_, err := api.Users.UsersOwnerGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersOwnerGroupsDeleteDefault); ok {
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
