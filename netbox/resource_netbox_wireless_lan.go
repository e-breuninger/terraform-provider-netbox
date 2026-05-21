package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxWirelessLANStatusOptions = []string{"active", "reserved", "disabled", "deprecated"}
var resourceNetboxWirelessLANAuthTypeOptions = []string{"open", "wep", "wpa-personal", "wpa-enterprise"}
var resourceNetboxWirelessLANAuthCipherOptions = []string{"auto", "tkip", "aes"}

func resourceNetboxWirelessLAN() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxWirelessLANCreate,
		Read:   resourceNetboxWirelessLANRead,
		Update: resourceNetboxWirelessLANUpdate,
		Delete: resourceNetboxWirelessLANDelete,

		Description: `:meta:subcategory:Wireless:

> A Wireless LAN represents a broadcast wireless network, identified by its SSID and optional authentication settings.`,

		Schema: map[string]*schema.Schema{
			"ssid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "active",
				ValidateFunc: validation.StringInSlice(resourceNetboxWirelessLANStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWirelessLANStatusOptions),
			},
			"group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auth_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxWirelessLANAuthTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWirelessLANAuthTypeOptions),
			},
			"auth_cipher": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxWirelessLANAuthCipherOptions, false),
				Description:  buildValidValueDescription(resourceNetboxWirelessLANAuthCipherOptions),
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

func resourceNetboxWirelessLANCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	ssid := d.Get("ssid").(string)
	status := d.Get("status").(string)

	data := &models.WritableWirelessLAN{
		Ssid:        &ssid,
		Status:      status,
		AuthType:    d.Get("auth_type").(string),
		AuthCipher:  d.Get("auth_cipher").(string),
		AuthPsk:     d.Get("auth_psk").(string),
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
	}

	if groupID, ok := d.GetOk("group_id"); ok {
		data.Group = int64ToPtr(int64(groupID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := wireless.NewWirelessWirelessLansCreateParams().WithData(data)
	res, err := api.Wireless.WirelessWirelessLansCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxWirelessLANRead(d, m)
}

func resourceNetboxWirelessLANRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := wireless.NewWirelessWirelessLansReadParams().WithID(id)
	res, err := api.Wireless.WirelessWirelessLansRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*wireless.WirelessWirelessLansReadDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	wlan := res.GetPayload()
	d.Set("ssid", wlan.Ssid)
	d.Set("description", wlan.Description)
	d.Set("comments", wlan.Comments)
	d.Set("auth_psk", wlan.AuthPsk)

	if wlan.Status != nil {
		d.Set("status", wlan.Status.Value)
	} else {
		d.Set("status", nil)
	}

	if wlan.AuthType != nil {
		d.Set("auth_type", wlan.AuthType.Value)
	} else {
		d.Set("auth_type", nil)
	}

	if wlan.AuthCipher != nil {
		d.Set("auth_cipher", wlan.AuthCipher.Value)
	} else {
		d.Set("auth_cipher", nil)
	}

	if wlan.Group != nil {
		d.Set("group_id", wlan.Group.ID)
	} else {
		d.Set("group_id", nil)
	}

	if wlan.Tenant != nil {
		d.Set("tenant_id", wlan.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if wlan.Vlan != nil {
		d.Set("vlan_id", wlan.Vlan.ID)
	} else {
		d.Set("vlan_id", nil)
	}

	cf := getCustomFields(wlan.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, wlan.Tags)

	return nil
}

func resourceNetboxWirelessLANUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	ssid := d.Get("ssid").(string)
	data := models.WritableWirelessLAN{
		Ssid:        &ssid,
		Status:      d.Get("status").(string),
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
	}

	overrideFields := make(map[string]any)

	groupID := d.Get("group_id").(int)
	if groupID != 0 {
		data.Group = int64ToPtr(int64(groupID))
	} else {
		overrideFields["group"] = nil
	}

	tenantID := d.Get("tenant_id").(int)
	if tenantID != 0 {
		data.Tenant = int64ToPtr(int64(tenantID))
	} else {
		overrideFields["tenant"] = nil
	}

	vlanID := d.Get("vlan_id").(int)
	if vlanID != 0 {
		data.Vlan = int64ToPtr(int64(vlanID))
	} else {
		overrideFields["vlan"] = nil
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

	params := wireless.NewWirelessWirelessLansPartialUpdateParams().WithID(id).WithData(&data)
	_, err = api.Wireless.WirelessWirelessLansPartialUpdate(params, nil, hackSerializeWirelessWithValues(overrideFields))
	if err != nil {
		return err
	}

	return resourceNetboxWirelessLANRead(d, m)
}

func resourceNetboxWirelessLANDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := wireless.NewWirelessWirelessLansDeleteParams().WithID(id)
	_, err := api.Wireless.WirelessWirelessLansDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*wireless.WirelessWirelessLansDeleteDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	return nil
}
