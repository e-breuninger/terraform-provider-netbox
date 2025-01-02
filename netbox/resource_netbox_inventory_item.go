package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxInventoryItem() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxInventoryItemCreate,
		Read:   resourceNetboxInventoryItemRead,
		Update: resourceNetboxInventoryItemUpdate,
		Delete: resourceNetboxInventoryItemDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/inventoryitem/):

> Inventory items represent hardware components installed within a device, such as a power supply or CPU or line card. They are intended to be used primarily for inventory purposes.`,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"part_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"component_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"dcim.powerport",
					"dcim.poweroutlet",
					"dcim.frontport",
					"dcim.rearport",
					"dcim.consoleserverport",
					"dcim.consoleport",
					"dcim.interface",
				}, false),
			},
			"component_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"component_type"},
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

func resourceNetboxInventoryItemCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableInventoryItem{
		Device:       int64ToPtr(int64(d.Get("device_id").(int))),
		Name:         strToPtr(d.Get("name").(string)),
		Parent:       getOptionalInt(d, "parent_id"),
		Label:        getOptionalStr(d, "label", false),
		Role:         getOptionalInt(d, "role_id"),
		Manufacturer: getOptionalInt(d, "manufacturer_id"),
		PartID:       getOptionalStr(d, "part_id", false),
		Serial:       getOptionalStr(d, "serial", false),
		Discovered:   d.Get("discovered").(bool),
		Description:  getOptionalStr(d, "description", false),
	}

	if assetTag := getOptionalStr(d, "asset_tag", false); assetTag != "" {
		data.AssetTag = &assetTag
	}

	if componentType := getOptionalStr(d, "component_type", false); componentType != "" {
		data.ComponentType = &componentType
		data.ComponentID = getOptionalInt(d, "component_id")
	}

	data.CustomFields = computeCustomFieldsModel(d)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimInventoryItemsCreateParams().WithData(&data)

	res, err := api.Dcim.DcimInventoryItemsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxInventoryItemRead(d, m)
}

func resourceNetboxInventoryItemRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInventoryItemsReadParams().WithID(id)

	res, err := api.Dcim.DcimInventoryItemsRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimInventoryItemsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	item := res.GetPayload()

	if item.Device != nil {
		d.Set("device_id", item.Device.ID)
	} else {
		d.Set("device_id", nil)
	}

	d.Set("name", item.Name)

	d.Set("parent_id", item.Parent)

	d.Set("label", item.Label)

	if item.Role != nil {
		d.Set("role_id", item.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	if item.Manufacturer != nil {
		d.Set("manufacturer_id", item.Manufacturer.ID)
	} else {
		d.Set("manufacturer_id", nil)
	}

	d.Set("part_id", item.PartID)
	d.Set("serial", item.Serial)
	d.Set("asset_tag", item.AssetTag)
	d.Set("discovered", item.Discovered)
	d.Set("description", item.Description)
	d.Set("component_type", item.ComponentType)
	d.Set("component_id", item.ComponentID)

	d.Set(customFieldsKey, res.GetPayload().CustomFields)

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxInventoryItemUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableInventoryItem{
		Device:       int64ToPtr(int64(d.Get("device_id").(int))),
		Name:         strToPtr(d.Get("name").(string)),
		Parent:       getOptionalInt(d, "parent_id"),
		Label:        getOptionalStr(d, "label", true),
		Role:         getOptionalInt(d, "role_id"),
		Manufacturer: getOptionalInt(d, "manufacturer_id"),
		PartID:       getOptionalStr(d, "part_id", true),
		Serial:       getOptionalStr(d, "serial", true),
		Discovered:   d.Get("discovered").(bool),
		Description:  getOptionalStr(d, "description", true),
	}

	if assetTag := getOptionalStr(d, "asset_tag", false); assetTag != "" {
		data.AssetTag = &assetTag
	}

	if componentType := getOptionalStr(d, "component_type", false); componentType != "" {
		data.ComponentType = &componentType
		data.ComponentID = getOptionalInt(d, "component_id")
	}

	data.CustomFields = computeCustomFieldsModel(d)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimInventoryItemsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimInventoryItemsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxInventoryItemRead(d, m)
}

func resourceNetboxInventoryItemDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInventoryItemsDeleteParams().WithID(id)

	_, err := api.Dcim.DcimInventoryItemsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
