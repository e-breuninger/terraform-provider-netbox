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

func resourceNetboxDeviceInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetboxDeviceInterfaceCreate,
		ReadContext:   resourceNetboxDeviceInterfaceRead,
		UpdateContext: resourceNetboxDeviceInterfaceUpdate,
		DeleteContext: resourceNetboxDeviceInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					models.InterfaceTypeValueVirtual,
					models.InterfaceTypeValueLag,
					models.InterfaceTypeValueNr100baseDashTx,
					models.InterfaceTypeValueNr1000baseDasht,
					models.InterfaceTypeValueNr2Dot5gbaseDasht,
					models.InterfaceTypeValueNr5gbaseDasht,
					models.InterfaceTypeValueNr10gbaseDasht,
					models.InterfaceTypeValueNr10gbaseDashCx4,
					models.InterfaceTypeValueNr1000baseDashxDashGbic,
					models.InterfaceTypeValueNr1000baseDashxDashSfp,
					models.InterfaceTypeValueNr10gbaseDashxDashSfpp,
					models.InterfaceTypeValueNr10gbaseDashxDashXfp,
					models.InterfaceTypeValueNr10gbaseDashxDashXenpak,
					models.InterfaceTypeValueNr10gbaseDashxDashX2,
					models.InterfaceTypeValueNr25gbaseDashxDashSfp28,
					models.InterfaceTypeValueNr50gbaseDashxDashSfp56,
					models.InterfaceTypeValueNr40gbaseDashxDashQsfpp,
					models.InterfaceTypeValueNr50gbaseDashxDashSfp28,
					models.InterfaceTypeValueNr100gbaseDashxDashCfp,
					models.InterfaceTypeValueNr100gbaseDashxDashCfp2,
					models.InterfaceTypeValueNr200gbaseDashxDashCfp2,
					models.InterfaceTypeValueNr100gbaseDashxDashCfp4,
					models.InterfaceTypeValueNr100gbaseDashxDashCpak,
					models.InterfaceTypeValueNr100gbaseDashxDashQsfp28,
					models.InterfaceTypeValueNr200gbaseDashxDashQsfp56,
					models.InterfaceTypeValueNr400gbaseDashxDashQsfpdd,
					models.InterfaceTypeValueNr400gbaseDashxDashOsfp,
					models.InterfaceTypeValueIeee802Dot11a,
					models.InterfaceTypeValueIeee802Dot11g,
					models.InterfaceTypeValueIeee802Dot11n,
					models.InterfaceTypeValueIeee802Dot11ac,
					models.InterfaceTypeValueIeee802Dot11ad,
					models.InterfaceTypeValueIeee802Dot11ax,
					models.InterfaceTypeValueGsm,
					models.InterfaceTypeValueCdma,
					models.InterfaceTypeValueLte,
					models.InterfaceTypeValueSonetDashOc3,
					models.InterfaceTypeValueSonetDashOc12,
					models.InterfaceTypeValueSonetDashOc48,
					models.InterfaceTypeValueSonetDashOc192,
					models.InterfaceTypeValueSonetDashOc768,
					models.InterfaceTypeValueSonetDashOc1920,
					models.InterfaceTypeValueSonetDashOc3840,
					models.InterfaceTypeValueNr1gfcDashSfp,
					models.InterfaceTypeValueNr2gfcDashSfp,
					models.InterfaceTypeValueNr4gfcDashSfp,
					models.InterfaceTypeValueNr8gfcDashSfpp,
					models.InterfaceTypeValueNr16gfcDashSfpp,
					models.InterfaceTypeValueNr32gfcDashSfp28,
					models.InterfaceTypeValueNr64gfcDashQsfpp,
					models.InterfaceTypeValueNr128gfcDashSfp28,
					models.InterfaceTypeValueInfinibandDashSdr,
					models.InterfaceTypeValueInfinibandDashDdr,
					models.InterfaceTypeValueInfinibandDashQdr,
					models.InterfaceTypeValueInfinibandDashFdr10,
					models.InterfaceTypeValueInfinibandDashFdr,
					models.InterfaceTypeValueInfinibandDashEdr,
					models.InterfaceTypeValueInfinibandDashHdr,
					models.InterfaceTypeValueInfinibandDashNdr,
					models.InterfaceTypeValueInfinibandDashXdr,
					models.InterfaceTypeValueT1,
					models.InterfaceTypeValueE1,
					models.InterfaceTypeValueT3,
					models.InterfaceTypeValueE3,
					models.InterfaceTypeValueCiscoDashStackwise,
					models.InterfaceTypeValueCiscoDashStackwiseDashPlus,
					models.InterfaceTypeValueCiscoDashFlexstack,
					models.InterfaceTypeValueCiscoDashFlexstackDashPlus,
					models.InterfaceTypeValueJuniperDashVcp,
					models.InterfaceTypeValueExtremeDashSummitstack,
					models.InterfaceTypeValueExtremeDashSummitstackDash128,
					models.InterfaceTypeValueExtremeDashSummitstackDash256,
					models.InterfaceTypeValueExtremeDashSummitstackDash512,
					models.InterfaceTypeValueOther,
				}, false),
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"connection_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"management_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					models.InterfaceModeValueAccess,
					models.InterfaceModeValueTagged,
					models.InterfaceModeValueTaggedDashAll,
				}, false),
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tagged_vlan": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"untagged_vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"tags": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Set:      schema.HashString,
			},
		},
	}
}

func resourceNetboxDeviceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	interfaceID := int64(d.Get("device_id").(int))
	interfaceType := d.Get("type").(string)
	name := d.Get("name").(string)
	tags, _ := getNestedTagListFromResourceDataSet(c, d.Get("tags"))

	params := &dcim.DcimInterfacesCreateParams{
		Context: ctx,
	}

	params.Data = &models.WritableInterface{
		Device:      &interfaceID,
		Type:        &interfaceType,
		Name:        &name,
		Tags:           tags,
		TaggedVlans: expandTaggedVlans(d.Get("tagged_vlan").([]interface{})),
	}

	//if v, ok := d.GetOk("connection_status"); ok {
	//connectionStatus := v.(bool)
	//params.Data.ConnectionStatus = &connectionStatus
	//}

	if v, ok := d.GetOk("enabled"); ok {
		params.Data.Enabled = v.(bool)
	}

	if v, ok := d.GetOk("management_only"); ok {
		params.Data.MgmtOnly = v.(bool)
	}

	if v, ok := d.GetOk("label"); ok {
		params.Data.Label = v.(string)
	}

	if v, ok := d.GetOk("mac_address"); ok {
		macAddress := v.(string)
		params.Data.MacAddress = &macAddress
	}

	if v, ok := d.GetOk("mode"); ok {
		params.Data.Mode = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		params.Data.Description = v.(string)
	}

	if v, ok := d.GetOk("untagged_vlan_id"); ok {
		untaggedVlan := int64(v.(int))
		params.Data.UntaggedVlan = &untaggedVlan
	}
	if v, ok := d.GetOk("mtu"); ok {
		mtu := int64(v.(int))
		params.Data.Mtu = &mtu
	}

	resp, err := c.Dcim.DcimInterfacesCreate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to create interface: %v", err)
	}

	d.SetId(strconv.FormatInt(resp.Payload.ID, 10))

	resourceNetboxDeviceInterfaceRead(ctx, d, m)

	return diags
}

func resourceNetboxDeviceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimInterfacesReadParams{
		Context: ctx,
		ID:      objectID,
	}

	resp, err := c.Dcim.DcimInterfacesRead(params, nil)
	if err != nil {
		if err.(*dcim.DcimInterfacesReadDefault).Code() == 404 {
			d.SetId("")
			return nil
		}

		return diag.Errorf("Unable to get interface: %v", err)
	}

	d.Set("device_id", resp.Payload.Device.ID)
	d.Set("type", resp.Payload.Type.Value)
	d.Set("name", resp.Payload.Name)

	//if resp.Payload.ConnectionStatus != nil {
	//	d.Set("connection_status", resp.Payload.ConnectionStatus.Value)
	//}

	if resp.Payload.Label != "" {
		d.Set("label", resp.Payload.Label)
	}

	if resp.Payload.MacAddress != nil {
		d.Set("mac_address", resp.Payload.MacAddress)
	}

	if resp.Payload.Mode != nil {
		d.Set("mode", resp.Payload.Mode.Value)
	}

	if resp.Payload.Description != "" {
		d.Set("description", resp.Payload.Description)
	}

	if resp.Payload.UntaggedVlan != nil {
		d.Set("untagged_vlan_id", resp.Payload.UntaggedVlan.ID)
	}

	if resp.Payload.Mtu != nil {
		d.Set("mtu", resp.Payload.Mtu)
	}

	d.Set("enabled", resp.Payload.Enabled)
	d.Set("management_only", resp.Payload.MgmtOnly)
	d.Set("tagged_vlan", flattenTaggedVlans(resp.Payload.TaggedVlans))
	d.Set("tags", getTagListFromNestedTagList(resp.GetPayload().Tags))

	return diags
}

func resourceNetboxDeviceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	deviceID := int64(d.Get("device_id").(int))
	interfaceType := d.Get("type").(string)
	name := d.Get("name").(string)
	tags, _ := getNestedTagListFromResourceDataSet(c, d.Get("tags"))

	params := &dcim.DcimInterfacesPartialUpdateParams{
		Context: ctx,
		ID:      objectID,
	}

	params.Data = &models.WritableInterface{
		Device:      &deviceID,
		Type:        &interfaceType,
		Name:        &name,
		TaggedVlans: expandTaggedVlans(d.Get("tagged_vlan").([]interface{})),
		Tags:           tags,
	}

	//	if d.HasChange("connection_status") {
	//	connectionStatus := d.Get("connection_status").(bool)
	//	params.Data.ConnectionStatus = &connectionStatus
	//}

	if d.HasChange("enabled") {
		params.Data.Enabled = d.Get("enabled").(bool)
	}

	if d.HasChange("management_only") {
		params.Data.MgmtOnly = d.Get("management_only").(bool)
	}

	if d.HasChange("label") {
		params.Data.Label = d.Get("label").(string)
	}

	if d.HasChange("mac_address") {
		macAddress := d.Get("mac_address").(string)
		params.Data.MacAddress = &macAddress
	}

	if d.HasChange("mode") {
		params.Data.Mode = d.Get("mode").(string)
	}

	if d.HasChange("description") {
		params.Data.Description = d.Get("description").(string)
	}

	if d.HasChange("untagged_vlan_id") {
		untaggedVlan := int64(d.Get("untagged_vlan_id").(int))
		params.Data.UntaggedVlan = &untaggedVlan
	}

	if d.HasChange("mtu") {
		mtu := int64(d.Get("mtu").(int))
		params.Data.Mtu = &mtu
	}

	_, err = c.Dcim.DcimInterfacesPartialUpdate(params, nil)
	if err != nil {
		return diag.Errorf("Unable to update interface: %v", err)
	}

	return resourceNetboxDeviceInterfaceRead(ctx, d, m)
}

func resourceNetboxDeviceInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.NetBoxAPI)

	var diags diag.Diagnostics

	objectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to parse ID: %v", err)
	}

	params := &dcim.DcimInterfacesDeleteParams{
		Context: ctx,
		ID:      objectID,
	}

	_, err = c.Dcim.DcimInterfacesDelete(params, nil)
	if err != nil {
		return diag.Errorf("Unable to delete interface: %v", err)
	}

	d.SetId("")

	return diags
}

func expandTags(input []interface{}) []*models.NestedTag {
	if len(input) == 0 {
		return nil
	}

	results := make([]*models.NestedTag, 0)

	for _, item := range input {
		values := item.(map[string]interface{})
		result := &models.NestedTag{}

		if v, ok := values["id"]; ok {
			result.ID = int64(v.(int))
		}

		if v, ok := values["name"]; ok {
			name := v.(string)
			result.Name = &name
		}

		if v, ok := values["slug"]; ok {
			slug := v.(string)
			result.Slug = &slug
		}

		if v, ok := values["color"]; ok {
			result.Color = v.(string)
		}

		results = append(results, result)
	}

	return results
}

func flattenTags(input []*models.NestedTag) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["name"] = item.Name
		values["slug"] = item.Slug
		values["color"] = item.Color

		result = append(result, values)
	}

	return result
}

func expandTaggedVlans(input []interface{}) []int64 {
	results := make([]int64, 0)

	for _, item := range input {
		value := item.(int)
		results = append(results, int64(value))
	}

	return results
}

func flattenTaggedVlans(input []*models.NestedVLAN) []interface{} {
	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["name"] = item.Name
		values["vid"] = item.Vid
		values["DisplayName"] = item.DisplayName

		result = append(result, values["id"])
	}

	return result
}
