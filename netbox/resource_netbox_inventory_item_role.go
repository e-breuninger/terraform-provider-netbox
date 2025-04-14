package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxInventoryItemRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxInventoryItemRoleCreate,
		Read:   resourceNetboxInventoryItemRoleRead,
		Update: resourceNetboxInventoryItemRoleUpdate,
		Delete: resourceNetboxInventoryItemRoleDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/inventoryitemrole/):

> Inventory items can be organized by functional roles, which are fully customizable by the user. For example, you might create roles for power supplies, fans, interface optics, etc.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"color_hex": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
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

func resourceNetboxInventoryItemRoleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.InventoryItemRole{
		Name:        strToPtr(d.Get("name").(string)),
		Slug:        strToPtr(d.Get("slug").(string)),
		Description: getOptionalStr(d, "description", false),
		Color:       getOptionalStr(d, "color_hex", false),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimInventoryItemRolesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimInventoryItemRolesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxInventoryItemRoleRead(d, m)
}

func resourceNetboxInventoryItemRoleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInventoryItemRolesReadParams().WithID(id)

	res, err := api.Dcim.DcimInventoryItemRolesRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimInventoryItemRolesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	role := res.GetPayload()
	d.Set("name", role.Name)
	d.Set("slug", role.Slug)
	d.Set("color_hex", role.Color)
	d.Set("description", role.Description)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxInventoryItemRoleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.InventoryItemRole{
		Name:        strToPtr(d.Get("name").(string)),
		Slug:        strToPtr(d.Get("slug").(string)),
		Description: getOptionalStr(d, "description", true),
		Color:       getOptionalStr(d, "color_hex", false),
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimInventoryItemRolesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimInventoryItemRolesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxInventoryItemRoleRead(d, m)
}

func resourceNetboxInventoryItemRoleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimInventoryItemRolesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimInventoryItemRolesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
