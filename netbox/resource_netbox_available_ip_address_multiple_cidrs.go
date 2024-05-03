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

func resourceNetboxAvailableIPAddressMultiplecidrs() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAvailableIPAddressMultiplecidrsCreate,
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

This resource will retrieve the next available IP address from a collection of given prefixes or IP ranges (specified by ID)`,

		Schema: map[string]*schema.Schema{
			"prefix_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				ExactlyOneOf: []string{"prefix_ids", "ip_range_ids"},
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"ip_range_ids": {
				Type:         schema.TypeList,
				Optional:     true,
				ExactlyOneOf: []string{"prefix_ids", "ip_range_ids"},
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// IF prefix_ids is given then prefix_id will be populated with the selected prefix_id
			"prefix_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			// IF ip_range_ids is given then ip_range_id will be populated with the selected ip_range_id
			"ip_range_id": {
				Type:     schema.TypeInt,
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

// TODO: unsure about compareable
func extractTFCollectionInt64(d *schema.ResourceData, id string) ([]int64, bool) {
	// TODO: Unsure if this works
	var numbers []int64
	var isSet bool
	var numberInterface interface{}
	if numberInterface, isSet = d.GetOk("prefix_ids"); isSet {
		numbers := numberInterface.([]interface{})
		for _, number := range numbers {
			numbers = append(numbers, int64(number.(int)))
		}
	}
	isSet = isSet && len(numbers) > 0
	return numbers, isSet
}

func resourceNetboxAvailableIPAddressMultiplecidrsCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	rangeIDs, rangeIDsIsSet := extractTFCollectionInt64(d, "ip_range_ids")
	prefixIDs, prefixIDsIsSet := extractTFCollectionInt64(d, "prefix_ids")

	vrfID := int64(int64(d.Get("vrf_id").(int)))
	nestedvrf := models.NestedVRF{
		ID: vrfID,
	}
	data := models.AvailableIP{
		Vrf: &nestedvrf,
	}

	if prefixIDsIsSet {
		var params *ipam.IpamPrefixesAvailableIpsCreateParams
		var prefixId int64
		for _, id := range prefixIDs {
			params = ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(id).WithData([]*models.AvailableIP{&data})
			if len(params.Data) != 0 {
				prefixId = id
				break
			}
		}
		if len(params.Data) == 0 {
			return fmt.Errorf("No avalible ip address found with the given prefix_ids")
		}

		res, err := api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
		if err != nil {
			return fmt.Errorf("Error creating ip address from prefixes: %s", err)
		}
		// Since we generated the ip_address, set that now
		d.SetId(strconv.FormatInt(res.Payload[0].ID, 10))
		d.Set("ip_address", *res.Payload[0].Address)
		d.Set("prefix_id", prefixId)
	}
	if rangeIDsIsSet {
		var params *ipam.IpamIPRangesAvailableIpsCreateParams
		var rangeId int64
		for _, id := range rangeIDs {
			params := ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(id).WithData([]*models.AvailableIP{&data})
			if len(params.Data) != 0 {
				rangeId = id
				break
			}
		}
		res, err := api.Ipam.IpamIPRangesAvailableIpsCreate(params, nil)
		if err != nil {
			return fmt.Errorf("Error creating ip address from ranges: %s", err)
		}
		// Since we generated the ip_address, set that now
		d.SetId(strconv.FormatInt(res.Payload[0].ID, 10))
		d.Set("ip_address", *res.Payload[0].Address)
		d.Set("ip_range_id", rangeId)
	}
	return resourceNetboxAvailableIPAddressUpdate(d, m)
}
