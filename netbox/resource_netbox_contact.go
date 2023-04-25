package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxContact() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxContactCreate,
		Read:   resourceNetboxContactRead,
		Update: resourceNetboxContactUpdate,
		Delete: resourceNetboxContactDelete,

		Description: `:meta:subcategory:Tenancy:From the [official documentation](https://docs.netbox.dev/en/stable/features/contacts/#contacts_1):

> A contact should represent an individual or permanent point of contact. Each contact must define a name, and may optionally include a title, phone number, email address, and related details.
>
> Contacts are reused for assignments, so each unique contact must be created only once and can be assigned to any number of NetBox objects, and there is no limit to the number of assigned contacts an object may have. Most core objects in NetBox can have contacts assigned to them.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			tagsKey: tagsSchema,
			"group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxContactCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	phone := d.Get("phone").(string)
	email := d.Get("email").(string)
	group_id := int64(d.Get("group_id").(int))

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	data := &models.WritableContact{}

	data.Name = &name
	data.Tags = tags
	data.Phone = phone
	data.Email = strfmt.Email(email)

	if group_id != 0 {
		data.Group = &group_id
	}

	params := tenancy.NewTenancyContactsCreateParams().WithData(data)

	res, err := api.Tenancy.TenancyContactsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxContactRead(d, m)
}

func resourceNetboxContactRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyContactsReadParams().WithID(id)

	res, err := api.Tenancy.TenancyContactsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*tenancy.TenancyContactsReadDefault); ok {
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
	d.Set("phone", res.GetPayload().Phone)
	d.Set("email", res.GetPayload().Email)
	if res.GetPayload().Group != nil {
		d.Set("group_id", res.GetPayload().Group.ID)
	}

	return nil
}

func resourceNetboxContactUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableContact{}

	name := d.Get("name").(string)
	phone := d.Get("phone").(string)
	email := d.Get("email").(string)
	group_id := int64(d.Get("group_id").(int))

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	data.Name = &name
	data.Tags = tags
	data.Phone = phone
	data.Email = strfmt.Email(email)
	if group_id != 0 {
		data.Group = &group_id
	}

	params := tenancy.NewTenancyContactsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Tenancy.TenancyContactsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxContactRead(d, m)
}

func resourceNetboxContactDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyContactsDeleteParams().WithID(id)

	_, err := api.Tenancy.TenancyContactsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*tenancy.TenancyContactsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
