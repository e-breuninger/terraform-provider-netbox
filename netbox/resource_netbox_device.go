package netbox

import (
	"context"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"

	"github.com/fbreckle/go-netbox/netbox/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNetboxDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceCreate,
		ReadContext:   resourceNetboxDeviceRead,
		UpdateContext: resourceNetboxDeviceUpdate,
		DeleteContext: resourceNetboxDeviceDelete,

		Schema: map[string]*schema.Schema{
			"device_type_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"device_role_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"comments": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 200),
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					models.DeviceStatusValueActive,
					models.DeviceStatusValueDecommissioning,
					models.DeviceStatusValueFailed,
					models.DeviceStatusValueInventory,
					models.DeviceStatusValueOffline,
					models.DeviceStatusValueStaged,
					models.DeviceStatusValuePlanned,
				}, false),
				Default: models.DeviceStatusValueActive,
			},

			"asset_tag": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},

			"cluster_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"serial": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},

			// "config_context": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },

			// "display_name": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },

			"face": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					models.DeviceFaceValueFront,
					models.DeviceFaceValueRear,
				}, false),
			},

			// "local_context_data": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },

			"name": {
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

			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},

			"custom_fields": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceNetboxDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	var deviceRoleID = int64(d.Get("device_role_id").(int))
	var deviceTypeID = int64(d.Get("device_type_id").(int))
	var siteID = int64(d.Get("site_id").(int))

	params := &dcim.DcimDevicesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableDeviceWithConfigContext{
		DeviceRole: &deviceRoleID,
		DeviceType: &deviceTypeID,
		Site:       &siteID,
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		tenantID := int64(v.(int))
		params.Data.Tenant = &tenantID
	}

	if v, ok := d.GetOk("comments"); ok {
		params.Data.Comments = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		params.Data.Status = v.(string)
	}

	if v, ok := d.GetOk("asset_tag"); ok {
		assetTag := v.(string)
		params.Data.AssetTag = &assetTag
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		clusterID := int64(v.(int))
		params.Data.Cluster = &clusterID
	}

	if v, ok := d.GetOk("serial"); ok {
		params.Data.Serial = v.(string)
	}

	// if v, ok := d.GetOk("config_context"); ok {
	// 	params.Data.ConfigContext = v.(map[string]string)
	// }

	// if v, ok := d.GetOk("display_name"); ok {
	// 	params.Data.DisplayName = v.(string)
	// }

	if v, ok := d.GetOk("face"); ok {
		params.Data.Face = v.(string)
	}

	// if v, ok := d.GetOk("local_context_data"); ok {
	// 	params.Data.LocalContextData = v.(*string)
	// }

	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		params.Data.Name = &name
	}

	if v, ok := d.GetOk("parent_device_id"); ok {
		params.Data.ParentDevice.ID = int64(v.(int))
	}

	if v, ok := d.GetOk("platform_id"); ok {
		platFormID := int64(v.(int))
		params.Data.Platform = &platFormID
	}

	if v, ok := d.GetOk("position_id"); ok {
		positionID := int64(v.(int))
		params.Data.Position = &positionID
	}
	if v, ok := d.GetOk("rack_id"); ok {
		rackID := int64(v.(int))
		params.Data.Rack = &rackID
	}
	if v, ok := d.GetOk("vc_position_id"); ok {
		vcPositionID := int64(v.(int))
		params.Data.VcPosition = &vcPositionID
	}
	if v, ok := d.GetOk("vc_priority_id"); ok {
		vcPriorityID := int64(v.(int))
		params.Data.VcPriority = &vcPriorityID
	}
	if v, ok := d.GetOk("virtual_chassis_id"); ok {
		vcID := int64(v.(int))
		params.Data.VirtualChassis = &vcID
	}

	if v, ok := d.GetOk("tags"); ok {
		tags, _ := getNestedTagListFromResourceDataSet(c, v)
		params.Data.Tags = tags
	}

	if v, ok := d.GetOk("custom_fields"); ok {
		params.Data.CustomFields = v.(map[string]interface{})
	}

	resp, err := c.Dcim.DcimDevicesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create rack: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceNetboxDeviceRead(ctx, d, m)

	return diags
}

func resourceNetboxDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimDevicesReadParams{
		Context: ctx,
		ID:      objectID,
	}

	resp, err := c.Dcim.DcimDevicesRead(params, nil)
	if err != nil {
		if err.(*dcim.DcimDevicesReadDefault).Code() == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get device: %v", err)
	}

	d.Set("site_id", resp.Payload.Site.ID)
	d.Set("device_type_id", resp.Payload.DeviceType.ID)
	d.Set("device_role_id", resp.Payload.DeviceRole.ID)

	if resp.Payload.Tenant != nil {
		d.Set("tenant_id", resp.Payload.Tenant.ID)
	}

	if resp.Payload.Comments != "" {
		d.Set("comments", resp.Payload.Comments)
	}

	if resp.Payload.Status != nil {
		d.Set("status", resp.Payload.Status.Value)
	}

	if resp.Payload.AssetTag != nil {
		d.Set("asset_tag", resp.Payload.AssetTag)
	}

	if resp.Payload.Cluster != nil {
		d.Set("cluster_id", resp.Payload.Cluster.ID)
	}

	if resp.Payload.Serial != "" {
		d.Set("serial", resp.Payload.Serial)
	}

	// if resp.Payload.ConfigContext != nil {
	// 	d.Set("config_context", resp.Payload.ConfigContext)
	// }

	// if resp.Payload.DisplayName != "" {
	// 	d.Set("display_name", resp.Payload.DisplayName)
	// }

	if resp.Payload.Face != nil {
		d.Set("face", resp.Payload.Face.Value)
	}

	// if resp.Payload.LocalContextData != nil {
	// 	d.Set("local_context_data", resp.Payload.LocalContextData)
	// }

	if resp.Payload.Name != nil {
		d.Set("name", resp.Payload.Name)
	}

	if resp.Payload.ParentDevice != nil {
		d.Set("parent_device_id", resp.Payload.ParentDevice.ID)
	}

	if resp.Payload.Platform != nil {
		d.Set("platform_id", resp.Payload.Platform.ID)
	}

	if resp.Payload.Position != nil {
		d.Set("position_id", resp.Payload.Position)
	}

	if resp.Payload.Rack != nil {
		d.Set("rack_id", resp.Payload.Rack.ID)
	}

	if resp.Payload.VcPosition != nil {
		d.Set("vc_position_id", resp.Payload.VcPosition)
	}

	if resp.Payload.VcPriority != nil {
		d.Set("vc_priority_id", resp.Payload.VcPriority)
	}

	if resp.Payload.VirtualChassis != nil {
		d.Set("virtual_chassis_id", resp.Payload.VirtualChassis.ID)
	}

	d.Set("tags", getTagListFromNestedTagList(resp.GetPayload().Tags))
	d.Set("custom_fields", resp.Payload.CustomFields)

	return diags
}

func resourceNetboxDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	deviceTypeID := int64(d.Get("device_type_id").(int))
	deviceRoleID := int64(d.Get("device_role_id").(int))
	siteID := int64(d.Get("site_id").(int))

	params := &dcim.DcimDevicesPartialUpdateParams{
		Context: ctx,
		ID:      objectID,
	}

	params.Data = &models.WritableDeviceWithConfigContext{
		DeviceType: &deviceTypeID,
		DeviceRole: &deviceRoleID,
		Site:       &siteID,
	}

	if d.HasChange("tenant_id") {
		tenantID := int64(d.Get("tenant_id").(int))
		params.Data.Tenant = &tenantID
	}

	if d.HasChange("comments") {
		params.Data.Comments = d.Get("comments").(string)
	}

	if d.HasChange("status") {
		params.Data.Status = d.Get("status").(string)
	}

	if d.HasChange("asset_tag") {
		aseetTag := d.Get("asset_tag").(string)
		params.Data.AssetTag = &aseetTag
	}

	if d.HasChange("cluster_id") {
		clusterID := int64(d.Get("cluster_id").(int))
		params.Data.Cluster = &clusterID
	}

	if d.HasChange("serial") {
		params.Data.Serial = d.Get("serial").(string)
	}

	// if d.HasChange("config_context") {
	// 	params.Data.ConfigContext = d.Get("config_context").(map[string]string)
	// }

	// if d.HasChange("display_name") {
	// 	params.Data.DisplayName = d.Get("display_name").(string)
	// }

	if d.HasChange("face") {
		params.Data.Face = d.Get("face").(string)
	}

	// if d.HasChange("local_context_data") {
	// 	localContextData := d.Get("local_context_data").(string)
	// 	params.Data.LocalContextData = &localContextData
	// }

	if d.HasChange("name") {
		name := d.Get("name").(string)
		params.Data.Name = &name
	}

	if d.HasChange("parent_device_id") {
		params.Data.ParentDevice.ID = int64(d.Get("parent_device_id").(int))
	}

	if d.HasChange("platform_id") {
		platformID := int64(d.Get("platform_id").(int))
		params.Data.Platform = &platformID
	}

	if d.HasChange("position_id") {
		positionID := int64(d.Get("parent_device_id").(int))
		params.Data.Position = &positionID
	}

	if d.HasChange("rack_id") {
		rackID := int64(d.Get("rack_id").(int))
		params.Data.Rack = &rackID
	}

	if d.HasChange("vc_position_id") {
		vcPositionID := int64(d.Get("vc_position_id").(int))
		params.Data.VcPosition = &vcPositionID
	}

	if d.HasChange("vc_priority_id") {
		vcPriorityID := int64(d.Get("vc_priority_id").(int))
		params.Data.VcPriority = &vcPriorityID
	}

	if d.HasChange("virtual_chassis_id") {
		vcID := int64(d.Get("virtual_chassis_id").(int))
		params.Data.VirtualChassis = &vcID
	}

	if d.HasChange("tags") {
		params.Data.Tags, _ = getNestedTagListFromResourceDataSet(c, d.Get("tags"))
	}

	if d.HasChange("custom_fields") {
		params.Data.CustomFields = d.Get("custom_fields").(map[string]interface{})
	}

	_, err = c.Dcim.DcimDevicesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update rack: %v", err)
	}

	return resourceNetboxDeviceRead(ctx, d, m)
}

func resourceNetboxDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimDevicesDeleteParams{
		Context: ctx,
		ID:      objectID,
	}

	_, err = c.Dcim.DcimDevicesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete device: %v", err)
	}

	d.SetId("")

	return diags
}
