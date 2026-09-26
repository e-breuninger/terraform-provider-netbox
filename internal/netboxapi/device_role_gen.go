// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DeviceRoleRequestDTO is the request DTO of the device_role resource; Payload writes it into
// the JSON request body.
type DeviceRoleRequestDTO struct {
	Name           *string           `json:"name,omitempty"`
	Slug           *string           `json:"slug,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Color          *string           `json:"color,omitempty"`
	VMRole         *bool             `json:"vm_role,omitempty"`
	Parent         *int64            `json:"parent,omitempty"`
	ConfigTemplate *int64            `json:"config_template,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the device_role resource.
func (requestDTO *DeviceRoleRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Name != nil {
		payload["slug"] = Slugify(*requestDTO.Name)
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Color != nil {
		payload["color"] = *requestDTO.Color
	} else {
		payload["color"] = ""
	}
	if requestDTO.VMRole != nil {
		payload["vm_role"] = *requestDTO.VMRole
	}
	if requestDTO.Parent != nil {
		payload["parent"] = *requestDTO.Parent
	} else {
		payload["parent"] = nil
	}
	if requestDTO.ConfigTemplate != nil {
		payload["config_template"] = *requestDTO.ConfigTemplate
	} else {
		payload["config_template"] = nil
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

// DeviceRoleResponseDTO is the response DTO of the device_role resource, built from
// the go-netbox DeviceRole by DeviceRoleResponseDTOFromGoNetbox.
type DeviceRoleResponseDTO struct {
	ID                  *int64            `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty"`
	Slug                *string           `json:"slug,omitempty"`
	Description         *string           `json:"description,omitempty"`
	Color               *string           `json:"color,omitempty"`
	VMRole              *bool             `json:"vm_role,omitempty"`
	Parent              *int64            `json:"parent,omitempty"`
	ConfigTemplate      *int64            `json:"config_template,omitempty"`
	Owner               *int64            `json:"owner,omitempty"`
	Created             *string           `json:"created,omitempty"`
	LastUpdated         *string           `json:"last_updated,omitempty"`
	URL                 *string           `json:"url,omitempty"`
	DeviceCount         *int64            `json:"device_count,omitempty"`
	VirtualmachineCount *int64            `json:"virtualmachine_count,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	TagsAll             []string          `json:"tags_all,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
}

// DeviceRoleResponseDTOFromGoNetbox converts a *models.DeviceRole to the response DTO.
func DeviceRoleResponseDTOFromGoNetbox(goNetboxModel *models.DeviceRole) *DeviceRoleResponseDTO {
	responseDTO := &DeviceRoleResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Color != "" {
		v := goNetboxModel.Color
		responseDTO.Color = &v
	}
	{
		v := goNetboxModel.VMRole
		responseDTO.VMRole = &v
	}
	if goNetboxModel.Parent != nil {
		v := goNetboxModel.Parent.ID
		responseDTO.Parent = &v
	}
	if goNetboxModel.ConfigTemplate != nil {
		v := goNetboxModel.ConfigTemplate.ID
		responseDTO.ConfigTemplate = &v
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
		v := goNetboxModel.VirtualmachineCount
		responseDTO.VirtualmachineCount = &v
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
