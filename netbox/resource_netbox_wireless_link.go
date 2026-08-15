package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxWirelessLinkStatusOptions = []string{"connected", "planned", "decommissioning"}
var resourceNetboxWirelessLinkAuthTypeOptions = []string{"open", "wep", "wpa-personal", "wpa-enterprise"}
var resourceNetboxWirelessLinkAuthCipherOptions = []string{"auto", "tkip", "aes"}

func resourceNetboxWirelessLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxWirelessLinkCreate,
		Read:   resourceNetboxWirelessLinkRead,
		Update: resourceNetboxWirelessLinkUpdate,
		Delete: resourceNetboxWirelessLinkDelete,

		Description: `:meta:subcategory:Wireless:

> A wireless link represents a point-to-point connection between two wireless interfaces.`,

		Schema: map[string]*schema.Schema{
			"interface_a_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"interface_b_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ssid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 32),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "connected",
				ValidateFunc: validation.StringInSlice(resourceNetboxWirelessLinkStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWirelessLinkStatusOptions),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxWirelessLinkAuthTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWirelessLinkAuthTypeOptions),
			},
			"auth_cipher": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxWirelessLinkAuthCipherOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWirelessLinkAuthCipherOptions),
			},
			"auth_psk": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringLenBetween(0, 64),
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxWirelessLinkCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	interfaceA := int64(d.Get("interface_a_id").(int))
	interfaceB := int64(d.Get("interface_b_id").(int))
	status := d.Get("status").(string)

	data := &models.WritableWirelessLink{
		Interfacea:  &interfaceA,
		Interfaceb:  &interfaceB,
		Ssid:        d.Get("ssid").(string),
		Status:      status,
		AuthType:    d.Get("auth_type").(string),
		AuthCipher:  d.Get("auth_cipher").(string),
		AuthPsk:     d.Get("auth_psk").(string),
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := wireless.NewWirelessWirelessLinksCreateParams().WithData(data)
	res, err := api.Wireless.WirelessWirelessLinksCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxWirelessLinkRead(d, m)
}

func resourceNetboxWirelessLinkRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := wireless.NewWirelessWirelessLinksReadParams().WithID(id)
	res, err := api.Wireless.WirelessWirelessLinksRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*wireless.WirelessWirelessLinksReadDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	link := res.GetPayload()
	d.Set("ssid", link.Ssid)
	d.Set("description", link.Description)
	d.Set("comments", link.Comments)
	d.Set("auth_psk", link.AuthPsk)

	if link.Interfacea != nil {
		d.Set("interface_a_id", link.Interfacea.ID)
	}

	if link.Interfaceb != nil {
		d.Set("interface_b_id", link.Interfaceb.ID)
	}

	if link.Status != nil {
		d.Set("status", link.Status.Value)
	} else {
		d.Set("status", nil)
	}

	if link.AuthType != nil {
		d.Set("auth_type", link.AuthType.Value)
	} else {
		d.Set("auth_type", nil)
	}

	if link.AuthCipher != nil {
		d.Set("auth_cipher", link.AuthCipher.Value)
	} else {
		d.Set("auth_cipher", nil)
	}

	if link.Tenant != nil {
		d.Set("tenant_id", link.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	cf := getCustomFields(link.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, link.Tags)

	return nil
}

func resourceNetboxWirelessLinkUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	interfaceA := int64(d.Get("interface_a_id").(int))
	interfaceB := int64(d.Get("interface_b_id").(int))

	data := models.WritableWirelessLink{
		Interfacea:  &interfaceA,
		Interfaceb:  &interfaceB,
		Status:      d.Get("status").(string),
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
	}

	overrideFields := make(map[string]any)

	if ssid, ok := d.GetOk("ssid"); ok {
		data.Ssid = ssid.(string)
	} else {
		overrideFields["ssid"] = ""
	}

	tenantID := d.Get("tenant_id").(int)
	if tenantID != 0 {
		data.Tenant = int64ToPtr(int64(tenantID))
	} else {
		overrideFields["tenant"] = nil
	}

	if authType, ok := d.GetOk("auth_type"); ok {
		data.AuthType = authType.(string)
	} else {
		overrideFields["auth_type"] = nil
	}

	if authCipher, ok := d.GetOk("auth_cipher"); ok {
		data.AuthCipher = authCipher.(string)
	} else {
		overrideFields["auth_cipher"] = nil
	}

	if authPSK, ok := d.GetOk("auth_psk"); ok {
		data.AuthPsk = authPSK.(string)
	} else {
		overrideFields["auth_psk"] = ""
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := wireless.NewWirelessWirelessLinksPartialUpdateParams().WithID(id).WithData(&data)
	_, err = api.Wireless.WirelessWirelessLinksPartialUpdate(params, nil, hackSerializeWirelessWithValues(overrideFields))
	if err != nil {
		return err
	}

	return resourceNetboxWirelessLinkRead(d, m)
}

func resourceNetboxWirelessLinkDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := wireless.NewWirelessWirelessLinksDeleteParams().WithID(id)
	_, err := api.Wireless.WirelessWirelessLinksDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*wireless.WirelessWirelessLinksDeleteDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	return nil
}
