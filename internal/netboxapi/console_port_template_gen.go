// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ConsolePortTemplateRequestDTO is the request DTO of the console_port_template resource; Payload writes it into
// the JSON request body.
type ConsolePortTemplateRequestDTO struct {
	DeviceType  *int64  `json:"device_type,omitempty"`
	ModuleType  *int64  `json:"module_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the console_port_template resource.
func (requestDTO *ConsolePortTemplateRequestDTO) Payload() map[string]any {
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
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	} else {
		payload["type"] = nil
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

// ConsolePortTemplateResponseDTO is the response DTO of the console_port_template resource, built from
// the go-netbox ConsolePortTemplate by ConsolePortTemplateResponseDTOFromGoNetbox.
type ConsolePortTemplateResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	DeviceType  *int64  `json:"device_type,omitempty"`
	ModuleType  *int64  `json:"module_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Created     *string `json:"created,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// ConsolePortTemplateResponseDTOFromGoNetbox converts a *models.ConsolePortTemplate to the response DTO.
func ConsolePortTemplateResponseDTOFromGoNetbox(goNetboxModel *models.ConsolePortTemplate) *ConsolePortTemplateResponseDTO {
	responseDTO := &ConsolePortTemplateResponseDTO{}
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
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
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
