package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxAvailableVLAN() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAvailableVLANCreate,
		Read:   resourceNetboxAvailableVLANRead,
		Update: resourceNetboxAvailableVLANUpdate,
		Delete: resourceNetboxAvailableVLANDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):Per [the docs](https://netbox.readthedocs.io/en/stable/models/ipam/vlan/):

> A VLAN represents an isolated Layer 2 domain identified by a numeric ID (1–4094). VLANs may be assigned to specific sites or marked as global.
> Optionally, they can be organized within VLAN groups to define scope and enforce uniqueness.
>
> Each VLAN can also be assigned an operational status and a functional role. Statuses are hard-coded in NetBox and include the following:
> * Active
> * Reserved
> * Deprecated

This resource will retrieve the next available VLAN ID from a given VLAN group (specified by ID).`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id": {
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vid": {
				Type:     schema.TypeInt,
				Computed: true, // it's auto-assigned by NetBox, not user-supplied
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			tagsKey: tagsSchema,
		},
	}
}

func resourceNetboxAvailableVLANCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	groupID := int64(d.Get("group_id").(int))

	tags, err := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	if err != nil {
		return err
	}
	data := &models.WritableCreateAvailableVLAN{
		Name:        strToPtr(d.Get("name").(string)),
		Description: getOptionalStr(d, "description", false),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Site:        getOptionalInt(d, "site_id"),
		Role:        getOptionalInt(d, "role_id"),
		Status:      d.Get("status").(string),
		Tags:        tags,
	}

	params := ipam.NewIpamVlanGroupsAvailableVlansCreateParams().WithID(groupID).WithData(data)
	resp, err := api.Ipam.IpamVlanGroupsAvailableVlansCreate(params, nil)
	if err != nil {
		return err
	}

	vlan := resp.Payload
	d.SetId(strconv.FormatInt(vlan.ID, 10))
	d.Set("vid", vlan.Vid)
	d.Set("name", vlan.Name)
	d.Set("group_id", vlan.Group.ID)
	return resourceNetboxAvailableVLANRead(d, m)
}

func resourceNetboxAvailableVLANRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamVlansReadParams().WithID(id)

	res, err := api.Ipam.IpamVlansRead(params, nil)
	if err != nil {
		if erresp, ok := err.(*ipam.IpamVlansReadDefault); ok {
			errorcode := erresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	vlan := res.GetPayload()
	// Required fields
	d.Set("vid", vlan.Vid)
	d.Set("name", vlan.Name)

	// Optional fields
	d.Set("description", vlan.Description)
	d.Set("comments", vlan.Comments)

	if vlan.Status != nil && vlan.Status.Value != nil {
		d.Set("status", *vlan.Status.Value)
	} else {
		d.Set("status", "")
	}

	if vlan.Tenant != nil {
		d.Set("tenant_id", vlan.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if vlan.Site != nil {
		d.Set("site_id", vlan.Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if vlan.Group != nil {
		d.Set("group_id", vlan.Group.ID)
	} else {
		d.Set("group_id", nil)
	}

	if vlan.Role != nil {
		d.Set("role_id", vlan.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	api.readTags(d, vlan.Tags)

	return nil
}

func resourceNetboxAvailableVLANUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	data := &models.WritableVLAN{
		Name:        strToPtr(d.Get("name").(string)),
		Description: getOptionalStr(d, "description", false),
		Tenant:      getOptionalInt(d, "tenant_id"),
		Site:        getOptionalInt(d, "site_id"),
		Group:       getOptionalInt(d, "group_id"),
		Role:        getOptionalInt(d, "role_id"),
		Status:      d.Get("status").(string),
		Vid:         int64ToPtr(int64(d.Get("vid").(int))),
	}

	var err_tags error
	data.Tags, err_tags = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	if err_tags != nil {
		return err_tags
	}

	params := ipam.NewIpamVlansUpdateParams().
		WithID(id).
		WithData(data)

	_, err := api.Ipam.IpamVlansUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxAvailableVLANRead(d, m)
}

func resourceNetboxAvailableVLANDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := ipam.NewIpamVlansDeleteParams().WithID(id)
	_, err := api.Ipam.IpamVlansDelete(params, nil)

	if err != nil {
		if errresp, ok := err.(*ipam.IpamVlansDeleteDefault); ok && errresp.Code() == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}
