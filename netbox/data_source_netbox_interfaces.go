package netbox

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxInterfaces() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxInterfaceRead,
		Description: `:meta:subcategory:Virtualization:`,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"limit": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
				Default:          0,
				Description:      "The limit of objects to return from the API lookup.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type:     schema.TypeString,
								Computed: true,
							},
						},
						"mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"tagged_vlans": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						// Do as a TypeList due to limitation of TypeMap
						"untagged_vlan": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"vm_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxInterfaceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := virtualization.NewVirtualizationInterfacesListParams()

	// Get user limit (0 = fetch all)
	var userLimit int64 = 0
	if limitValue, ok := d.GetOk("limit"); ok {
		userLimit = int64(limitValue.(int))
	}

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "cluster_id":
				params.ClusterID = &vString
			case "mac_address":
				params.MacAddress = &vString
			case "name":
				params.Name = &vString
			case "tag":
				params.Tag = []string{vString} //TODO: switch schema to list?
			case "vm_id":
				params.VirtualMachineID = &vString
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}

	// Fetch all pages with pagination (fetch all when name_regex is used)
	paginationHelper := NewPaginationHelper(FetchAll)
	var allInterfaces []*models.VMInterface

	pageSize := paginationHelper.GetPageSize()
	for {
		currentOffset := paginationHelper.CurrentOffset()
		params.Limit = &pageSize
		params.Offset = &currentOffset

		res, err := api.Virtualization.VirtualizationInterfacesList(params, nil)
		if err != nil {
			return fmt.Errorf("failed to fetch interfaces at offset %d: %w", currentOffset, err)
		}

		payload := res.GetPayload()
		allInterfaces = append(allInterfaces, payload.Results...)

		if len(payload.Results) == 0 {
			break
		}

		if !paginationHelper.ShouldContinuePaging(int64(len(allInterfaces)), payload.Next) {
			break
		}

		paginationHelper.Advance(int64(len(payload.Results)))
	}

	// Apply name_regex filter
	var filteredInterfaces []*models.VMInterface
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, vmInterface := range allInterfaces {
			if r.MatchString(*vmInterface.Name) {
				filteredInterfaces = append(filteredInterfaces, vmInterface)
			}
		}
	} else {
		filteredInterfaces = allInterfaces
	}

	// Apply user limit to filtered results
	if userLimit > 0 && int64(len(filteredInterfaces)) > userLimit {
		filteredInterfaces = filteredInterfaces[:userLimit]
	}

	if len(filteredInterfaces) == 0 {
		return errors.New("no result")
	}

	var s []map[string]interface{}
	for _, v := range filteredInterfaces {
		var mapping = make(map[string]interface{})
		mapping["id"] = v.ID
		if v.Description != "" {
			mapping["description"] = v.Description
		}
		mapping["enabled"] = v.Enabled
		if v.MacAddress != nil {
			mapping["mac_address"] = *v.MacAddress
		}
		if v.Mode != nil {
			mapping["mode"] = map[string]string{
				"label": *v.Mode.Label,
				"value": *v.Mode.Value,
			}
		}
		if v.Mtu != nil {
			mapping["mtu"] = *v.Mtu
		}
		if v.Name != nil {
			mapping["name"] = *v.Name
		}
		if v.TaggedVlans != nil {
			mapping["tagged_vlans"] = flattenVlanAttributes(v.TaggedVlans)
		}
		if v.Tags != nil {
			var tags []int64
			for _, t := range v.Tags {
				tags = append(tags, t.ID)
			}
			mapping["tag_ids"] = tags
		}
		if v.UntaggedVlan != nil {
			vlanSlice := []*models.NestedVLAN{v.UntaggedVlan}
			mapping["untagged_vlan"] = flattenVlanAttributes(vlanSlice)
		}

		mapping["vm_id"] = v.VirtualMachine.ID

		if v.Type != nil {
			mapping["type"] = *v.Type.Value
		}

		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("interfaces", s)
}

func flattenVlanAttributes(vlans []*models.NestedVLAN) []map[string]interface{} {
	var mappedVlans []map[string]interface{}
	for _, vlan := range vlans {
		v := *vlan
		mappedVlan := map[string]interface{}{
			"id":   v.ID,
			"vid":  *v.Vid,
			"name": *v.Name,
		}
		mappedVlans = append(mappedVlans, mappedVlan)
	}
	return mappedVlans
}
