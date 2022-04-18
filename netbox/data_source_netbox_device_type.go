package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxDeviceType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetboxDeviceTypeRead,
		Schema: map[string]*schema.Schema{
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	model := d.Get("model").(string)
	params := dcim.NewDcimDeviceTypesListParams()
	params.Model = &model
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := api.Dcim.DcimDeviceTypesList(params, nil)
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
	d.Set("model", result.Model)
	d.Set("slug", result.Slug)
	return nil
}
