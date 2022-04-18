package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceTypeCreate,
		Read:   resourceNetboxDeviceTypeRead,
		Update: resourceNetboxDeviceTypeUpdate,
		Delete: resourceNetboxDeviceTypeDelete,

		Schema: map[string]*schema.Schema{
			"manufacturer_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"u_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"is_full_depth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	manufacturer_id := int64(d.Get("manufacturer_id").(int))
	model := d.Get("model").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to model if not given
	if !slugOk {
		slug = model
	} else {
		slug = slugValue.(string)
	}

	u_height := int64(d.Get("u_height").(int))
	is_full_depth := d.Get("is_full_depth").(bool)
	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := dcim.NewDcimDeviceTypesCreateParams().WithData(
		&models.WritableDeviceType{
			Manufacturer: &manufacturer_id,
			Model:        &model,
			Slug:         &slug,
			UHeight:      &u_height,
			IsFullDepth:  is_full_depth,
			Tags:         tags,
		},
	)

	res, err := api.Dcim.DcimDeviceTypesCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
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

	d.Set("manufacturer_id", res.GetPayload().Manufacturer.ID)
	d.Set("model", res.GetPayload().Model)
	d.Set("slug", res.GetPayload().Slug)
	d.Set("u_height", res.GetPayload().UHeight)
	d.Set("is_full_depth", res.GetPayload().IsFullDepth)
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxDeviceTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDeviceType{}

	manufacturer_id := int64(d.Get("manufacturer_id").(int))
	model := d.Get("model").(string)
	u_height := int64(d.Get("u_height").(int))
	is_full_depth := d.Get("is_full_depth").(bool)

	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to model if not given
	if !slugOk {
		slug = model
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Model = &model
	data.Manufacturer = &manufacturer_id
	data.UHeight = &u_height
	data.IsFullDepth = is_full_depth
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
