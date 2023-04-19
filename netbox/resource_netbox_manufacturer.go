package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxManufacturer() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxManufacturerCreate,
		Read:   resourceNetboxManufacturerRead,
		Update: resourceNetboxManufacturerUpdate,
		Delete: resourceNetboxManufacturerDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device-types/#manufacturers):

> A manufacturer represents the "make" of a device; e.g. Cisco or Dell. Each device type must be assigned to a manufacturer. (Inventory items and platforms may also be associated with manufacturers.) Each manufacturer must have a unique name and may have a description assigned to it.`,

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

func resourceNetboxManufacturerCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.Manufacturer{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Tags = []*models.NestedTag{}

	params := dcim.NewDcimManufacturersCreateParams().WithData(&data)

	res, err := api.Dcim.DcimManufacturersCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxManufacturerRead(d, m)
}

func resourceNetboxManufacturerRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimManufacturersReadParams().WithID(id)

	res, err := api.Dcim.DcimManufacturersRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimManufacturersReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)

	return nil
}

func resourceNetboxManufacturerUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.Manufacturer{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(name))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	data.Tags = []*models.NestedTag{}

	params := dcim.NewDcimManufacturersPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimManufacturersPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxManufacturerRead(d, m)
}

func resourceNetboxManufacturerDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimManufacturersDeleteParams().WithID(id)

	_, err := api.Dcim.DcimManufacturersDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimManufacturersDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
