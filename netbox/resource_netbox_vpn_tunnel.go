package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceNetboxVpnTunnelEncapsulationOptions = []string{"ipsec-transport", "ipsec-tunnel", "ip-ip", "gre"}
var resourceNetboxVpnTunnelStatusOptions = []string{"planned", "active", "disabled"}

func resourceNetboxVpnTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVpnTunnelCreate,
		Read:   resourceNetboxVpnTunnelRead,
		Update: resourceNetboxVpnTunnelUpdate,
		Delete: resourceNetboxVpnTunnelDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/features/vpn-tunnels/):

> NetBox can model private tunnels formed among virtual termination points across your network. Typical tunnel implementations include GRE, IP-in-IP, and IPSec. A tunnel may be terminated to two or more device or virtual machine interfaces. For convenient organization, tunnels may be assigned to user-defined groups.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"encapsulation": {
				Type:        schema.TypeString,
				Required:    true,
				Description: buildValidValueDescription(resourceNetboxVpnTunnelEncapsulationOptions),
			},
			"status": {
				Type:        schema.TypeString,
				Required:    true,
				Description: buildValidValueDescription(resourceNetboxVpnTunnelStatusOptions),
			},
			"tunnel_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxVpnTunnelCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableTunnel{}

	data.Name = strToPtr(d.Get("name").(string))
	data.Encapsulation = strToPtr(d.Get("encapsulation").(string))
	data.Status = strToPtr(d.Get("status").(string))
	data.Group = int64ToPtr(int64(d.Get("tunnel_group_id").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Tenant = getOptionalInt(d, "tenant_id")
	data.TunnelID = getOptionalInt(d, "tunnel_id")

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags

	params := vpn.NewVpnTunnelsCreateParams().WithData(&data)

	res, err := api.Vpn.VpnTunnelsCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVpnTunnelRead(d, m)
}

func resourceNetboxVpnTunnelRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnTunnelsReadParams().WithID(id)

	res, err := api.Vpn.VpnTunnelsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnTunnelsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	tunnel := res.GetPayload()
	d.Set("name", tunnel.Name)
	d.Set("encapsulation", tunnel.Encapsulation.Value)
	d.Set("status", tunnel.Status.Value)

	if tunnel.Group != nil {
		d.Set("tunnel_group_id", tunnel.Group.ID)
	} else {
		d.Set("tunnel_group_id", nil)
	}

	if tunnel.Tenant != nil {
		d.Set("tenant_id", tunnel.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	d.Set("tunnel_id", tunnel.TunnelID)

	d.Set("description", tunnel.Description)

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxVpnTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableTunnel{}

	data.Name = strToPtr(d.Get("name").(string))
	data.Encapsulation = strToPtr(d.Get("encapsulation").(string))
	data.Status = strToPtr(d.Get("status").(string))
	data.Group = int64ToPtr(int64(d.Get("tunnel_group_id").(int)))

	data.Description = getOptionalStr(d, "description", false)
	data.Tenant = getOptionalInt(d, "tenant_id")
	data.TunnelID = getOptionalInt(d, "tunnel_id")

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
	data.Tags = tags

	params := vpn.NewVpnTunnelsUpdateParams().WithID(id).WithData(&data)

	_, err := api.Vpn.VpnTunnelsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVpnTunnelRead(d, m)
}

func resourceNetboxVpnTunnelDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnTunnelsDeleteParams().WithID(id)

	_, err := api.Vpn.VpnTunnelsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnTunnelsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
