package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxPowerFeed() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPowerFeedCreate,
		Read:   resourceNetboxPowerFeedRead,
		Update: resourceNetboxPowerFeedUpdate,
		Delete: resourceNetboxPowerFeedDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerfeed/):

> A power feed represents the distribution of power from a power panel to a particular device, typically a power distribution unit (PDU). The power port (inlet) on a device can be connected via a cable to a power feed. A power feed may optionally be assigned to a rack to allow more easily tracking the distribution of power among racks.`,

		Schema: map[string]*schema.Schema{
			"power_panel_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [offline, active, planned, failed]",
				ValidateFunc: validation.StringInSlice([]string{"offline", "active", "planned", "failed"}, false),
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [primary, redundant]",
				ValidateFunc: validation.StringInSlice([]string{"primary", "redundant"}, false),
			},
			"supply": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [single-phase, three-phase]",
				ValidateFunc: validation.StringInSlice([]string{"single-phase", "three-phase"}, false),
			},
			"phase": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [primary, redundant]",
				ValidateFunc: validation.StringInSlice([]string{"primary", "redundant"}, false),
			},
			"voltage": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"amperage": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_percent_utilization": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 100),
			},
			"rack_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mark_connected": {
				Type:    schema.TypeBool,
				Default: false,
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

func resourceNetboxPowerFeedCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritablePowerFeed{
		PowerPanel:     int64ToPtr(int64(d.Get("power_panel_id").(int))),
		Name:           strToPtr(d.Get("name").(string)),
		Status:         d.Get("status").(string),
		Type:           d.Get("type").(string),
		Supply:         d.Get("supply").(string),
		Phase:          d.Get("phase").(string),
		Voltage:        int64ToPtr(int64(d.Get("voltage").(int))),
		Amperage:       int64(d.Get("amperage").(int)),
		MaxUtilization: int64(d.Get("max_percent_utilization").(int)),
		Rack:           getOptionalInt(d, "rack_id"),
		MarkConnected:  d.Get("mark_connected").(bool),
		Description:    getOptionalStr(d, "description", false),
		Comments:       getOptionalStr(d, "comments", false),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerFeedsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimPowerFeedsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxPowerFeedRead(d, m)
}

func resourceNetboxPowerFeedRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerFeedsReadParams().WithID(id)

	res, err := api.Dcim.DcimPowerFeedsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimPowerFeedsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	powerFeed := res.GetPayload()

	if powerFeed.PowerPanel != nil {
		d.Set("power_panel_id", powerFeed.PowerPanel.ID)
	} else {
		d.Set("power_panel_id", nil)
	}

	d.Set("name", powerFeed.Name)

	if powerFeed.Status != nil {
		d.Set("status", powerFeed.Status.Value)
	} else {
		d.Set("status", nil)
	}

	if powerFeed.Type != nil {
		d.Set("type", powerFeed.Type.Value)
	} else {
		d.Set("type", nil)
	}

	if powerFeed.Supply != nil {
		d.Set("supply", powerFeed.Supply.Value)
	} else {
		d.Set("supply", nil)
	}

	if powerFeed.Phase != nil {
		d.Set("phase", powerFeed.Phase.Value)
	} else {
		d.Set("phase", nil)
	}

	d.Set("voltage", powerFeed.Voltage)
	d.Set("amperage", powerFeed.Amperage)
	d.Set("max_percent_utilization", powerFeed.MaxUtilization)

	if powerFeed.Rack != nil {
		d.Set("rack_id", powerFeed.Rack.ID)
	} else {
		d.Set("rack_id", nil)
	}

	d.Set("mark_connected", powerFeed.MarkConnected)
	d.Set("description", powerFeed.Description)
	d.Set("comments", powerFeed.Comments)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxPowerFeedUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritablePowerFeed{
		PowerPanel:     int64ToPtr(int64(d.Get("power_panel_id").(int))),
		Name:           strToPtr(d.Get("name").(string)),
		Status:         d.Get("status").(string),
		Type:           d.Get("type").(string),
		Supply:         d.Get("supply").(string),
		Phase:          d.Get("phase").(string),
		Voltage:        int64ToPtr(int64(d.Get("voltage").(int))),
		Amperage:       int64(d.Get("amperage").(int)),
		MaxUtilization: int64(d.Get("max_percent_utilization").(int)),
		Rack:           getOptionalInt(d, "rack_id"),
		MarkConnected:  d.Get("mark_connected").(bool),
		Description:    getOptionalStr(d, "description", false),
		Comments:       getOptionalStr(d, "comments", false),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimPowerFeedsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimPowerFeedsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxPowerFeedRead(d, m)
}

func resourceNetboxPowerFeedDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimPowerFeedsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimPowerFeedsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
