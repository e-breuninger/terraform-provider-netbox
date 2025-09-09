package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceNetboxVpnTunnelTerminationRoleOptions = []string{"peer", "hub", "spoke"}

func resourceNetboxVpnTunnelTermination() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxVpnTunnelTerminationCreate,
		Read:   resourceNetboxVpnTunnelTerminationRead,
		Update: resourceNetboxVpnTunnelTerminationUpdate,
		Delete: resourceNetboxVpnTunnelTerminationDelete,

		Description: `:meta:subcategory:VPN Tunnels:From the [official documentation](https://docs.netbox.dev/en/stable/features/vpn-tunnels/):

> NetBox can model private tunnels formed among virtual termination points across your network. Typical tunnel implementations include GRE, IP-in-IP, and IPSec. A tunnel may be terminated to two or more device or virtual machine interfaces. For convenient organization, tunnels may be assigned to user-defined groups.`,

		Schema: map[string]*schema.Schema{
			"tunnel_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"role": {
				Type:        schema.TypeString,
				Required:    true,
				Description: buildValidValueDescription(resourceNetboxVpnTunnelTerminationRoleOptions),
			},
			"virtual_machine_interface_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"virtual_machine_interface_id", "device_interface_id"},
			},
			"device_interface_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"virtual_machine_interface_id", "device_interface_id"},
			},
			"outside_ip_address_id": {
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

func resourceNetboxVpnTunnelTerminationCreate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	data := models.WritableTunnelTermination{}

	tunnelID := int64(d.Get("tunnel_id").(int))

	data.Tunnel = &tunnelID
	data.Role = d.Get("role").(string)

	vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
	deviceInterfaceID := getOptionalInt(d, "device_interface_id")

	switch {
	case vmInterfaceID != nil:
		data.TerminationType = strToPtr("virtualization.vminterface")
		data.TerminationID = *vmInterfaceID
	case deviceInterfaceID != nil:
		data.TerminationType = strToPtr("dcim.interface")
		data.TerminationID = *deviceInterfaceID
	}

	data.OutsideIP = getOptionalInt(d, "outside_ip_address_id")

	tags, _ := getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	data.Tags = tags

	params := vpn.NewVpnTunnelTerminationsCreateParams().WithData(&data)

	res, err := api.Vpn.VpnTunnelTerminationsCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxVpnTunnelTerminationRead(d, m)
}

func resourceNetboxVpnTunnelTerminationRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnTunnelTerminationsReadParams().WithID(id)

	res, err := api.Vpn.VpnTunnelTerminationsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnTunnelTerminationsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	tunnelTermination := res.GetPayload()
	d.Set("tunnel_id", tunnelTermination.Tunnel.ID)
	d.Set("role", tunnelTermination.Role.Value)

	vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
	deviceInterfaceID := getOptionalInt(d, "device_interface_id")

	switch {
	case vmInterfaceID != nil:
		d.Set("virtual_machine_interface_id", tunnelTermination.TerminationID)
	case deviceInterfaceID != nil:
		d.Set("device_interface_id", tunnelTermination.TerminationID)
	}

	if tunnelTermination.OutsideIP != nil {
		d.Set("outside_ip_address_id", tunnelTermination.OutsideIP.ID)
	}

	state.readTags(d, res.GetPayload().Tags)
	return nil
}

func resourceNetboxVpnTunnelTerminationUpdate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableTunnelTermination{}

	tunnelID := int64(d.Get("tunnel_id").(int))
	data.Tunnel = &tunnelID
	data.Role = d.Get("role").(string)

	vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
	deviceInterfaceID := getOptionalInt(d, "device_interface_id")

	switch {
	case vmInterfaceID != nil:
		data.TerminationType = strToPtr("virtualization.vminterface")
		data.TerminationID = *vmInterfaceID
	case deviceInterfaceID != nil:
		data.TerminationType = strToPtr("dcim.interface")
		data.TerminationID = *deviceInterfaceID
	}

	data.OutsideIP = getOptionalInt(d, "outside_ip_address_id")

	tags, _ := getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	data.Tags = tags

	params := vpn.NewVpnTunnelTerminationsUpdateParams().WithID(id).WithData(&data)

	_, err := api.Vpn.VpnTunnelTerminationsUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxVpnTunnelTerminationRead(d, m)
}

func resourceNetboxVpnTunnelTerminationDelete(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := vpn.NewVpnTunnelTerminationsDeleteParams().WithID(id)

	_, err := api.Vpn.VpnTunnelTerminationsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*vpn.VpnTunnelTerminationsDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
