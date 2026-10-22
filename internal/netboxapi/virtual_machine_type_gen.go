// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualMachineTypeRequestDTO is the request DTO of the virtual_machine_type resource; Payload writes it into
// the JSON request body.
type VirtualMachineTypeRequestDTO struct {
	Name            *string           `json:"name,omitempty"`
	Slug            *string           `json:"slug,omitempty"`
	DefaultPlatform *int64            `json:"default_platform,omitempty"`
	DefaultVcpus    *float64          `json:"default_vcpus,omitempty"`
	DefaultMemory   *int64            `json:"default_memory,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Comments        *string           `json:"comments,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_machine_type resource.
func (requestDTO *VirtualMachineTypeRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Name != nil {
		payload["slug"] = Slugify(*requestDTO.Name)
	}
	if requestDTO.DefaultPlatform != nil {
		payload["default_platform"] = *requestDTO.DefaultPlatform
	} else {
		payload["default_platform"] = nil
	}
	if requestDTO.DefaultVcpus != nil {
		payload["default_vcpus"] = *requestDTO.DefaultVcpus
	} else {
		payload["default_vcpus"] = nil
	}
	if requestDTO.DefaultMemory != nil {
		payload["default_memory"] = *requestDTO.DefaultMemory
	} else {
		payload["default_memory"] = nil
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

// VirtualMachineTypeResponseDTO is the response DTO of the virtual_machine_type resource, built from
// the go-netbox VirtualMachineType by VirtualMachineTypeResponseDTOFromGoNetbox.
type VirtualMachineTypeResponseDTO struct {
	ID                  *int64            `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty"`
	Slug                *string           `json:"slug,omitempty"`
	DefaultPlatform     *int64            `json:"default_platform,omitempty"`
	DefaultVcpus        *float64          `json:"default_vcpus,omitempty"`
	DefaultMemory       *int64            `json:"default_memory,omitempty"`
	Description         *string           `json:"description,omitempty"`
	Comments            *string           `json:"comments,omitempty"`
	Owner               *int64            `json:"owner,omitempty"`
	Created             *string           `json:"created,omitempty"`
	LastUpdated         *string           `json:"last_updated,omitempty"`
	URL                 *string           `json:"url,omitempty"`
	VirtualMachineCount *int64            `json:"virtual_machine_count,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	TagsAll             []string          `json:"tags_all,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
}

// VirtualMachineTypeResponseDTOFromGoNetbox converts a *models.VirtualMachineType to the response DTO.
func VirtualMachineTypeResponseDTOFromGoNetbox(goNetboxModel *models.VirtualMachineType) *VirtualMachineTypeResponseDTO {
	responseDTO := &VirtualMachineTypeResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.DefaultPlatform != nil {
		v := goNetboxModel.DefaultPlatform.ID
		responseDTO.DefaultPlatform = &v
	}
	responseDTO.DefaultVcpus = goNetboxModel.DefaultVcpus
	responseDTO.DefaultMemory = goNetboxModel.DefaultMemory
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
		v := goNetboxModel.VirtualMachineCount
		responseDTO.VirtualMachineCount = &v
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
