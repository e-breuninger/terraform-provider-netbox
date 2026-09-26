// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// RirRequestDTO is the request DTO of the rir resource; Payload writes it into
// the JSON request body.
type RirRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Description  *string           `json:"description,omitempty"`
	IsPrivate    *bool             `json:"is_private,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the rir resource.
func (requestDTO *RirRequestDTO) Payload() map[string]any {
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
	if requestDTO.IsPrivate != nil {
		payload["is_private"] = *requestDTO.IsPrivate
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

// RirResponseDTO is the response DTO of the rir resource, built from
// the go-netbox RIR by RirResponseDTOFromGoNetbox.
type RirResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Slug           *string           `json:"slug,omitempty"`
	Description    *string           `json:"description,omitempty"`
	IsPrivate      *bool             `json:"is_private,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Created        *string           `json:"created,omitempty"`
	LastUpdated    *string           `json:"last_updated,omitempty"`
	URL            *string           `json:"url,omitempty"`
	AggregateCount *int64            `json:"aggregate_count,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	TagsAll        []string          `json:"tags_all,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// RirResponseDTOFromGoNetbox converts a *models.RIR to the response DTO.
func RirResponseDTOFromGoNetbox(goNetboxModel *models.RIR) *RirResponseDTO {
	responseDTO := &RirResponseDTO{}
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
	{
		v := goNetboxModel.IsPrivate
		responseDTO.IsPrivate = &v
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
		v := goNetboxModel.AggregateCount
		responseDTO.AggregateCount = &v
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
