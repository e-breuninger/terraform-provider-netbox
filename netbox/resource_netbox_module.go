package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxModule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxModuleCreate,
		Read:   resourceNetboxModuleRead,
		Update: resourceNetboxModuleUpdate,
		Delete: resourceNetboxModuleDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/module/):

> A module is a field-replaceable hardware component installed within a device which houses its own child components. The most common example is a chassis-based router or switch.

Similar to devices, modules are instantiated from module types, and any components associated with the module type are automatically instantiated on the new model. Each module must be installed within a module bay on a device, and each module bay may have only one module installed in it.`,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"module_bay_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"module_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [offline, active, planned, staged, failed, decommissioning]",
				ValidateFunc: validation.StringInSlice([]string{"offline", "active", "planned", "staged", "failed", "decommissioning"}, false),
			},
			"serial": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
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

func resourceNetboxModuleCreate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	data := models.WritableModule{
		Device:      int64ToPtr(int64(d.Get("device_id").(int))),
		ModuleBay:   int64ToPtr(int64(d.Get("module_bay_id").(int))),
		ModuleType:  int64ToPtr(int64(d.Get("module_type_id").(int))),
		Status:      d.Get("status").(string),
		Serial:      getOptionalStr(d, "serial", false),
		Description: getOptionalStr(d, "description", false),
		Comments:    getOptionalStr(d, "comments", false),
	}

	if assetTag := getOptionalStr(d, "asset_tag", false); assetTag != "" {
		data.AssetTag = &assetTag
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimModulesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimModulesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxModuleRead(d, m)
}

func resourceNetboxModuleRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModulesReadParams().WithID(id)

	res, err := api.Dcim.DcimModulesRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimModulesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	module := res.GetPayload()

	if module.Device != nil {
		d.Set("device_id", module.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	if module.ModuleBay != nil {
		d.Set("module_bay_id", module.ModuleBay.ID)
	} else {
		d.Set("module_bay_id", nil)
	}

	if module.ModuleType != nil {
		d.Set("module_type_id", module.ModuleType.ID)
	} else {
		d.Set("module_type_id", nil)
	}

	if module.Status != nil {
		d.Set("status", module.Status.Value)
	} else {
		d.Set("status", nil)
	}

	d.Set("serial", module.Serial)
	d.Set("asset_tag", module.AssetTag)
	d.Set("description", module.Description)
	d.Set("comments", module.Comments)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	state.readTags(d, res.GetPayload().Tags)

	return nil
}

func resourceNetboxModuleUpdate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableModule{
		Device:      int64ToPtr(int64(d.Get("device_id").(int))),
		ModuleBay:   int64ToPtr(int64(d.Get("module_bay_id").(int))),
		ModuleType:  int64ToPtr(int64(d.Get("module_type_id").(int))),
		Status:      d.Get("status").(string),
		Serial:      getOptionalStr(d, "serial", true),
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
	}

	if assetTag := getOptionalStr(d, "asset_tag", false); assetTag != "" {
		data.AssetTag = &assetTag
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimModulesPartialUpdateParams().WithID(id).WithData(&data)

	_, err = api.Dcim.DcimModulesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxModuleRead(d, m)
}

func resourceNetboxModuleDelete(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModulesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimModulesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
