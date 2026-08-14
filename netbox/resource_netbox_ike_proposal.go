package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIkeProposalAuthenticationMethodOptions = []string{"preshared-keys", "certificates", "rsa-signatures", "dsa-signatures"}
var resourceNetboxIkeProposalEncryptionAlgorithmOptions = []string{"aes-128-cbc", "aes-128-gcm", "aes-192-cbc", "aes-192-gcm", "aes-256-cbc", "aes-256-gcm", "3des-cbc", "des-cbc"}
var resourceNetboxIkeProposalAuthenticationAlgorithmOptions = []string{"hmac-sha1", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-md5"}
var resourceNetboxIkeProposalGroupOptions = []int{1, 2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34}

func resourceNetboxIkeProposal() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIkeProposalCreate,
		Read:   resourceNetboxIkeProposalRead,
		Update: resourceNetboxIkeProposalUpdate,
		Delete: resourceNetboxIkeProposalDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/models/vpn/ikeproposal/):

> An Internet Key Exchange (IKE) proposal defines a set of parameters used to establish a secure bidirectional connection across an untrusted medium, such as the Internet. IKE proposals defined in NetBox can be referenced by IKE policies, which are in turn employed by IPSec profiles.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_method": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIkeProposalAuthenticationMethodOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIkeProposalAuthenticationMethodOptions),
			},
			"encryption_algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIkeProposalEncryptionAlgorithmOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIkeProposalEncryptionAlgorithmOptions),
			},
			"authentication_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIkeProposalAuthenticationAlgorithmOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIkeProposalAuthenticationAlgorithmOptions),
			},
			"group": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice(resourceNetboxIkeProposalGroupOptions),
				Description:  "The Diffie-Hellman group supported by the proposal. Valid values are 1, 2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32 and 34.",
			},
			"sa_lifetime": {
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

func resourceNetboxIkeProposalCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableIKEProposal{}

	data.Name = strToPtr(d.Get("name").(string))
	data.AuthenticationMethod = strToPtr(d.Get("authentication_method").(string))
	data.EncryptionAlgorithm = strToPtr(d.Get("encryption_algorithm").(string))
	data.Group = int64ToPtr(int64(d.Get("group").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Comments = getOptionalStr(d, "comments", false)
	data.SaLifetime = getOptionalInt(d, "sa_lifetime")

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

	params := vpn.NewVpnIkeProposalsCreateParams().WithData(&data)

	res, err := api.Vpn.VpnIkeProposalsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIkeProposalRead(d, m)
}

func resourceNetboxIkeProposalRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIkeProposalsReadParams().WithID(id)

	res, err := api.Vpn.VpnIkeProposalsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIkeProposalsReadDefault); ok {
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
	d.Set("sa_lifetime", proposal.SaLifetime)

	if proposal.AuthenticationMethod != nil {
		d.Set("authentication_method", proposal.AuthenticationMethod.Value)
	} else {
		d.Set("authentication_method", nil)
	}

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

	if proposal.Group != nil {
		d.Set("group", proposal.Group.Value)
	} else {
		d.Set("group", nil)
	}

	cf := getCustomFields(proposal.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, proposal.Tags)

	return nil
}

func resourceNetboxIkeProposalUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIKEProposal{}

	data.Name = strToPtr(d.Get("name").(string))
	data.AuthenticationMethod = strToPtr(d.Get("authentication_method").(string))
	data.EncryptionAlgorithm = strToPtr(d.Get("encryption_algorithm").(string))
	data.Group = int64ToPtr(int64(d.Get("group").(int)))

	data.Description = getOptionalStr(d, "description", true)
	data.Comments = getOptionalStr(d, "comments", true)
	data.SaLifetime = getOptionalInt(d, "sa_lifetime")

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

	params := vpn.NewVpnIkeProposalsUpdateParams().WithID(id).WithData(&data)

	_, err = api.Vpn.VpnIkeProposalsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIkeProposalRead(d, m)
}

func resourceNetboxIkeProposalDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnIkeProposalsDeleteParams().WithID(id)

	_, err := api.Vpn.VpnIkeProposalsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnIkeProposalsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
