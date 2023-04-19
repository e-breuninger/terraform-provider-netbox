package netbox

import (
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
				ExactlyOneOf: []string{"prefix_id", "ip_range_id"},
			},
			"ip_range_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface_id": {
				Type:     schema.TypeInt,
				Optional: true,
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
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "dhcp", "slaac"}, false),
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
				ValidateFunc: validation.StringInSlice([]string{"loopback", "secondary", "anycast", "vip", "vrrp", "hsrp", "glbp", "carp"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxAvailableIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	prefixId := int64(d.Get("prefix_id").(int))
	vrfId := int64(int64(d.Get("vrf_id").(int)))
	rangeId := int64(d.Get("ip_range_id").(int))
	nestedvrf := models.NestedVRF{
		ID: vrfId,
	}
	data := models.AvailableIP{
		Vrf: &nestedvrf,
	}
	if prefixId != 0 {
		params := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(prefixId).WithData([]*models.AvailableIP{&data})
		res, _ := api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
		// Since we generated the ip_address set that now
		d.SetId(strconv.FormatInt(res.Payload[0].ID, 10))
		d.Set("ip_address", *res.Payload[0].Address)
	}
	if rangeId != 0 {
		params := ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(rangeId).WithData([]*models.AvailableIP{&data})
		res, _ := api.Ipam.IpamIPRangesAvailableIpsCreate(params, nil)
		// Since we generated the ip_address set that now
		d.SetId(strconv.FormatInt(res.Payload[0].ID, 10))
		d.Set("ip_address", *res.Payload[0].Address)
	}
	return resourceNetboxAvailableIPAddressUpdate(d, m)
}

func resourceNetboxAvailableIPAddressRead(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		errorcode := err.(*ipam.IpamIPAddressesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	if res.GetPayload().AssignedObjectID != nil {
		d.Set("interface_id", res.GetPayload().AssignedObjectID)
	} else {
		d.Set("interface_id", nil)
	}

	if res.GetPayload().Vrf != nil {
		d.Set("vrf_id", res.GetPayload().Vrf.ID)
	} else {
		d.Set("vrf_id", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if res.GetPayload().DNSName != "" {
		d.Set("dns_name", res.GetPayload().DNSName)
	}

	d.Set("ip_address", res.GetPayload().Address)
	d.Set("description", res.GetPayload().Description)
	d.Set("status", res.GetPayload().Status.Value)
	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxAvailableIPAddressUpdate(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPAddress{}

	ipAddress := d.Get("ip_address").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	role := d.Get("role").(string)

	data.Status = status
	data.Description = description
	data.Address = &ipAddress
	data.Role = role

	if d.HasChange("dns_name") {
		// WritableIPAddress omits empty values so set to ' '
		if dnsName := d.Get("dns_name"); dnsName.(string) == "" {
			data.DNSName = " "
		} else {
			data.DNSName = dnsName.(string)
		}
	}

	if interfaceID, ok := d.GetOk("interface_id"); ok {
		// The other possible type is dcim.interface for devices
		data.AssignedObjectType = strToPtr("virtualization.vminterface")
		data.AssignedObjectID = int64ToPtr(int64(interfaceID.(int)))
	}

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
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
