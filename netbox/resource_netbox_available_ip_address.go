package netbox

import (
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxAvailableIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAvailableIPAddressCreate,
		Read:   resourceNetboxAvailableIPAddressRead,
		Update: resourceNetboxAvailableIPAddressUpdate,
		Delete: resourceNetboxAvailableIPAddressDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):Per [the docs](https://netbox.readthedocs.io/en/stable/models/ipam/ipaddress/):

> An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.
> Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the "global" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchya.
>
> Each IP address can also be assigned an operational status and a functional role. Statuses are hard-coded in NetBox and include the following:
> * Active
> * Reserved
> * Deprecated
> * DHCP
> * SLAAC (IPv6 Stateless Address Autoconfiguration)

This resource will retrieve the next available IP address from a given prefix or IP range (specified by ID)`,

		Schema: map[string]*schema.Schema{
			"prefix_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"prefix_id", "ip_range_id", "prefix_ids", "ip_range_ids"},
			},
			"prefix_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				ExactlyOneOf: []string{"prefix_id", "ip_range_id", "prefix_ids", "ip_range_ids"},
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"ip_range_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"prefix_id", "ip_range_id", "prefix_ids", "ip_range_ids"},
			},
			"ip_range_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				ExactlyOneOf: []string{"prefix_id", "ip_range_id", "prefix_ids", "ip_range_ids"},
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"selected_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the prefix or IP range that was used to generate the IP address.",
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
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
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIPAddressStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIPAddressStatusOptions),
				Default:      "active",
			},
			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxIPAddressRoleOptions, false),
				Description:  buildValidValueDescription(resourceNetboxIPAddressRoleOptions),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

type ipamIPAvailableIpsCreateCreated interface {
	GetPayload() []*models.IPAddress
}

func createPayloadHandler(res ipamIPAvailableIpsCreateCreated) (int64, string, error) {
	if res == nil {
		// Ranges causes issues here.
		return 0, "", fmt.Errorf("payload is nil")
	}
	if len(res.GetPayload()) != 1 {
		return 0, "", fmt.Errorf("expected 1 ip address, got %d", len(res.GetPayload()))
	}
	return res.GetPayload()[0].ID, *res.GetPayload()[0].Address, nil
}

func resourceNetboxAvailableIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	var selectedID int64
	prefixID := int64(d.Get("prefix_id").(int))
	rangeID := int64(d.Get("ip_range_id").(int))
	prefixIDs, err := assertInterfaceToInt64Slice(d.Get("prefix_ids"))
	if err != nil {
		return fmt.Errorf("unable to convert prefixIDs to []int64: %w", err)
	}
	rangeIDs, err := assertInterfaceToInt64Slice(d.Get("ip_range_ids"))
	if err != nil {
		return fmt.Errorf("unable to convert rangeIDs to []int64: %w", err)
	}
	vrfID := int64(d.Get("vrf_id").(int))
	nestedvrf := models.NestedVRF{
		ID: vrfID,
	}
	data := models.AvailableIP{
		Vrf: &nestedvrf,
	}

	var res ipamIPAvailableIpsCreateCreated
	if prefixID != 0 {
		params := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(prefixID).WithData([]*models.AvailableIP{&data})
		res, _ = api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
		selectedID = prefixID
	}
	if rangeID != 0 {
		selectedID = rangeID
		params := ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(rangeID).WithData([]*models.AvailableIP{&data})
		res, err = api.Ipam.IpamIPRangesAvailableIpsCreate(params, nil)
		// Since we generated the ip_address, set that now
	}
	if len(prefixIDs) > 0 {

		for _, selectedID := range prefixIDs {
			// q: Ask for forgivnes or check first?
			params := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(selectedID).WithData([]*models.AvailableIP{&data})
			res, err = api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
			if err == nil {
				// There is avalible ips
				break
			}
		}
		if err != nil {
			return fmt.Errorf("unable to create a ip address for prefixes: %v, err: %w", prefixIDs, err)
		}
	}
	if len(rangeIDs) > 0 {
		// Try Ranges until one does not return an error
		for _, selectedID = range prefixIDs {
			params := ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(selectedID).WithData([]*models.AvailableIP{&data})
			res, err = api.Ipam.IpamIPRangesAvailableIpsCreate(params, nil)
			if err == nil {
				// There is avalible ips
				break
			}
		}
		if err != nil {
			return fmt.Errorf("unable to create a ip address for ip Ranges: %v, err: %w", prefixIDs, err)
		}

	}
	netboxID, ipaddress, err := createPayloadHandler(res)
	if err != nil {
		return fmt.Errorf("unable to handle payload: %w", err)
	}
	d.SetId(strconv.FormatInt(netboxID, 10))
	d.Set("ip_address", ipaddress)
	d.Set("selected_id", selectedID)
	return resourceNetboxAvailableIPAddressUpdate(d, m)
}

func assertInterfaceToInt64Slice(x interface{}) ([]int64, error) {
	var intSlice []int64
	var number int
	var ok bool
	xSlice, ok := x.([]interface{})
	if !ok {
		return nil, fmt.Errorf("assertInterfaceToInt64Slice: Unable to convert x:%v to []interface{}", x)
	}
	for _, v := range xSlice {
		if number, ok = v.(int); !ok {
			return nil, fmt.Errorf("assertSliceInterfaceToInt64: Unable to convert number:%v to int", v)
		}
		intSlice = append(intSlice, int64(number))
	}
	return intSlice, nil
}

func resourceNetboxAvailableIPAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
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
		case deviceInterfaceID != nil:
			d.Set("device_interface_id", ipAddress.AssignedObjectID)
		// if interfaceID is given, object_type must be set as well
		case interfaceID != nil:
			d.Set("object_type", ipAddress.AssignedObjectType)
			d.Set("interface_id", ipAddress.AssignedObjectID)
		}
	} else {
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

	d.Set("ip_address", ipAddress.Address)
	d.Set("description", ipAddress.Description)
	d.Set("status", ipAddress.Status.Value)
	d.Set(tagsKey, getTagListFromNestedTagList(ipAddress.Tags))
	return nil
}

func resourceNetboxAvailableIPAddressUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPAddress{}

	data.Address = strToPtr(d.Get("ip_address").(string))
	data.Status = d.Get("status").(string)

	data.Description = getOptionalStr(d, "description", false)
	data.Role = getOptionalStr(d, "role", false)
	data.DNSName = getOptionalStr(d, "dns_name", false)
	data.Vrf = getOptionalInt(d, "vrf_id")
	data.Tenant = getOptionalInt(d, "tenant_id")

	if interfaceID, ok := d.GetOk("interface_id"); ok {
		// The other possible type is dcim.interface for devices
		data.AssignedObjectType = strToPtr("virtualization.vminterface")
		data.AssignedObjectID = int64ToPtr(int64(interfaceID.(int)))
	}

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

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := ipam.NewIpamIPAddressesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Ipam.IpamIPAddressesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxAvailableIPAddressRead(d, m)
}

func resourceNetboxAvailableIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

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
