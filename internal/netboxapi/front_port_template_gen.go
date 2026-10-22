// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// FrontPortTemplateRequestDTO is the request DTO of the front_port_template resource; Payload writes it into
// the JSON request body.
type FrontPortTemplateRequestDTO struct {
	DeviceType       *int64                                  `json:"device_type,omitempty"`
	ModuleType       *int64                                  `json:"module_type,omitempty"`
	Name             *string                                 `json:"name,omitempty"`
	Type             *string                                 `json:"type,omitempty"`
	Positions        *int64                                  `json:"positions,omitempty"`
	RearPorts        []*FrontPortTemplateRequestDTORearPorts `json:"rear_ports,omitempty"`
	RearPortID       *int64                                  `json:"rear_port_id,omitempty"`
	RearPortPosition *int64                                  `json:"rear_port_position,omitempty"`
	Color            *string                                 `json:"color,omitempty"`
	Label            *string                                 `json:"label,omitempty"`
	Description      *string                                 `json:"description,omitempty"`
}

// FrontPortTemplateRequestDTORearPorts is the request DTO of the rear_ports object; it is written into the
// request body as JSON.
type FrontPortTemplateRequestDTORearPorts struct {
	Position         *int64 `json:"position,omitempty"`
	RearPort         *int64 `json:"rear_port,omitempty"`
	RearPortPosition *int64 `json:"rear_port_position,omitempty"`
}

// Payload returns the JSON request body of the front_port_template resource.
func (requestDTO *FrontPortTemplateRequestDTO) Payload() map[string]any {
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
	}
	if requestDTO.Positions != nil {
		payload["positions"] = *requestDTO.Positions
	}
	if requestDTO.RearPorts != nil {
		payload["rear_ports"] = requestDTO.RearPorts
	}
	if requestDTO.Color != nil {
		payload["color"] = *requestDTO.Color
	} else {
		payload["color"] = ""
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

// FrontPortTemplateResponseDTO is the response DTO of the front_port_template resource, built from
// the go-netbox FrontPortTemplate by FrontPortTemplateResponseDTOFromGoNetbox.
type FrontPortTemplateResponseDTO struct {
	ID               *int64                                   `json:"id,omitempty"`
	DeviceType       *int64                                   `json:"device_type,omitempty"`
	ModuleType       *int64                                   `json:"module_type,omitempty"`
	Name             *string                                  `json:"name,omitempty"`
	Type             *string                                  `json:"type,omitempty"`
	Positions        *int64                                   `json:"positions,omitempty"`
	RearPorts        []*FrontPortTemplateResponseDTORearPorts `json:"rear_ports,omitempty"`
	RearPortID       *int64                                   `json:"rear_port_id,omitempty"`
	RearPortPosition *int64                                   `json:"rear_port_position,omitempty"`
	Color            *string                                  `json:"color,omitempty"`
	Label            *string                                  `json:"label,omitempty"`
	Description      *string                                  `json:"description,omitempty"`
	Created          *string                                  `json:"created,omitempty"`
	LastUpdated      *string                                  `json:"last_updated,omitempty"`
	URL              *string                                  `json:"url,omitempty"`
}

// FrontPortTemplateResponseDTORearPorts is the response DTO of the rear_ports object, decoded from the
// go-netbox value as JSON.
type FrontPortTemplateResponseDTORearPorts struct {
	Position         *int64 `json:"position,omitempty"`
	RearPort         *int64 `json:"rear_port,omitempty"`
	RearPortPosition *int64 `json:"rear_port_position,omitempty"`
}

// FrontPortTemplateResponseDTOFromGoNetbox converts a *models.FrontPortTemplate to the response DTO.
func FrontPortTemplateResponseDTOFromGoNetbox(goNetboxModel *models.FrontPortTemplate) *FrontPortTemplateResponseDTO {
	responseDTO := &FrontPortTemplateResponseDTO{}
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
	{
		v := goNetboxModel.Positions
		responseDTO.Positions = &v
	}
	if goNetboxModel.RearPorts != nil {
		var v []*FrontPortTemplateResponseDTORearPorts
		if err := reencode(goNetboxModel.RearPorts, &v); err == nil {
			responseDTO.RearPorts = v
		}
	}
	if goNetboxModel.Color != "" {
		v := goNetboxModel.Color
		responseDTO.Color = &v
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
