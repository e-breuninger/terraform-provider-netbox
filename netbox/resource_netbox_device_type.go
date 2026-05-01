package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxDeviceTypeAirflowOptions = []string{"front-to-rear", "rear-to-front", "left-to-right", "right-to-left", "side-to-rear", "passive", "mixed"}
var resourceNetboxDeviceTypeWeightUnitOptions = []string{"kg", "g", "lb", "oz"}

func resourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDeviceTypeCreate,
		Read:   resourceNetboxDeviceTypeRead,
		Update: resourceNetboxDeviceTypeUpdate,
		Delete: resourceNetboxDeviceTypeDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/device-types/#device-types_1):

> A device type represents a particular make and model of hardware that exists in the real world. Device types define the physical attributes of a device (rack height and depth) and its individual components (console, power, network interfaces, and so on).`,

		Schema: map[string]*schema.Schema{
			"model": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Marketing name of the model.",
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
				Description:  "URL-safe identifier for the device type. Defaults to a slugified `model` if not given.",
			},
			"manufacturer_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ID of the `netbox_manufacturer` this device type belongs to.",
			},
			"part_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Manufacturer part number / SKU.",
			},
			"u_height": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Default:     "1.0",
				Description: "Rack height in U. Defaults to `1.0`.",
			},
			"is_full_depth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the device occupies the full rack depth.",
			},
			"subdevice_role": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "For chassis-style devices: `parent` for the chassis, `child` for the modules. Leave unset for a single-piece device.",
			},
			"airflow": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxDeviceTypeAirflowOptions, false),
				Description:  "Default airflow direction. " + buildValidValueDescription(resourceNetboxDeviceTypeAirflowOptions),
			},
			"weight": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "Numeric weight of the device. Pair with `weight_unit`.",
			},
			"weight_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"weight"},
				ValidateFunc: validation.StringInSlice(resourceNetboxDeviceTypeWeightUnitOptions, false),
				Description:  "Unit for `weight`. " + buildValidValueDescription(resourceNetboxDeviceTypeWeightUnitOptions),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
				Description:  "Short description (max 200 chars).",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Free-form Markdown comments.",
			},
			"default_platform_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the `netbox_platform` to default new devices of this type to.",
			},
			"exclude_from_utilization": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set, devices of this type are excluded from rack utilization calculations.",
			},
			customFieldsKey: customFieldsSchema,
			tagsKey:         tagsSchema,

			// Nested template lifecycle. Each block is hash-keyed by `name`.
			// See netbox/device_type_templates.go for the full per-type schema
			// and the orchestrator that runs them in dependency order.
			powerPortTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        powerPortTemplateSchema(),
				Set:         templateNameHash,
				Description: "Power port templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/powerporttemplate/).",
			},
			interfaceTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        interfaceTemplateSchema(),
				Set:         templateNameHash,
				Description: "Network interface templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/interfacetemplate/).",
			},
			consolePortTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        consolePortTemplateSchema(),
				Set:         templateNameHash,
				Description: "Console port templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/consoleporttemplate/).",
			},
			consoleServerPortTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        consoleServerPortTemplateSchema(),
				Set:         templateNameHash,
				Description: "Console server port templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/consoleserverporttemplate/).",
			},
			rearPortTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        rearPortTemplateSchema(),
				Set:         templateNameHash,
				Description: "Rear port templates instantiated on every device of this type. Front ports reference these by name. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/rearporttemplate/).",
			},
			deviceBayTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        deviceBayTemplateSchema(),
				Set:         templateNameHash,
				Description: "Device bay templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/devicebaytemplate/).",
			},
			moduleBayTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        moduleBayTemplateSchema(),
				Set:         templateNameHash,
				Description: "Module bay templates instantiated on every device of this type. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/modulebaytemplate/).",
			},
			powerOutletTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        powerOutletTemplateSchema(),
				Set:         templateNameHash,
				Description: "Power outlet templates instantiated on every device of this type. May reference a sibling `power_port_templates` block by name. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/poweroutlettemplate/).",
			},
			frontPortTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        frontPortTemplateSchema(),
				Set:         templateNameHash,
				Description: "Front port templates instantiated on every device of this type. Each must reference a sibling `rear_port_templates` block by name. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/frontporttemplate/).",
			},
			inventoryItemTemplatesKey: {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        inventoryItemTemplateSchema(),
				Set:         templateNameHash,
				Description: "Inventory item templates instantiated on every device of this type. Supports a parent tree via the `parent` field and an optional polymorphic FK via `component_type`/`component_id`. See [the NetBox docs](https://docs.netbox.dev/en/stable/models/dcim/inventoryitemtemplate/).",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableDeviceType{}

	model := d.Get("model").(string)
	data.Model = &model

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(model))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	manufacturerIDValue, ok := d.GetOk("manufacturer_id")
	if ok {
		data.Manufacturer = int64ToPtr(int64(manufacturerIDValue.(int)))
	}

	if partNo, ok := d.GetOk("part_number"); ok {
		data.PartNumber = partNo.(string)
	}

	//Needed to account for 0 u_height values
	uHeightValue := d.Get("u_height")
	data.UHeight = float64ToPtr(uHeightValue.(float64))

	if isFullDepthValue, ok := d.GetOk("is_full_depth"); ok {
		data.IsFullDepth = isFullDepthValue.(bool)
	}

	if subdeviceRoleValue, ok := d.GetOk("subdevice_role"); ok {
		data.SubdeviceRole = subdeviceRoleValue.(string)
	}

	data.Airflow = getOptionalStr(d, "airflow", false)
	data.Weight = getOptionalFloat(d, "weight")
	data.WeightUnit = getOptionalStr(d, "weight_unit", false)
	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.DefaultPlatform = getOptionalInt(d, "default_platform_id")
	data.ExcludeFromUtilization = d.Get("exclude_from_utilization").(bool)

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = getCustomFields(cf)
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := dcim.NewDcimDeviceTypesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDeviceTypesCreate(params, nil)
	if err != nil {
		return err
	}

	deviceTypeID := res.GetPayload().ID
	d.SetId(strconv.FormatInt(deviceTypeID, 10))

	if err := syncDeviceTypeTemplates(d, api, deviceTypeID); err != nil {
		return err
	}

	return resourceNetboxDeviceTypeRead(d, m)
}

func resourceNetboxDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceTypesReadParams().WithID(id)

	res, err := api.Dcim.DcimDeviceTypesRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*dcim.DcimDeviceTypesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	deviceType := res.GetPayload()
	d.Set("model", deviceType.Model)
	d.Set("slug", deviceType.Slug)
	d.Set("manufacturer_id", deviceType.Manufacturer.ID)
	d.Set("part_number", deviceType.PartNumber)
	d.Set("u_height", deviceType.UHeight)
	d.Set("is_full_depth", deviceType.IsFullDepth)
	if deviceType.SubdeviceRole != nil && deviceType.SubdeviceRole.Value != nil {
		d.Set("subdevice_role", *deviceType.SubdeviceRole.Value)
	} else {
		d.Set("subdevice_role", "")
	}

	if deviceType.Airflow != nil && deviceType.Airflow.Value != nil {
		d.Set("airflow", *deviceType.Airflow.Value)
	} else {
		d.Set("airflow", "")
	}
	d.Set("weight", deviceType.Weight)
	if deviceType.WeightUnit != nil && deviceType.WeightUnit.Value != nil {
		d.Set("weight_unit", *deviceType.WeightUnit.Value)
	} else {
		d.Set("weight_unit", "")
	}
	d.Set("description", deviceType.Description)
	d.Set("comments", deviceType.Comments)
	if deviceType.DefaultPlatform != nil {
		d.Set("default_platform_id", deviceType.DefaultPlatform.ID)
	} else {
		d.Set("default_platform_id", nil)
	}
	d.Set("exclude_from_utilization", deviceType.ExcludeFromUtilization)
	d.Set(customFieldsKey, readCustomFields(deviceType.CustomFields))

	api.readTags(d, deviceType.Tags)

	if err := readDeviceTypeTemplates(d, api, id); err != nil {
		return err
	}

	return nil
}

func resourceNetboxDeviceTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDeviceType{}

	model := d.Get("model").(string)
	data.Model = &model

	slugValue, slugOk := d.GetOk("slug")
	// Default slug to generated slug if not given
	if !slugOk {
		data.Slug = strToPtr(getSlug(model))
	} else {
		data.Slug = strToPtr(slugValue.(string))
	}

	manufacturerIDValue, ok := d.GetOk("manufacturer_id")
	if ok {
		data.Manufacturer = int64ToPtr(int64(manufacturerIDValue.(int)))
	}

	if partNo, ok := d.GetOk("part_number"); ok {
		data.PartNumber = partNo.(string)
	}

	uHeightValue := d.Get("u_height")
	data.UHeight = float64ToPtr(uHeightValue.(float64))

	if isFullDepthValue, ok := d.GetOk("is_full_depth"); ok {
		data.IsFullDepth = isFullDepthValue.(bool)
	}

	if subdeviceRoleValue, ok := d.GetOk("subdevice_role"); ok {
		data.SubdeviceRole = subdeviceRoleValue.(string)
	}

	data.Airflow = getOptionalStr(d, "airflow", false)
	data.Weight = getOptionalFloat(d, "weight")
	data.WeightUnit = getOptionalStr(d, "weight_unit", false)
	data.Description = getOptionalStr(d, "description", true)
	data.Comments = getOptionalStr(d, "comments", true)
	data.DefaultPlatform = getOptionalInt(d, "default_platform_id")
	data.ExcludeFromUtilization = d.Get("exclude_from_utilization").(bool)

	// CustomFields needs the diff-based clearing helper so keys removed from
	// HCL get explicit JSON null; see customFieldsForUpdate for full rationale.
	data.CustomFields = customFieldsForUpdate(d)

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := dcim.NewDcimDeviceTypesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimDeviceTypesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	if err := syncDeviceTypeTemplates(d, api, id); err != nil {
		return err
	}

	return resourceNetboxDeviceTypeRead(d, m)
}

func resourceNetboxDeviceTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDeviceTypesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDeviceTypesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDeviceTypesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
