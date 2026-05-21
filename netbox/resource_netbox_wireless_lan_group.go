package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxWirelessLANGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxWirelessLANGroupCreate,
		Read:   resourceNetboxWirelessLANGroupRead,
		Update: resourceNetboxWirelessLANGroupUpdate,
		Delete: resourceNetboxWirelessLANGroupDelete,

		Description: `:meta:subcategory:Wireless:

> A Wireless LAN Group is used to organize wireless LANs into a recursive hierarchy.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			customFieldsKey: customFieldsSchema,
			tagsKey:         tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxWirelessLANGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	slugValue, slugOk := d.GetOk("slug")
	slug := getSlug(name)
	if slugOk {
		slug = slugValue.(string)
	}

	data := &models.WritableWirelessLANGroup{
		Name:        &name,
		Slug:        &slug,
		Description: description,
	}

	if parentID, ok := d.GetOk("parent_id"); ok {
		data.Parent = int64ToPtr(int64(parentID.(int)))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := wireless.NewWirelessWirelessLanGroupsCreateParams().WithData(data)
	res, err := api.Wireless.WirelessWirelessLanGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxWirelessLANGroupRead(d, m)
}

func resourceNetboxWirelessLANGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := wireless.NewWirelessWirelessLanGroupsReadParams().WithID(id)
	res, err := api.Wireless.WirelessWirelessLanGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*wireless.WirelessWirelessLanGroupsReadDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	group := res.GetPayload()
	d.Set("name", group.Name)
	d.Set("slug", group.Slug)
	d.Set("description", group.Description)

	if group.Parent != nil {
		d.Set("parent_id", group.Parent.ID)
	} else {
		d.Set("parent_id", nil)
	}

	cf := getCustomFields(group.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, group.Tags)

	return nil
}

func resourceNetboxWirelessLANGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)

	slugValue, slugOk := d.GetOk("slug")
	slug := getSlug(name)
	if slugOk {
		slug = slugValue.(string)
	}

	data := models.WritableWirelessLANGroup{
		Name:        &name,
		Slug:        &slug,
		Description: getOptionalStr(d, "description", true),
	}

	var nullFields []string

	parentID := d.Get("parent_id").(int)
	if parentID != 0 {
		data.Parent = int64ToPtr(int64(parentID))
	} else {
		nullFields = append(nullFields, "parent")
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := wireless.NewWirelessWirelessLanGroupsPartialUpdateParams().WithID(id).WithData(&data)
	_, err = api.Wireless.WirelessWirelessLanGroupsPartialUpdate(params, nil, hackSerializeWirelessAsNull(nullFields...))
	if err != nil {
		return err
	}

	return resourceNetboxWirelessLANGroupRead(d, m)
}

func resourceNetboxWirelessLANGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := wireless.NewWirelessWirelessLanGroupsDeleteParams().WithID(id)
	_, err := api.Wireless.WirelessWirelessLanGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*wireless.WirelessWirelessLanGroupsDeleteDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	return nil
}
