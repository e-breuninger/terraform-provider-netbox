package netbox

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxIPAddressObjectTypeOptions = []string{"virtualization.vminterface", "dcim.interface", "ipam.fhrpgroup"}
var resourceNetboxIPAddressStatusOptions = []string{"active", "reserved", "deprecated", "dhcp", "slaac"}
var resourceNetboxIPAddressRoleOptions = []string{"loopback", "secondary", "anycast", "vip", "vrrp", "hsrp", "glbp", "carp"}
var resourceNetboxIPAddressAllocationSources = []string{"ip_address", "prefix_id", "ip_range_id"}

// forceNewOnIPAddressAllocationSourceMove replaces the address only when
// prefix_id/ip_range_id really moves from one prefix or range to another.
//
// Netbox does not report which prefix or range an address was allocated from,
// so an imported address has 0 in state for both. An old value of 0 therefore
// means "source unknown", not "moved", and neither does a new value that is not
// known until apply. See forceNewOnAllocationSourceMove on
// netbox_available_ip_address, which this mirrors.
func forceNewOnIPAddressAllocationSourceMove(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	for _, key := range []string{"prefix_id", "ip_range_id"} {
		oldValue, newValue := d.GetChange(key)
		if !d.NewValueKnown(key) || oldValue.(int) == 0 || oldValue.(int) == newValue.(int) {
			continue
		}
		if err := d.ForceNew(key); err != nil {
			return err
		}
	}
	return nil
}

func resourceNetboxIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIPAddressCreate,
		Read:   resourceNetboxIPAddressRead,
		Update: resourceNetboxIPAddressUpdate,
		Delete: resourceNetboxIPAddressDelete,

		CustomizeDiff: forceNewOnIPAddressAllocationSourceMove,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses):

> An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
>
> Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchy.

