// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// AggregateRequestDTO is the request DTO of the aggregate resource; Payload writes it into
// the JSON request body.
type AggregateRequestDTO struct {
	Prefix       *string           `json:"prefix,omitempty"`
	Rir          *int64            `json:"rir,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	DateAdded    *string           `json:"date_added,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the aggregate resource.
func (requestDTO *AggregateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Prefix != nil {
		payload["prefix"] = *requestDTO.Prefix
	}
	if requestDTO.Rir != nil {
		payload["rir"] = *requestDTO.Rir
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.DateAdded != nil {
		payload["date_added"] = *requestDTO.DateAdded
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

// AggregateResponseDTO is the response DTO of the aggregate resource, built from
// the go-netbox Aggregate by AggregateResponseDTOFromGoNetbox.
type AggregateResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Prefix       *string           `json:"prefix,omitempty"`
	Rir          *int64            `json:"rir,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	DateAdded    *string           `json:"date_added,omitempty"`
	Family       *int64            `json:"family,omitempty"`
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

// AggregateResponseDTOFromGoNetbox converts a *models.Aggregate to the response DTO.
func AggregateResponseDTOFromGoNetbox(goNetboxModel *models.Aggregate) *AggregateResponseDTO {
	responseDTO := &AggregateResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Prefix = goNetboxModel.Prefix
	if goNetboxModel.Rir != nil {
		v := goNetboxModel.Rir.ID
		responseDTO.Rir = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.DateAdded != nil {
		v := goNetboxModel.DateAdded.String()
		responseDTO.DateAdded = &v
	}
	if goNetboxModel.Family != nil {
		responseDTO.Family = choiceValue[int64](goNetboxModel.Family.Value)
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
