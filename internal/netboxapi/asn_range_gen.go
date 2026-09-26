// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// AsnRangeRequestDTO is the request DTO of the asn_range resource; Payload writes it into
// the JSON request body.
type AsnRangeRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Rir          *int64            `json:"rir,omitempty"`
	Start        *int64            `json:"start,omitempty"`
	End          *int64            `json:"end,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the asn_range resource.
func (requestDTO *AsnRangeRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Name != nil {
		payload["slug"] = Slugify(*requestDTO.Name)
	}
	if requestDTO.Rir != nil {
		payload["rir"] = *requestDTO.Rir
	}
	if requestDTO.Start != nil {
		payload["start"] = *requestDTO.Start
	}
	if requestDTO.End != nil {
		payload["end"] = *requestDTO.End
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
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

// AsnRangeResponseDTO is the response DTO of the asn_range resource, built from
// the go-netbox ASNRange by AsnRangeResponseDTOFromGoNetbox.
type AsnRangeResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Rir          *int64            `json:"rir,omitempty"`
	Start        *int64            `json:"start,omitempty"`
	End          *int64            `json:"end,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	AsnCount     *int64            `json:"asn_count,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// AsnRangeResponseDTOFromGoNetbox converts a *models.ASNRange to the response DTO.
func AsnRangeResponseDTOFromGoNetbox(goNetboxModel *models.ASNRange) *AsnRangeResponseDTO {
	responseDTO := &AsnRangeResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.Rir != nil {
		v := goNetboxModel.Rir.ID
		responseDTO.Rir = &v
	}
	responseDTO.Start = goNetboxModel.Start
	responseDTO.End = goNetboxModel.End
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
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
	{
		v := goNetboxModel.AsnCount
		responseDTO.AsnCount = &v
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
