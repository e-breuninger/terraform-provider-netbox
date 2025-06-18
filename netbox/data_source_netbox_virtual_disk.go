package netbox

import (
	"fmt"
	"regexp"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceNetboxVirtualDisk() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxVirtualDiskRead,
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
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_disks": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_mb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"virtual_machine_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						tagsKey: {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						customFieldsKey: {
							Type:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetboxVirtualDiskRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	params := virtualization.NewVirtualizationVirtualDisksListParams()

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		var tags []string
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			vString := v.(string)
			switch k {
			case "name":
				params.Name = &vString
			case "tag":
				tags = append(tags, vString)
				params.Tag = tags
			// The fbreckle fork may not support virtual_machine_id filter directly
			default:
				return fmt.Errorf("'%s' is not a supported filter parameter", k)
			}
		}
	}
	if limit, ok := d.GetOk("limit"); ok {
		limitInt := int64(limit.(int))
		params.Limit = &limitInt
	}

	res, err := api.Virtualization.VirtualizationVirtualDisksList(params, nil)
	if err != nil {
		return err
	}

	var filteredDisks []*models.VirtualDisk
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, disk := range res.GetPayload().Results {
			if disk.Name != nil && r.MatchString(*disk.Name) {
				filteredDisks = append(filteredDisks, disk)
			}
		}
	} else {
		filteredDisks = res.GetPayload().Results
	}

	var s []map[string]interface{}
	for _, v := range filteredDisks {
		var mapping = make(map[string]interface{})
		if v.ID != 0 {
			mapping["id"] = v.ID
		}
		if v.Name != nil {
			mapping["name"] = *v.Name
		}
		if v.Description != "" {
			mapping["description"] = v.Description
		}
		if v.Size != nil {
			mapping["size_mb"] = *v.Size
		}
		if v.VirtualMachine != nil {
			mapping["virtual_machine_id"] = v.VirtualMachine.ID
		}
		if v.CustomFields != nil {
			mapping["custom_fields"] = v.CustomFields
		}
		if v.Tags != nil {
			tags := []string{}
			for _, t := range v.Tags {
				if t.Name != nil {
					tags = append(tags, *t.Name)
				}
			}
			mapping["tags"] = tags
		}
		s = append(s, mapping)
	}

	d.SetId(id.UniqueId())
	return d.Set("virtual_disks", s)
}
