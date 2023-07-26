package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxPowerPanel() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPowerPanelCreate,
		Read:   resourceNetboxPowerPanelRead,
		Update: resourceNetboxPowerPanelUpdate,
		Delete: resourceNetboxPowerPanelDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerpanel/):

> A power panel represents the origin point in NetBox for electrical power being disseminated by one or more power feeds. In a data center environment, one power panel often serves a group of racks, with an individual power feed extending to each rack, though this is not always the case. It is common to have two sets of panels and feeds arranged in parallel to provide redundant power to each rack.`,

		Schema: map[string]*schema.Schema{
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceNetboxPowerPanelCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritablePowerPanel{
		Site:        int64ToPtr(d.Get("site_id").(int64)),
		Name:        strToPtr(d.Get("name").(string)),
		Location:    getOptionalInt(d, "location_id"),
		Description: getOptionalStr(d, "description", false),
		Comments:    getOptionalStr(d, "comments", false),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerPanelsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimPowerPanelsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxPowerPanelRead(d, m)
}

func resourceNetboxPowerPanelRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerPanelsReadParams().WithID(id)

	res, err := api.Dcim.DcimPowerPanelsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimPowerPanelsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	powerPanel := res.GetPayload()

	if powerPanel.Site != nil {
		d.Set("site_id", powerPanel.Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	d.Set("name", powerPanel.Name)

	if powerPanel.Location != nil {
		d.Set("location_id", powerPanel.Location.ID)
	} else {
		d.Set("location_id", nil)
	}

	d.Set("description", powerPanel.Description)
	d.Set("comments", powerPanel.Comments)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxPowerPanelUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritablePowerPanel{
		Site:        int64ToPtr(d.Get("site_id").(int64)),
		Name:        strToPtr(d.Get("name").(string)),
		Location:    getOptionalInt(d, "location_id"),
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerPanelsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimPowerPanelsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxPowerPanelRead(d, m)
}

func resourceNetboxPowerPanelDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerPanelsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimPowerPanelsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
