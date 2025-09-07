package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxClusterCreate,
		Read:   resourceNetboxClusterRead,
		Update: resourceNetboxClusterUpdate,
		Delete: resourceNetboxClusterDelete,

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://netboxlabs.com/docs/netbox/models/virtualization/cluster/):
> A cluster is a logical grouping of physical resources within which virtual machines run. Physical devices may be associated with clusters as hosts. This allows users to track on which host(s) a particular virtual machine may reside.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"site_id", "site_group_id", "region_id"},
			},
			"site_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"location_id", "site_group_id", "region_id"},
			},
			"site_group_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"location_id", "site_id", "region_id"},
			},
			"region_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"location_id", "site_id", "site_group_id"},
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxClusterCreate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	data := models.WritableCluster{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterTypeID := int64(d.Get("cluster_type_id").(int))
	data.Type = &clusterTypeID

	siteID := getOptionalInt(d, "site_id")
	siteGroupID := getOptionalInt(d, "site_group_id")
	locationID := getOptionalInt(d, "location_id")
	regionID := getOptionalInt(d, "region_id")

	switch {
	case siteID != nil:
		data.ScopeType = strToPtr("dcim.site")
		data.ScopeID = siteID
	case siteGroupID != nil:
		data.ScopeType = strToPtr("dcim.sitegroup")
		data.ScopeID = siteGroupID
	case locationID != nil:
		data.ScopeType = strToPtr("dcim.location")
		data.ScopeID = locationID
	case regionID != nil:
		data.ScopeType = strToPtr("dcim.region")
		data.ScopeID = regionID
	default:
		data.ScopeType = nil
		data.ScopeID = nil
	}

	if clusterGroupIDValue, ok := d.GetOk("cluster_group_id"); ok {
		clusterGroupID := int64(clusterGroupIDValue.(int))
		data.Group = &clusterGroupID
	}

	data.Comments = getOptionalStr(d, "comments", false)
	data.Description = getOptionalStr(d, "description", false)

	if tenantIDValue, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	tags, _ := getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	data.Tags = tags

	params := virtualization.NewVirtualizationClustersCreateParams().WithData(&data)

	res, err := api.Virtualization.VirtualizationClustersCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxClusterRead(d, m)
}

func resourceNetboxClusterRead(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationClustersReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationClustersRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationClustersReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	cluster := res.GetPayload()

	d.Set("name", cluster.Name)
	d.Set("cluster_type_id", cluster.Type.ID)

	if cluster.Group != nil {
		d.Set("cluster_group_id", cluster.Group.ID)
	} else {
		d.Set("cluster_group_id", nil)
	}

	d.Set("comments", cluster.Comments)
	d.Set("description", cluster.Description)

	if cluster.Tenant != nil {
		d.Set("tenant_id", cluster.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	d.Set("site_id", nil)
	d.Set("site_group_id", nil)
	d.Set("location_id", nil)
	d.Set("region_id", nil)

	if cluster.ScopeType != nil && cluster.ScopeID != nil {
		scopeID := cluster.ScopeID
		switch scopeType := cluster.ScopeType; *scopeType {
		case "dcim.site":
			d.Set("site_id", scopeID)
		case "dcim.sitegroup":
			d.Set("site_group_id", scopeID)
		case "dcim.location":
			d.Set("location_id", scopeID)
		case "dcim.region":
			d.Set("region_id", scopeID)
		}
	}

	state.readTags(d, res.GetPayload().Tags)
	return nil
}

func resourceNetboxClusterUpdate(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCluster{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterTypeID := int64(d.Get("cluster_type_id").(int))
	data.Type = &clusterTypeID

	if clusterGroupIDValue, ok := d.GetOk("cluster_group_id"); ok {
		clusterGroupID := int64(clusterGroupIDValue.(int))
		data.Group = &clusterGroupID
	}

	data.Comments = getOptionalStr(d, "comments", true)
	data.Description = getOptionalStr(d, "description", true)

	if tenantIDValue, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	siteID := getOptionalInt(d, "site_id")
	siteGroupID := getOptionalInt(d, "site_group_id")
	locationID := getOptionalInt(d, "location_id")
	regionID := getOptionalInt(d, "region_id")

	switch {
	case siteID != nil:
		data.ScopeType = strToPtr("dcim.site")
		data.ScopeID = siteID
	case siteGroupID != nil:
		data.ScopeType = strToPtr("dcim.sitegroup")
		data.ScopeID = siteGroupID
	case locationID != nil:
		data.ScopeType = strToPtr("dcim.location")
		data.ScopeID = locationID
	case regionID != nil:
		data.ScopeType = strToPtr("dcim.region")
		data.ScopeID = regionID
	default:
		data.ScopeType = nil
		data.ScopeID = nil
	}

	tags, _ := getNestedTagListFromResourceDataSet(state, d.Get(tagsAllKey))
	data.Tags = tags

	params := virtualization.NewVirtualizationClustersPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationClustersPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxClusterRead(d, m)
}

func resourceNetboxClusterDelete(d *schema.ResourceData, m interface{}) error {
	state := m.(*providerState)
	api := state.legacyAPI

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationClustersDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationClustersDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*virtualization.VirtualizationClustersDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}
