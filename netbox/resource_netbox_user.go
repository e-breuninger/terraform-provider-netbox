package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxUserCreate,
		Read:   resourceNetboxUserRead,
		Update: resourceNetboxUserUpdate,
		Delete: resourceNetboxUserDelete,

		Description: `:meta:subcategory:Authentication:This resource is used to manage users.`,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"staff": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxUserCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableUser{}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	active := d.Get("active").(bool)
	staff := d.Get("staff").(bool)

	data.Username = &username
	data.Password = &password
	data.IsActive = active
	data.IsStaff = staff

	data.Groups = []int64{}

	params := users.NewUsersUsersCreateParams().WithData(&data)
	res, err := api.Users.UsersUsersCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxUserRead(d, m)
}

func resourceNetboxUserRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersUsersReadParams().WithID(id)

	res, err := api.Users.UsersUsersRead(params, nil)
	if err != nil {
		errorcode := err.(*users.UsersUsersReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	if res.GetPayload().Username != nil {
		d.Set("username", res.GetPayload().Username)
	}

	d.Set("staff", res.GetPayload().IsStaff)
	d.Set("active", res.GetPayload().IsActive)

	// Passwords cannot be set and not read

	return nil
}

func resourceNetboxUserUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableUser{}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	active := d.Get("active").(bool)
	staff := d.Get("staff").(bool)

	data.Username = &username
	data.Password = &password
	data.IsActive = active
	data.IsStaff = staff

	data.Groups = []int64{}

	params := users.NewUsersUsersUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersUsersUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxUserRead(d, m)
}

func resourceNetboxUserDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersUsersDeleteParams().WithID(id)
	_, err := api.Users.UsersUsersDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersUsersDeleteDefault); ok {
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
