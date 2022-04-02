package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceTypeCreate,
		Read:   resourceNetboxDeviceTypeRead,
		Update: resourceNetboxDeviceTypeUpdate,
		Delete: resourceNetboxDeviceTypeDelete,

		Schema: map[string]*schema.Schema{
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			"manufacturer_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxDeviceTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableDeviceType{}

	model := d.Get("model").(string)
	data.Model = &model

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to model if not given
	if !slugOk {
		data.Slug = strToPtr(model)
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	manufacturerIDValue, ok := d.GetOk("manufacturer_id")
	if ok {
		data.Manufacturer = int64ToPtr(int64(manufacturerIDValue.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := dcim.NewDcimDeviceTypesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDeviceTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceTypeRead(d, m)
}

func resourceNetboxDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceTypesReadParams().WithID(id)

	res, err := api.Dcim.DcimDeviceTypesRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimDeviceTypesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("model", res.GetPayload().Model)
	d.Set("slug", res.GetPayload().Slug)
	d.Set("manufacturer_id", res.GetPayload().Manufacturer.ID)
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	
	return nil
}

func resourceNetboxDeviceTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDeviceType{}

	model := d.Get("model").(string)
	data.Model = &model

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to model if not given
	if !slugOk {
		data.Slug = strToPtr(model)
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	manufacturerIDValue, ok := d.GetOk("manufacturer_id")
	if ok {
		data.Manufacturer = int64ToPtr(int64(manufacturerIDValue.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := dcim.NewDcimDeviceTypesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimDeviceTypesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDeviceTypeRead(d, m)
}

func resourceNetboxDeviceTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceTypesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDeviceTypesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
