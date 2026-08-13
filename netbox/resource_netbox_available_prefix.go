package netbox

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxAvailablePrefix() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxAvailablePrefixCreate,
		Read:          resourceNetboxPrefixRead,
		Update:        resourceNetboxPrefixUpdate,
		Delete:        resourceNetboxPrefixDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):This resource allocates the next available prefix from a parent prefix.

Allocations from the same parent prefix are serialized inside the provider, so several of these resources can be created in the same apply without racing each other. Conflicts caused by anything outside the current provider process, such as a second ` + "`terraform apply`" + ` or somebody working in the Netbox UI, are retried for a short while but can still fail if contention keeps up.`,

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
			"location_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"site_id", "site_group_id", "region_id"},
			},
			"site_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"location_id", "site_group_id", "region_id"},
			},
			"site_group_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"location_id", "site_id", "region_id"},
			},
			"region_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"location_id", "site_id", "site_group_id"},
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			customFieldsKey: customFieldsSchema,
			tagsKey:         tagsSchema,
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

func resourceNetboxAvailablePrefixCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*providerState)

	parentPrefixID := int64(d.Get("parent_prefix_id").(int))
	prefixLength := int64(d.Get("prefix_length").(int))
	data := models.PrefixLength{
		PrefixLength: &prefixLength,
	}
	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}
	params := ipam.NewIpamPrefixesAvailablePrefixesCreateParams().WithID(parentPrefixID).WithData(&data)

	// Only the allocation itself is serialized and retried. Everything after it
	// operates on the prefix we just got, which nothing else is competing for.
	unlock, err := api.lockAllocation(ctx, allocationLockKeyPrefix(parentPrefixID))
	if err != nil {
		return diag.FromErr(err)
	}
	var prefixID int64
	var prefix *string
	err = retryAllocation(ctx, func() error {
		res, err := api.Ipam.IpamPrefixesAvailablePrefixesCreate(params, nil)
		if err != nil {
			return err
		}
		payload := res.GetPayload()
		prefixID = payload.ID
		prefix = payload.Prefix
		return nil
	})
	unlock()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(prefixID, 10))
	d.Set("prefix", prefix)

	return diag.FromErr(resourceNetboxPrefixUpdate(d, m))
}
