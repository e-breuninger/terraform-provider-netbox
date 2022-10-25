package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxConfigContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxConfigContextCreate,
		Read:   resourceNetboxConfigContextRead,
		Update: resourceNetboxConfigContextUpdate,
		Delete: resourceNetboxConfigContextDelete,

		Description: `:meta:subcategory:Extras:From the [official documentation](https://docs.netbox.dev/en/stable/models/extras/configcontext/):

> Context data is made available to devices and/or virtual machines based on their relationships to other objects in NetBox. For example, context data can be associated only with devices assigned to a particular site, or only to virtual machines in a certain cluster.`,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"data": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				//				ValidateFunc: validation.StringLenBetween(0, 30),
			},
			"cluster_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"cluster_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"clusters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"device_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"locations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"platforms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"regions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"site_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"sites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tenant_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tenants": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxConfigContextCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name
	dataJson := d.Get("data").(map[string]interface{})
	data.Data = &dataJson

	cluster_groups := d.Get("cluster_groups").([]interface{})
	data.ClusterGroups = make([]int64, len(cluster_groups))
	for i, v := range cluster_groups {
		data.ClusterGroups[i] = int64(v.(int))
	}
	cluster_types := d.Get("cluster_types").([]interface{})
	data.ClusterTypes = make([]int64, len(cluster_types))
	for i, v := range cluster_types {
		data.ClusterTypes[i] = int64(v.(int))
	}
	clusters := d.Get("clusters").([]interface{})
	data.Clusters = make([]int64, len(clusters))
	for i, v := range clusters {
		data.Clusters[i] = int64(v.(int))
	}
	device_types := d.Get("device_types").([]interface{})
	data.DeviceTypes = make([]int64, len(device_types))
	for i, v := range device_types {
		data.DeviceTypes[i] = int64(v.(int))
	}
	locations := d.Get("locations").([]interface{})
	data.Locations = make([]int64, len(locations))
	for i, v := range locations {
		data.Locations[i] = int64(v.(int))
	}
	platforms := d.Get("platforms").([]interface{})
	data.Platforms = make([]int64, len(platforms))
	for i, v := range platforms {
		data.Platforms[i] = int64(v.(int))
	}
	regions := d.Get("regions").([]interface{})
	data.Regions = make([]int64, len(regions))
	for i, v := range regions {
		data.Regions[i] = int64(v.(int))
	}
	roles := d.Get("roles").([]interface{})
	data.Roles = make([]int64, len(roles))
	for i, v := range roles {
		data.Roles[i] = int64(v.(int))
	}
	site_groups := d.Get("site_groups").([]interface{})
	data.SiteGroups = make([]int64, len(site_groups))
	for i, v := range site_groups {
		data.SiteGroups[i] = int64(v.(int))
	}
	sites := d.Get("sites").([]interface{})
	data.Sites = make([]int64, len(sites))
	for i, v := range sites {
		data.Sites[i] = int64(v.(int))
	}
	tenant_groups := d.Get("tenant_groups").([]interface{})
	data.TenantGroups = make([]int64, len(tenant_groups))
	for i, v := range tenant_groups {
		data.TenantGroups[i] = int64(v.(int))
	}
	tenants := d.Get("tenants").([]interface{})
	data.Tenants = make([]int64, len(tenants))
	for i, v := range tenants {
		data.Tenants[i] = int64(v.(int))
	}
	tags := d.Get("tags").([]interface{})
	data.Tags = make([]string, len(tags))
	for i, v := range tags {
		data.Tags[i] = v.(string)
	}

	weightValue, weightOk := d.GetOk("weight")
	var weight int64
	// Default weight to 1000 if not given
	if !weightOk {
		weight = 1000
	} else {
		weight = int64(weightValue.(int))
	}
	data.Weight = &weight

	params := extras.NewExtrasConfigContextsCreateParams().WithData(&data)

	res, err := api.Extras.ExtrasConfigContextsCreate(params, nil)
	if err != nil {
		//return errors.New(getTextFromError(err))
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxConfigContextRead(d, m)
}

func resourceNetboxConfigContextRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasConfigContextsReadParams().WithID(id)

	res, err := api.Extras.ExtrasConfigContextsRead(params, nil)

	if err != nil {
		errorcode := err.(*extras.ExtrasConfigContextsReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("weight", res.GetPayload().Weight)
	d.Set("data", res.GetPayload().Data)

	cluster_groups := res.GetPayload().ClusterGroups
	clusterGroupsSlice := make([]int64, len(cluster_groups))
	for i, v := range cluster_groups {
		clusterGroupsSlice[i] = int64(v.ID)
	}
	d.Set("cluster_groups", clusterGroupsSlice)

	cluster_types := res.GetPayload().ClusterTypes
	clusterTypesSlice := make([]int64, len(cluster_types))
	for i, v := range cluster_types {
		clusterTypesSlice[i] = int64(v.ID)
	}
	d.Set("cluster_types", clusterTypesSlice)

	clusters := res.GetPayload().Clusters
	clustersSlice := make([]int64, len(clusters))
	for i, v := range clusters {
		clustersSlice[i] = int64(v.ID)
	}
	d.Set("clusters", clustersSlice)

	device_types := res.GetPayload().DeviceTypes
	deviceTypesSlice := make([]int64, len(device_types))
	for i, v := range device_types {
		deviceTypesSlice[i] = int64(v.ID)
	}
	d.Set("device_types", deviceTypesSlice)

	locations := res.GetPayload().Locations
	locationsSlice := make([]int64, len(locations))
	for i, v := range locations {
		locationsSlice[i] = int64(v.ID)
	}
	d.Set("locations", locationsSlice)

	platforms := res.GetPayload().Platforms
	platformsSlice := make([]int64, len(platforms))
	for i, v := range platforms {
		platformsSlice[i] = int64(v.ID)
	}
	d.Set("platforms", platformsSlice)

	regions := res.GetPayload().Regions
	regionsSlice := make([]int64, len(regions))
	for i, v := range regions {
		regionsSlice[i] = int64(v.ID)
	}
	d.Set("regions", regionsSlice)

	roles := res.GetPayload().Roles
	rolesSlice := make([]int64, len(roles))
	for i, v := range roles {
		rolesSlice[i] = int64(v.ID)
	}
	d.Set("roles", rolesSlice)

	site_groups := res.GetPayload().SiteGroups
	siteGroupsSlice := make([]int64, len(site_groups))
	for i, v := range site_groups {
		siteGroupsSlice[i] = int64(v.ID)
	}
	d.Set("site_groups", siteGroupsSlice)

	sites := res.GetPayload().Sites
	sitesSlice := make([]int64, len(sites))
	for i, v := range sites {
		sitesSlice[i] = int64(v.ID)
	}
	d.Set("sites", sitesSlice)

	tags := res.GetPayload().Tags
	tagsSlice := make([]string, len(tags))
	for i, v := range tags {
		tagsSlice[i] = string(v)
	}
	d.Set("tags", tagsSlice)

	tenant_groups := res.GetPayload().TenantGroups
	tenantGroupsSlice := make([]int64, len(tenant_groups))
	for i, v := range tenant_groups {
		tenantGroupsSlice[i] = int64(v.ID)
	}
	d.Set("tenant_groups", tenantGroupsSlice)

	tenants := res.GetPayload().Tenants
	tenantsSlice := make([]int64, len(tenants))
	for i, v := range tenants {
		tenantsSlice[i] = int64(v.ID)
	}
	d.Set("tenants", tenantsSlice)

	return nil
}

func resourceNetboxConfigContextUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	/*name := d.Get("name").(string)
	datajson := d.Get("data").(string)
	cluster_groups := d.Get("cluster_groups").([]int64)
	cluster_types := d.Get("cluster_types").([]int64)
	clusters := d.Get("clusters").([]int64)
	device_types := d.Get("device_types").([]int64)
	locations := d.Get("locations").([]int64)
	platforms := d.Get("platforms").([]int64)
	regions := d.Get("regions").([]int64)
	roles := d.Get("roles").([]int64)
	site_groups := d.Get("site_groups").([]int64)
	sites := d.Get("sites").([]int64)
	tenant_groups := d.Get("tenant_groups").([]int64)
	tenants := d.Get("tenants").([]int64)
	tags := d.Get("tags").([]string)

	weightValue, weightOk := d.GetOk("weight")
	var weight int64
	// Default weight to 1000 if not given
	if !weightOk {
		weight = 1000
	} else {
		weight = weightValue.(int64)
	}

	data.Weight = &weight
	data.Name = &name
	data.Data = &datajson
	data.ClusterGroups = cluster_groups
	data.ClusterTypes = cluster_types
	data.Clusters = clusters
	data.DeviceTypes = device_types
	data.Locations = locations
	data.Platforms = platforms
	data.Regions = regions
	data.Roles = roles
	data.SiteGroups = site_groups
	data.Sites = sites
	data.TenantGroups = tenant_groups
	data.Tenants = tenants
	data.Tags = tags
	*/
	data := models.WritableConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name
	dataJson := d.Get("data").(map[string]interface{})
	data.Data = &dataJson

	cluster_groups := d.Get("cluster_groups").([]interface{})
	data.ClusterGroups = make([]int64, len(cluster_groups))
	for i, v := range cluster_groups {
		data.ClusterGroups[i] = int64(v.(int))
	}
	cluster_types := d.Get("cluster_types").([]interface{})
	data.ClusterTypes = make([]int64, len(cluster_types))
	for i, v := range cluster_types {
		data.ClusterTypes[i] = int64(v.(int))
	}
	clusters := d.Get("clusters").([]interface{})
	data.Clusters = make([]int64, len(clusters))
	for i, v := range clusters {
		data.Clusters[i] = int64(v.(int))
	}
	device_types := d.Get("device_types").([]interface{})
	data.DeviceTypes = make([]int64, len(device_types))
	for i, v := range device_types {
		data.DeviceTypes[i] = int64(v.(int))
	}
	locations := d.Get("locations").([]interface{})
	data.Locations = make([]int64, len(locations))
	for i, v := range locations {
		data.Locations[i] = int64(v.(int))
	}
	platforms := d.Get("platforms").([]interface{})
	data.Platforms = make([]int64, len(platforms))
	for i, v := range platforms {
		data.Platforms[i] = int64(v.(int))
	}
	regions := d.Get("regions").([]interface{})
	data.Regions = make([]int64, len(regions))
	for i, v := range regions {
		data.Regions[i] = int64(v.(int))
	}
	roles := d.Get("roles").([]interface{})
	data.Roles = make([]int64, len(roles))
	for i, v := range roles {
		data.Roles[i] = int64(v.(int))
	}
	site_groups := d.Get("site_groups").([]interface{})
	data.SiteGroups = make([]int64, len(site_groups))
	for i, v := range site_groups {
		data.SiteGroups[i] = int64(v.(int))
	}
	sites := d.Get("sites").([]interface{})
	data.Sites = make([]int64, len(sites))
	for i, v := range sites {
		data.Sites[i] = int64(v.(int))
	}
	tenant_groups := d.Get("tenant_groups").([]interface{})
	data.TenantGroups = make([]int64, len(tenant_groups))
	for i, v := range tenant_groups {
		data.TenantGroups[i] = int64(v.(int))
	}
	tenants := d.Get("tenants").([]interface{})
	data.Tenants = make([]int64, len(tenants))
	for i, v := range tenants {
		data.Tenants[i] = int64(v.(int))
	}
	tags := d.Get("tags").([]interface{})
	data.Tags = make([]string, len(tags))
	for i, v := range tags {
		data.Tags[i] = v.(string)
	}

	weightValue, weightOk := d.GetOk("weight")
	var weight int64
	// Default weight to 1000 if not given
	if !weightOk {
		weight = 1000
	} else {
		weight = int64(weightValue.(int))
	}
	data.Weight = &weight

	params := extras.NewExtrasConfigContextsPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Extras.ExtrasConfigContextsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxConfigContextRead(d, m)
}

func resourceNetboxConfigContextDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := extras.NewExtrasConfigContextsDeleteParams().WithID(id)

	_, err := api.Extras.ExtrasConfigContextsDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
