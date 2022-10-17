package netbox

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var resourceNetboxDeviceStatusOptions = []string{"offline", "active", "planned", "staged", "failed", "inventory"}
var resourceNetboxDeviceRackFaceOptions = []string{"front", "rear"}

func resourceNetboxDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceCreate,
		ReadContext:   resourceNetboxDeviceRead,
		UpdateContext: resourceNetboxDeviceUpdate,
		DeleteContext: resourceNetboxDeviceDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/features/devices/#devices):

> Every piece of hardware which is installed within a site or rack exists in NetBox as a device. Devices are measured in rack units (U) and can be half depth or full depth. A device may have a height of 0U: These devices do not consume vertical rack space and cannot be assigned to a particular rack unit. A common example of a 0U device is a vertically-mounted PDU.`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cluster_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"serial": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_context_data": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"primary_ipv4": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary_ipv6": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(resourceNetboxDeviceStatusOptions, false),
				Description:  buildValidValueDescription(resourceNetboxDeviceStatusOptions),
				Default:      "active",
			},
			"rack_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rack_face": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"rack_position"},
				ValidateFunc: validation.StringInSlice(resourceNetboxDeviceRackFaceOptions, false),
				Description:  buildValidValueDescription(resourceNetboxDeviceRackFaceOptions),
			},
			"rack_position": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"virtual_chassis_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"virtual_chassis_master", "virtual_chassis_id"},
			},
			"virtual_chassis_position": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_chassis_priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_chassis_master": {
				Type:         schema.TypeBool,
				Optional:     true,
				RequiredWith: []string{"virtual_chassis_master", "virtual_chassis_id"},
			},
			"local_context_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This is best managed through the use of `jsonencode` and a map of settings.",
			},
			customFieldsKey: customFieldsSchema,
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetboxDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	name := d.Get("name").(string)

	data := models.WritableDeviceWithConfigContext{
		Name: &name,
	}

	typeIDValue, ok := d.GetOk("device_type_id")
	if ok {
		typeID := int64(typeIDValue.(int))
		data.DeviceType = &typeID
	}

	if assetTagValue, ok := d.GetOk("asset_tag"); ok {
		assetTag := string(assetTagValue.(string))
		data.AssetTag = &assetTag
	}

	data.Comments = d.Get("comments").(string)

	data.Description = d.Get("description").(string)

	local_context_data := d.Get("local_context_data").(map[string]interface{})
	data.LocalContextData = local_context_data

	data.Serial = d.Get("serial").(string)

	data.Status = d.Get("status").(string)

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	locationIDValue, ok := d.GetOk("location_id")
	if ok {
		locationID := int64(locationIDValue.(int))
		data.Location = &locationID
	}

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.Role = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	data.Rack = getOptionalInt(d, "rack_id")
	data.Face = getOptionalStr(d, "rack_face", false)

	rackPosition, ok := d.GetOk("rack_position")
	if ok && rackPosition.(float64) > 0 {
		data.Position = float64ToPtr(rackPosition.(float64))
	} else {
		data.Position = nil
	}

	data.VirtualChassis = getOptionalInt(d, "virtual_chassis_id")
	data.VcPosition = getOptionalInt(d, "virtual_chassis_position")
	data.VcPriority = getOptionalInt(d, "virtual_chassis_priority")

	localContextValue, ok := d.GetOk("local_context_data")
	if ok {
		var jsonObj any
		localContextBA := []byte(localContextValue.(string))
		if err := json.Unmarshal(localContextBA, &jsonObj); err == nil {
			data.LocalContextData = jsonObj
		}
	}

	ct, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = ct
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimDevicesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDevicesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	if vcMaster, ok := d.GetOk("virtual_chassis_master"); ok {
		var err error
		if vcMaster.(bool) {
			err = virtualChassisUpdateMaster(api, *data.VirtualChassis, &(res.GetPayload().ID))
		} else {
			err = virtualChassisUpdateMaster(api, *data.VirtualChassis, nil)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimDevicesReadParams().WithID(id)

	res, err := api.Dcim.DcimDevicesRead(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDevicesReadDefault); ok {
			errorcode := errresp.Code()
			if errorcode == 404 {
				// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	device := res.GetPayload()

	d.Set("name", device.Name)

	if device.DeviceType != nil {
		d.Set("device_type_id", device.DeviceType.ID)
	}

	if device.PrimaryIp4 != nil {
		d.Set("primary_ipv4", device.PrimaryIp4.ID)
	} else {
		d.Set("primary_ipv4", nil)
	}

	if device.PrimaryIp6 != nil {
		d.Set("primary_ipv6", device.PrimaryIp6.ID)
	} else {
		d.Set("primary_ipv6", nil)
	}

	if device.Tenant != nil {
		d.Set("tenant_id", device.Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if device.Platform != nil {
		d.Set("platform_id", device.Platform.ID)
	} else {
		d.Set("platform_id", nil)
	}

	if device.Location != nil {
		d.Set("location_id", device.Location.ID)
	} else {
		d.Set("location_id", nil)
	}

	if device.Cluster != nil {
		d.Set("cluster_id", device.Cluster.ID)
	} else {
		d.Set("cluster_id", nil)
	}

	if device.Role != nil {
		d.Set("role_id", device.Role.ID)
	} else {
		d.Set("role_id", nil)
	}

	if device.Site != nil {
		d.Set("site_id", device.Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	cf := getCustomFields(res.GetPayload().CustomFields)
	if cf != nil {
		d.Set(customFieldsKey, cf)
	}

	d.Set("asset_tag", device.AssetTag)

	d.Set("comments", device.Comments)

	d.Set("local_context_data", device.LocalContextData)

	d.Set("description", device.Description)

	d.Set("serial", device.Serial)

	d.Set("status", device.Status.Value)

	if device.Rack != nil {
		d.Set("rack_id", device.Rack.ID)
	} else {
		d.Set("rack_id", nil)
	}

	if device.Face != nil {
		d.Set("rack_face", device.Face.Value)
	} else {
		d.Set("rack_face", nil)
	}

	d.Set("rack_position", device.Position)

	if device.VirtualChassis != nil {
		d.Set("virtual_chassis_id", device.VirtualChassis.ID)
		d.Set(
			"virtual_chassis_master",
			device.VirtualChassis.Master != nil && device.VirtualChassis.Master.ID == device.ID,
		)
	} else {
		d.Set("virtual_chassis_id", 0)
		d.Set("virtual_chassis_master", false)
	}
	d.Set("virtual_chassis_position", device.VcPosition)
	d.Set("virtual_chassis_priority", device.VcPriority)

	if device.LocalContextData != nil {
		if jsonArr, err := json.Marshal(device.LocalContextData); err == nil {
			d.Set("local_context_data", string(jsonArr))
		}
	} else {
		d.Set("local_context_data", nil)
	}

	d.Set(tagsKey, getTagListFromNestedTagList(device.Tags))
	return diags
}

func resourceNetboxDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDeviceWithConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

	status := d.Get("status").(string)
	data.Status = status

	typeIDValue, ok := d.GetOk("device_type_id")
	if ok {
		typeID := int64(typeIDValue.(int))
		data.DeviceType = &typeID
	}

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	locationIDValue, ok := d.GetOk("location_id")
	if ok {
		locationID := int64(locationIDValue.(int))
		data.Location = &locationID
	}

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	roleIDValue, ok := d.GetOk("role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.Role = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	primaryIP4Value, ok := d.GetOk("primary_ipv4")
	if ok {
		primaryIP4 := int64(primaryIP4Value.(int))
		data.PrimaryIp4 = &primaryIP4
	}

	primaryIP6Value, ok := d.GetOk("primary_ipv6")
	if ok {
		primaryIP6 := int64(primaryIP6Value.(int))
		data.PrimaryIp6 = &primaryIP6
	}

	data.Rack = getOptionalInt(d, "rack_id")
	data.Face = getOptionalStr(d, "rack_face", false)
	data.Position = getOptionalFloat(d, "rack_position")

	data.VirtualChassis = getOptionalInt(d, "virtual_chassis_id")
	data.VcPosition = getOptionalInt(d, "virtual_chassis_position")
	data.VcPriority = getOptionalInt(d, "virtual_chassis_priority")

	localContextValue, ok := d.GetOk("local_context_data")
	if ok {
		var jsonObj any
		localContextBA := []byte(localContextValue.(string))
		if err := json.Unmarshal(localContextBA, &jsonObj); err == nil {
			data.LocalContextData = jsonObj
		}
	}

	LocalContextDataValue, ok := d.GetOk("local_context_data")
	if ok {
		local_context_data := LocalContextDataValue.(map[string]interface{})
		data.LocalContextData = local_context_data
	} else {
		local_context_data := map[string]string{}
		data.LocalContextData = local_context_data
	}

	cf, ok := d.GetOk(customFieldsKey)
	if ok {
		data.CustomFields = cf
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	if d.HasChanges("asset_tag") {
		if assetTagValue, ok := d.GetOk("asset_tag"); ok {
			assetTag := assetTagValue.(string)
			data.AssetTag = &assetTag
		} else {
			assetTag := " "
			data.AssetTag = &assetTag
		}
	}

	if d.HasChanges("comments") {
		// check if comment is set
		if commentsValue, ok := d.GetOk("comments"); ok {
			data.Comments = commentsValue.(string)
		} else {
			data.Comments = " "
		}
	}
	if d.HasChanges("description") {
		// check if description is set
		if descriptionValue, ok := d.GetOk("description"); ok {
			data.Description = descriptionValue.(string)
		} else {
			data.Description = " "
		}
	}

	if d.HasChanges("serial") {
		// check if serial is set
		serialValue, ok := d.GetOk("serial")
		serial := ""
		if !ok {
			// Setting an space string deletes the serial
			serial = " "
		} else {
			serial = serialValue.(string)
		}
		data.Serial = serial
	}

	params := dcim.NewDcimDevicesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimDevicesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("virtual_chassis_master") && data.VirtualChassis != nil {
		var err error
		if vcMaster, ok := d.GetOk("virtual_chassis_master"); ok {
			if vcMaster.(bool) {
				err = virtualChassisUpdateMaster(api, *data.VirtualChassis, &id)
			} else {
				err = virtualChassisUpdateMaster(api, *data.VirtualChassis, nil)
			}
		} else {
			// It was set before, but no longer set, remove it as master
			err = virtualChassisUpdateMaster(api, *data.VirtualChassis, nil)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	// If the device is member of a virtual chassis and it's the master, we cannot
	// delete it directly. We first need to update it to not be the master.
	if virtualChassisIDValue, ok := d.GetOk("virtual_chassis_id"); ok {
		if d.Get("virtual_chassis_master").(bool) {
			virtualChassisID := int64(virtualChassisIDValue.(int))
			err := virtualChassisUpdateMaster(api, virtualChassisID, nil)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	params := dcim.NewDcimDevicesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDevicesDelete(params, nil)
	if err != nil {
		if errresp, ok := err.(*dcim.DcimDevicesDeleteDefault); ok {
			if errresp.Code() == 404 {
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}
	return diags
}
