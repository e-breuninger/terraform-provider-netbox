// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// PowerPortTemplateRequestDTO is the request DTO of the power_port_template resource; Payload writes it into
// the JSON request body.
type PowerPortTemplateRequestDTO struct {
	DeviceType    *int64  `json:"device_type,omitempty"`
	ModuleType    *int64  `json:"module_type,omitempty"`
	Name          *string `json:"name,omitempty"`
	Type          *string `json:"type,omitempty"`
	MaximumDraw   *int64  `json:"maximum_draw,omitempty"`
	AllocatedDraw *int64  `json:"allocated_draw,omitempty"`
	Label         *string `json:"label,omitempty"`
	Description   *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the power_port_template resource.
func (requestDTO *PowerPortTemplateRequestDTO) Payload() map[string]any {
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
	if requestDTO.MaximumDraw != nil {
		payload["maximum_draw"] = *requestDTO.MaximumDraw
	} else {
		payload["maximum_draw"] = nil
	}
	if requestDTO.AllocatedDraw != nil {
		payload["allocated_draw"] = *requestDTO.AllocatedDraw
	} else {
		payload["allocated_draw"] = nil
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

// PowerPortTemplateResponseDTO is the response DTO of the power_port_template resource, built from
// the go-netbox PowerPortTemplate by PowerPortTemplateResponseDTOFromGoNetbox.
type PowerPortTemplateResponseDTO struct {
	ID            *int64  `json:"id,omitempty"`
	DeviceType    *int64  `json:"device_type,omitempty"`
	ModuleType    *int64  `json:"module_type,omitempty"`
	Name          *string `json:"name,omitempty"`
	Type          *string `json:"type,omitempty"`
	MaximumDraw   *int64  `json:"maximum_draw,omitempty"`
	AllocatedDraw *int64  `json:"allocated_draw,omitempty"`
	Label         *string `json:"label,omitempty"`
	Description   *string `json:"description,omitempty"`
	Created       *string `json:"created,omitempty"`
	LastUpdated   *string `json:"last_updated,omitempty"`
	URL           *string `json:"url,omitempty"`
}

// PowerPortTemplateResponseDTOFromGoNetbox converts a *models.PowerPortTemplate to the response DTO.
func PowerPortTemplateResponseDTOFromGoNetbox(goNetboxModel *models.PowerPortTemplate) *PowerPortTemplateResponseDTO {
	responseDTO := &PowerPortTemplateResponseDTO{}
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
	responseDTO.MaximumDraw = goNetboxModel.MaximumDraw
	responseDTO.AllocatedDraw = goNetboxModel.AllocatedDraw
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
