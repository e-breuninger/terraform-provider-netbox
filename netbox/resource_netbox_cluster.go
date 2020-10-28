package netbox

import (
	"errors"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
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
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNetboxClusterCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	clusterTypeID := int64(d.Get("cluster_type_id").(int))
	tags, _ := getNestedTagListFromResourceDataSet(api, d.Get("tags"))

	params := virtualization.NewVirtualizationClustersCreateParams().WithData(
		&models.WritableCluster{
			Name: &name,
			Type: &clusterTypeID,
			Tags: tags,
		},
	)

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
		var apiError *runtime.APIError
		if errors.As(err, &apiError) {
			errorcode := err.(*runtime.APIError).Response.(runtime.ClientResponse).Code()
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
	d.Set("tags", getTagListFromNestedTagList(res.GetPayload().Tags))
	return nil
}

func resourceNetboxClusterUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableCluster{}

	name := d.Get("name").(string)
	clusterTypeID := int64(d.Get("cluster_type_id").(int))

	data.Name = &name
	data.Type = &clusterTypeID

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
