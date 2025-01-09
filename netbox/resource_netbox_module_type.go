package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxModuleType() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxModuleTypeCreate,
		Read:   resourceNetboxModuleTypeRead,
		Update: resourceNetboxModuleTypeUpdate,
		Delete: resourceNetboxModuleTypeDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):

> A module type represents a specific make and model of hardware component which is installable within a device's module bay and has its own child components. For example, consider a chassis-based switch or router with a number of field-replaceable line cards. Each line card has its own model number and includes a certain set of components such as interfaces. Each module type may have a manufacturer, model number, and part number assigned to it.`,

		Schema: map[string]*schema.Schema{
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"weight_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"weight"},
				Description:  "One of [kg, g, lb, oz]",
				ValidateFunc: validation.StringInSlice([]string{"kg", "g", "lb", "oz"}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			customFieldsKey: customFieldsSchema,
			tagsKey:         tagsSchema,
		},
		CustomizeDiff: customFieldsDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxModuleTypeCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableModuleType{
		Manufacturer: int64ToPtr(int64(d.Get("manufacturer_id").(int))),
		Model:        strToPtr(d.Get("model").(string)),
		PartNumber:   getOptionalStr(d, "part_number", false),
		Weight:       getOptionalFloat(d, "weight"),
		WeightUnit:   getOptionalStr(d, "weight_unit", false),
		Description:  getOptionalStr(d, "description", false),
		Comments:     getOptionalStr(d, "comments", false),
	}

	data.CustomFields = computeCustomFieldsModel(d)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimModuleTypesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimModuleTypesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxModuleTypeRead(d, m)
}

func resourceNetboxModuleTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModuleTypesReadParams().WithID(id)

	res, err := api.Dcim.DcimModuleTypesRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimModuleTypesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	moduleType := res.GetPayload()

	if moduleType.Manufacturer != nil {
		d.Set("manufacturer_id", moduleType.Manufacturer.ID)
	} else {
		d.Set("manufacturer_id", nil)
	}

	d.Set("model", moduleType.Model)
	d.Set("part_number", moduleType.PartNumber)
	d.Set("weight", moduleType.Weight)

	if moduleType.WeightUnit != nil {
		d.Set("weight_unit", moduleType.WeightUnit.Value)
	} else {
		d.Set("weight_unit", nil)
	}

	d.Set("description", moduleType.Description)
	d.Set("comments", moduleType.Comments)

	d.Set(customFieldsKey, res.GetPayload().CustomFields)

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxModuleTypeUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableModuleType{
		Manufacturer: int64ToPtr(int64(d.Get("manufacturer_id").(int))),
		Model:        strToPtr(d.Get("model").(string)),
		PartNumber:   getOptionalStr(d, "part_number", true),
		Weight:       getOptionalFloat(d, "weight"),
		WeightUnit:   getOptionalStr(d, "weight_unit", false),
		Description:  getOptionalStr(d, "description", true),
		Comments:     getOptionalStr(d, "comments", true),
	}

	data.CustomFields = computeCustomFieldsModel(d)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimModuleTypesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimModuleTypesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxModuleTypeRead(d, m)
}

func resourceNetboxModuleTypeDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimModuleTypesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimModuleTypesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
