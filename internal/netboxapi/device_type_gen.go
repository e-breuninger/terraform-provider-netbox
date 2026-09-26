// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DeviceTypeRequestDTO is the request DTO of the device_type resource; Payload writes it into
// the JSON request body.
type DeviceTypeRequestDTO struct {
	Manufacturer  *int64            `json:"manufacturer,omitempty"`
	Model         *string           `json:"model,omitempty"`
	Slug          *string           `json:"slug,omitempty"`
	PartNumber    *string           `json:"part_number,omitempty"`
	UHeight       *float64          `json:"u_height,omitempty"`
	IsFullDepth   *bool             `json:"is_full_depth,omitempty"`
	SubdeviceRole *string           `json:"subdevice_role,omitempty"`
	Airflow       *string           `json:"airflow,omitempty"`
	Weight        *float64          `json:"weight,omitempty"`
	WeightUnit    *string           `json:"weight_unit,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the device_type resource.
func (requestDTO *DeviceTypeRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Manufacturer != nil {
		payload["manufacturer"] = *requestDTO.Manufacturer
	}
	if requestDTO.Model != nil {
		payload["model"] = *requestDTO.Model
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Model != nil {
		payload["slug"] = Slugify(*requestDTO.Model)
	}
	if requestDTO.PartNumber != nil {
		payload["part_number"] = *requestDTO.PartNumber
	} else {
		payload["part_number"] = ""
	}
	if requestDTO.UHeight != nil {
		payload["u_height"] = *requestDTO.UHeight
	}
	if requestDTO.IsFullDepth != nil {
		payload["is_full_depth"] = *requestDTO.IsFullDepth
	}
	if requestDTO.SubdeviceRole != nil {
		payload["subdevice_role"] = *requestDTO.SubdeviceRole
	} else {
		payload["subdevice_role"] = nil
	}
	if requestDTO.Airflow != nil {
		payload["airflow"] = *requestDTO.Airflow
	} else {
		payload["airflow"] = nil
	}
	if requestDTO.Weight != nil {
		payload["weight"] = *requestDTO.Weight
	} else {
		payload["weight"] = nil
	}
	if requestDTO.WeightUnit != nil {
		payload["weight_unit"] = *requestDTO.WeightUnit
	} else {
		payload["weight_unit"] = nil
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

// DeviceTypeResponseDTO is the response DTO of the device_type resource, built from
// the go-netbox DeviceType by DeviceTypeResponseDTOFromGoNetbox.
type DeviceTypeResponseDTO struct {
	ID                             *int64            `json:"id,omitempty"`
	Manufacturer                   *int64            `json:"manufacturer,omitempty"`
	Model                          *string           `json:"model,omitempty"`
	Slug                           *string           `json:"slug,omitempty"`
	PartNumber                     *string           `json:"part_number,omitempty"`
	UHeight                        *float64          `json:"u_height,omitempty"`
	IsFullDepth                    *bool             `json:"is_full_depth,omitempty"`
	SubdeviceRole                  *string           `json:"subdevice_role,omitempty"`
	Airflow                        *string           `json:"airflow,omitempty"`
	Weight                         *float64          `json:"weight,omitempty"`
	WeightUnit                     *string           `json:"weight_unit,omitempty"`
	Description                    *string           `json:"description,omitempty"`
	Comments                       *string           `json:"comments,omitempty"`
	Owner                          *int64            `json:"owner,omitempty"`
	Created                        *string           `json:"created,omitempty"`
	LastUpdated                    *string           `json:"last_updated,omitempty"`
	URL                            *string           `json:"url,omitempty"`
	DeviceCount                    *int64            `json:"device_count,omitempty"`
	ConsolePortTemplateCount       *int64            `json:"console_port_template_count,omitempty"`
	ConsoleServerPortTemplateCount *int64            `json:"console_server_port_template_count,omitempty"`
	PowerPortTemplateCount         *int64            `json:"power_port_template_count,omitempty"`
	PowerOutletTemplateCount       *int64            `json:"power_outlet_template_count,omitempty"`
	InterfaceTemplateCount         *int64            `json:"interface_template_count,omitempty"`
	FrontPortTemplateCount         *int64            `json:"front_port_template_count,omitempty"`
	RearPortTemplateCount          *int64            `json:"rear_port_template_count,omitempty"`
	DeviceBayTemplateCount         *int64            `json:"device_bay_template_count,omitempty"`
	ModuleBayTemplateCount         *int64            `json:"module_bay_template_count,omitempty"`
	InventoryItemTemplateCount     *int64            `json:"inventory_item_template_count,omitempty"`
	Tags                           []string          `json:"tags,omitempty"`
	TagsAll                        []string          `json:"tags_all,omitempty"`
	CustomFields                   map[string]string `json:"custom_fields,omitempty"`
}

// DeviceTypeResponseDTOFromGoNetbox converts a *models.DeviceType to the response DTO.
func DeviceTypeResponseDTOFromGoNetbox(goNetboxModel *models.DeviceType) *DeviceTypeResponseDTO {
	responseDTO := &DeviceTypeResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Manufacturer != nil {
		v := goNetboxModel.Manufacturer.ID
		responseDTO.Manufacturer = &v
	}
	responseDTO.Model = goNetboxModel.Model
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.PartNumber != "" {
		v := goNetboxModel.PartNumber
		responseDTO.PartNumber = &v
	}
	responseDTO.UHeight = goNetboxModel.UHeight
	{
		v := goNetboxModel.IsFullDepth
		responseDTO.IsFullDepth = &v
	}
	if goNetboxModel.SubdeviceRole != nil {
		responseDTO.SubdeviceRole = choiceValue[string](goNetboxModel.SubdeviceRole.Value)
	}
	if goNetboxModel.Airflow != nil {
		responseDTO.Airflow = choiceValue[string](goNetboxModel.Airflow.Value)
	}
	responseDTO.Weight = goNetboxModel.Weight
	if goNetboxModel.WeightUnit != nil {
		responseDTO.WeightUnit = choiceValue[string](goNetboxModel.WeightUnit.Value)
	}
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
		v := goNetboxModel.DeviceCount
		responseDTO.DeviceCount = &v
	}
	{
		v := goNetboxModel.ConsolePortTemplateCount
		responseDTO.ConsolePortTemplateCount = &v
	}
	{
		v := goNetboxModel.ConsoleServerPortTemplateCount
		responseDTO.ConsoleServerPortTemplateCount = &v
	}
	{
		v := goNetboxModel.PowerPortTemplateCount
		responseDTO.PowerPortTemplateCount = &v
	}
	{
		v := goNetboxModel.PowerOutletTemplateCount
		responseDTO.PowerOutletTemplateCount = &v
	}
	{
		v := goNetboxModel.InterfaceTemplateCount
		responseDTO.InterfaceTemplateCount = &v
	}
	{
		v := goNetboxModel.FrontPortTemplateCount
		responseDTO.FrontPortTemplateCount = &v
	}
	{
		v := goNetboxModel.RearPortTemplateCount
		responseDTO.RearPortTemplateCount = &v
	}
	{
		v := goNetboxModel.DeviceBayTemplateCount
		responseDTO.DeviceBayTemplateCount = &v
	}
	{
		v := goNetboxModel.ModuleBayTemplateCount
		responseDTO.ModuleBayTemplateCount = &v
	}
	{
		v := goNetboxModel.InventoryItemTemplateCount
		responseDTO.InventoryItemTemplateCount = &v
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
