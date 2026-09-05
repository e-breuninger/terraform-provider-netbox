package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIpsecProfileModeOptions = []string{"esp", "ah"}

func resourceNetboxIpsecProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpsecProfileCreate,
		Read:   resourceNetboxIpsecProfileRead,
		Update: resourceNetboxIpsecProfileUpdate,
		Delete: resourceNetboxIpsecProfileDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/models/vpn/ipsecprofile/):

> An IPSec profile defines an IKE policy, IPSec policy, and IPSec mode used for establishing an IPSec tunnel.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIpsecProfileModeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIpsecProfileModeOptions),
			},
			"ike_policy_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ipsec_policy_id": {
				Type:     schema.TypeInt,
				Required: true,
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

func resourceNetboxIpsecProfileCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableIPSecProfile{}

	data.Name = strToPtr(d.Get("name").(string))
	data.Mode = strToPtr(d.Get("mode").(string))
	data.IkePolicy = int64ToPtr(int64(d.Get("ike_policy_id").(int)))
	data.IpsecPolicy = int64ToPtr(int64(d.Get("ipsec_policy_id").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIpsecProfilesCreateParams().WithData(&data)

	res, err := api.Vpn.VpnIpsecProfilesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIpsecProfileRead(d, m)
}

func resourceNetboxIpsecProfileRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIpsecProfilesReadParams().WithID(id)

	res, err := api.Vpn.VpnIpsecProfilesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIpsecProfilesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	profile := res.GetPayload()
	d.Set("name", profile.Name)
	d.Set("description", profile.Description)
	d.Set("comments", profile.Comments)

	if profile.Mode != nil {
		d.Set("mode", profile.Mode.Value)
	} else {
		d.Set("mode", nil)
	}

	if profile.IkePolicy != nil {
		d.Set("ike_policy_id", profile.IkePolicy.ID)
	} else {
		d.Set("ike_policy_id", nil)
	}

	if profile.IpsecPolicy != nil {
		d.Set("ipsec_policy_id", profile.IpsecPolicy.ID)
	} else {
		d.Set("ipsec_policy_id", nil)
	}

	cf := getCustomFields(profile.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, profile.Tags)

	return nil
}

func resourceNetboxIpsecProfileUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPSecProfile{}

	data.Name = strToPtr(d.Get("name").(string))
	data.Mode = strToPtr(d.Get("mode").(string))
	data.IkePolicy = int64ToPtr(int64(d.Get("ike_policy_id").(int)))
	data.IpsecPolicy = int64ToPtr(int64(d.Get("ipsec_policy_id").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIpsecProfilesUpdateParams().WithID(id).WithData(&data)

	_, err = api.Vpn.VpnIpsecProfilesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpsecProfileRead(d, m)
}

func resourceNetboxIpsecProfileDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIpsecProfilesDeleteParams().WithID(id)

	_, err := api.Vpn.VpnIpsecProfilesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIpsecProfilesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
