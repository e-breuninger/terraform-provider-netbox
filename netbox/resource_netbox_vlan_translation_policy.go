package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVlanTranslationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVlanTranslationPolicyCreate,
		Read:   resourceNetboxVlanTranslationPolicyRead,
		Update: resourceNetboxVlanTranslationPolicyUpdate,
		Delete: resourceNetboxVlanTranslationPolicyDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/vlantranslationpolicy/):

> A VLAN translation policy represents a collection of VLAN translation rules that can be applied to an interface. Multiple interfaces can share the same VLAN translation policy.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
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
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVlanTranslationPolicyCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.VLANTranslationPolicy{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	params := ipam.NewIpamVlanTranslationPoliciesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamVlanTranslationPoliciesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVlanTranslationPolicyRead(d, m)
}

func resourceNetboxVlanTranslationPolicyRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlanTranslationPoliciesReadParams().WithID(id)

	res, err := api.Ipam.IpamVlanTranslationPoliciesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamVlanTranslationPoliciesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
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

	return nil
}

func resourceNetboxVlanTranslationPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.VLANTranslationPolicy{}

	name := d.Get("name").(string)
	data.Name = &name
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	params := ipam.NewIpamVlanTranslationPoliciesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamVlanTranslationPoliciesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxVlanTranslationPolicyRead(d, m)
}

func resourceNetboxVlanTranslationPolicyDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlanTranslationPoliciesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamVlanTranslationPoliciesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamVlanTranslationPoliciesDeleteDefault); ok {
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
