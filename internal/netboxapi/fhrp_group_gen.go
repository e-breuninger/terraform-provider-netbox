// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// FhrpGroupRequestDTO is the request DTO of the fhrp_group resource; Payload writes it into
// the JSON request body.
type FhrpGroupRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Protocol     *string           `json:"protocol,omitempty"`
	GroupID      *int64            `json:"group_id,omitempty"`
	AuthType     *string           `json:"auth_type,omitempty"`
	AuthKey      *string           `json:"auth_key,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the fhrp_group resource.
func (requestDTO *FhrpGroupRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	} else {
		payload["name"] = ""
	}
	if requestDTO.Protocol != nil {
		payload["protocol"] = *requestDTO.Protocol
	}
	if requestDTO.GroupID != nil {
		payload["group_id"] = *requestDTO.GroupID
	}
	if requestDTO.AuthType != nil {
		payload["auth_type"] = *requestDTO.AuthType
	} else {
		payload["auth_type"] = nil
	}
	if requestDTO.AuthKey != nil {
		payload["auth_key"] = *requestDTO.AuthKey
	} else {
		payload["auth_key"] = ""
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

// FhrpGroupResponseDTO is the response DTO of the fhrp_group resource, built from
// the go-netbox FHRPGroup by FhrpGroupResponseDTOFromGoNetbox.
type FhrpGroupResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Protocol     *string           `json:"protocol,omitempty"`
	GroupID      *int64            `json:"group_id,omitempty"`
	AuthType     *string           `json:"auth_type,omitempty"`
	AuthKey      *string           `json:"auth_key,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// FhrpGroupResponseDTOFromGoNetbox converts a *models.FHRPGroup to the response DTO.
func FhrpGroupResponseDTOFromGoNetbox(goNetboxModel *models.FHRPGroup) *FhrpGroupResponseDTO {
	responseDTO := &FhrpGroupResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Name != "" {
		v := goNetboxModel.Name
		responseDTO.Name = &v
	}
	responseDTO.Protocol = goNetboxModel.Protocol
	responseDTO.GroupID = goNetboxModel.GroupID
	responseDTO.AuthType = goNetboxModel.AuthType
	if goNetboxModel.AuthKey != "" {
		v := goNetboxModel.AuthKey
		responseDTO.AuthKey = &v
	}
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
