// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ModuleBayTemplateRequestDTO is the request DTO of the module_bay_template resource; Payload writes it into
// the JSON request body.
type ModuleBayTemplateRequestDTO struct {
	DeviceType  *int64  `json:"device_type,omitempty"`
	ModuleType  *int64  `json:"module_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Position    *string `json:"position,omitempty"`
	Label       *string `json:"label,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the module_bay_template resource.
func (requestDTO *ModuleBayTemplateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.DeviceType != nil {
		payload["device_type"] = *requestDTO.DeviceType
	} else {
		payload["device_type"] = nil
	}
	if requestDTO.ModuleType != nil {
		payload["module_type"] = *requestDTO.ModuleType
	} else {
		payload["module_type"] = nil
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Position != nil {
		payload["position"] = *requestDTO.Position
	} else {
		payload["position"] = ""
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	return payload
}

// ModuleBayTemplateResponseDTO is the response DTO of the module_bay_template resource, built from
// the go-netbox ModuleBayTemplate by ModuleBayTemplateResponseDTOFromGoNetbox.
type ModuleBayTemplateResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	DeviceType  *int64  `json:"device_type,omitempty"`
	ModuleType  *int64  `json:"module_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Position    *string `json:"position,omitempty"`
	Label       *string `json:"label,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Description *string `json:"description,omitempty"`
	Created     *string `json:"created,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// ModuleBayTemplateResponseDTOFromGoNetbox converts a *models.ModuleBayTemplate to the response DTO.
func ModuleBayTemplateResponseDTOFromGoNetbox(goNetboxModel *models.ModuleBayTemplate) *ModuleBayTemplateResponseDTO {
	responseDTO := &ModuleBayTemplateResponseDTO{}
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
	if goNetboxModel.ModuleType != nil {
		v := goNetboxModel.ModuleType.ID
		responseDTO.ModuleType = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Position != "" {
		v := goNetboxModel.Position
		responseDTO.Position = &v
	}
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
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
