// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DeviceRequestDTO is the request DTO of the device resource; Payload writes it into
// the JSON request body.
type DeviceRequestDTO struct {
	Name             *string           `json:"name,omitempty"`
	DeviceType       *int64            `json:"device_type,omitempty"`
	Role             *int64            `json:"role,omitempty"`
	Site             *int64            `json:"site,omitempty"`
	Location         *int64            `json:"location,omitempty"`
	Rack             *int64            `json:"rack,omitempty"`
	Position         *float64          `json:"position,omitempty"`
	Face             *string           `json:"face,omitempty"`
	Status           *string           `json:"status,omitempty"`
	Airflow          *string           `json:"airflow,omitempty"`
	Tenant           *int64            `json:"tenant,omitempty"`
	Platform         *int64            `json:"platform,omitempty"`
	Cluster          *int64            `json:"cluster,omitempty"`
	VirtualChassis   *int64            `json:"virtual_chassis,omitempty"`
	VcPosition       *int64            `json:"vc_position,omitempty"`
	VcPriority       *int64            `json:"vc_priority,omitempty"`
	ConfigTemplate   *int64            `json:"config_template,omitempty"`
	Serial           *string           `json:"serial,omitempty"`
	AssetTag         *string           `json:"asset_tag,omitempty"`
	LocalContextData *string           `json:"local_context_data,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Comments         *string           `json:"comments,omitempty"`
	Owner            *int64            `json:"owner,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	CustomFields     map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the device resource.
func (requestDTO *DeviceRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	} else {
		payload["name"] = nil
	}
	if requestDTO.DeviceType != nil {
		payload["device_type"] = *requestDTO.DeviceType
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	}
	if requestDTO.Site != nil {
		payload["site"] = *requestDTO.Site
	}
	if requestDTO.Location != nil {
		payload["location"] = *requestDTO.Location
	} else {
		payload["location"] = nil
	}
	if requestDTO.Rack != nil {
		payload["rack"] = *requestDTO.Rack
	} else {
		payload["rack"] = nil
	}
	if requestDTO.Position != nil {
		payload["position"] = *requestDTO.Position
	} else {
		payload["position"] = nil
	}
	if requestDTO.Face != nil {
		payload["face"] = *requestDTO.Face
	} else {
		payload["face"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Airflow != nil {
		payload["airflow"] = *requestDTO.Airflow
	} else {
		payload["airflow"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Platform != nil {
		payload["platform"] = *requestDTO.Platform
	} else {
		payload["platform"] = nil
	}
	if requestDTO.Cluster != nil {
		payload["cluster"] = *requestDTO.Cluster
	} else {
		payload["cluster"] = nil
	}
	if requestDTO.VirtualChassis != nil {
		payload["virtual_chassis"] = *requestDTO.VirtualChassis
	} else {
		payload["virtual_chassis"] = nil
	}
	if requestDTO.VcPosition != nil {
		payload["vc_position"] = *requestDTO.VcPosition
	} else {
		payload["vc_position"] = nil
	}
	if requestDTO.VcPriority != nil {
		payload["vc_priority"] = *requestDTO.VcPriority
	} else {
		payload["vc_priority"] = nil
	}
	if requestDTO.ConfigTemplate != nil {
		payload["config_template"] = *requestDTO.ConfigTemplate
	} else {
		payload["config_template"] = nil
	}
	if requestDTO.Serial != nil {
		payload["serial"] = *requestDTO.Serial
	} else {
		payload["serial"] = ""
	}
	if requestDTO.AssetTag != nil {
		payload["asset_tag"] = *requestDTO.AssetTag
	} else {
		payload["asset_tag"] = nil
	}
	if requestDTO.LocalContextData != nil {
		payload["local_context_data"] = parseJSONText(*requestDTO.LocalContextData)
	} else {
		payload["local_context_data"] = nil
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Comments != nil {
		payload["comments"] = *requestDTO.Comments
	} else {
		payload["comments"] = ""
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// DeviceResponseDTO is the response DTO of the device resource, built from
// the go-netbox DeviceWithConfigContext by DeviceResponseDTOFromGoNetbox.
type DeviceResponseDTO struct {
	ID                     *int64            `json:"id,omitempty"`
	Name                   *string           `json:"name,omitempty"`
	DeviceType             *int64            `json:"device_type,omitempty"`
	Role                   *int64            `json:"role,omitempty"`
	Site                   *int64            `json:"site,omitempty"`
	Location               *int64            `json:"location,omitempty"`
	Rack                   *int64            `json:"rack,omitempty"`
	Position               *float64          `json:"position,omitempty"`
	Face                   *string           `json:"face,omitempty"`
	Status                 *string           `json:"status,omitempty"`
	Airflow                *string           `json:"airflow,omitempty"`
	Tenant                 *int64            `json:"tenant,omitempty"`
	Platform               *int64            `json:"platform,omitempty"`
	Cluster                *int64            `json:"cluster,omitempty"`
	VirtualChassis         *int64            `json:"virtual_chassis,omitempty"`
	VcPosition             *int64            `json:"vc_position,omitempty"`
	VcPriority             *int64            `json:"vc_priority,omitempty"`
	ConfigTemplate         *int64            `json:"config_template,omitempty"`
	Serial                 *string           `json:"serial,omitempty"`
	AssetTag               *string           `json:"asset_tag,omitempty"`
	PrimaryIp4             *int64            `json:"primary_ip4,omitempty"`
	PrimaryIp6             *int64            `json:"primary_ip6,omitempty"`
	OobIP                  *int64            `json:"oob_ip,omitempty"`
	LocalContextData       *string           `json:"local_context_data,omitempty"`
	Description            *string           `json:"description,omitempty"`
	Comments               *string           `json:"comments,omitempty"`
	Owner                  *int64            `json:"owner,omitempty"`
	Created                *string           `json:"created,omitempty"`
	LastUpdated            *string           `json:"last_updated,omitempty"`
	URL                    *string           `json:"url,omitempty"`
	ConsolePortCount       *int64            `json:"console_port_count,omitempty"`
	ConsoleServerPortCount *int64            `json:"console_server_port_count,omitempty"`
	PowerPortCount         *int64            `json:"power_port_count,omitempty"`
	PowerOutletCount       *int64            `json:"power_outlet_count,omitempty"`
	InterfaceCount         *int64            `json:"interface_count,omitempty"`
	FrontPortCount         *int64            `json:"front_port_count,omitempty"`
	RearPortCount          *int64            `json:"rear_port_count,omitempty"`
	DeviceBayCount         *int64            `json:"device_bay_count,omitempty"`
	ModuleBayCount         *int64            `json:"module_bay_count,omitempty"`
	InventoryItemCount     *int64            `json:"inventory_item_count,omitempty"`
	Tags                   []string          `json:"tags,omitempty"`
	TagsAll                []string          `json:"tags_all,omitempty"`
	CustomFields           map[string]string `json:"custom_fields,omitempty"`
}

// DeviceResponseDTOFromGoNetbox converts a *models.DeviceWithConfigContext to the response DTO.
func DeviceResponseDTOFromGoNetbox(goNetboxModel *models.DeviceWithConfigContext) *DeviceResponseDTO {
	responseDTO := &DeviceResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.DeviceType != nil {
		v := goNetboxModel.DeviceType.ID
		responseDTO.DeviceType = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.Site != nil {
		v := goNetboxModel.Site.ID
		responseDTO.Site = &v
	}
	if goNetboxModel.Location != nil {
		v := goNetboxModel.Location.ID
		responseDTO.Location = &v
	}
	if goNetboxModel.Rack != nil {
		v := goNetboxModel.Rack.ID
		responseDTO.Rack = &v
	}
	responseDTO.Position = goNetboxModel.Position
	if goNetboxModel.Face != nil {
		responseDTO.Face = choiceValue[string](goNetboxModel.Face.Value)
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Airflow != nil {
		responseDTO.Airflow = choiceValue[string](goNetboxModel.Airflow.Value)
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Platform != nil {
		v := goNetboxModel.Platform.ID
		responseDTO.Platform = &v
	}
	if goNetboxModel.Cluster != nil {
		v := goNetboxModel.Cluster.ID
		responseDTO.Cluster = &v
	}
	if goNetboxModel.VirtualChassis != nil {
		v := goNetboxModel.VirtualChassis.ID
		responseDTO.VirtualChassis = &v
	}
	responseDTO.VcPosition = goNetboxModel.VcPosition
	responseDTO.VcPriority = goNetboxModel.VcPriority
	if goNetboxModel.ConfigTemplate != nil {
		v := goNetboxModel.ConfigTemplate.ID
		responseDTO.ConfigTemplate = &v
	}
	if goNetboxModel.Serial != "" {
		v := goNetboxModel.Serial
		responseDTO.Serial = &v
	}
	responseDTO.AssetTag = goNetboxModel.AssetTag
	if goNetboxModel.PrimaryIp4 != nil {
		v := goNetboxModel.PrimaryIp4.ID
		responseDTO.PrimaryIp4 = &v
	}
	if goNetboxModel.PrimaryIp6 != nil {
		v := goNetboxModel.PrimaryIp6.ID
		responseDTO.PrimaryIp6 = &v
	}
	if goNetboxModel.OobIP != nil {
		v := goNetboxModel.OobIP.ID
		responseDTO.OobIP = &v
	}
	responseDTO.LocalContextData = jsonText(goNetboxModel.LocalContextData)
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Comments != "" {
		v := goNetboxModel.Comments
		responseDTO.Comments = &v
	}
	if goNetboxModel.Owner != nil {
		v := goNetboxModel.Owner.ID
		responseDTO.Owner = &v
	}
	if goNetboxModel.Created != nil {
		v := goNetboxModel.Created.String()
		responseDTO.Created = &v
	}
	if goNetboxModel.LastUpdated != nil {
		v := goNetboxModel.LastUpdated.String()
		responseDTO.LastUpdated = &v
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	{
		v := goNetboxModel.ConsolePortCount
		responseDTO.ConsolePortCount = &v
	}
	{
		v := goNetboxModel.ConsoleServerPortCount
		responseDTO.ConsoleServerPortCount = &v
	}
	{
		v := goNetboxModel.PowerPortCount
		responseDTO.PowerPortCount = &v
	}
	{
		v := goNetboxModel.PowerOutletCount
		responseDTO.PowerOutletCount = &v
	}
	{
		v := goNetboxModel.InterfaceCount
		responseDTO.InterfaceCount = &v
	}
	{
		v := goNetboxModel.FrontPortCount
		responseDTO.FrontPortCount = &v
	}
	{
		v := goNetboxModel.RearPortCount
		responseDTO.RearPortCount = &v
	}
	{
		v := goNetboxModel.DeviceBayCount
		responseDTO.DeviceBayCount = &v
	}
	{
		v := goNetboxModel.ModuleBayCount
		responseDTO.ModuleBayCount = &v
	}
	{
		v := goNetboxModel.InventoryItemCount
		responseDTO.InventoryItemCount = &v
	}
	if goNetboxModel.Tags != nil {
		responseDTO.Tags = []string{}
		for _, tag := range goNetboxModel.Tags {
			if tag != nil && tag.Slug != nil {
				responseDTO.Tags = append(responseDTO.Tags, *tag.Slug)
			}
		}
	}
	responseDTO.CustomFields = customFieldValues(goNetboxModel.CustomFields)
	return responseDTO
}
