// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DeviceBayRequestDTO is the request DTO of the device_bay resource; Payload writes it into
// the JSON request body.
type DeviceBayRequestDTO struct {
	Device          *int64            `json:"device,omitempty"`
	Name            *string           `json:"name,omitempty"`
	Label           *string           `json:"label,omitempty"`
	InstalledDevice *int64            `json:"installed_device,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the device_bay resource.
func (requestDTO *DeviceBayRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.InstalledDevice != nil {
		payload["installed_device"] = *requestDTO.InstalledDevice
	} else {
		payload["installed_device"] = nil
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

// DeviceBayResponseDTO is the response DTO of the device_bay resource, built from
// the go-netbox DeviceBay by DeviceBayResponseDTOFromGoNetbox.
type DeviceBayResponseDTO struct {
	ID              *int64            `json:"id,omitempty"`
	Device          *int64            `json:"device,omitempty"`
	Name            *string           `json:"name,omitempty"`
	Label           *string           `json:"label,omitempty"`
	InstalledDevice *int64            `json:"installed_device,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Created         *string           `json:"created,omitempty"`
	LastUpdated     *string           `json:"last_updated,omitempty"`
	URL             *string           `json:"url,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	TagsAll         []string          `json:"tags_all,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// DeviceBayResponseDTOFromGoNetbox converts a *models.DeviceBay to the response DTO.
func DeviceBayResponseDTOFromGoNetbox(goNetboxModel *models.DeviceBay) *DeviceBayResponseDTO {
	responseDTO := &DeviceBayResponseDTO{}
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
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	if goNetboxModel.InstalledDevice != nil {
		v := goNetboxModel.InstalledDevice.ID
		responseDTO.InstalledDevice = &v
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
