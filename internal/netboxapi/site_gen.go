// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// SiteRequestDTO is the request DTO of the site resource; Payload writes it into
// the JSON request body.
type SiteRequestDTO struct {
	Name            *string           `json:"name,omitempty"`
	Slug            *string           `json:"slug,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Facility        *string           `json:"facility,omitempty"`
	PhysicalAddress *string           `json:"physical_address,omitempty"`
	ShippingAddress *string           `json:"shipping_address,omitempty"`
	Comments        *string           `json:"comments,omitempty"`
	TimeZone        *string           `json:"time_zone,omitempty"`
	Latitude        *float64          `json:"latitude,omitempty"`
	Longitude       *float64          `json:"longitude,omitempty"`
	Region          *int64            `json:"region,omitempty"`
	Tenant          *int64            `json:"tenant,omitempty"`
	Group           *int64            `json:"group,omitempty"`
	Asns            []int64           `json:"asns,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the site resource.
func (requestDTO *SiteRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Name != nil {
		payload["slug"] = Slugify(*requestDTO.Name)
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Facility != nil {
		payload["facility"] = *requestDTO.Facility
	} else {
		payload["facility"] = ""
	}
	if requestDTO.PhysicalAddress != nil {
		payload["physical_address"] = *requestDTO.PhysicalAddress
	} else {
		payload["physical_address"] = ""
	}
	if requestDTO.ShippingAddress != nil {
		payload["shipping_address"] = *requestDTO.ShippingAddress
	} else {
		payload["shipping_address"] = ""
	}
	if requestDTO.Comments != nil {
		payload["comments"] = *requestDTO.Comments
	} else {
		payload["comments"] = ""
	}
	if requestDTO.TimeZone != nil {
		payload["time_zone"] = *requestDTO.TimeZone
	} else {
		payload["time_zone"] = nil
	}
	if requestDTO.Latitude != nil {
		payload["latitude"] = *requestDTO.Latitude
	} else {
		payload["latitude"] = nil
	}
	if requestDTO.Longitude != nil {
		payload["longitude"] = *requestDTO.Longitude
	} else {
		payload["longitude"] = nil
	}
	if requestDTO.Region != nil {
		payload["region"] = *requestDTO.Region
	} else {
		payload["region"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	} else {
		payload["group"] = nil
	}
	if requestDTO.Asns != nil {
		payload["asns"] = requestDTO.Asns
	} else {
		payload["asns"] = []any{}
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

// SiteResponseDTO is the response DTO of the site resource, built from
// the go-netbox Site by SiteResponseDTOFromGoNetbox.
type SiteResponseDTO struct {
	ID                  *int64            `json:"id,omitempty"`
	Name                *string           `json:"name,omitempty"`
	Slug                *string           `json:"slug,omitempty"`
	Status              *string           `json:"status,omitempty"`
	Description         *string           `json:"description,omitempty"`
	Facility            *string           `json:"facility,omitempty"`
	PhysicalAddress     *string           `json:"physical_address,omitempty"`
	ShippingAddress     *string           `json:"shipping_address,omitempty"`
	Comments            *string           `json:"comments,omitempty"`
	TimeZone            *string           `json:"time_zone,omitempty"`
	Latitude            *float64          `json:"latitude,omitempty"`
	Longitude           *float64          `json:"longitude,omitempty"`
	Region              *int64            `json:"region,omitempty"`
	Tenant              *int64            `json:"tenant,omitempty"`
	Group               *int64            `json:"group,omitempty"`
	Asns                []int64           `json:"asns,omitempty"`
	Owner               *int64            `json:"owner,omitempty"`
	Created             *string           `json:"created,omitempty"`
	LastUpdated         *string           `json:"last_updated,omitempty"`
	URL                 *string           `json:"url,omitempty"`
	CircuitCount        *int64            `json:"circuit_count,omitempty"`
	DeviceCount         *int64            `json:"device_count,omitempty"`
	PrefixCount         *int64            `json:"prefix_count,omitempty"`
	RackCount           *int64            `json:"rack_count,omitempty"`
	VirtualmachineCount *int64            `json:"virtualmachine_count,omitempty"`
	VlanCount           *int64            `json:"vlan_count,omitempty"`
	Tags                []string          `json:"tags,omitempty"`
	TagsAll             []string          `json:"tags_all,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
}

// SiteResponseDTOFromGoNetbox converts a *models.Site to the response DTO.
func SiteResponseDTOFromGoNetbox(goNetboxModel *models.Site) *SiteResponseDTO {
	responseDTO := &SiteResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Facility != "" {
		v := goNetboxModel.Facility
		responseDTO.Facility = &v
	}
	if goNetboxModel.PhysicalAddress != "" {
		v := goNetboxModel.PhysicalAddress
		responseDTO.PhysicalAddress = &v
	}
	if goNetboxModel.ShippingAddress != "" {
		v := goNetboxModel.ShippingAddress
		responseDTO.ShippingAddress = &v
	}
	if goNetboxModel.Comments != "" {
		v := goNetboxModel.Comments
		responseDTO.Comments = &v
	}
	responseDTO.TimeZone = goNetboxModel.TimeZone
	responseDTO.Latitude = goNetboxModel.Latitude
	responseDTO.Longitude = goNetboxModel.Longitude
	if goNetboxModel.Region != nil {
		v := goNetboxModel.Region.ID
		responseDTO.Region = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	if goNetboxModel.Asns != nil {
		responseDTO.Asns = []int64{}
		for _, ref := range goNetboxModel.Asns {
			if ref != nil {
				responseDTO.Asns = append(responseDTO.Asns, ref.ID)
			}
		}
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
		v := goNetboxModel.PrefixCount
		responseDTO.PrefixCount = &v
	}
	{
		v := goNetboxModel.RackCount
		responseDTO.RackCount = &v
	}
	{
		v := goNetboxModel.VirtualmachineCount
		responseDTO.VirtualmachineCount = &v
	}
	{
		v := goNetboxModel.VlanCount
		responseDTO.VlanCount = &v
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
