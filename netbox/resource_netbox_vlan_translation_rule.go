package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVlanTranslationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVlanTranslationRuleCreate,
		Read:   resourceNetboxVlanTranslationRuleRead,
		Update: resourceNetboxVlanTranslationRuleUpdate,
		Delete: resourceNetboxVlanTranslationRuleDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/vlantranslationrule/):

> A VLAN translation rule represents a mapping of a local VLAN ID to a remote VLAN ID within a VLAN translation policy. VLAN translation rules are grouped into policies, which may then be applied to interfaces.`,

		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"local_vid": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 4094),
			},
			"remote_vid": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 4094),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVlanTranslationRuleCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.VLANTranslationRule{}

	policyID := int64(d.Get("policy_id").(int))
	data.Policy = &policyID

	localVid := int64(d.Get("local_vid").(int))
	data.LocalVid = &localVid

	remoteVid := int64(d.Get("remote_vid").(int))
	data.RemoteVid = &remoteVid

	data.Description = d.Get("description").(string)

	params := ipam.NewIpamVlanTranslationRulesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamVlanTranslationRulesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVlanTranslationRuleRead(d, m)
}

func resourceNetboxVlanTranslationRuleRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlanTranslationRulesReadParams().WithID(id)

	res, err := api.Ipam.IpamVlanTranslationRulesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamVlanTranslationRulesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	rule := res.GetPayload()

	if rule.Policy != nil {
		d.Set("policy_id", *rule.Policy)
	}
	if rule.LocalVid != nil {
		d.Set("local_vid", *rule.LocalVid)
	}
	if rule.RemoteVid != nil {
		d.Set("remote_vid", *rule.RemoteVid)
	}
	d.Set("description", rule.Description)

	return nil
}

func resourceNetboxVlanTranslationRuleUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.VLANTranslationRule{}

	policyID := int64(d.Get("policy_id").(int))
	data.Policy = &policyID

	localVid := int64(d.Get("local_vid").(int))
	data.LocalVid = &localVid

	remoteVid := int64(d.Get("remote_vid").(int))
	data.RemoteVid = &remoteVid

	data.Description = d.Get("description").(string)

	params := ipam.NewIpamVlanTranslationRulesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamVlanTranslationRulesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxVlanTranslationRuleRead(d, m)
}

func resourceNetboxVlanTranslationRuleDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlanTranslationRulesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamVlanTranslationRulesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamVlanTranslationRulesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.SetId("")
	return nil
}
