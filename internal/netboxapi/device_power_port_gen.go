// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DevicePowerPortRequestDTO is the request DTO of the device_power_port resource; Payload writes it into
// the JSON request body.
type DevicePowerPortRequestDTO struct {
	Device        *int64            `json:"device,omitempty"`
	Module        *int64            `json:"module,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Type          *string           `json:"type,omitempty"`
	MaximumDraw   *int64            `json:"maximum_draw,omitempty"`
	AllocatedDraw *int64            `json:"allocated_draw,omitempty"`
	Label         *string           `json:"label,omitempty"`
	MarkConnected *bool             `json:"mark_connected,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the device_power_port resource.
func (requestDTO *DevicePowerPortRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	}
	if requestDTO.Module != nil {
		payload["module"] = *requestDTO.Module
	} else {
		payload["module"] = nil
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
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
	if requestDTO.MarkConnected != nil {
		payload["mark_connected"] = *requestDTO.MarkConnected
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
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

// DevicePowerPortResponseDTO is the response DTO of the device_power_port resource, built from
// the go-netbox PowerPort by DevicePowerPortResponseDTOFromGoNetbox.
type DevicePowerPortResponseDTO struct {
	ID            *int64            `json:"id,omitempty"`
	Device        *int64            `json:"device,omitempty"`
	Module        *int64            `json:"module,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Type          *string           `json:"type,omitempty"`
	MaximumDraw   *int64            `json:"maximum_draw,omitempty"`
	AllocatedDraw *int64            `json:"allocated_draw,omitempty"`
	Label         *string           `json:"label,omitempty"`
	MarkConnected *bool             `json:"mark_connected,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Created       *string           `json:"created,omitempty"`
	LastUpdated   *string           `json:"last_updated,omitempty"`
	URL           *string           `json:"url,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	TagsAll       []string          `json:"tags_all,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// DevicePowerPortResponseDTOFromGoNetbox converts a *models.PowerPort to the response DTO.
func DevicePowerPortResponseDTOFromGoNetbox(goNetboxModel *models.PowerPort) *DevicePowerPortResponseDTO {
	responseDTO := &DevicePowerPortResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Device != nil {
		v := goNetboxModel.Device.ID
		responseDTO.Device = &v
	}
	if goNetboxModel.Module != nil {
		v := goNetboxModel.Module.ID
		responseDTO.Module = &v
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
	{
		v := goNetboxModel.MarkConnected
		responseDTO.MarkConnected = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
