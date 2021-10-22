package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVlanCreate,
		Read:   resourceNetboxVlanRead,
		Update: resourceNetboxVlanUpdate,
		Delete: resourceNetboxVlanDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated"}, false),
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				Set:      schema.HashString,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxVlanCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableVLAN{}

	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))

	data.Name = &name
	data.Vid = &vid
	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	if v, ok := d.GetOk("site_id"); ok {
		siteID := int64(v.(int))
		data.Site = &siteID
	}
	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		data.Tenant = &tenantID
	}
	if v, ok := d.GetOk("status"); ok {
		data.Status = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		data.Description = v.(string)
	}
	if v, ok := d.GetOk("role_id"); ok {
		roleID := int64(v.(int))
		data.Role = &roleID
	}

	params := ipam.NewIpamVlansCreateParams().WithData(&data)
	res, err := api.Ipam.IpamVlansCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVlanUpdate(d, m)
}

func resourceNetboxVlanRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	params := ipam.NewIpamVlansReadParams().WithID(id)

	res, err := api.Ipam.IpamVlansRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamVlansReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("vid", res.GetPayload().Vid)
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	}

	if res.GetPayload().Description != "" {
		d.Set("description", res.GetPayload().Description)
	}

	if res.GetPayload().Status != nil {
		d.Set("status", res.GetPayload().Status.Value)
	}

	if res.GetPayload().Role != nil {
		d.Set("role_id", res.GetPayload().Role.ID)
	}

	return nil
}

func resourceNetboxVlanUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err1 := strconv.ParseInt(d.Id(), 10, 64)
	if err1 != nil {
		return err1
	}
	data := models.WritableVLAN{}

	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))

	data.Name = &name
	data.Vid = &vid
	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		data.Site = &siteID
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		data.Tenant = &tenantID
	}

	if d.HasChange("status") {
		data.Status = d.Get("status").(string)
	}

	if d.HasChange("description") {
		data.Description = d.Get("description").(string)
	}

	if d.HasChange("role_id") {
		roleID := int64(d.Get("role_id").(int))
		data.Role = &roleID
	}

	params := ipam.NewIpamVlansUpdateParams().WithID(id).WithData(&data)
	_, err2 := api.Ipam.IpamVlansUpdate(params, nil)
	if err2 != nil {
		return err2
	}

	return resourceNetboxVlanRead(d, m)
}

func resourceNetboxVlanDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, err1 := strconv.ParseInt(d.Id(), 10, 64)
	if err1 != nil {
		return err1
	}

	params := ipam.NewIpamVlansDeleteParams().WithID(id)
	_, err2 := api.Ipam.IpamVlansDelete(params, nil)
	if err2 != nil {
		return err2
	}

	return nil
}
