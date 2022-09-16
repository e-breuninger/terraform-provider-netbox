package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetboxDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceCreate,
		ReadContext:   resourceNetboxDeviceRead,
		UpdateContext: resourceNetboxDeviceUpdate,
		DeleteContext: resourceNetboxDeviceDelete,

		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):From the [official documentation](https://docs.netbox.dev/en/stable/core-functionality/devices/#devices):

> Every piece of hardware which is installed within a site or rack exists in NetBox as a device. Devices are measured in rack units (U) and can be half depth or full depth. A device may have a height of 0U: These devices do not consume vertical rack space and cannot be assigned to a particular rack unit. A common example of a 0U device is a vertically-mounted PDU.`,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"device_type_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"device_role_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Default:  models.DeviceStatusValueActive,
				Optional: true,
			},
			"cluster_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"face": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_device_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"position_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rack_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vc_position_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vc_priority_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_chassis_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			tagsKey: tagsSchema,
			"primary_ipv4": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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

	comments := d.Get("comments").(string)
	data.Comments = comments

	serial := d.Get("serial").(string)
	data.Serial = serial

	tenantIDValue, ok := d.GetOk("tenant_id")
	if ok {
		tenantID := int64(tenantIDValue.(int))
		data.Tenant = &tenantID
	}

	locationIDValue, ok := d.GetOk("location_id")
	if ok {
		locationID := int64(locationIDValue.(int))
		data.Location = &locationID
	}

	roleIDValue, ok := d.GetOk("device_role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.DeviceRole = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	assetTagValue, ok := d.GetOk("asset_tag")
	if ok {
		assetTag := assetTagValue.(string)
		data.AssetTag = &assetTag
	}

	clusterIDValue, ok := d.GetOk("cluster_id")
	if ok {
		clusterID := int64(clusterIDValue.(int))
		data.Cluster = &clusterID
	}

	face := d.Get("face").(string)
	data.AssetTag = &face

	parentDeviceIDValue, ok := d.GetOk("parent_device_id")
	if ok {
		parentDeviceID := int64(parentDeviceIDValue.(int))
		data.ParentDevice.ID = parentDeviceID
	}

	platformIDValue, ok := d.GetOk("platform_id")
	if ok {
		platformID := int64(platformIDValue.(int))
		data.Platform = &platformID
	}

	positionIDValue, ok := d.GetOk("position_id")
	if ok {
		positionID := int64(positionIDValue.(int))
		data.Position = &positionID
	}

	rackIDValue, ok := d.GetOk("rack_id")
	if ok {
		rackID := int64(rackIDValue.(int))
		data.Rack = &rackID
	}

	vcPositionIDValue, ok := d.GetOk("vc_position_id")
	if ok {
		vcPositionID := int64(vcPositionIDValue.(int))
		data.VcPosition = &vcPositionID
	}

	vcPriorityIDValue, ok := d.GetOk("vc_priority_id")
	if ok {
		vcPriorityID := int64(vcPriorityIDValue.(int))
		data.VcPriority = &vcPriorityID
	}

	vcIDValue, ok := d.GetOk("virtual_chassis_id")
	if ok {
		vcID := int64(vcIDValue.(int))
		data.VirtualChassis = &vcID
	}

	cfValue, ok := d.GetOk("custom_fields")
	if ok {
		data.CustomFields = cfValue.(map[string]interface{})
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	params := dcim.NewDcimDevicesCreateParams().WithData(&data)

	res, err := api.Dcim.DcimDevicesCreate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(res.GetPayload().ID, 10))

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)

	params := dcim.NewDcimDevicesReadParams().WithID(id)

	res, err := api.Dcim.DcimDevicesRead(params, nil)
	if err != nil {
		errorcode := err.(*dcim.DcimDevicesReadDefault).Code()
		if errorcode == 404 {
			// If the ID is updated to blank, this tells Terraform the resource no longer exists (maybe it was destroyed out of band). Just like the destroy callback, the Read function should gracefully handle this case. https://www.terraform.io/docs/extend/writing-custom-providers.html
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("name", res.GetPayload().Name)

	if res.GetPayload().DeviceType != nil {
		d.Set("device_type_id", res.GetPayload().DeviceType.ID)
	}

	if res.GetPayload().PrimaryIp4 != nil {
		d.Set("primary_ipv4", res.GetPayload().PrimaryIp4.ID)
	} else {
		d.Set("primary_ipv4", nil)
	}

	if res.GetPayload().Tenant != nil {
		d.Set("tenant_id", res.GetPayload().Tenant.ID)
	} else {
		d.Set("tenant_id", nil)
	}

	if res.GetPayload().Location != nil {
		d.Set("location_id", res.GetPayload().Location.ID)
	} else {
		d.Set("location_id", nil)
	}

	if res.GetPayload().DeviceRole != nil {
		d.Set("device_role_id", res.GetPayload().DeviceRole.ID)
	} else {
		d.Set("device_role_id", nil)
	}

	if res.GetPayload().Site != nil {
		d.Set("site_id", res.GetPayload().Site.ID)
	} else {
		d.Set("site_id", nil)
	}

	if res.GetPayload().ParentDevice != nil {
		d.Set("parent_device_id", res.GetPayload().ParentDevice.ID)
	} else {
		d.Set("parent_device_id", nil)
	}

	if res.GetPayload().Cluster != nil {
		d.Set("cluster_id", res.GetPayload().Cluster.ID)
	} else {
		d.Set("cluster_id", nil)
	}

	if res.GetPayload().Platform != nil {
		d.Set("platform_id", res.GetPayload().Platform.ID)
	} else {
		d.Set("platform_id", nil)
	}

	if res.GetPayload().Rack != nil {
		d.Set("rack_id", res.GetPayload().Rack.ID)
	} else {
		d.Set("rack_id", nil)
	}

	if res.GetPayload().VirtualChassis != nil {
		d.Set("virtual_chassis_id", res.GetPayload().VirtualChassis.ID)
	} else {
		d.Set("virtual_chassis_id", nil)
	}

	if res.GetPayload().Status != nil {
		d.Set("status", res.GetPayload().Status.Value)
	} else {
		d.Set("status", nil)
	}

	if res.GetPayload().AssetTag != nil {
		d.Set("asset_tag", res.GetPayload().AssetTag)
	}

	if res.GetPayload().Face != nil {
		d.Set("face", res.GetPayload().Face.Value)
	} else {
		d.Set("face", nil)
	}

	d.Set("position_id", res.GetPayload().Position)

	d.Set("vc_position_id", res.GetPayload().VcPosition)

	d.Set("vc_priority_id", res.GetPayload().VcPriority)

	d.Set("comments", res.GetPayload().Comments)

	d.Set("serial", res.GetPayload().Serial)

	d.Set("custom_fields", res.GetPayload().CustomFields)

	d.Set(tagsKey, getTagListFromNestedTagList(res.GetPayload().Tags))
	return diags
}

func resourceNetboxDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	data := models.WritableDeviceWithConfigContext{}

	name := d.Get("name").(string)
	data.Name = &name

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

	locationIDValue, ok := d.GetOk("location_id")
	if ok {
		locationID := int64(locationIDValue.(int))
		data.Location = &locationID
	}

	roleIDValue, ok := d.GetOk("device_role_id")
	if ok {
		roleID := int64(roleIDValue.(int))
		data.DeviceRole = &roleID
	}

	siteIDValue, ok := d.GetOk("site_id")
	if ok {
		siteID := int64(siteIDValue.(int))
		data.Site = &siteID
	}

	commentsValue, ok := d.GetOk("comments")
	if ok {
		comments := commentsValue.(string)
		data.Comments = comments
	} else {
		comments := " "
		data.Comments = comments
	}

	primaryIPValue, ok := d.GetOk("primary_ipv4")
	if ok {
		primaryIP := int64(primaryIPValue.(int))
		data.PrimaryIp4 = &primaryIP
	}

	data.Tags, _ = getNestedTagListFromResourceDataSet(api, d.Get(tagsKey))

	if d.HasChanges("comments") {
		// check if comment is set
		commentsValue, ok := d.GetOk("comments")
		comments := ""
		if !ok {
			// Setting an space string deletes the comment
			comments = " "
		} else {
			comments = commentsValue.(string)
		}
		data.Comments = comments
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

	if d.HasChange("status") {
		status := d.Get("status").(string)
		data.Status = status
	}

	if d.HasChange("asset_tag") {
		assetTag := d.Get("asset_tag").(string)
		data.AssetTag = &assetTag
	}

	if d.HasChange("cluster_id") {
		clusterID := int64(d.Get("cluster_id").(int))
		data.Cluster = &clusterID
	}

	if d.HasChange("face") {
		face := d.Get("face").(string)
		data.Face = &face
	}

	if d.HasChange("parent_device_id") {
		data.ParentDevice.ID = int64(d.Get("parent_device_id").(int))
	}

	if d.HasChange("platform_id") {
		platformID := int64(d.Get("platform_id").(int))
		data.Platform = &platformID
	}

	if d.HasChange("position_id") {
		positionID := int64(d.Get("parent_device_id").(int))
		data.Position = &positionID
	}

	if d.HasChange("rack_id") {
		rackID := int64(d.Get("rack_id").(int))
		data.Rack = &rackID
	}

	if d.HasChange("vc_position_id") {
		vcPositionID := int64(d.Get("vc_position_id").(int))
		data.VcPosition = &vcPositionID
	}

	if d.HasChange("vc_priority_id") {
		vcPriorityID := int64(d.Get("vc_priority_id").(int))
		data.VcPriority = &vcPriorityID
	}

	if d.HasChange("virtual_chassis_id") {
		vcID := int64(d.Get("virtual_chassis_id").(int))
		data.VirtualChassis = &vcID
	}

	if d.HasChange("custom_fields") {
		data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	params := dcim.NewDcimDevicesUpdateParams().WithID(id).WithData(&data)

	_, err := api.Dcim.DcimDevicesUpdate(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	id, _ := strconv.ParseInt(d.Id(), 10, 64)
	params := dcim.NewDcimDevicesDeleteParams().WithID(id)

	_, err := api.Dcim.DcimDevicesDelete(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
