package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxTenantGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxTenantGroupRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNetboxTenantGroupRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)
	params := tenancy.NewTenancyTenantGroupsListParams()
	params.Name = &name
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Tenancy.TenancyTenantGroupsList(params, nil)
	if err != nil {
		return err
	}

	if *res.GetPayload().Count > int64(1) {
		return errors.New("More than one result. Specify a more narrow filter")
	}
	if *res.GetPayload().Count == int64(0) {
		return errors.New("No result")
	}
	result := res.GetPayload().Results[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("name", result.Name)
	d.Set("slug", result.Slug)
	d.Set("description", result.Description)
	d.Set("parent_id", result.Parent.ID)
	return nil
}
