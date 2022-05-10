package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTokenCreate,
		Read:   resourceNetboxTokenRead,
		Update: resourceNetboxTokenUpdate,
		Delete: resourceNetboxTokenDelete,

		Schema: map[string]*schema.Schema{
			"user_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				Sensitive:    true,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(40, 256),
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxTokenCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableToken{}

	userid := int64(d.Get("user_id").(int))

	key := d.Get("key").(string)

	data.User = &userid
	data.Key = key

	params := users.NewUsersTokensCreateParams().WithData(&data)
	res, err := api.Users.UsersTokensCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxTokenUpdate(d, m)
}

func resourceNetboxTokenRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersTokensReadParams().WithID(id)

	res, err := api.Users.UsersTokensRead(params, nil)
	if err != nil {
		errorcode := err.(*users.UsersTokensReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	if res.GetPayload().User != nil {
		d.Set("user_id", res.GetPayload().User.ID)
	}

	d.Set("key", res.GetPayload().Key)

	return nil
}

func resourceNetboxTokenUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableToken{}

	userid := int64(d.Get("user_id").(int))
	key := d.Get("key").(string)
	
	data.User = &userid
	data.Key = key

	params := users.NewUsersTokensUpdateParams().WithID(id).WithData(&data)
	_, err := api.Users.UsersTokensUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxTokenRead(d, m)
}

func resourceNetboxTokenDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := users.NewUsersTokensDeleteParams().WithID(id)
	_, err := api.Users.UsersTokensDelete(params, nil)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
