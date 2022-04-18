package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxManufacturer() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxManufacturerCreate,
		Read:   resourceNetboxManufacturerRead,
		Update: resourceNetboxManufacturerUpdate,
		Delete: resourceNetboxManufacturerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxManufacturerCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to name attribute if not given
	if !slugOk {
		slug = name
	} else {
		slug = slugValue.(string)
	}

	description := d.Get("description").(string)
	params := dcim.NewDcimManufacturersCreateParams().WithData(
		&models.Manufacturer{
			Name:        &name,
			Slug:        &slug,
			Description: description,
		},
	)

	res, err := api.Dcim.DcimManufacturersCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
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
	d.Set("description", res.GetPayload().Description)
	return nil
}

func resourceNetboxManufacturerUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.Manufacturer{}

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to name if not given
	if !slugOk {
		slug = name
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Name = &name
	data.Description = description

	params := dcim.NewDcimManufacturersUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimManufacturersUpdate(params, nil)
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
		return err
	}
	return nil
}
