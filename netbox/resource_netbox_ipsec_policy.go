package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIpsecPolicyPfsGroupOptions = []int{1, 2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34}

func resourceNetboxIpsecPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpsecPolicyCreate,
		Read:   resourceNetboxIpsecPolicyRead,
		Update: resourceNetboxIpsecPolicyUpdate,
		Delete: resourceNetboxIpsecPolicyDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/models/vpn/ipsecpolicy/):

> An IPSec policy defines a set of proposals to be used in the formation of IPSec tunnels. A perfect forward secrecy (PFS) group may optionally also be defined. These policies are referenced by IPSec profiles.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"proposal_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"pfs_group": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntInSlice(resourceNetboxIpsecPolicyPfsGroupOptions),
				Description:  "The Diffie-Hellman group used for perfect forward secrecy. Valid values are 1, 2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32 and 34.",
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

func resourceNetboxIpsecPolicyCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableIPSecPolicy{}

	data.Name = strToPtr(d.Get("name").(string))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.PfsGroup = getOptionalInt(d, "pfs_group")

	data.Proposals = toInt64List(d.Get("proposal_ids"))

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIpsecPoliciesCreateParams().WithData(&data)

	res, err := api.Vpn.VpnIpsecPoliciesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIpsecPolicyRead(d, m)
}

func resourceNetboxIpsecPolicyRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIpsecPoliciesReadParams().WithID(id)

	res, err := api.Vpn.VpnIpsecPoliciesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIpsecPoliciesReadDefault); ok {
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

	if policy.PfsGroup != nil {
		d.Set("pfs_group", policy.PfsGroup.Value)
	} else {
		d.Set("pfs_group", nil)
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

func resourceNetboxIpsecPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPSecPolicy{}

	data.Name = strToPtr(d.Get("name").(string))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.PfsGroup = getOptionalInt(d, "pfs_group")

	data.Proposals = toInt64List(d.Get("proposal_ids"))

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIpsecPoliciesUpdateParams().WithID(id).WithData(&data)

	_, err = api.Vpn.VpnIpsecPoliciesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpsecPolicyRead(d, m)
}

func resourceNetboxIpsecPolicyDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIpsecPoliciesDeleteParams().WithID(id)

	_, err := api.Vpn.VpnIpsecPoliciesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIpsecPoliciesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
