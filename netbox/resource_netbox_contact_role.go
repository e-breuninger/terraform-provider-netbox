package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxContactRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxContactRoleCreate,
		Read:   resourceNetboxContactRoleRead,
		Update: resourceNetboxContactRoleUpdate,
		Delete: resourceNetboxContactRoleDelete,

		Description: `:meta:subcategory:Tenancy:From the [official documentation](https://docs.netbox.dev/en/stable/features/contacts/#contactroles):

> A contact role defines the relationship of a contact to an assigned object. For example, you might define roles for administrative, operational, and emergency contacts`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxContactRoleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	data := &models.ContactRole{}

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Name = &name
	data.Tags = []*models.NestedTag{}

	params := tenancy.NewTenancyContactRolesCreateParams().WithData(data)

	res, err := api.Tenancy.TenancyContactRolesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxContactRoleRead(d, m)
}

func resourceNetboxContactRoleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyContactRolesReadParams().WithID(id)

	res, err := api.Tenancy.TenancyContactRolesRead(params, nil)

	if err != nil {
		errorcode := err.(*tenancy.TenancyContactRolesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	contactrole := res.GetPayload()
	d.Set("name", contactrole.Name)
	d.Set("slug", contactrole.Slug)

	return nil
}

func resourceNetboxContactRoleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.ContactRole{}

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Name = &name
	data.Tags = []*models.NestedTag{}

	params := tenancy.NewTenancyContactRolesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Tenancy.TenancyContactRolesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxContactRoleRead(d, m)
}

func resourceNetboxContactRoleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyContactRolesDeleteParams().WithID(id)

	_, err := api.Tenancy.TenancyContactRolesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
