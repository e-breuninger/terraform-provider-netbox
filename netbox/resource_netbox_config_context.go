package netbox

import (
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"data": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
			},
			"cluster_groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"cluster_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"clusters": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"device_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"locations": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"platforms": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"regions": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"site_groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"sites": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tenant_groups": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tenants": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Default:  nil,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxConfigContextCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)
	data := models.WritableConfigContext{}
	data.Name = strToPtr(d.Get("name").(string))

	dataJSON, ok := d.GetOk("data")
	if ok {
		var jsonObj any
		localContextBA := []byte(dataJSON.(string))
		if err := json.Unmarshal(localContextBA, &jsonObj); err == nil {
			data.Data = jsonObj
		}
	}
	data.Description = d.Get("description").(string)
	data.ClusterGroups = toInt64List(d.Get("cluster_groups"))
	data.ClusterTypes = toInt64List(d.Get("cluster_types"))
	data.Clusters = toInt64List(d.Get("clusters"))
	data.DeviceTypes = toInt64List(d.Get("device_types"))
	data.Locations = toInt64List(d.Get("locations"))
	data.Platforms = toInt64List(d.Get("platforms"))
	data.Regions = toInt64List(d.Get("regions"))
	data.Roles = toInt64List(d.Get("roles"))
	data.SiteGroups = toInt64List(d.Get("site_groups"))
	data.Sites = toInt64List(d.Get("sites"))
	data.TenantGroups = toInt64List(d.Get("tenant_groups"))
	data.Tenants = toInt64List(d.Get("tenants"))
	data.Tags = toStringList(d.Get("tags"))
	data.Weight = int64ToPtr(int64(d.Get("weight").(int)))

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
	d.Set("description", res.GetPayload().Description)
	d.Set("weight", res.GetPayload().Weight)

	if res.GetPayload().Data != nil {
		if jsonArr, err := json.Marshal(res.GetPayload().Data); err == nil {
			d.Set("data", string(jsonArr))
		}
	} else {
		d.Set("data", nil)
	}

	clusterGroups := res.GetPayload().ClusterGroups
	clusterGroupsSlice := make([]int64, len(clusterGroups))
	for i, v := range clusterGroups {
		clusterGroupsSlice[i] = int64(v.ID)
	}
	d.Set("cluster_groups", clusterGroupsSlice)

	clusterTypes := res.GetPayload().ClusterTypes
	clusterTypesSlice := make([]int64, len(clusterTypes))
	for i, v := range clusterTypes {
		clusterTypesSlice[i] = int64(v.ID)
	}
	d.Set("cluster_types", clusterTypesSlice)

	clusters := res.GetPayload().Clusters
	clustersSlice := make([]int64, len(clusters))
	for i, v := range clusters {
		clustersSlice[i] = int64(v.ID)
	}
	d.Set("clusters", clustersSlice)

	deviceTypes := res.GetPayload().DeviceTypes
	deviceTypesSlice := make([]int64, len(deviceTypes))
	for i, v := range deviceTypes {
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

	siteGroups := res.GetPayload().SiteGroups
	siteGroupsSlice := make([]int64, len(siteGroups))
	for i, v := range siteGroups {
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

	tenantGroups := res.GetPayload().TenantGroups
	tenantGroupsSlice := make([]int64, len(tenantGroups))
	for i, v := range tenantGroups {
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

	data := models.WritableConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

	dataValue, ok := d.GetOk("data")
	if ok {
		var jsonObj any
		localContextBA := []byte(dataValue.(string))
		if err := json.Unmarshal(localContextBA, &jsonObj); err == nil {
			data.Data = jsonObj
		}
	}
	data.Description = d.Get("description").(string)
	data.ClusterGroups = toInt64List(d.Get("cluster_groups"))
	data.ClusterTypes = toInt64List(d.Get("cluster_types"))
	data.Clusters = toInt64List(d.Get("clusters"))
	data.DeviceTypes = toInt64List(d.Get("device_types"))
	data.Locations = toInt64List(d.Get("locations"))
	data.Platforms = toInt64List(d.Get("platforms"))
	data.Regions = toInt64List(d.Get("regions"))
	data.Roles = toInt64List(d.Get("roles"))
	data.SiteGroups = toInt64List(d.Get("site_groups"))
	data.Sites = toInt64List(d.Get("sites"))
	data.TenantGroups = toInt64List(d.Get("tenant_groups"))
	data.Tenants = toInt64List(d.Get("tenants"))
	data.Tags = toStringList(d.Get("tags"))
	data.Weight = int64ToPtr(int64(d.Get("weight").(int)))

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
