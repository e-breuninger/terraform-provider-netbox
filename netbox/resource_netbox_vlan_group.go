package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVlanGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVlanGroupCreate,
		Read:   resourceNetboxVlanGroupRead,
		Update: resourceNetboxVlanGroupUpdate,
		Delete: resourceNetboxVlanGroupDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):

> A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"min_vid": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 4093),
			},
			"max_vid": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(2, 4094),
			},
			"scope_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"dcim.location", "dcim.site", "dcim.sitegroup", "dcim.region", "dcim.rack", "virtualization.cluster", "virtualization.clustergroup"}, false),
			},
			"scope_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"scope_type"},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVlanGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.VLANGroup{}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)
	min_vid := int64(d.Get("min_vid").(int))
	max_vid := int64(d.Get("max_vid").(int))
	description := d.Get("description").(string)

	data.Name = &name
	data.Slug = &slug
	data.MinVid = min_vid
	data.MaxVid = max_vid
	data.Description = description

	if scopeType, ok := d.GetOk("scope_type"); ok {
		data.ScopeType = strToPtr(scopeType.(string))
	}

	if scopeID, ok := d.GetOk("scope_id"); ok {
		data.ScopeID = int64ToPtr(int64(scopeID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamVlanGroupsCreateParams().WithData(&data)
	res, err := api.Ipam.IpamVlanGroupsCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVlanGroupRead(d, m)
}

func resourceNetboxVlanGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlanGroupsReadParams().WithID(id)

	res, err := api.Ipam.IpamVlanGroupsRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamVlanGroupsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	vlanGroup := res.GetPayload()

	d.Set("name", vlanGroup.Name)
	d.Set("slug", vlanGroup.Slug)
	d.Set("min_vid", vlanGroup.MinVid)
	d.Set("max_vid", vlanGroup.MaxVid)
	d.Set("description", vlanGroup.Description)
	d.Set(tagsKey, getTagListFromNestedTagList(vlanGroup.Tags))

	if vlanGroup.ScopeType != nil {
		d.Set("scope_type", vlanGroup.ScopeType)
	}

	if vlanGroup.ScopeID != nil {
		d.Set("scope_id", vlanGroup.ScopeID)
	}

	return nil
}

func resourceNetboxVlanGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.VLANGroup{}

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)
	min_vid := int64(d.Get("min_vid").(int))
	max_vid := int64(d.Get("max_vid").(int))
	description := d.Get("description").(string)

	data.Name = &name
	data.Slug = &slug
	data.MinVid = min_vid
	data.MaxVid = max_vid
	data.Description = description

	if scopeType, ok := d.GetOk("scope_type"); ok {
		data.ScopeType = strToPtr(scopeType.(string))
	}

	if scopeID, ok := d.GetOk("scope_id"); ok {
		data.ScopeID = int64ToPtr(int64(scopeID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamVlanGroupsUpdateParams().WithID(id).WithData(&data)
	_, err := api.Ipam.IpamVlanGroupsUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxVlanGroupRead(d, m)
}

func resourceNetboxVlanGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlanGroupsDeleteParams().WithID(id)
	_, err := api.Ipam.IpamVlanGroupsDelete(params, nil)
	if err != nil {
		return err
	}

	return nil
}
