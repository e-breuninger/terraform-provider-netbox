package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxFhrpGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxFhrpGroupCreate,
		Read:   resourceNetboxFhrpGroupRead,
		Update: resourceNetboxFhrpGroupUpdate,
		Delete: resourceNetboxFhrpGroupDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/fhrpgroup/):`,

		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Protocol",
			},
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Group ID",
			},
			"auth_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Auth type",
			},
			"auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Authentication key",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments",
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxFhrpGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.FHRPGroup{}

	protocol := d.Get("protocol").(string)
	data.Protocol = &protocol

	group_id := int64(d.Get("group_id").(int))
	data.GroupID = &group_id

	data.AuthType = d.Get("auth_type").(string)
	data.AuthKey = d.Get("auth_key").(string)
	data.Name = d.Get("name").(string)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)
	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := ipam.NewIpamFhrpGroupsCreateParams().WithData(&data)

	res, err := api.Ipam.IpamFhrpGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxFhrpGroupRead(d, m)
}

func resourceNetboxFhrpGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamFhrpGroupsReadParams().WithID(id)

	res, err := api.Ipam.IpamFhrpGroupsRead(params, nil)

	if err != nil {
		if errresp, ok := err.(*ipam.IpamAsnsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	fhrpgroup := res.GetPayload()

	d.Set("protocol", fhrpgroup.Protocol)
	d.Set("group_id", fhrpgroup.GroupID)
	d.Set("auth_type", fhrpgroup.AuthType)
	d.Set("auth_key", fhrpgroup.AuthKey)
	d.Set("name", fhrpgroup.Name)
	d.Set("description", fhrpgroup.Description)
	d.Set("comments", fhrpgroup.Comments)
	api.readTags(d, fhrpgroup.Tags)

	return nil
}

func resourceNetboxFhrpGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.FHRPGroup{}

	protocol := d.Get("protocol").(string)
	data.Protocol = &protocol

	group_id := int64(d.Get("group_id").(int))
	data.GroupID = &group_id

	data.AuthType = d.Get("auth_type").(string)
	data.AuthKey = d.Get("auth_key").(string)
	data.Name = d.Get("name").(string)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := ipam.NewIpamFhrpGroupsUpdateParams().WithID(id).WithData(&data)

	_, err = api.Ipam.IpamFhrpGroupsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxFhrpGroupRead(d, m)
}

func resourceNetboxFhrpGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamFhrpGroupsDeleteParams().WithID(id)

	_, err := api.Ipam.IpamFhrpGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamFhrpGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
