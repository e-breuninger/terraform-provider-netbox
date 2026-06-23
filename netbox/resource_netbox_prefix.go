package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxPrefixStatusOptions = []string{"active", "container", "reserved", "deprecated"}

func resourceNetboxPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxPrefixCreate,
		Read:   resourceNetboxPrefixRead,
		Update: resourceNetboxPrefixUpdate,
		Delete: resourceNetboxPrefixDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#prefixes):

> A prefix is an IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. 192.0.2.0/24). A prefix entails only the "network portion" of an IP address: All bits in the address not covered by the mask must be zero. (In other words, a prefix cannot be a specific IP address.)
>
> Prefixes are automatically organized by their parent aggregates. Additionally, each prefix can be assigned to a particular site and virtual routing and forwarding instance (VRF). Each VRF represents a separate IP space or routing table. All prefixes not assigned to a VRF are considered to be in the "global" table.`,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsCIDR,
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
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceNetboxPrefixCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	data := models.WritablePrefix{}

	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	description := d.Get("description").(string)
	isPool := d.Get("is_pool").(bool)
	markUtilized := d.Get("mark_utilized").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.Description = description
	data.IsPool = isPool

	data.MarkUtilized = markUtilized

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	siteID := getOptionalInt(d, "site_id")
	siteGroupID := getOptionalInt(d, "site_group_id")
	locationID := getOptionalInt(d, "location_id")
	regionID := getOptionalInt(d, "region_id")

	switch {
	case siteID != nil:
		data.ScopeType = strToPtr("dcim.site")
		data.ScopeID = siteID
	case siteGroupID != nil:
		data.ScopeType = strToPtr("dcim.sitegroup")
		data.ScopeID = siteGroupID
	case locationID != nil:
		data.ScopeType = strToPtr("dcim.location")
		data.ScopeID = locationID
	case regionID != nil:
		data.ScopeType = strToPtr("dcim.region")
		data.ScopeID = regionID
	default:
		data.ScopeType = nil
		data.ScopeID = nil
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := ipam.NewIpamPrefixesCreateParams().WithData(&data)
	res, err := api.Ipam.IpamPrefixesCreate(params, nil)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamPrefixesReadParams().WithID(id)

	res, err := api.Ipam.IpamPrefixesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamPrefixesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	prefix := res.GetPayload()
	d.Set("description", prefix.Description)
	d.Set("is_pool", prefix.IsPool)
	d.Set("mark_utilized", prefix.MarkUtilized)
	if prefix.Status != nil {
		d.Set("status", prefix.Status.Value)
	}
	if prefix.Prefix != nil {
		d.Set("prefix", prefix.Prefix)
	}

	if prefix.Vrf != nil {
		d.Set("vrf_id", prefix.Vrf.ID)
	} else {
		d.Set("vrf_id", nil)
	}

	if prefix.Tenant != nil {
		d.Set("tenant_id", prefix.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if prefix.Vlan != nil {
		d.Set("vlan_id", prefix.Vlan.ID)
	} else {
		d.Set("vlan_id", nil)
	}

	if prefix.Role != nil {
		d.Set("role_id", prefix.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	d.Set("site_id", nil)
	d.Set("site_group_id", nil)
	d.Set("location_id", nil)
	d.Set("region_id", nil)

	if prefix.ScopeType != nil && prefix.ScopeID != nil {
		scopeID := prefix.ScopeID
		switch scopeType := prefix.ScopeType; *scopeType {
		case "dcim.site":
			d.Set("site_id", scopeID)
		case "dcim.sitegroup":
			d.Set("site_group_id", scopeID)
		case "dcim.location":
			d.Set("location_id", scopeID)
		case "dcim.region":
			d.Set("region_id", scopeID)
		}
	}
	cf := flattenCustomFields(prefix.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	api.readTags(d, prefix.Tags)
	// FIGURE OUT NESTED VRF AND NESTED VLAN (from maybe interfaces?)

	return nil
}

func resourceNetboxPrefixUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritablePrefix{}
	prefix := d.Get("prefix").(string)
	status := d.Get("status").(string)
	isPool := d.Get("is_pool").(bool)
	markUtilized := d.Get("mark_utilized").(bool)

	data.Prefix = &prefix
	data.Status = status

	data.IsPool = isPool
	data.MarkUtilized = markUtilized

	if description, ok := d.GetOk("description"); ok {
		data.Description = description.(string)
	} else {
		data.Description = " "
	}

	if vrfID, ok := d.GetOk("vrf_id"); ok {
		data.Vrf = int64ToPtr(int64(vrfID.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	if vlanID, ok := d.GetOk("vlan_id"); ok {
		data.Vlan = int64ToPtr(int64(vlanID.(int)))
	}

	if roleID, ok := d.GetOk("role_id"); ok {
		data.Role = int64ToPtr(int64(roleID.(int)))
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	siteID := getOptionalInt(d, "site_id")
	siteGroupID := getOptionalInt(d, "site_group_id")
	locationID := getOptionalInt(d, "location_id")
	regionID := getOptionalInt(d, "region_id")

	switch {
	case siteID != nil:
		data.ScopeType = strToPtr("dcim.site")
		data.ScopeID = siteID
	case siteGroupID != nil:
		data.ScopeType = strToPtr("dcim.sitegroup")
		data.ScopeID = siteGroupID
	case locationID != nil:
		data.ScopeType = strToPtr("dcim.location")
		data.ScopeID = locationID
	case regionID != nil:
		data.ScopeType = strToPtr("dcim.region")
		data.ScopeID = regionID
	default:
		data.ScopeType = nil
		data.ScopeID = nil
	}

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	params := ipam.NewIpamPrefixesUpdateParams().WithID(id).WithData(&data)
	_, err = api.Ipam.IpamPrefixesUpdate(params, nil)
	if err != nil {
		return err
	}
	return resourceNetboxPrefixRead(d, m)
}

func resourceNetboxPrefixDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamPrefixesDeleteParams().WithID(id)
	_, err := api.Ipam.IpamPrefixesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamPrefixesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.SetId("")
	return nil
}
