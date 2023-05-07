package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxCluster() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxClusterRead,
		Description: `:meta:subcategory:Virtualization:`,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site_id": {
				Type:         schema.TypeInt,
				Computed:     true,
				Optional:     true,
				AtLeastOneOf: []string{"name", "site_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "site_id"},
			},
			"cluster_type_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cluster_group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			tagsKey: tagsSchemaRead,
		},
	}
}

func dataSourceNetboxClusterRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	params := virtualization.NewVirtualizationClustersListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if site_id, ok := d.Get("site_id").(int); ok && site_id != 0 {
		params.SiteID = strToPtr(strconv.FormatInt(int64(site_id), 10))
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Virtualization.VirtualizationClustersList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no result")
	}
	result := res.GetPayload().Results[0]
	d.Set("cluster_id", result.ID)
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("cluster_type_id", result.Type.ID)

	if result.Group != nil {
		d.Set("cluster_group_id", result.Group.ID)
	} else {
		d.Set("cluster_group_id", nil)
	}

	if result.Site != nil {
		d.Set("site_id", result.Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(result.Tags))
	return nil
}
