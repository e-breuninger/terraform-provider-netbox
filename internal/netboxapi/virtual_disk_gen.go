// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualDiskRequestDTO is the request DTO of the virtual_disk resource; Payload writes it into
// the JSON request body.
type VirtualDiskRequestDTO struct {
	VirtualMachine *int64            `json:"virtual_machine,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Size           *int64            `json:"size,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_disk resource.
func (requestDTO *VirtualDiskRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.VirtualMachine != nil {
		payload["virtual_machine"] = *requestDTO.VirtualMachine
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Size != nil {
		payload["size"] = *requestDTO.Size
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

// VirtualDiskResponseDTO is the response DTO of the virtual_disk resource, built from
// the go-netbox VirtualDisk by VirtualDiskResponseDTOFromGoNetbox.
type VirtualDiskResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	VirtualMachine *int64            `json:"virtual_machine,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Size           *int64            `json:"size,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Created        *string           `json:"created,omitempty"`
	LastUpdated    *string           `json:"last_updated,omitempty"`
	URL            *string           `json:"url,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	TagsAll        []string          `json:"tags_all,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// VirtualDiskResponseDTOFromGoNetbox converts a *models.VirtualDisk to the response DTO.
func VirtualDiskResponseDTOFromGoNetbox(goNetboxModel *models.VirtualDisk) *VirtualDiskResponseDTO {
	responseDTO := &VirtualDiskResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.VirtualMachine != nil {
		v := goNetboxModel.VirtualMachine.ID
		responseDTO.VirtualMachine = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Size = goNetboxModel.Size
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