Either set ` + "`ip_address`" + ` directly, or set ` + "`prefix_id`" + `/` + "`ip_range_id`" + ` to have Netbox allocate the next free address from that prefix or range (mirrors ` + "`netbox_available_ip_address`" + `).`,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsCIDR,
				ExactlyOneOf: resourceNetboxIPAddressAllocationSources,
			},
			// Not ForceNew: replacement is conditional, see
			// forceNewOnIPAddressAllocationSourceMove. Write-only allocation
			// instructions, not attributes of the result - Netbox cannot report
			// which prefix or range an address came from, so Read never sets
			// these.
			"prefix_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: resourceNetboxIPAddressAllocationSources,
			},
			"ip_range_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: resourceNetboxIPAddressAllocationSources,
			},
			"interface_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"object_type"},
			},
			"object_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIPAddressObjectTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIPAddressObjectTypeOptions),
				RequiredWith: []string{"interface_id"},
			},
			"virtual_machine_interface_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"interface_id", "device_interface_id"},
			},
			"device_interface_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"interface_id", "virtual_machine_interface_id"},
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIPAddressStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIPAddressStatusOptions),
			},
			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
				// NetBox always converts DNS names to lowercase
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
			tagsKey: tagsSchema,
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIPAddressRoleOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIPAddressRoleOptions),
			},
			"nat_inside_address_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nat_outside_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address_family": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// allocateIPAddress takes the next free address from prefix_id/ip_range_id via
// Netbox's available-ips endpoint - the same call netbox_available_ip_address's
// Create makes - and seeds d.Id()/ip_address from it. The caller still has to
// apply every other attribute with an Update.
func allocateIPAddress(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.AvailableIP{
		Vrf: &models.NestedVRF{ID: int64(d.Get("vrf_id").(int))},
	}

	if prefixID := int64(d.Get("prefix_id").(int)); prefixID != 0 {
		params := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(prefixID).WithData([]*models.AvailableIP{&data})
		res, err := api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
		if err != nil {
			return err
		}
		if len(res.Payload) == 0 {
			return fmt.Errorf("no available IP addresses in prefix %d", prefixID)
		}
		d.SetId(strconv.FormatInt(res.Payload[0].ID, 10))
		d.Set("ip_address", *res.Payload[0].Address)
		return nil
	}

	rangeID := int64(d.Get("ip_range_id").(int))
	params := ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(rangeID).WithData([]*models.AvailableIP{&data})
	res, err := api.Ipam.IpamIPRangesAvailableIpsCreate(params, nil)
	if err != nil {
		return err
	}
	if len(res.Payload) == 0 {
		return fmt.Errorf("no available IP addresses in IP range %d", rangeID)
	}
	d.SetId(strconv.FormatInt(res.Payload[0].ID, 10))
	d.Set("ip_address", *res.Payload[0].Address)
	return nil
}

func resourceNetboxIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	if d.Get("ip_address").(string) == "" {
		if err := allocateIPAddress(d, m); err != nil {
			return err
		}
		return resourceNetboxIPAddressUpdate(d, m)
	}

	api := m.(*providerState)

	data := models.WritableIPAddress{}

	data.Address = strToPtr(d.Get("ip_address").(string))
	data.Status = d.Get("status").(string)

	data.Description = getOptionalStr(d, "description", false)
	data.Role = getOptionalStr(d, "role", false)
	data.DNSName = getOptionalStr(d, "dns_name", false)
	data.Vrf = getOptionalInt(d, "vrf_id")
	data.Tenant = getOptionalInt(d, "tenant_id")
	data.NatInside = getOptionalInt(d, "nat_inside_address_id")

	vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
	deviceInterfaceID := getOptionalInt(d, "device_interface_id")
	interfaceID := getOptionalInt(d, "interface_id")

	switch {
	case vmInterfaceID != nil:
		data.AssignedObjectType = strToPtr("virtualization.vminterface")
		data.AssignedObjectID = vmInterfaceID
	case deviceInterfaceID != nil:
		data.AssignedObjectType = strToPtr("dcim.interface")
		data.AssignedObjectID = deviceInterfaceID
	// if interfaceID is given, object_type must be set as well
	case interfaceID != nil:
		data.AssignedObjectType = strToPtr(d.Get("object_type").(string))
		data.AssignedObjectID = interfaceID
	// default = ip is not linked to anything
	default:
		data.AssignedObjectType = strToPtr("")
		data.AssignedObjectID = nil
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamIPAddressesCreateParams().WithData(&data)

	res, err := api.Ipam.IpamIPAddressesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxIPAddressRead(d, m)
}

func resourceNetboxIPAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamIPAddressesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	ipAddress := res.GetPayload()
	if ipAddress.AssignedObjectID != nil {
		vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
		deviceInterfaceID := getOptionalInt(d, "device_interface_id")
		interfaceID := getOptionalInt(d, "interface_id")

		switch {
		case vmInterfaceID != nil:
			d.Set("virtual_machine_interface_id", ipAddress.AssignedObjectID)
			d.Set("interface_id", nil)
			d.Set("object_type", "")
		case deviceInterfaceID != nil:
			d.Set("device_interface_id", ipAddress.AssignedObjectID)
			d.Set("interface_id", nil)
			d.Set("object_type", "")
		// if interfaceID is given, object_type must be set as well
		case interfaceID != nil:
			d.Set("interface_id", ipAddress.AssignedObjectID)
			d.Set("object_type", ipAddress.AssignedObjectType)
		default:
			// Set changes made to the ip address outside of Terraform and update the state accordingly.
			d.Set("interface_id", ipAddress.AssignedObjectID)
			d.Set("object_type", ipAddress.AssignedObjectType)
		}
	} else {
		d.Set("virtual_machine_interface_id", nil)
		d.Set("device_interface_id", nil)
		d.Set("interface_id", nil)
		d.Set("object_type", "")
	}

	if ipAddress.Vrf != nil {
		d.Set("vrf_id", ipAddress.Vrf.ID)
	} else {
		d.Set("vrf_id", nil)
	}

	if ipAddress.Tenant != nil {
		d.Set("tenant_id", ipAddress.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if ipAddress.DNSName != "" {
		d.Set("dns_name", ipAddress.DNSName)
	}

	if ipAddress.Role != nil {
		d.Set("role", ipAddress.Role.Value)
	} else {
		d.Set("role", nil)
	}

	if ipAddress.NatInside != nil {
		d.Set("nat_inside_address_id", ipAddress.NatInside.ID)
	} else {
		d.Set("nat_inside_address_id", nil)
	}

	if ipAddress.NatOutside != nil {
		natOutsideIPAddresses := ipAddress.NatOutside

		var s []map[string]interface{}
		for _, v := range natOutsideIPAddresses {
			var mapping = make(map[string]interface{})

			mapping["id"] = v.ID
			mapping["ip_address"] = v.Address
			mapping["address_family"] = v.Family

			s = append(s, mapping)
		}
		d.Set("nat_outside_addresses", s)
	} else {
		d.Set("nat_outside_addresses", nil)
	}

	d.Set("ip_address", ipAddress.Address)
	d.Set("description", ipAddress.Description)
	d.Set("status", ipAddress.Status.Value)
	api.readTags(d, ipAddress.Tags)
	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	return nil
}

func resourceNetboxIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPAddress{}

	data.Address = strToPtr(d.Get("ip_address").(string))
	data.Status = d.Get("status").(string)

	data.Description = getOptionalStr(d, "description", true)
	data.Role = getOptionalStr(d, "role", false)
	data.DNSName = getOptionalStr(d, "dns_name", true)
	data.Vrf = getOptionalInt(d, "vrf_id")
	data.Tenant = getOptionalInt(d, "tenant_id")
	data.NatInside = getOptionalInt(d, "nat_inside_address_id")

	vmInterfaceID := getOptionalInt(d, "virtual_machine_interface_id")
	deviceInterfaceID := getOptionalInt(d, "device_interface_id")
	interfaceID := getOptionalInt(d, "interface_id")

	switch {
	case vmInterfaceID != nil:
		data.AssignedObjectType = strToPtr("virtualization.vminterface")
		data.AssignedObjectID = vmInterfaceID
	case deviceInterfaceID != nil:
		data.AssignedObjectType = strToPtr("dcim.interface")
		data.AssignedObjectID = deviceInterfaceID
	// if interfaceID is given, object_type must be set as well
	case interfaceID != nil:
		data.AssignedObjectType = strToPtr(d.Get("object_type").(string))
		data.AssignedObjectID = interfaceID
	// default = ip is not linked to anything
	default:
		data.AssignedObjectType = strToPtr("")
		data.AssignedObjectID = nil
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamIPAddressesUpdateParams().WithID(id).WithData(&data)

	_, err = api.Ipam.IpamIPAddressesUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIPAddressRead(d, m)
}

func resourceNetboxIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesDeleteParams().WithID(id)

	_, err := api.Ipam.IpamIPAddressesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamIPAddressesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
