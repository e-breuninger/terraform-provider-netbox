package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxRackTypeFormFactorOptions = []string{"2-post-frame", "4-post-frame", "4-post-cabinet", "wall-frame", "wall-frame-vertical", "wall-cabinet", "wall-cabinet-vertical"}
var resourceNetboxRackTypeWeightUnitOptions = []string{"kg", "g", "lb", "oz"}
var resourceNetboxRackTypeOuterUnitOptions = []string{"mm", "in"}
var resourceNetboxRackTypeWidthOptions = []int{10, 19, 21, 23}

func resourceNetboxRackType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxRackTypeCreate,
		Read:   resourceNetboxRackTypeRead,
		Update: resourceNetboxRackTypeUpdate,
		Delete: resourceNetboxRackTypeDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://netboxlabs.com/docs/netbox/en/stable/models/dcim/racktype/):

> A rack type defines the physical characteristics of a particular model of rack.`,

		Schema: map[string]*schema.Schema{
			"model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"form_factor": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxRackTypeFormFactorOptions, false),
				Description:  buildValidValueDescription(resourceNetboxRackTypeFormFactorOptions),
			},
			"width": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice(resourceNetboxRackTypeWidthOptions),
				Description:  "Valid values are `10`, `19`, `21` and `23`",
			},
			"u_height": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 100),
			},
			"starting_unit": {
				Type:     schema.TypeInt,
				Required: true,
			},
			tagsKey: tagsSchema,
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
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
				ValidateFunc: validation.StringInSlice(resourceNetboxRackTypeOuterUnitOptions, false),
				Description:  buildValidValueDescription(resourceNetboxRackTypeOuterUnitOptions),
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
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
				ValidateFunc: validation.StringInSlice(resourceNetboxRackTypeWeightUnitOptions, false),
				Description:  buildValidValueDescription(resourceNetboxRackTypeWeightUnitOptions),
			},
			"mounting_depth_mm": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validatePositiveInt16,
			},
			//			"tenant_id": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//			},
			//			"facility_id": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			"location_id": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//			},
			//			"role_id": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//			},
			//			"serial": {
			//				Type:         schema.TypeString,
			//				Optional:     true,
			//				ValidateFunc: validation.StringLenBetween(0, 50),
			//			},
			//			"asset_tag": {
			//				Type:         schema.TypeString,
			//				Optional:     true,
			//				ValidateFunc: validation.StringLenBetween(0, 50),
			//			},
			//			"type": {
			//				Type:         schema.TypeString,
			//				Optional:     true,
			//				ValidateFunc: validation.StringInSlice(resourceNetboxRackTypeTypeOptions, false),
			//				Description:  buildValidValueDescription(resourceNetboxRackTypeTypeOptions),
			//			},
			//			"desc_units": {
			//				Type:        schema.TypeBool,
			//				Optional:    true,
			//				Description: "If rack units are descending",
			//				Default:     false,
			//			},
			//			"outer_width": {
			//				Type:         schema.TypeInt,
			//				Optional:     true,
			//				ValidateFunc: validatePositiveInt16,
			//			},
			//			"outer_depth": {
			//				Type:         schema.TypeInt,
			//				Optional:     true,
			//				ValidateFunc: validatePositiveInt16,
			//			},
			//			"outer_unit": {
			//				Type:         schema.TypeString,
			//				Optional:     true,
			//				RequiredWith: []string{"outer_width", "outer_depth"},
			//				ValidateFunc: validation.StringInSlice(resourceNetboxRackTypeOuterUnitOptions, false),
			//				Description:  buildValidValueDescription(resourceNetboxRackTypeOuterUnitOptions),
			//			},
			//			"mounting_depth": {
			//				Type:         schema.TypeInt,
			//				Optional:     true,
			//				ValidateFunc: validatePositiveInt16,
			//			},
			//			"description": {
			//				Type:         schema.TypeString,
			//				Optional:     true,
			//				ValidateFunc: validation.StringLenBetween(0, 200),
			//			},
			//			"comments": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxRackTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	model := d.Get("model").(string)
	formFactor := d.Get("form_factor").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(model)
	} else {
		slug = slugValue.(string)
	}
	manufacturerID := int64(d.Get("manufacturer_id").(int))
	width := int64(d.Get("width").(int))
	uHeight := int64(d.Get("u_height").(int))

	data := models.WritableRackTypeRequest{
		Model:         &model,
		FormFactor:    &formFactor,
		Slug:          &slug,
		Manufacturer:  &manufacturerID,
		Width:         &width,
		UHeight:       uHeight,
		Description:   getOptionalStr(d, "description", false),
		OuterWidth:    getOptionalInt(d, "outer_width"),
		OuterDepth:    getOptionalInt(d, "outer_depth"),
		OuterUnit:     getOptionalStr(d, "outer_unit", false),
		Comments:      getOptionalStr(d, "comments", false),
		Weight:        getOptionalFloat(d, "weight"),
		MaxWeight:     getOptionalInt(d, "max_weight"),
		WeightUnit:    getOptionalStr(d, "weight_unit", false),
		MountingDepth: getOptionalInt(d, "mounting_depth_mm"),
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := dcim.NewDcimRackTypesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimRackTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxRackTypeRead(d, m)
}

func resourceNetboxRackTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackTypesReadParams().WithID(id)

	res, err := api.Dcim.DcimRackTypesRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackTypesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	rackType := res.GetPayload()

	d.Set("model", rackType.Model)
	d.Set("form_factor", rackType.FormFactor.Value)
	d.Set("starting_unit", rackType.StartingUnit)
	d.Set("manufacturer_id", rackType.Manufacturer.ID)

	if rackType.Width != nil {
		d.Set("width", rackType.Width.Value)
	} else {
		d.Set("width", nil)
	}

	d.Set("u_height", rackType.UHeight)
	api.readTags(d, res.GetPayload().Tags)
	d.Set("description", rackType.Description)
	d.Set("comments", rackType.Comments)

	d.Set("outer_width", rackType.OuterWidth)
	d.Set("outer_depth", rackType.OuterDepth)

	if rackType.OuterUnit != nil {
		d.Set("outer_unit", rackType.OuterUnit.Value)
	} else {
		d.Set("outer_unit", nil)
	}

	d.Set("weight", rackType.Weight)
	d.Set("max_weight", rackType.MaxWeight)

	if rackType.WeightUnit != nil {
		d.Set("weight_unit", rackType.WeightUnit.Value)
	} else {
		d.Set("weight_unit", nil)
	}

	d.Set("mounting_depth_mm", rackType.MountingDepth)

	return nil
}

func resourceNetboxRackTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	model := d.Get("model").(string)

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(model)
	} else {
		slug = slugValue.(string)
	}
	manufacturerID := int64(d.Get("manufacturer_id").(int))
	width := int64(d.Get("width").(int))
	uHeight := int64(d.Get("u_height").(int))

	data := models.WritableRackTypeRequest{
		Model:         &model,
		Slug:          &slug,
		Manufacturer:  &manufacturerID,
		Width:         &width,
		UHeight:       uHeight,
		Description:   getOptionalStr(d, "description", true),
		OuterWidth:    getOptionalInt(d, "outer_width"),
		OuterDepth:    getOptionalInt(d, "outer_depth"),
		OuterUnit:     getOptionalStr(d, "outer_unit", true),
		Comments:      getOptionalStr(d, "comments", true),
		Weight:        getOptionalFloat(d, "weight"),
		MaxWeight:     getOptionalInt(d, "max_weight"),
		WeightUnit:    getOptionalStr(d, "weight_unit", true),
		MountingDepth: getOptionalInt(d, "mounting_depth_mm"),
	}

	params := dcim.NewDcimRackTypesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimRackTypesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxRackTypeRead(d, m)
}

func resourceNetboxRackTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimRackTypesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimRackTypesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimRackTypesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
