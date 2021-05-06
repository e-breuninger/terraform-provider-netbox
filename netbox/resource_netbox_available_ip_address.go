package netbox

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxAvailableIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAvailableIPAddressCreate,
		Read:   resourceNetboxAvailableIPAddressRead,
		Update: resourceNetboxAvailableIPAddressUpdate,
		Delete: resourceNetboxAvailableIPAddressDelete,

		Schema: map[string]*schema.Schema{
			"prefix_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "reserved", "deprecated", "dhcp"}, false),
			},
			"dns_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxAvailableIPAddressCreate(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.NetBoxAPI)
	prefix_id := int64(d.Get("prefix_id").(int))

	// Get all the available IPs from the given prefix
	ipsread := &ipam.IpamPrefixesAvailableIpsReadParams{
		ID:      prefix_id,
		Context: context.Background(),
	}

	resp, err := api.Ipam.IpamPrefixesAvailableIpsRead(ipsread, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("Problem getting IPs from prefix %d: %v", prefix_id, err))
	}

	data := models.WritableIPAddress{}
	// Use the first Payload Address as our data.Address value
	data.Address = &resp.Payload[0].Address
	data.Status = d.Get("status").(string)

	if dnsName, ok := d.GetOk("dns_name"); ok {
		data.DNSName = dnsName.(string)
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := ipam.NewIpamIPAddressesCreateParams().WithData(&data)

	res, err := api.Ipam.IpamIPAddressesCreate(params, nil)
	if err != nil {
		return err
	}

	// Since we generated the ip_address set that now
	d.Set("ip_address", res.Payload.Address)
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxAvailableIPAddressUpdate(d, m)
}

func resourceNetboxAvailableIPAddressRead(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamIPAddressesReadParams().WithID(id)

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		var apiError *runtime.APIError
		if errors.As(err, &apiError) {
			errorcode := err.(*runtime.APIError).Response.(runtime.ClientResponse).Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if res.GetPayload().AssignedObject != nil {
		d.Set("interface_id", res.GetPayload().AssignedObject.ID)
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
	d.Set("status", res.GetPayload().Status.Value)
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxAvailableIPAddressUpdate(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableIPAddress{}

	ipAddress := d.Get("ip_address").(string)
	status := d.Get("status").(string)

	data.Status = status
	data.Address = &ipAddress

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
		data.AssignedObjectType = "virtualization.vminterface"
		data.AssignedObjectID = int64ToPtr(int64(interfaceID.(int)))
	}

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get("tags"))

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
		return err
	}
	return nil
}
