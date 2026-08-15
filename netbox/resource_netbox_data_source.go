package netbox

import (
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/core"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxDataSourceTypeOptions = []string{"local", "git", "amazon-s3"}

var resourceNetboxDataSourceSyncIntervalOptions = []int{1, 60, 720, 1440, 10080, 43200}

func resourceNetboxDataSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxDataSourceCreate,
		Read:   resourceNetboxDataSourceRead,
		Update: resourceNetboxDataSourceUpdate,
		Delete: resourceNetboxDataSourceDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/integrations/synchronized-data/):

> NetBox can be configured to retrieve data from a remote source, such as a git repository or an S3 bucket, and use it to populate certain models such as config contexts and config templates. A data source represents such a remote source of data.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxDataSourceTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxDataSourceTypeOptions),
			},
			"source_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_rules": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntInSlice(resourceNetboxDataSourceSyncIntervalOptions),
			},
			"parameters": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					equal, _ := jsonSemanticCompare(oldValue, newValue)
					return equal
				},
				DiffSuppressOnRefresh: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_synced": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDataSourceCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	data := models.WritableDataSource{}

	name := d.Get("name").(string)
	data.Name = &name

	sourceURL := d.Get("source_url").(string)
	data.SourceURL = &sourceURL

	dsType := d.Get("type").(string)
	data.Type = &dsType

	data.Enabled = d.Get("enabled").(bool)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)
	data.IgnoreRules = d.Get("ignore_rules").(string)
	data.SyncInterval = getOptionalInt(d, "sync_interval")

	parametersValue, ok := d.GetOk("parameters")
	if ok {
		var jsonObj any
		parametersBA := []byte(parametersValue.(string))
		if err := json.Unmarshal(parametersBA, &jsonObj); err == nil {
			data.Parameters = jsonObj
		}
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := core.NewCoreDataSourcesCreateParams().WithData(&data)

	res, err := api.Core.CoreDataSourcesCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDataSourceRead(d, m)
}

func resourceNetboxDataSourceRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := core.NewCoreDataSourcesReadParams().WithID(id)

	res, err := api.Core.CoreDataSourcesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*core.CoreDataSourcesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	dataSource := res.GetPayload()
	d.Set("name", dataSource.Name)
	d.Set("source_url", dataSource.SourceURL)

	if dataSource.Type != nil {
		d.Set("type", dataSource.Type.Value)
	}

	d.Set("enabled", dataSource.Enabled)
	d.Set("description", dataSource.Description)
	d.Set("comments", dataSource.Comments)
	d.Set("ignore_rules", dataSource.IgnoreRules)
	d.Set("sync_interval", dataSource.SyncInterval)
	d.Set("file_count", dataSource.FileCount)

	if dataSource.Status != nil {
		d.Set("status", dataSource.Status.Value)
	}

	if dataSource.LastSynced != nil {
		d.Set("last_synced", dataSource.LastSynced.String())
	} else {
		d.Set("last_synced", nil)
	}

	if dataSource.Parameters != nil {
		if jsonArr, err := json.Marshal(dataSource.Parameters); err == nil {
			d.Set("parameters", string(jsonArr))
		}
	} else {
		d.Set("parameters", nil)
	}

	cf := getCustomFields(dataSource.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	return nil
}

func resourceNetboxDataSourceUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDataSource{}

	name := d.Get("name").(string)
	data.Name = &name

	sourceURL := d.Get("source_url").(string)
	data.SourceURL = &sourceURL

	dsType := d.Get("type").(string)
	data.Type = &dsType

	data.Enabled = d.Get("enabled").(bool)
	data.Description = d.Get("description").(string)
	data.Comments = d.Get("comments").(string)
	data.IgnoreRules = d.Get("ignore_rules").(string)
	data.SyncInterval = getOptionalInt(d, "sync_interval")

	parametersValue, ok := d.GetOk("parameters")
	if ok {
		var jsonObj any
		parametersBA := []byte(parametersValue.(string))
		if err := json.Unmarshal(parametersBA, &jsonObj); err == nil {
			data.Parameters = jsonObj
		}
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	params := core.NewCoreDataSourcesPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Core.CoreDataSourcesPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxDataSourceRead(d, m)
}

func resourceNetboxDataSourceDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := core.NewCoreDataSourcesDeleteParams().WithID(id)

	_, err := api.Core.CoreDataSourcesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*core.CoreDataSourcesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
