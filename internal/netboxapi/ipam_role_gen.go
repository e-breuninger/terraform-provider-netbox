// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamRoleRequestDTO is the request DTO of the ipam_role resource; Payload writes it into
// the JSON request body.
type IpamRoleRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Weight       *int64            `json:"weight,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ipam_role resource.
func (requestDTO *IpamRoleRequestDTO) Payload() map[string]any {
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
	if requestDTO.Weight != nil {
		payload["weight"] = *requestDTO.Weight
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

// IpamRoleResponseDTO is the response DTO of the ipam_role resource, built from
// the go-netbox Role by IpamRoleResponseDTOFromGoNetbox.
type IpamRoleResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Weight       *int64            `json:"weight,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	PrefixCount  *int64            `json:"prefix_count,omitempty"`
	VlanCount    *int64            `json:"vlan_count,omitempty"`
	AsnCount     *int64            `json:"asn_count,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// IpamRoleResponseDTOFromGoNetbox converts a *models.Role to the response DTO.
func IpamRoleResponseDTOFromGoNetbox(goNetboxModel *models.Role) *IpamRoleResponseDTO {
	responseDTO := &IpamRoleResponseDTO{}
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
	responseDTO.Weight = goNetboxModel.Weight
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
		v := goNetboxModel.PrefixCount
		responseDTO.PrefixCount = &v
	}
	{
		v := goNetboxModel.VlanCount
		responseDTO.VlanCount = &v
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
