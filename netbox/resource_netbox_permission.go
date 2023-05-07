package netbox

import (
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxPermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPermissionCreate,
		Read:   resourceNetboxPermissionRead,
		Update: resourceNetboxPermissionUpdate,
		Delete: resourceNetboxPermissionDelete,
		Description: `:meta:subcategory:Authentication:This resource manages the object-based permissions for Netbox users, built into the application.

> Object-based permissions enable an administrator to grant users or groups the ability to perform an action on arbitrary subsets of objects in NetBox, rather than all objects of a certain type.
> For more information, see the [Netbox Object-Based Permissions Docs.](https://docs.netbox.dev/en/stable/administration/permissions/)`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the permission object.",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "The description of the permission object.",
				Optional:    true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Whether the permission object is enabled or not.",
				Optional:    true,
				Default:     true,
			},
			"object_types": {
				Type: schema.TypeSet,
				Description: "A list of object types that the permission object allows access to. Should be in a form " +
					"the API can accept. For example: `circuits.provider`, `dcim.inventoryitem`, etc.",
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of group IDs that have been assigned to this permission object.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"users": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of user IDs that have been assigned to this permission object.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"actions": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list actions that are allowed on the object types. Acceptable values are `view`, `add`, `change`, or `delete`.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"constraints": {
				Type: schema.TypeString,
				Description: "A JSON string of an arbitrary filter used to limit the granted action(s) to a specific subset of objects. " +
					"For more information on correct syntax, see https://docs.netbox.dev/en/stable/administration/permissions/#constraints ",
				Optional:     true,
				Default:      nil,
				ValidateFunc: validation.StringIsJSON,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxPermissionCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableObjectPermission{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = d.Get("description").(string)
	data.Enabled = d.Get("enabled").(bool)

	data.ObjectTypes = toStringList(d.Get("object_types"))
	data.Groups = toInt64List(d.Get("groups"))
	data.Users = toInt64List(d.Get("users"))
	data.Actions = toStringList(d.Get("actions"))

	var constraints interface{}
	c := d.Get("constraints").(string)
	if c == "" {
		data.Constraints = nil
	} else {
		err := json.Unmarshal([]byte(c), &constraints)
		if err != nil {
			return err
		}
		switch v := constraints.(type) {
		case []interface{}:
			data.Constraints = v
		case map[string]interface{}:
			data.Constraints = v

		}
	}

	params := users.NewUsersPermissionsCreateParams().WithData(&data)
	res, err := api.Users.UsersPermissionsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxPermissionRead(d, m)
}

func resourceNetboxPermissionRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersPermissionsReadParams().WithID(id)

	res, err := api.Users.UsersPermissionsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersPermissionsReadDefault); ok {
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
	d.Set("enabled", res.GetPayload().Enabled)
	d.Set("object_types", res.GetPayload().ObjectTypes)

	var groups []int
	for _, v := range res.GetPayload().Groups {
		groups = append(groups, int(v.ID))
	}
	d.Set("groups", groups)

	var users []int
	for _, v := range res.GetPayload().Users {
		users = append(users, int(v.ID))
	}
	d.Set("users", users)

	d.Set("actions", res.GetPayload().Actions)

	b, err := json.Marshal(res.GetPayload().Constraints)
	if err != nil {
		return err
	}
	d.Set("constraints", string(b))

	return nil
}

func resourceNetboxPermissionUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableObjectPermission{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = d.Get("description").(string)
	data.Enabled = d.Get("enabled").(bool)

	data.ObjectTypes = toStringList(d.Get("object_types"))
	data.Groups = toInt64List(d.Get("groups"))
	data.Users = toInt64List(d.Get("users"))
	data.Actions = toStringList(d.Get("actions"))

	var constraints interface{}
	c := d.Get("constraints").(string)
	if c == "" {
		data.Constraints = nil
	} else {
		err := json.Unmarshal([]byte(c), &constraints)
		if err != nil {
			return err
		}
		switch v := constraints.(type) {
		case []interface{}:
			data.Constraints = v
		case map[string]interface{}:
			data.Constraints = v
		}
	}
	params := users.NewUsersPermissionsUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersPermissionsUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxPermissionRead(d, m)
}

func resourceNetboxPermissionDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersPermissionsDeleteParams().WithID(id)
	_, err := api.Users.UsersPermissionsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*users.UsersPermissionsDeleteDefault); ok {
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
