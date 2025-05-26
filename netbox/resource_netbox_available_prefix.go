package netbox

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

		Description: `:meta:subcategory:IP Address Management (IPAM):`,

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
				ValidateFunc: validation.StringInSlice(resourceNetboxPrefixStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxPrefixStatusOptions),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mark_utilized": {
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
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			tagsKey: tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(c context.Context, rd *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				parentPrefixID, prefixID, prefixLength, err := resourceNetboxAvailablePrefixParseImport(rd.Id())
				if err != nil {
					return nil, err
				}

				rd.Set("parent_prefix_id", parentPrefixID)
				rd.Set("prefix_length", prefixLength)
				rd.SetId(prefixID)

				return []*schema.ResourceData{rd}, nil
			},
		},
	}
}

func resourceNetboxAvailablePrefixParseImport(importStr string) (int, string, int, error) {
	parts := strings.SplitN(importStr, " ", 3)

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return 0, "", 0, fmt.Errorf("unexpected format of (%s), expected 'parent_prefix_id prefix_id prefix_length'", importStr)
	}

	parentID, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parent_id (%s) is not an integer", parts[0])
	}
	prefixLength, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("prefix_length (%s) is not an integer", parts[1])
	}

	return parentID, parts[1], prefixLength, nil
}

func resourceNetboxAvailablePrefixCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	parentPrefixID := int64(d.Get("parent_prefix_id").(int))
	prefixLength := int64(d.Get("prefix_length").(int))
	data := models.PrefixLength{
		PrefixLength: &prefixLength,
	}
	params := ipam.NewIpamPrefixesAvailablePrefixesCreateParams().WithID(parentPrefixID).WithData(&data)

	res, err := api.Ipam.IpamPrefixesAvailablePrefixesCreate(params, nil)
	if err != nil {
		return err
	}

	payload := res.GetPayload()
	d.SetId(strconv.FormatInt(payload.ID, 10))
	d.Set("prefix", payload.Prefix)

	return resourceNetboxPrefixUpdate(d, m)
}
