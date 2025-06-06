package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/strfmt"
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
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
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
			"group_ids": {
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
func resourceNetboxUserCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.WritableUser{}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	email := d.Get("email").(string)
	firstName := d.Get("first_name").(string)
	lastName := d.Get("last_name").(string)
	active := d.Get("active").(bool)
	staff := d.Get("staff").(bool)
	groupIDs := toInt64List(d.Get("group_ids"))

	data.Username = &username
	data.Password = &password
	data.Email = strfmt.Email(email)
	data.FirstName = firstName
	data.LastName = lastName
	data.IsActive = active
	data.IsStaff = staff
	data.Groups = groupIDs

	params := users.NewUsersUsersCreateParams().WithData(&data)
	res, err := api.Users.UsersUsersCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxUserRead(d, m)
}

func resourceNetboxUserRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersUsersReadParams().WithID(id)

	res, err := api.Users.UsersUsersRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersUsersReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if res.GetPayload().Username != nil {
		d.Set("username", res.GetPayload().Username)
	}

	d.Set("email", res.GetPayload().Email)
	d.Set("first_name", res.GetPayload().FirstName)
	d.Set("last_name", res.GetPayload().LastName)

	d.Set("staff", res.GetPayload().IsStaff)
	d.Set("active", res.GetPayload().IsActive)
	d.Set("group_ids", getIDsFromNestedGroup(res.GetPayload().Groups))

	// Passwords cannot be set and not read

	return nil
}

func resourceNetboxUserUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableUser{}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	email := d.Get("email").(string)
	firstName := d.Get("first_name").(string)
	lastName := d.Get("last_name").(string)
	active := d.Get("active").(bool)
	staff := d.Get("staff").(bool)
	groupIDs := toInt64List(d.Get("group_ids"))

	data.Username = &username
	data.Password = &password
	data.Email = strfmt.Email(email)
	data.FirstName = firstName
	data.LastName = lastName
	data.IsActive = active
	data.IsStaff = staff
	data.Groups = groupIDs

	params := users.NewUsersUsersUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersUsersUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxUserRead(d, m)
}

func resourceNetboxUserDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
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

func getIDsFromNestedGroup(nestedGroups []*models.NestedGroup) []int64 {
	var groupIDs []int64
	for _, group := range nestedGroups {
		groupIDs = append(groupIDs, group.ID)
	}
	return groupIDs
}
