// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// AvailablePrefixRequestDTO is the request DTO of the available_prefix resource; Payload writes it into
// the JSON request body.
type AvailablePrefixRequestDTO struct {
	ParentPrefixID *int64            `json:"parent_prefix_id,omitempty"`
	PrefixLength   *int64            `json:"prefix_length,omitempty"`
	Status         *string           `json:"status,omitempty"`
	Role           *int64            `json:"role,omitempty"`
	Tenant         *int64            `json:"tenant,omitempty"`
	Vlan           *int64            `json:"vlan,omitempty"`
	IsPool         *bool             `json:"is_pool,omitempty"`
	MarkUtilized   *bool             `json:"mark_utilized,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Vrf            *int64            `json:"vrf,omitempty"`
	ScopeType      *string           `json:"scope_type,omitempty"`
	ScopeID        *int64            `json:"scope_id,omitempty"`
	SiteID         *int64            `json:"site_id,omitempty"`
	LocationID     *int64            `json:"location_id,omitempty"`
	RegionID       *int64            `json:"region_id,omitempty"`
	SiteGroupID    *int64            `json:"site_group_id,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the available_prefix resource.
func (requestDTO *AvailablePrefixRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	} else {
		payload["role"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Vlan != nil {
		payload["vlan"] = *requestDTO.Vlan
	} else {
		payload["vlan"] = nil
	}
	if requestDTO.IsPool != nil {
		payload["is_pool"] = *requestDTO.IsPool
	}
	if requestDTO.MarkUtilized != nil {
		payload["mark_utilized"] = *requestDTO.MarkUtilized
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Vrf != nil {
		payload["vrf"] = *requestDTO.Vrf
	}
	if requestDTO.ScopeType != nil {
		payload["scope_type"] = *requestDTO.ScopeType
	} else {
		payload["scope_type"] = nil
	}
	if requestDTO.ScopeID != nil {
		payload["scope_id"] = *requestDTO.ScopeID
	} else {
		payload["scope_id"] = nil
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

// AvailablePrefixResponseDTO is the response DTO of the available_prefix resource, built from
// the go-netbox Prefix by AvailablePrefixResponseDTOFromGoNetbox.
type AvailablePrefixResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Prefix       *string           `json:"prefix,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Role         *int64            `json:"role,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Vlan         *int64            `json:"vlan,omitempty"`
	IsPool       *bool             `json:"is_pool,omitempty"`
	MarkUtilized *bool             `json:"mark_utilized,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Vrf          *int64            `json:"vrf,omitempty"`
	ScopeType    *string           `json:"scope_type,omitempty"`
	ScopeID      *int64            `json:"scope_id,omitempty"`
	SiteID       *int64            `json:"site_id,omitempty"`
	LocationID   *int64            `json:"location_id,omitempty"`
	RegionID     *int64            `json:"region_id,omitempty"`
	SiteGroupID  *int64            `json:"site_group_id,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// AvailablePrefixResponseDTOFromGoNetbox converts a *models.Prefix to the response DTO.
func AvailablePrefixResponseDTOFromGoNetbox(goNetboxModel *models.Prefix) *AvailablePrefixResponseDTO {
	responseDTO := &AvailablePrefixResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Prefix = goNetboxModel.Prefix
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Vlan != nil {
		v := goNetboxModel.Vlan.ID
		responseDTO.Vlan = &v
	}
	{
		v := goNetboxModel.IsPool
		responseDTO.IsPool = &v
	}
	{
		v := goNetboxModel.MarkUtilized
		responseDTO.MarkUtilized = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Vrf != nil {
		v := goNetboxModel.Vrf.ID
		responseDTO.Vrf = &v
	}
	responseDTO.ScopeType = goNetboxModel.ScopeType
	responseDTO.ScopeID = goNetboxModel.ScopeID
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
