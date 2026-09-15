package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxWirelessLink() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxWirelessLinkRead,
		Description: `:meta:subcategory:Wireless:`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"interface_a_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface_b_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_cipher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_psk": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			customFieldsKey: {
				Type:     schema.TypeMap,
				Computed: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxWirelessLinkRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id := int64(d.Get("id").(int))
	params := wireless.NewWirelessWirelessLinksReadParams().WithID(id)

	res, err := api.Wireless.WirelessWirelessLinksRead(params, nil)
	if err != nil {
		return err
	}

	result := res.GetPayload()
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("ssid", result.Ssid)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set("auth_psk", result.AuthPsk)

	if result.Interfacea != nil {
		d.Set("interface_a_id", result.Interfacea.ID)
	}

	if result.Interfaceb != nil {
		d.Set("interface_b_id", result.Interfaceb.ID)
	}

	if result.Status != nil {
		d.Set("status", result.Status.Value)
	}

	if result.AuthType != nil {
		d.Set("auth_type", result.AuthType.Value)
	}

	if result.AuthCipher != nil {
		d.Set("auth_cipher", result.AuthCipher.Value)
	}

	if result.Tenant != nil {
		d.Set("tenant_id", result.Tenant.ID)
	}

	if result.CustomFields != nil {
		d.Set(customFieldsKey, flattenCustomFields(result.CustomFields))
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
