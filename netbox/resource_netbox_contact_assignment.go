package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxContactAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxContactAssignmentCreate,
		Read:   resourceNetboxContactAssignmentRead,
		Update: resourceNetboxContactAssignmentUpdate,
		Delete: resourceNetboxContactAssignmentDelete,

		Description: `:meta:subcategory:Tenancy:From the [official documentation](https://docs.netbox.dev/en/stable/features/contacts#contactassignments_1):

> Much like tenancy, contact assignment enables you to track ownership of resources modeled in NetBox.`,

		Schema: map[string]*schema.Schema{
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"contact_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxContactAssignmentCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	content_type := d.Get("content_type").(string)
	object_id := int64(d.Get("object_id").(int))
	contact_id := int64(d.Get("contact_id").(int))
	role_id := int64(d.Get("role_id").(int))

	data := &models.WritableContactAssignment{}

	data.ContentType = &content_type
	data.ObjectID = &object_id
	data.Contact = &contact_id
	data.Role = &role_id

	params := tenancy.NewTenancyContactAssignmentsCreateParams().WithData(data)

	res, err := api.Tenancy.TenancyContactAssignmentsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxContactAssignmentRead(d, m)
}

func resourceNetboxContactAssignmentRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyContactAssignmentsReadParams().WithID(id)

	res, err := api.Tenancy.TenancyContactAssignmentsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*tenancy.TenancyContactAssignmentsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("content_type", res.GetPayload().ContentType)

	if res.GetPayload().ObjectID != nil {
		d.Set("object_id", res.GetPayload().ObjectID)
	}
	if res.GetPayload().Contact != nil {
		d.Set("contact_id", res.GetPayload().Contact.ID)
	}
	if res.GetPayload().Role != nil {
		d.Set("role_id", res.GetPayload().Role.ID)
	}

	return nil
}

func resourceNetboxContactAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableContactAssignment{}

	content_type := d.Get("content_type").(string)
	object_id := int64(d.Get("object_id").(int))
	contact_id := int64(d.Get("contact_id").(int))
	role_id := int64(d.Get("role_id").(int))

	data.ContentType = &content_type
	if object_id != 0 {
		data.ObjectID = &object_id
	}
	if contact_id != 0 {
		data.Contact = &contact_id
	}
	if role_id != 0 {
		data.Role = &role_id
	}

	params := tenancy.NewTenancyContactAssignmentsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Tenancy.TenancyContactAssignmentsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxContactAssignmentRead(d, m)
}

func resourceNetboxContactAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := tenancy.NewTenancyContactAssignmentsDeleteParams().WithID(id)

	_, err := api.Tenancy.TenancyContactAssignmentsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*tenancy.TenancyContactAssignmentsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
