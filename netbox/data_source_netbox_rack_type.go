package netbox

import (
	"errors"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetboxRackType() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceNetboxRackTypeRead,
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):Look up a rack type by ` + "`model`" + ` or ` + "`slug`" + `.`,
		Schema: map[string]*schema.Schema{
			"model": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"model", "slug"},
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"model", "slug"},
			},
			"manufacturer_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"form_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"width": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"u_height": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"starting_unit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"outer_width": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outer_depth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outer_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"max_weight": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"weight_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mounting_depth_mm": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetboxRackTypeRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	model := d.Get("model").(string)
	slug := d.Get("slug").(string)
	if model == "" && slug == "" {
		return errors.New("either 'model' or 'slug' must be specified")
	}

	// The rack-types list endpoint has no server-side model/slug filter, so we
	// list the rack types and match exactly client-side.
	params := dcim.NewDcimRackTypesListParams()
	params.Limit = int64ToPtr(0)

	res, err := api.Dcim.DcimRackTypesList(params, nil)
	if err != nil {
		return err
	}

	var matches []*models.RackType
	for _, rackType := range res.GetPayload().Results {
		if model != "" && (rackType.Model == nil || *rackType.Model != model) {
			continue
		}
		if slug != "" && (rackType.Slug == nil || *rackType.Slug != slug) {
			continue
		}
		matches = append(matches, rackType)
	}

	if len(matches) == 0 {
		return errors.New("no rack type found matching filter")
	}
	if len(matches) > 1 {
		return errors.New("more than one rack type returned, specify a more narrow filter")
	}

	result := matches[0]
	d.SetId(strconv.FormatInt(result.ID, 10))
	d.Set("model", result.Model)
	d.Set("slug", result.Slug)

	if result.Manufacturer != nil {
		d.Set("manufacturer_id", result.Manufacturer.ID)
	}

	if result.FormFactor != nil {
		d.Set("form_factor", result.FormFactor.Value)
	}

	if result.Width != nil {
		d.Set("width", result.Width.Value)
	}

	d.Set("u_height", result.UHeight)
	d.Set("starting_unit", result.StartingUnit)
	d.Set("description", result.Description)
	d.Set("outer_width", result.OuterWidth)
	d.Set("outer_depth", result.OuterDepth)

	if result.OuterUnit != nil {
		d.Set("outer_unit", result.OuterUnit.Value)
	}

	d.Set("weight", result.Weight)
	d.Set("max_weight", result.MaxWeight)

	if result.WeightUnit != nil {
		d.Set("weight_unit", result.WeightUnit.Value)
	}

	d.Set("mounting_depth_mm", result.MountingDepth)
	d.Set("comments", result.Comments)

	return nil
}
