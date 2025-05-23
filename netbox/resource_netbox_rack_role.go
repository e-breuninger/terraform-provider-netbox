package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxRackRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRackRoleCreate,
		Read:   resourceNetboxRackRoleRead,
		Update: resourceNetboxRackRoleUpdate,
		Delete: resourceNetboxRackRoleDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackrole/):

> Each rack can optionally be assigned a user-defined functional role. For example, you might designate a rack for compute or storage resources, or to house colocated customer devices.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"color_hex": {
				Type:     schema.TypeString,
				Required: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRackRoleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	color := d.Get("color_hex").(string)
	description := getOptionalStr(d, "description", false)

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))

	params := dcim.NewDcimRackRolesCreateParams().WithData(
		&models.RackRole{
			Name:        &name,
			Slug:        &slug,
			Color:       color,
			Description: description,
			Tags:        tags,
		},
	)

	res, err := api.Dcim.DcimRackRolesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRackRoleRead(d, m)
}

func resourceNetboxRackRoleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackRolesReadParams().WithID(id)

	res, err := api.Dcim.DcimRackRolesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackRolesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	rackRole := res.GetPayload()

	d.Set("name", rackRole.Name)
	d.Set("slug", rackRole.Slug)
	d.Set("description", rackRole.Description)
	d.Set("color_hex", rackRole.Color)
	api.readTags(d, res.GetPayload().Tags)
	return nil
}

func resourceNetboxRackRoleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.RackRole{}

	name := d.Get("name").(string)
	color := d.Get("color_hex").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string

	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}

	data.Slug = &slug
	data.Name = &name
	data.Description = getOptionalStr(d, "description", true)
	data.Color = color

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	data.Tags = tags

	params := dcim.NewDcimRackRolesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimRackRolesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxRackRoleRead(d, m)
}

func resourceNetboxRackRoleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackRolesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRackRolesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackRolesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
