package netbox

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxAvailablePrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxAvailablePrefixCreate,
		Read:   resourceNetboxPrefixRead,
		Update: resourceNetboxPrefixUpdate,
		Delete: resourceNetboxPrefixDelete,

		Schema: map[string]*schema.Schema{
			"parent_prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"prefix_length": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(0, 128),
			},
			"prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"active", "container", "reserved", "deprecated"}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(c context.Context, rd *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				parent_prefix_id, prefix_id, prefix_length, err := resourceNetboxAvailablePrefixParseImport(rd.Id())
				if err != nil {
					return nil, err
				}

				rd.Set("parent_prefix_id", parent_prefix_id)
				rd.Set("prefix_length", prefix_length)
				rd.SetId(prefix_id)

				return []*schema.ResourceData{rd}, nil
			},
		},
	}
}

func resourceNetboxAvailablePrefixParseImport(import_str string) (int, string, int, error) {
	parts := strings.SplitN(import_str, " ", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return 0, "", 0, fmt.Errorf("unexpected format of (%s), expected 'parent_prefix_id prefix_id prefix_length'", import_str)
	}

	parent_id, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parent_id (%s) is not an integer", parts[0])
	}
	prefix_length, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("prefix_length (%s) is not an integer", parts[1])
	}

	return parent_id, parts[1], prefix_length, nil
}

func resourceNetboxAvailablePrefixCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*client.NetBoxAPI)

	parent_prefix_id := int64(d.Get("parent_prefix_id").(int))
	prefix_length := int64(d.Get("prefix_length").(int))
	data := models.PrefixLength{
		PrefixLength: &prefix_length,
	}
	params := ipam.NewIpamPrefixesAvailablePrefixesCreateParams().WithID(parent_prefix_id).WithData(&data)

	res, err := api.Ipam.IpamPrefixesAvailablePrefixesCreate(params, nil)
	if err != nil {
		return err
	}

	payload := res.GetPayload()
	d.SetId(strconv.FormatInt(payload.ID, 10))
	d.Set("prefix", payload.Prefix)

	return resourceNetboxPrefixUpdate(d, m)
}
