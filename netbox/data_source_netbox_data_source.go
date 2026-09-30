package netbox

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDataSource() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxDataSourceRead,
		Description: `:meta:subcategory:Extras:`,
		Schema: map[string]*schema.Schema{
			"name": {
				AtLeastOneOf: []string{"id", "name"},
				Optional:     true,
				Type:         schema.TypeString,
			},
			"id": {
				AtLeastOneOf: []string{"id", "name"},
				Computed:     true,
				Optional:     true,
				Type:         schema.TypeString,
			},
			"type": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"source_url": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"enabled": {
				Computed: true,
				Type:     schema.TypeBool,
			},
			"description": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"comments": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"ignore_rules": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"sync_interval": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"parameters": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"status": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"last_synced": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"file_count": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"custom_fields": {
				Computed: true,
				Type:     schema.TypeMap,
			},
		},
	}
}

func dataSourceNetboxDataSourceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	params := core.NewCoreDataSourcesListParams()

	if name, ok := d.Get("name").(string); ok && name != "" {
		params.Name = &name
	}

	if id, ok := d.Get("id").(string); ok && id != "0" && id != "" {
		params.SetID(&id)
	}

	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Core.CoreDataSourcesList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("more than one result, specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("no data source found matching filter")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("source_url", result.SourceURL)

	if result.Type != nil {
		d.Set("type", result.Type.Value)
	}

	d.Set("enabled", result.Enabled)
	d.Set("description", result.Description)
	d.Set("comments", result.Comments)
	d.Set("ignore_rules", result.IgnoreRules)
	d.Set("sync_interval", result.SyncInterval)
	d.Set("file_count", result.FileCount)

	if result.Status != nil {
		d.Set("status", result.Status.Value)
	}

	if result.LastSynced != nil {
		d.Set("last_synced", result.LastSynced.String())
	} else {
		d.Set("last_synced", nil)
	}

	if result.Parameters != nil {
		if jsonArr, err := json.Marshal(result.Parameters); err == nil {
			d.Set("parameters", string(jsonArr))
		}
	} else {
		d.Set("parameters", nil)
	}

	if result.CustomFields != nil {
		d.Set("custom_fields", flattenCustomFields(result.CustomFields))
	}

	return nil
}
