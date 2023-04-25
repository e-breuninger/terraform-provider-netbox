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

		Description: `:meta:subcategory:Virtualization:From the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#clusters):

> A cluster is a logical grouping of physical resources within which virtual machines run. A cluster must be assigned a type (technological classification), and may optionally be assigned to a cluster group, site, and/or tenant. Each cluster must have a unique name within its assigned group and/or site, if any.
>
> Physical devices may be associated with clusters as hosts. This allows users to track on which host(s) a particular virtual machine may reside. However, NetBox does not support pinning a specific VM within a cluster to a particular host device.`,

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
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
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

	if tenantIDValue, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
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

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
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

	if tenantIDValue, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))
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
