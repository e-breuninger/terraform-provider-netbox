package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPrefixCreate,
		Read:   resourceNetboxPrefixRead,
		Update: resourceNetboxPrefixUpdate,
		Delete: resourceNetboxPrefixDelete,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "container"}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mark_utilized": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}
func resourceNetboxPrefixCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritablePrefix{}

	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	is_pool := d.Get("is_pool").(bool)
	mark_utilized := d.Get("mark_utilized").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = is_pool

	data.MarkUtilized = mark_utilized

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))
	data.CustomFields = d.Get(customFieldsKey)

	params := ipam.NewIpamPrefixesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamPrefixesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))
	d.Set(customFieldsKey, res.GetPayload().CustomFields)

	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamPrefixesReadParams().WithID(id)

	res, err := api.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamPrefixesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("description", res.GetPayload().Description)
	d.Set("is_pool", res.GetPayload().IsPool)
	d.Set("mark_utilized", res.GetPayload().MarkUtilized)
	if res.GetPayload().Status != nil {
		d.Set("status", res.GetPayload().Status.Value)
	}
	if res.GetPayload().Prefix != nil {
		d.Set("prefix", res.GetPayload().Prefix)
	}

	if res.GetPayload().Vrf != nil {
		d.Set("vrf_id", res.GetPayload().Vrf.ID)
	} else {
		d.Set("vrf_id", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if res.GetPayload().Vlan != nil {
		d.Set("vlan_id", res.GetPayload().Vlan.ID)
	} else {
		d.Set("vlan_id", nil)
	}

	if res.GetPayload().Role != nil {
		d.Set("role_id", res.GetPayload().Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	d.Set(customFieldsKey, res.GetPayload().CustomFields)
	// FIGURE OUT NESTED VRF AND NESTED VLAN (from maybe interfaces?)

	return nil
}

func resourceNetboxPrefixUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritablePrefix{}
	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	is_pool := d.Get("is_pool").(bool)
	mark_utilized := d.Get("mark_utilized").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = is_pool
	data.MarkUtilized = mark_utilized

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))
	data.CustomFields = d.Get(customFieldsKey)

	params := ipam.NewIpamPrefixesUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamPrefixesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamPrefixesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamPrefixesDelete(params, nil)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
