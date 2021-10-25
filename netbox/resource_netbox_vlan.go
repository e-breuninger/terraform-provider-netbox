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
	status := d.Get("status").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Vid = &vid
	data.Status = status
	data.Description = description

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

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
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
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

	if res.GetPayload().Name != nil {
		d.Set("name", res.GetPayload().Name)
	}

	if res.GetPayload().Vid != nil {
		d.Set("vid", res.GetPayload().Vid)
	}

	if res.GetPayload().Status != nil {
		d.Set("status", res.GetPayload().Status.Value)
	}

	if res.GetPayload().Description != "" {
		d.Set("description", res.GetPayload().Description)
	}

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	}

	if res.GetPayload().Role != nil {
		d.Set("role_id", res.GetPayload().Role.ID)
	}

	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))

	return nil
}

func resourceNetboxVlanUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableVLAN{}
	name := d.Get("name").(string)
	vid := int64(d.Get("vid").(int))
	status := d.Get("status").(string)
	description := d.Get("description").(string)

	data.Name = &name
	data.Vid = &vid

	data.Status = status
	data.Description = description

	if siteID, ok := d.GetOk("site_id"); ok {
		data.Site = int64ToPtr(int64(siteID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := ipam.NewIpamVlansUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamVlansUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxVlanRead(d, m)
}

func resourceNetboxVlanDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlansDeleteParams().WithID(id)
	_, err := api.Ipam.IpamVlansDelete(params, nil)
	if err != nil {
		return err
	}

	return nil
}
