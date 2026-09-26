// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ConfigContextProfileRequestDTO is the request DTO of the config_context_profile resource; Payload writes it into
// the JSON request body.
type ConfigContextProfileRequestDTO struct {
	Name        *string  `json:"name,omitempty"`
	Schema      *string  `json:"schema,omitempty"`
	Description *string  `json:"description,omitempty"`
	Comments    *string  `json:"comments,omitempty"`
	Owner       *int64   `json:"owner,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// Payload returns the JSON request body of the config_context_profile resource.
func (requestDTO *ConfigContextProfileRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Schema != nil {
		payload["schema"] = parseJSONText(*requestDTO.Schema)
	} else {
		payload["schema"] = nil
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
	return payload
}

// ConfigContextProfileResponseDTO is the response DTO of the config_context_profile resource, built from
// the go-netbox ConfigContextProfile by ConfigContextProfileResponseDTOFromGoNetbox.
type ConfigContextProfileResponseDTO struct {
	ID          *int64   `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Schema      *string  `json:"schema,omitempty"`
	Description *string  `json:"description,omitempty"`
	Comments    *string  `json:"comments,omitempty"`
	Owner       *int64   `json:"owner,omitempty"`
	Created     *string  `json:"created,omitempty"`
	LastUpdated *string  `json:"last_updated,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TagsAll     []string `json:"tags_all,omitempty"`
}

// ConfigContextProfileResponseDTOFromGoNetbox converts a *models.ConfigContextProfile to the response DTO.
func ConfigContextProfileResponseDTOFromGoNetbox(goNetboxModel *models.ConfigContextProfile) *ConfigContextProfileResponseDTO {
	responseDTO := &ConfigContextProfileResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Schema = jsonText(goNetboxModel.Schema)
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
	return responseDTO
}
