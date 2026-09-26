// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualCircuitTypeRequestDTO is the request DTO of the virtual_circuit_type resource; Payload writes it into
// the JSON request body.
type VirtualCircuitTypeRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_circuit_type resource.
func (requestDTO *VirtualCircuitTypeRequestDTO) Payload() map[string]any {
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

// VirtualCircuitTypeResponseDTO is the response DTO of the virtual_circuit_type resource, built from
// the go-netbox VirtualCircuitType by VirtualCircuitTypeResponseDTOFromGoNetbox.
type VirtualCircuitTypeResponseDTO struct {
	ID                  *int64            `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty"`
	Slug                *string           `json:"slug,omitempty"`
	Description         *string           `json:"description,omitempty"`
	Comments            *string           `json:"comments,omitempty"`
	Owner               *int64            `json:"owner,omitempty"`
	Created             *string           `json:"created,omitempty"`
	LastUpdated         *string           `json:"last_updated,omitempty"`
	URL                 *string           `json:"url,omitempty"`
	VirtualCircuitCount *int64            `json:"virtual_circuit_count,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	TagsAll             []string          `json:"tags_all,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
}

// VirtualCircuitTypeResponseDTOFromGoNetbox converts a *models.VirtualCircuitType to the response DTO.
func VirtualCircuitTypeResponseDTOFromGoNetbox(goNetboxModel *models.VirtualCircuitType) *VirtualCircuitTypeResponseDTO {
	responseDTO := &VirtualCircuitTypeResponseDTO{}
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
		v := goNetboxModel.VirtualCircuitCount
		responseDTO.VirtualCircuitCount = &v
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
