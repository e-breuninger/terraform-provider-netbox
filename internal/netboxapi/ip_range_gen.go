// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IPRangeRequestDTO is the request DTO of the ip_range resource; Payload writes it into
// the JSON request body.
type IPRangeRequestDTO struct {
	StartAddress  *string           `json:"start_address,omitempty"`
	EndAddress    *string           `json:"end_address,omitempty"`
	Vrf           *int64            `json:"vrf,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	Role          *int64            `json:"role,omitempty"`
	Status        *string           `json:"status,omitempty"`
	MarkUtilized  *bool             `json:"mark_utilized,omitempty"`
	MarkPopulated *bool             `json:"mark_populated,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ip_range resource.
func (requestDTO *IPRangeRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.StartAddress != nil {
		payload["start_address"] = *requestDTO.StartAddress
	}
	if requestDTO.EndAddress != nil {
		payload["end_address"] = *requestDTO.EndAddress
	}
	if requestDTO.Vrf != nil {
		payload["vrf"] = *requestDTO.Vrf
	} else {
		payload["vrf"] = nil
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
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.MarkUtilized != nil {
		payload["mark_utilized"] = *requestDTO.MarkUtilized
	}
	if requestDTO.MarkPopulated != nil {
		payload["mark_populated"] = *requestDTO.MarkPopulated
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

// IPRangeResponseDTO is the response DTO of the ip_range resource, built from
// the go-netbox IPRange by IPRangeResponseDTOFromGoNetbox.
type IPRangeResponseDTO struct {
	ID            *int64            `json:"id,omitempty"`
	StartAddress  *string           `json:"start_address,omitempty"`
	EndAddress    *string           `json:"end_address,omitempty"`
	Vrf           *int64            `json:"vrf,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	Role          *int64            `json:"role,omitempty"`
	Status        *string           `json:"status,omitempty"`
	MarkUtilized  *bool             `json:"mark_utilized,omitempty"`
	MarkPopulated *bool             `json:"mark_populated,omitempty"`
	Size          *int64            `json:"size,omitempty"`
	Family        *int64            `json:"family,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Created       *string           `json:"created,omitempty"`
	LastUpdated   *string           `json:"last_updated,omitempty"`
	URL           *string           `json:"url,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	TagsAll       []string          `json:"tags_all,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// IPRangeResponseDTOFromGoNetbox converts a *models.IPRange to the response DTO.
func IPRangeResponseDTOFromGoNetbox(goNetboxModel *models.IPRange) *IPRangeResponseDTO {
	responseDTO := &IPRangeResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.StartAddress = goNetboxModel.StartAddress
	responseDTO.EndAddress = goNetboxModel.EndAddress
	if goNetboxModel.Vrf != nil {
		v := goNetboxModel.Vrf.ID
		responseDTO.Vrf = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	{
		v := goNetboxModel.MarkUtilized
		responseDTO.MarkUtilized = &v
	}
	{
		v := goNetboxModel.MarkPopulated
		responseDTO.MarkPopulated = &v
	}
	{
		v := goNetboxModel.Size
		responseDTO.Size = &v
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
