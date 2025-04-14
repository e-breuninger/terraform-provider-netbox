package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxGroupCreate,
		Read:   resourceNetboxGroupRead,
		Update: resourceNetboxGroupUpdate,
		Delete: resourceNetboxGroupDelete,

		Description: `:meta:subcategory:Authentication:This resource is used to manage groups.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.Group{}

	name := d.Get("name").(string)

	data.Name = &name

	params := users.NewUsersGroupsCreateParams().WithData(&data)
	res, err := api.Users.UsersGroupsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxGroupRead(d, m)
}

func resourceNetboxGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersGroupsReadParams().WithID(id)

	res, err := api.Users.UsersGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersGroupsReadDefault); ok {
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

	return nil
}

func resourceNetboxGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.Group{}

	name := d.Get("name").(string)

	data.Name = &name

	params := users.NewUsersGroupsUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersGroupsUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxGroupRead(d, m)
}

func resourceNetboxGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersGroupsDeleteParams().WithID(id)
	_, err := api.Users.UsersGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersGroupsDeleteDefault); ok {
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
