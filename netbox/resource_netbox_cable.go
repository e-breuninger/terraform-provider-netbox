package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxCable() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxCableCreate,
		Read:   resourceNetboxCableRead,
		Update: resourceNetboxCableUpdate,
		Delete: resourceNetboxCableDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/cable/):

> All connections between device components in NetBox are represented using cables. A cable represents a direct physical connection between two sets of endpoints (A and B), such as a console port and a patch panel port, or between two network interfaces.`,

		Schema: map[string]*schema.Schema{
			"a_termination": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     genericObjectSchema,
			},
			"b_termination": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     genericObjectSchema,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "One of [connected, planned, decommissioning]",
				ValidateFunc: validation.StringInSlice([]string{"connected", "planned", "decommissioning"}, false),
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "One of [cat3, cat5, cat5e, cat6, cat6a, cat7, cat7a, cat8, dac-active, dac-passive, mrj21-trunk, coaxial, mmf, mmf-om1, mmf-om2, mmf-om3, mmf-om4, mmf-om5, smf, smf-os1, smf-os2, aoc, power]",
				ValidateFunc: validation.StringInSlice([]string{
					"cat3", "cat5", "cat5e", "cat6", "cat6a", "cat7", "cat7a", "cat8", "dac-active",
					"dac-passive", "mrj21-trunk", "coaxial", "mmf", "mmf-om1", "mmf-om2", "mmf-om3",
					"mmf-om4", "mmf-om5", "smf", "smf-os1", "smf-os2", "aoc", "power",
				}, false),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"color_hex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"length": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"length_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"length"},
				Description:  "One of [km, m, cm, mi, ft, in]",
				ValidateFunc: validation.StringInSlice([]string{"km", "m", "cm", "mi", "ft", "in"}, false),
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
			tagsKey:         tagsSchema,
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxCableCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableCable{
		Status:      d.Get("status").(string),
		Type:        getOptionalStr(d, "type", false),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Label:       getOptionalStr(d, "label", false),
		Color:       getOptionalStr(d, "color_hex", false),
		Length:      getOptionalFloat(d, "length"),
		LengthUnit:  getOptionalStr(d, "length_unit", false),
		Description: getOptionalStr(d, "description", false),
		Comments:    getOptionalStr(d, "comments", false),
	}

	aTerminations := d.Get("a_termination").(*schema.Set)
	data.ATerminations = getGenericObjectsFromSchemaSet(aTerminations)

	bTerminations := d.Get("b_termination").(*schema.Set)
	data.BTerminations = getGenericObjectsFromSchemaSet(bTerminations)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimCablesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimCablesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxCableRead(d, m)
}

func resourceNetboxCableRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimCablesReadParams().WithID(id)

	res, err := api.Dcim.DcimCablesRead(params, nil)

	if err != nil {
		errorcode := err.(*dcim.DcimCablesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	cable := res.GetPayload()

	d.Set("a_termination", getSchemaSetFromGenericObjects(cable.ATerminations))
	d.Set("b_termination", getSchemaSetFromGenericObjects(cable.BTerminations))

	if cable.Status != nil {
		d.Set("status", cable.Status.Value)
	} else {
		d.Set("status", nil)
	}

	d.Set("type", cable.Type)

	if cable.Tenant != nil {
		d.Set("tenant_id", cable.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	d.Set("label", cable.Label)
	d.Set("color_hex", cable.Color)
	d.Set("length", cable.Length)

	if cable.LengthUnit != nil {
		d.Set("length_unit", cable.LengthUnit.Value)
	} else {
		d.Set("length_unit", nil)
	}

	d.Set("description", cable.Description)
	d.Set("comments", cable.Comments)

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxCableUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := models.WritableCable{
		Status:      d.Get("status").(string),
		Type:        getOptionalStr(d, "type", false),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Label:       getOptionalStr(d, "label", true),
		Color:       getOptionalStr(d, "color_hex", false),
		Length:      getOptionalFloat(d, "length"),
		LengthUnit:  getOptionalStr(d, "length_unit", false),
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
	}

	aTerminations := d.Get("a_termination").(*schema.Set)
	data.ATerminations = getGenericObjectsFromSchemaSet(aTerminations)

	bTerminations := d.Get("b_termination").(*schema.Set)
	data.BTerminations = getGenericObjectsFromSchemaSet(bTerminations)

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	params := dcim.NewDcimCablesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimCablesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxCableRead(d, m)
}

func resourceNetboxCableDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimCablesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimCablesDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
