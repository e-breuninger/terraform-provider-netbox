// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DeviceBayTemplateRequestDTO is the request DTO of the device_bay_template resource; Payload writes it into
// the JSON request body.
type DeviceBayTemplateRequestDTO struct {
	DeviceType  *int64  `json:"device_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the device_bay_template resource.
func (requestDTO *DeviceBayTemplateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.DeviceType != nil {
		payload["device_type"] = *requestDTO.DeviceType
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	return payload
}

// DeviceBayTemplateResponseDTO is the response DTO of the device_bay_template resource, built from
// the go-netbox DeviceBayTemplate by DeviceBayTemplateResponseDTOFromGoNetbox.
type DeviceBayTemplateResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	DeviceType  *int64  `json:"device_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Created     *string `json:"created,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// DeviceBayTemplateResponseDTOFromGoNetbox converts a *models.DeviceBayTemplate to the response DTO.
func DeviceBayTemplateResponseDTOFromGoNetbox(goNetboxModel *models.DeviceBayTemplate) *DeviceBayTemplateResponseDTO {
	responseDTO := &DeviceBayTemplateResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.DeviceType != nil {
		v := goNetboxModel.DeviceType.ID
		responseDTO.DeviceType = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
	return responseDTO
}
