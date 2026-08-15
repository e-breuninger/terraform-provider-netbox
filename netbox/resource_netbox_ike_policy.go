package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIkePolicyVersionOptions = []int{1, 2}
var resourceNetboxIkePolicyModeOptions = []string{"aggressive", "main"}

func resourceNetboxIkePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIkePolicyCreate,
		Read:   resourceNetboxIkePolicyRead,
		Update: resourceNetboxIkePolicyUpdate,
		Delete: resourceNetboxIkePolicyDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/models/vpn/ikepolicy/):

> An Internet Key Exchange (IKE) policy defines an IKE version, mode, and set of proposals to be used in IKE negotiation. These policies are referenced by IPSec profiles.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice(resourceNetboxIkePolicyVersionOptions),
				Description:  "The IKE version. Valid values are 1 and 2.",
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIkePolicyModeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIkePolicyModeOptions),
			},
			"proposal_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"preshared_key": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
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

func resourceNetboxIkePolicyCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableIKEPolicy{}

	data.Name = strToPtr(d.Get("name").(string))
	data.Version = int64ToPtr(int64(d.Get("version").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.PresharedKey = getOptionalStr(d, "preshared_key", false)

	if mode, ok := d.GetOk("mode"); ok {
		data.Mode = strToPtr(mode.(string))
	}

	data.Proposals = toInt64List(d.Get("proposal_ids"))

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIkePoliciesCreateParams().WithData(&data)

	res, err := api.Vpn.VpnIkePoliciesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIkePolicyRead(d, m)
}

func resourceNetboxIkePolicyRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIkePoliciesReadParams().WithID(id)

	res, err := api.Vpn.VpnIkePoliciesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIkePoliciesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	policy := res.GetPayload()
	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("comments", policy.Comments)
	d.Set("preshared_key", policy.PresharedKey)

	if policy.Version != nil {
		d.Set("version", policy.Version.Value)
	} else {
		d.Set("version", nil)
	}

	if policy.Mode != nil {
		d.Set("mode", policy.Mode.Value)
	} else {
		d.Set("mode", nil)
	}

	proposals := policy.Proposals
	proposalIDs := make([]int64, len(proposals))
	for i, v := range proposals {
		proposalIDs[i] = v.ID
	}
	d.Set("proposal_ids", proposalIDs)

	cf := getCustomFields(policy.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, policy.Tags)

	return nil
}

func resourceNetboxIkePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIKEPolicy{}

	data.Name = strToPtr(d.Get("name").(string))
	data.Version = int64ToPtr(int64(d.Get("version").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.PresharedKey = getOptionalStr(d, "preshared_key", false)

	if mode, ok := d.GetOk("mode"); ok {
		data.Mode = strToPtr(mode.(string))
	}

	data.Proposals = toInt64List(d.Get("proposal_ids"))

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIkePoliciesUpdateParams().WithID(id).WithData(&data)

	_, err = api.Vpn.VpnIkePoliciesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIkePolicyRead(d, m)
}

func resourceNetboxIkePolicyDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIkePoliciesDeleteParams().WithID(id)

	_, err := api.Vpn.VpnIkePoliciesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIkePoliciesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
