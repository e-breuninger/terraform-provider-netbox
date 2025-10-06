package netbox

import (
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strconv"
)

func dataSourceNetboxIPAddress() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxIPAddressRead,
		Description: `:meta:subcategory:IP Address Management (IPAM):`,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_fields": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"address_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tenant": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"assigned_object": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id":          {Type: schema.TypeInt, Computed: true},
						"url":         {Type: schema.TypeString, Computed: true},
						"display":     {Type: schema.TypeString, Computed: true},
						"name":        {Type: schema.TypeString, Computed: true},
						"description": {Type: schema.TypeString, Computed: true},
						"device": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id":          {Type: schema.TypeInt, Computed: true},
									"type":        {Type: schema.TypeString, Computed: true},
									"url":         {Type: schema.TypeString, Computed: true},
									"display":     {Type: schema.TypeString, Computed: true},
									"name":        {Type: schema.TypeString, Computed: true},
									"description": {Type: schema.TypeString, Computed: true},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxIPAddressRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id := d.Get("id").(int)

	params := ipam.NewIpamIPAddressesReadParams()
	params.SetID(int64(id))

	res, err := api.Ipam.IpamIPAddressesRead(params, nil)
	if err != nil {
		return err
	}

	result := res.GetPayload()

	d.SetId(strconv.FormatInt(result.ID, 10))

	d.Set("id", result.ID)
	d.Set("ip_address", result.Address)
	d.Set("description", result.Description)
	d.Set("created", result.Created.String())
	d.Set("last_updated", result.LastUpdated.String())
	d.Set("custom_fields", result.CustomFields)
	d.Set("address_family", result.Family.Label)
	d.Set("status", result.Status.Value)
	d.Set("dns_name", result.DNSName)

	if result.Role != nil {
		d.Set("role", result.Role.Value)
	}

	var tenant []map[string]interface{}
	if result.Tenant != nil {
		var mapping = make(map[string]interface{})
		mapping["id"] = result.Tenant.ID
		mapping["name"] = result.Tenant.Name
		mapping["slug"] = result.Tenant.Slug
		tenant = append(tenant, mapping)
	}
	d.Set("tenant", tenant)

	var tags []map[string]interface{}
	for _, t := range result.Tags {
		var tagmapping = make(map[string]interface{})
		tagmapping["name"] = t.Name
		tagmapping["display"] = t.Display
		tagmapping["slug"] = t.Slug
		tagmapping["id"] = t.ID
		tags = append(tags, tagmapping)
	}
	d.Set("tags", tags)

	if result.AssignedObject != nil {
		if aoMap, ok := result.AssignedObject.(map[string]interface{}); ok {
			var deviceList []interface{}
			for _, objType := range []string{"device", "virtual_machine"} {
				if obj, ok := aoMap[objType].(map[string]interface{}); ok {
					deviceList = append(deviceList, map[string]interface{}{
						"id":          obj["id"],
						"type":        objType,
						"url":         obj["url"],
						"display":     obj["display"],
						"name":        obj["name"],
						"description": obj["description"],
					})
				}
			}
			assignedObject := []interface{}{
				map[string]interface{}{
					"id":          aoMap["id"],
					"url":         aoMap["url"],
					"display":     aoMap["display"],
					"name":        aoMap["name"],
					"description": aoMap["description"],
					"device":      deviceList,
				},
			}
			d.Set("assigned_object", assignedObject)
		}
	}

	return nil
}
