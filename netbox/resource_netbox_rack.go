package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxRack() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRackCreate,
		Read:   resourceNetboxRackRead,
		Update: resourceNetboxRackUpdate,
		Delete: resourceNetboxRackDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rack/):

> The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.

Rack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.

Each rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, "M204.313") whereas internally you refer to is simply as "R113." A unique serial number and asset tag may also be associated with each rack.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [reserved available planned active deprecated]",
				ValidateFunc: validation.StringInSlice([]string{"reserved", "available", "planned", "active", "deprecated"}, false),
			},
			"width": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "One of [10, 19, 21, 23]",
				ValidateFunc: validation.IntInSlice([]int{10, 19, 21, 23}),
			},
			"u_height": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 100),
			},
			tagsKey: tagsSchema,
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"facility_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"serial": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"asset_tag": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "One of [2-post-frame 4-post-frame 4-post-cabinet wall-frame wall-frame-vertical wall-cabinet wall-cabinet-vertical]",
				ValidateFunc: validation.StringInSlice([]string{"2-post-frame", "4-post-frame", "4-post-cabinet", "wall-frame", "wall-frame-vertical", "wall-cabinet", "wall-cabinet-vertical"}, false),
			},
			"weight": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_weight": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validatePositiveInt32,
			},
			"weight_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"weight", "max_weight"},
				Description:  "One of [kg, g, lb, oz]",
				ValidateFunc: validation.StringInSlice([]string{"kg", "g", "lb", "oz"}, false),
			},
			"desc_units": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If rack units are descending",
				Default:     false,
			},
			"outer_width": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validatePositiveInt16,
			},
			"outer_depth": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validatePositiveInt16,
			},
			"outer_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"outer_width", "outer_depth"},
				Description:  "One of [mm, in]",
				ValidateFunc: validation.StringInSlice([]string{"mm", "in"}, false),
			},
			"mounting_depth": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validatePositiveInt16,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRackCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	siteID := int64(d.Get("site_id").(int))
	status := d.Get("status").(string)
	width := int64(d.Get("width").(int))
	uHeight := int64(d.Get("u_height").(int))

	data := models.WritableRack{
		Name:    &name,
		Site:    &siteID,
		Status:  status,
		Width:   width,
		UHeight: uHeight,
	}

	data.Tenant = getOptionalInt(d, "tenant_id")
	if facilityId := getOptionalStr(d, "facility_id", false); facilityId != "" {
		data.FacilityID = strToPtr(facilityId)
	}
	data.Location = getOptionalInt(d, "location_id")
	data.Role = getOptionalInt(d, "role_id")
	data.Serial = getOptionalStr(d, "serial", false)
	if assetTag := getOptionalStr(d, "asset_tag", false); assetTag != "" {
		data.AssetTag = &assetTag
	}
	data.Type = getOptionalStr(d, "type", false)
	data.Weight = getOptionalFloat(d, "weight")
	data.MaxWeight = getOptionalInt(d, "max_weight")
	data.WeightUnit = getOptionalStr(d, "weight_unit", false)

	if descUnits, ok := d.GetOk("desc_units"); ok {
		data.DescUnits = descUnits.(bool)
	}

	data.OuterWidth = getOptionalInt(d, "outer_width")
	data.OuterDepth = getOptionalInt(d, "outer_depth")
	data.OuterUnit = getOptionalStr(d, "outer_unit", false)
	data.MountingDepth = getOptionalInt(d, "mounting_depth")
	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimRacksCreateParams().WithData(&data)

	res, err := api.Dcim.DcimRacksCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRackRead(d, m)
}

func resourceNetboxRackRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRacksReadParams().WithID(id)

	res, err := api.Dcim.DcimRacksRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*dcim.DcimRacksReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	rack := res.GetPayload()

	d.Set("name", rack.Name)

	if rack.Site != nil {
		d.Set("site_id", rack.Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if rack.Status != nil {
		d.Set("status", rack.Status.Value)
	} else {
		d.Set("status", nil)
	}

	if rack.Width != nil {
		d.Set("width", rack.Width.Value)
	} else {
		d.Set("width", nil)
	}

	d.Set("u_height", rack.UHeight)

	if rack.Tenant != nil {
		d.Set("tenant_id", rack.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	d.Set("facility_id", rack.FacilityID)

	if rack.Location != nil {
		d.Set("location_id", rack.Location.ID)
	} else {
		d.Set("location_id", nil)
	}

	if rack.Role != nil {
		d.Set("role_id", rack.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	d.Set("serial", rack.Serial)
	d.Set("asset_tag", rack.AssetTag)

	if rack.Type != nil {
		d.Set("type", rack.Type.Value)
	} else {
		d.Set("type", nil)
	}

	d.Set("weight", rack.Weight)
	d.Set("max_weight", rack.MaxWeight)

	if rack.WeightUnit != nil {
		d.Set("weight_unit", rack.WeightUnit.Value)
	} else {
		d.Set("weight_unit", nil)
	}

	d.Set("desc_units", rack.DescUnits)
	d.Set("outer_width", rack.OuterWidth)
	d.Set("outer_depth", rack.OuterDepth)

	if rack.OuterUnit != nil {
		d.Set("outer_unit", rack.OuterUnit.Value)
	} else {
		d.Set("outer_unit", nil)
	}

	d.Set("mounting_depth", rack.MountingDepth)
	d.Set("description", rack.Description)
	d.Set("comments", rack.Comments)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxRackUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	siteID := int64(d.Get("site_id").(int))
	status := d.Get("status").(string)
	width := int64(d.Get("width").(int))
	uHeight := int64(d.Get("u_height").(int))

	data := models.WritableRack{
		Name:    &name,
		Site:    &siteID,
		Status:  status,
		Width:   width,
		UHeight: uHeight,
	}

	data.Tenant = getOptionalInt(d, "tenant_id")

	if facilityId := getOptionalStr(d, "facility_id", false); facilityId != "" {
		data.FacilityID = strToPtr(facilityId)
	}

	data.Location = getOptionalInt(d, "location_id")
	data.Role = getOptionalInt(d, "role_id")
	data.Serial = getOptionalStr(d, "serial", true)
	if assetTag := getOptionalStr(d, "asset_tag", false); assetTag != "" {
		data.AssetTag = &assetTag
	}
	data.Type = getOptionalStr(d, "type", false)
	data.Weight = getOptionalFloat(d, "weight")
	data.MaxWeight = getOptionalInt(d, "max_weight")
	data.WeightUnit = getOptionalStr(d, "weight_unit", false)

	if descUnits, ok := d.GetOk("desc_units"); ok {
		data.DescUnits = descUnits.(bool)
	}

	data.OuterWidth = getOptionalInt(d, "outer_width")
	data.OuterDepth = getOptionalInt(d, "outer_depth")
	data.OuterUnit = getOptionalStr(d, "outer_unit", false)
	data.MountingDepth = getOptionalInt(d, "mounting_depth")
	data.Description = getOptionalStr(d, "description", true)
	data.Comments = getOptionalStr(d, "comments", true)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := dcim.NewDcimRacksPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimRacksPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxRackRead(d, m)
}

func resourceNetboxRackDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRacksDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRacksDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRacksDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
