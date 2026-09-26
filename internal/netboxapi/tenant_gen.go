// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenantRequestDTO is the request DTO of the tenant resource; Payload writes it into
// the JSON request body.
type TenantRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Group        *int64            `json:"group,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the tenant resource.
func (requestDTO *TenantRequestDTO) Payload() map[string]any {
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
	if requestDTO.Comments != nil {
		payload["comments"] = *requestDTO.Comments
	} else {
		payload["comments"] = ""
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	} else {
		payload["group"] = nil
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

// TenantResponseDTO is the response DTO of the tenant resource, built from
// the go-netbox Tenant by TenantResponseDTOFromGoNetbox.
type TenantResponseDTO struct {
	ID                  *int64            `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty"`
	Slug                *string           `json:"slug,omitempty"`
	Description         *string           `json:"description,omitempty"`
	Comments            *string           `json:"comments,omitempty"`
	Group               *int64            `json:"group,omitempty"`
	Owner               *int64            `json:"owner,omitempty"`
	Created             *string           `json:"created,omitempty"`
	LastUpdated         *string           `json:"last_updated,omitempty"`
	URL                 *string           `json:"url,omitempty"`
	CircuitCount        *int64            `json:"circuit_count,omitempty"`
	DeviceCount         *int64            `json:"device_count,omitempty"`
	IpaddressCount      *int64            `json:"ipaddress_count,omitempty"`
	PrefixCount         *int64            `json:"prefix_count,omitempty"`
	RackCount           *int64            `json:"rack_count,omitempty"`
	SiteCount           *int64            `json:"site_count,omitempty"`
	VirtualmachineCount *int64            `json:"virtualmachine_count,omitempty"`
	VlanCount           *int64            `json:"vlan_count,omitempty"`
	VrfCount            *int64            `json:"vrf_count,omitempty"`
	ClusterCount        *int64            `json:"cluster_count,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	TagsAll             []string          `json:"tags_all,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
}

// TenantResponseDTOFromGoNetbox converts a *models.Tenant to the response DTO.
func TenantResponseDTOFromGoNetbox(goNetboxModel *models.Tenant) *TenantResponseDTO {
	responseDTO := &TenantResponseDTO{}
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
	if goNetboxModel.Comments != "" {
		v := goNetboxModel.Comments
		responseDTO.Comments = &v
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
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
		v := goNetboxModel.CircuitCount
		responseDTO.CircuitCount = &v
	}
	{
		v := goNetboxModel.DeviceCount
		responseDTO.DeviceCount = &v
	}
	{
		v := goNetboxModel.IpaddressCount
		responseDTO.IpaddressCount = &v
	}
	{
		v := goNetboxModel.PrefixCount
		responseDTO.PrefixCount = &v
	}
	{
		v := goNetboxModel.RackCount
		responseDTO.RackCount = &v
	}
	{
		v := goNetboxModel.SiteCount
		responseDTO.SiteCount = &v
	}
	{
		v := goNetboxModel.VirtualmachineCount
		responseDTO.VirtualmachineCount = &v
	}
	{
		v := goNetboxModel.VlanCount
		responseDTO.VlanCount = &v
	}
	{
		v := goNetboxModel.VrfCount
		responseDTO.VrfCount = &v
	}
	{
		v := goNetboxModel.ClusterCount
		responseDTO.ClusterCount = &v
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
