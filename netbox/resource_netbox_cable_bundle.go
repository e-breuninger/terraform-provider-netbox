package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxCableBundle() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCableBundleCreate,
		Read:   resourceNetboxCableBundleRead,
		Update: resourceNetboxCableBundleUpdate,
		Delete: resourceNetboxCableBundleDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/cablebundle/):

> This model allows grouping of cables which are physically bundled together, e.g. running through the same conduit or in the same cable tray.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cable_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCableBundleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.CableBundle{
		Name:        strToPtr(d.Get("name").(string)),
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
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

	params := dcim.NewDcimCableBundlesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimCableBundlesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCableBundleRead(d, m)
}

func resourceNetboxCableBundleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimCableBundlesReadParams().WithID(id)

	res, err := api.Dcim.DcimCableBundlesRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*dcim.DcimCableBundlesReadDefault); ok {
			if errresp.Code() == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	bundle := res.GetPayload()

	d.Set("name", bundle.Name)
	d.Set("description", bundle.Description)
	d.Set("comments", bundle.Comments)
	d.Set("cable_count", bundle.CableCount)

	cf := getCustomFields(bundle.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, bundle.Tags)

	return nil
}

func resourceNetboxCableBundleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.CableBundle{
		Name:        strToPtr(d.Get("name").(string)),
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
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

	params := dcim.NewDcimCableBundlesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimCableBundlesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCableBundleRead(d, m)
}

func resourceNetboxCableBundleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimCableBundlesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimCableBundlesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
