package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxVpnTunnelGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVpnTunnelGroupCreate,
		Read:   resourceNetboxVpnTunnelGroupRead,
		Update: resourceNetboxVpnTunnelGroupUpdate,
		Delete: resourceNetboxVpnTunnelGroupDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/features/vpn-tunnels/):

> NetBox can model private tunnels formed among virtual termination points across your network. Typical tunnel implementations include GRE, IP-in-IP, and IPSec. A tunnel may be terminated to two or more device or virtual machine interfaces. For convenient organization, tunnels may be assigned to user-defined groups.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringLenBetween(0, 100),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVpnTunnelGroupCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.TunnelGroup{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	data.Slug = &slug

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	}

	data.Tags = []*models.NestedTag{}

	params := vpn.NewVpnTunnelGroupsCreateParams().WithData(&data)

	res, err := api.Vpn.VpnTunnelGroupsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVpnTunnelGroupRead(d, m)
}

func resourceNetboxVpnTunnelGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnTunnelGroupsReadParams().WithID(id)

	res, err := api.Vpn.VpnTunnelGroupsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnTunnelGroupsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("slug", res.GetPayload().Slug)
	d.Set("description", res.GetPayload().Description)
	return nil
}

func resourceNetboxVpnTunnelGroupUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.TunnelGroup{}

	name := d.Get("name").(string)
	data.Name = &name

	slugValue, slugOk := d.GetOk("slug")
	var slug string
	// Default slug to generated slug if not given
	if !slugOk {
		slug = getSlug(name)
	} else {
		slug = slugValue.(string)
	}
	data.Slug = &slug

	if d.HasChange("description") {
		// description omits empty values so set to ' '
		if description := d.Get("description"); description.(string) == "" {
			data.Description = " "
		} else {
			data.Description = description.(string)
		}
	}

	data.Tags = []*models.NestedTag{}

	params := vpn.NewVpnTunnelGroupsUpdateParams().WithID(id).WithData(&data)

	_, err := api.Vpn.VpnTunnelGroupsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVpnTunnelGroupRead(d, m)
}

func resourceNetboxVpnTunnelGroupDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnTunnelGroupsDeleteParams().WithID(id)

	_, err := api.Vpn.VpnTunnelGroupsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnTunnelGroupsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
