package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxRackGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRackGroupCreate,
		Read:   resourceNetboxRackGroupRead,
		Update: resourceNetboxRackGroupUpdate,
		Delete: resourceNetboxRackGroupDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackgroup/):

> Racks can be grouped by site and/or location. This provides for a hierarchical organization of racks within NetBox, for example to represent the physical arrangement of racks within a facility.`,

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
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRackGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	data := &models.RackGroup{
		Name:        &name,
		Slug:        &slug,
		Description: getOptionalStr(d, "description", false),
		Comments:    getOptionalStr(d, "comments", false),
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimRackGroupsCreateParams().WithData(data)

	res, err := api.Dcim.DcimRackGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRackGroupRead(d, m)
}

func resourceNetboxRackGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimRackGroupsReadParams().WithID(id)

	res, err := api.Dcim.DcimRackGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackGroupsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	rackGroup := res.GetPayload()
	d.Set("name", rackGroup.Name)
	d.Set("slug", rackGroup.Slug)
	d.Set("description", rackGroup.Description)
	d.Set("comments", rackGroup.Comments)

	cf := getCustomFields(rackGroup.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, rackGroup.Tags)

	return nil
}

func resourceNetboxRackGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.RackGroup{}

	name := d.Get("name").(string)
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
	data.Comments = getOptionalStr(d, "comments", true)

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimRackGroupsPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimRackGroupsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxRackGroupRead(d, m)
}

func resourceNetboxRackGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackGroupsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRackGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
