package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIpsecProposalEncryptionAlgorithmOptions = []string{"aes-128-cbc", "aes-128-gcm", "aes-192-cbc", "aes-192-gcm", "aes-256-cbc", "aes-256-gcm", "3des-cbc", "des-cbc"}
var resourceNetboxIpsecProposalAuthenticationAlgorithmOptions = []string{"hmac-sha1", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-md5"}

func resourceNetboxIpsecProposal() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpsecProposalCreate,
		Read:   resourceNetboxIpsecProposalRead,
		Update: resourceNetboxIpsecProposalUpdate,
		Delete: resourceNetboxIpsecProposalDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/models/vpn/ipsecproposal/):

> An IPSec proposal defines a set of parameters used in negotiating security associations for IPSec tunnels. IPSec proposals defined in NetBox can be referenced by IPSec policies, which are in turn employed by IPSec profiles.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIpsecProposalEncryptionAlgorithmOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIpsecProposalEncryptionAlgorithmOptions),
			},
			"authentication_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIpsecProposalAuthenticationAlgorithmOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIpsecProposalAuthenticationAlgorithmOptions),
			},
			"sa_lifetime_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sa_lifetime_data": {
				Type:     schema.TypeInt,
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

func resourceNetboxIpsecProposalCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableIPSecProposal{}

	data.Name = strToPtr(d.Get("name").(string))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.SaLifetimeSeconds = getOptionalInt(d, "sa_lifetime_seconds")
	data.SaLifetimeData = getOptionalInt(d, "sa_lifetime_data")

	if encAlg, ok := d.GetOk("encryption_algorithm"); ok {
		data.EncryptionAlgorithm = strToPtr(encAlg.(string))
	}

	if authAlg, ok := d.GetOk("authentication_algorithm"); ok {
		data.AuthenticationAlgorithm = strToPtr(authAlg.(string))
	}

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIpsecProposalsCreateParams().WithData(&data)

	res, err := api.Vpn.VpnIpsecProposalsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIpsecProposalRead(d, m)
}

func resourceNetboxIpsecProposalRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIpsecProposalsReadParams().WithID(id)

	res, err := api.Vpn.VpnIpsecProposalsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIpsecProposalsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	proposal := res.GetPayload()
	d.Set("name", proposal.Name)
	d.Set("description", proposal.Description)
	d.Set("comments", proposal.Comments)
	d.Set("sa_lifetime_seconds", proposal.SaLifetimeSeconds)
	d.Set("sa_lifetime_data", proposal.SaLifetimeData)

	if proposal.EncryptionAlgorithm != nil {
		d.Set("encryption_algorithm", proposal.EncryptionAlgorithm.Value)
	} else {
		d.Set("encryption_algorithm", nil)
	}

	if proposal.AuthenticationAlgorithm != nil {
		d.Set("authentication_algorithm", proposal.AuthenticationAlgorithm.Value)
	} else {
		d.Set("authentication_algorithm", nil)
	}

	cf := getCustomFields(proposal.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, proposal.Tags)

	return nil
}

func resourceNetboxIpsecProposalUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPSecProposal{}

	data.Name = strToPtr(d.Get("name").(string))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.SaLifetimeSeconds = getOptionalInt(d, "sa_lifetime_seconds")
	data.SaLifetimeData = getOptionalInt(d, "sa_lifetime_data")

	if encAlg, ok := d.GetOk("encryption_algorithm"); ok {
		data.EncryptionAlgorithm = strToPtr(encAlg.(string))
	}

	if authAlg, ok := d.GetOk("authentication_algorithm"); ok {
		data.AuthenticationAlgorithm = strToPtr(authAlg.(string))
	}

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}
	data.Tags = tags

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := vpn.NewVpnIpsecProposalsUpdateParams().WithID(id).WithData(&data)

	_, err = api.Vpn.VpnIpsecProposalsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpsecProposalRead(d, m)
}

func resourceNetboxIpsecProposalDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIpsecProposalsDeleteParams().WithID(id)

	_, err := api.Vpn.VpnIpsecProposalsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIpsecProposalsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
