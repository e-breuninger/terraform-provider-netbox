package netbox

import (
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxL2VPNTypeOptions = []string{"vpws", "vpls", "vxlan", "vxlan-evpn", "mpls-evpn", "pbb-evpn", "epl", "evpl", "ep-lan", "evp-lan", "ep-tree", "evp-tree"}

func resourceNetboxL2VPN() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxL2VPNCreate,
		Read:   resourceNetboxL2VPNRead,
		Update: resourceNetboxL2VPNUpdate,
		Delete: resourceNetboxL2VPNDelete,

		Description: `:meta:subcategory:IP Address Management (IPAM):From the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/l2vpn/):

> This model represents a Layer 2 VPN, such as VPLS/VPWS, VXLAN, or MPLS EVPN. Each L2VPN can be assigned to multiple import and export route targets, however this is optional. Once created, an L2VPN can be assigned to a VLAN or an interface.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxL2VPNTypeOptions, false),
				Description:  buildValidValueDescription(resourceNetboxL2VPNTypeOptions),
			},
			"identifier": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"import_targets": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"export_targets": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			customFieldsKey: customFieldsSchema,
			tagsKey:         tagsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func getIDsFromNestedRouteTargetList(nestedRouteTargets []*models.NestedRouteTarget) []int64 {
	var routeTargets []int64
	for _, rt := range nestedRouteTargets {
		routeTargets = append(routeTargets, rt.ID)
	}
	return routeTargets
}

func resourceNetboxL2VPNCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)
	vpnType := d.Get("type").(string)

	data := models.WritableL2VPN{
		Name:        &name,
		Slug:        &slug,
		Type:        &vpnType,
		Description: d.Get("description").(string),
		Comments:    d.Get("comments").(string),
	}

	if identifier, ok := d.GetOk("identifier"); ok {
		data.Identifier = int64ToPtr(int64(identifier.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	data.ImportTargets = toInt64List(d.Get("import_targets"))
	data.ExportTargets = toInt64List(d.Get("export_targets"))

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamL2vpnsCreateParams().WithData(&data)
	res, err := api.Ipam.IpamL2vpnsCreate(params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxL2VPNRead(d, m)
}

func resourceNetboxL2VPNRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamL2vpnsReadParams().WithID(id)

	res, err := api.Ipam.IpamL2vpnsRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamL2vpnsReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return err
	}

	l2vpn := res.GetPayload()
	d.Set("name", l2vpn.Name)
	d.Set("slug", l2vpn.Slug)
	d.Set("description", l2vpn.Description)
	d.Set("comments", l2vpn.Comments)

	if l2vpn.Type != nil {
		d.Set("type", l2vpn.Type.Value)
	} else {
		d.Set("type", nil)
	}

	if l2vpn.Identifier != nil {
		d.Set("identifier", l2vpn.Identifier)
	} else {
		d.Set("identifier", nil)
	}

	if l2vpn.Tenant != nil {
		d.Set("tenant_id", l2vpn.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	d.Set("import_targets", getIDsFromNestedRouteTargetList(l2vpn.ImportTargets))
	d.Set("export_targets", getIDsFromNestedRouteTargetList(l2vpn.ExportTargets))

	cf := getCustomFields(l2vpn.CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}
	api.readTags(d, l2vpn.Tags)

	return nil
}

func resourceNetboxL2VPNUpdate(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)
	vpnType := d.Get("type").(string)

	data := models.WritableL2VPN{
		Name:        &name,
		Slug:        &slug,
		Type:        &vpnType,
		Description: getOptionalStr(d, "description", true),
		Comments:    getOptionalStr(d, "comments", true),
	}

	if identifier, ok := d.GetOk("identifier"); ok {
		data.Identifier = int64ToPtr(int64(identifier.(int)))
	}

	if tenantID, ok := d.GetOk("tenant_id"); ok {
		data.Tenant = int64ToPtr(int64(tenantID.(int)))
	}

	data.ImportTargets = toInt64List(d.Get("import_targets"))
	data.ExportTargets = toInt64List(d.Get("export_targets"))

	var err error
	data.Tags, err = getNestedTagListFromResourceDataSet(api, d.Get(tagsAllKey))
	if err != nil {
		return err
	}

	if cf, ok := d.GetOk(customFieldsKey); ok {
		data.CustomFields = cf
	}

	params := ipam.NewIpamL2vpnsPartialUpdateParams().WithID(id).WithData(&data)
	_, err = api.Ipam.IpamL2vpnsPartialUpdate(params, nil)
	if err != nil {
		return err
	}

	return resourceNetboxL2VPNRead(d, m)
}

func resourceNetboxL2VPNDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*providerState)
	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := ipam.NewIpamL2vpnsDeleteParams().WithID(id)

	_, err := api.Ipam.IpamL2vpnsDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*ipam.IpamL2vpnsDeleteDefault); ok {
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
