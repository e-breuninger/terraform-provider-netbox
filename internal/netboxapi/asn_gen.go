// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// AsnRequestDTO is the request DTO of the asn resource; Payload writes it into
// the JSON request body.
type AsnRequestDTO struct {
	Asn          *int64            `json:"asn,omitempty"`
	Rir          *int64            `json:"rir,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Role         *int64            `json:"role,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the asn resource.
func (requestDTO *AsnRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Asn != nil {
		payload["asn"] = *requestDTO.Asn
	}
	if requestDTO.Rir != nil {
		payload["rir"] = *requestDTO.Rir
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	} else {
		payload["role"] = nil
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

// AsnResponseDTO is the response DTO of the asn resource, built from
// the go-netbox ASN by AsnResponseDTOFromGoNetbox.
type AsnResponseDTO struct {
	ID            *int64            `json:"id,omitempty"`
	Asn           *int64            `json:"asn,omitempty"`
	Rir           *int64            `json:"rir,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	Role          *int64            `json:"role,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Created       *string           `json:"created,omitempty"`
	LastUpdated   *string           `json:"last_updated,omitempty"`
	URL           *string           `json:"url,omitempty"`
	SiteCount     *int64            `json:"site_count,omitempty"`
	ProviderCount *int64            `json:"provider_count,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	TagsAll       []string          `json:"tags_all,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// AsnResponseDTOFromGoNetbox converts a *models.ASN to the response DTO.
func AsnResponseDTOFromGoNetbox(goNetboxModel *models.ASN) *AsnResponseDTO {
	responseDTO := &AsnResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Asn = goNetboxModel.Asn
	if goNetboxModel.Rir != nil {
		v := goNetboxModel.Rir.ID
		responseDTO.Rir = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
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
		v := goNetboxModel.SiteCount
		responseDTO.SiteCount = &v
	}
	{
		v := goNetboxModel.ProviderCount
		responseDTO.ProviderCount = &v
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
