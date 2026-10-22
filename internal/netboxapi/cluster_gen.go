// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ClusterRequestDTO is the request DTO of the cluster resource; Payload writes it into
// the JSON request body.
type ClusterRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Type         *int64            `json:"type,omitempty"`
	Group        *int64            `json:"group,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Description  *string           `json:"description,omitempty"`
	ScopeType    *string           `json:"scope_type,omitempty"`
	ScopeID      *int64            `json:"scope_id,omitempty"`
	SiteID       *int64            `json:"site_id,omitempty"`
	LocationID   *int64            `json:"location_id,omitempty"`
	RegionID     *int64            `json:"region_id,omitempty"`
	SiteGroupID  *int64            `json:"site_group_id,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the cluster resource.
func (requestDTO *ClusterRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	} else {
		payload["group"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
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

// ClusterResponseDTO is the response DTO of the cluster resource, built from
// the go-netbox Cluster by ClusterResponseDTOFromGoNetbox.
type ClusterResponseDTO struct {
	ID                  *int64            `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty"`
	Type                *int64            `json:"type,omitempty"`
	Group               *int64            `json:"group,omitempty"`
	Tenant              *int64            `json:"tenant,omitempty"`
	Status              *string           `json:"status,omitempty"`
	Description         *string           `json:"description,omitempty"`
	ScopeType           *string           `json:"scope_type,omitempty"`
	ScopeID             *int64            `json:"scope_id,omitempty"`
	SiteID              *int64            `json:"site_id,omitempty"`
	LocationID          *int64            `json:"location_id,omitempty"`
	RegionID            *int64            `json:"region_id,omitempty"`
	SiteGroupID         *int64            `json:"site_group_id,omitempty"`
	Comments            *string           `json:"comments,omitempty"`
	Owner               *int64            `json:"owner,omitempty"`
	Created             *string           `json:"created,omitempty"`
	LastUpdated         *string           `json:"last_updated,omitempty"`
	URL                 *string           `json:"url,omitempty"`
	DeviceCount         *int64            `json:"device_count,omitempty"`
	VirtualmachineCount *int64            `json:"virtualmachine_count,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	TagsAll             []string          `json:"tags_all,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
}

// ClusterResponseDTOFromGoNetbox converts a *models.Cluster to the response DTO.
func ClusterResponseDTOFromGoNetbox(goNetboxModel *models.Cluster) *ClusterResponseDTO {
	responseDTO := &ClusterResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Type != nil {
		v := goNetboxModel.Type.ID
		responseDTO.Type = &v
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	responseDTO.ScopeType = goNetboxModel.ScopeType
	responseDTO.ScopeID = goNetboxModel.ScopeID
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
		v := goNetboxModel.DeviceCount
		responseDTO.DeviceCount = &v
	}
	{
		v := goNetboxModel.VirtualmachineCount
		responseDTO.VirtualmachineCount = &v
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
