package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
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

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_type_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxClusterCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	data := models.WritableCluster{}

	name := d.Get("name").(string)
	data.Name = &name

	clusterTypeID := int64(d.Get("cluster_type_id").(int))
	data.Type = &clusterTypeID

	if clusterGroupIDValue, ok := d.GetOk("cluster_group_id"); ok {
		clusterGroupID := int64(clusterGroupIDValue.(int))
		data.Group = &clusterGroupID
	}

	if siteIDValue, ok := d.GetOk("site_id"); ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))
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
	api := m.(*client.NetBoxAPI)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationClustersReadParams().WithID(id)

	res, err := api.Virtualization.VirtualizationClustersRead(params, nil)
	if err != nil {
		errorcode := err.(*virtualization.VirtualizationClustersReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", res.GetPayload().Name)
	d.Set("cluster_type_id", res.GetPayload().Type.ID)

	if res.GetPayload().Group != nil {
		d.Set("cluster_group_id", res.GetPayload().Group.ID)
	} else {
		d.Set("cluster_group_id", nil)
	}

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxClusterUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

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

	if siteIDValue, ok := d.GetOk("site_id"); ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))
	data.Tags = tags

	params := virtualization.NewVirtualizationClustersPartialUpdateParams().WithID(id).WithData(&data)

	_, err := api.Virtualization.VirtualizationClustersPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxClusterRead(d, m)
}

func resourceNetboxClusterDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := virtualization.NewVirtualizationClustersDeleteParams().WithID(id)

	_, err := api.Virtualization.VirtualizationClustersDelete(params, nil)
	if err != nil {
		return err
	}
	return nil
}
