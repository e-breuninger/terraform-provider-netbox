// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VlanGroupRequestDTO is the request DTO of the vlan_group resource; Payload writes it into
// the JSON request body.
type VlanGroupRequestDTO struct {
	Name         *string                         `json:"name,omitempty"`
	Slug         *string                         `json:"slug,omitempty"`
	Description  *string                         `json:"description,omitempty"`
	Comments     *string                         `json:"comments,omitempty"`
	VidRanges    []*VlanGroupRequestDTOVidRanges `json:"vid_ranges,omitempty"`
	Tenant       *int64                          `json:"tenant,omitempty"`
	ScopeType    *string                         `json:"scope_type,omitempty"`
	ScopeID      *int64                          `json:"scope_id,omitempty"`
	SiteID       *int64                          `json:"site_id,omitempty"`
	LocationID   *int64                          `json:"location_id,omitempty"`
	RegionID     *int64                          `json:"region_id,omitempty"`
	SiteGroupID  *int64                          `json:"site_group_id,omitempty"`
	Owner        *int64                          `json:"owner,omitempty"`
	Tags         []string                        `json:"tags,omitempty"`
	CustomFields map[string]string               `json:"custom_fields,omitempty"`
}

// VlanGroupRequestDTOVidRanges is the request DTO of the vid_ranges object; it is written into the
// request body as JSON.
type VlanGroupRequestDTOVidRanges struct {
	Start *int64 `json:"start,omitempty"`
	End   *int64 `json:"end,omitempty"`
}

// Payload returns the JSON request body of the vlan_group resource.
func (requestDTO *VlanGroupRequestDTO) Payload() map[string]any {
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
	if requestDTO.VidRanges != nil {
		ranges := make([][]int64, 0, len(requestDTO.VidRanges))
		for _, intRange := range requestDTO.VidRanges {
			s, e := int64(0), int64(0)
			if intRange.Start != nil {
				s = *intRange.Start
			}
			if intRange.End != nil {
				e = *intRange.End
			}
			ranges = append(ranges, []int64{s, e})
		}
		payload["vid_ranges"] = ranges
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
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

// VlanGroupResponseDTO is the response DTO of the vlan_group resource, built from
// the go-netbox VLANGroup by VlanGroupResponseDTOFromGoNetbox.
type VlanGroupResponseDTO struct {
	ID           *int64                           `json:"id,omitempty"`
	Name         *string                          `json:"name,omitempty"`
	Slug         *string                          `json:"slug,omitempty"`
	Description  *string                          `json:"description,omitempty"`
	Comments     *string                          `json:"comments,omitempty"`
	VidRanges    []*VlanGroupResponseDTOVidRanges `json:"vid_ranges,omitempty"`
	Tenant       *int64                           `json:"tenant,omitempty"`
	ScopeType    *string                          `json:"scope_type,omitempty"`
	ScopeID      *int64                           `json:"scope_id,omitempty"`
	SiteID       *int64                           `json:"site_id,omitempty"`
	LocationID   *int64                           `json:"location_id,omitempty"`
	RegionID     *int64                           `json:"region_id,omitempty"`
	SiteGroupID  *int64                           `json:"site_group_id,omitempty"`
	Owner        *int64                           `json:"owner,omitempty"`
	Created      *string                          `json:"created,omitempty"`
	LastUpdated  *string                          `json:"last_updated,omitempty"`
	URL          *string                          `json:"url,omitempty"`
	VlanCount    *int64                           `json:"vlan_count,omitempty"`
	Utilization  *string                          `json:"utilization,omitempty"`
	Tags         []string                         `json:"tags,omitempty"`
	TagsAll      []string                         `json:"tags_all,omitempty"`
	CustomFields map[string]string                `json:"custom_fields,omitempty"`
}

// VlanGroupResponseDTOVidRanges is the response DTO of the vid_ranges object, decoded from the
// go-netbox value as JSON.
type VlanGroupResponseDTOVidRanges struct {
	Start *int64 `json:"start,omitempty"`
	End   *int64 `json:"end,omitempty"`
}

// VlanGroupResponseDTOFromGoNetbox converts a *models.VLANGroup to the response DTO.
func VlanGroupResponseDTOFromGoNetbox(goNetboxModel *models.VLANGroup) *VlanGroupResponseDTO {
	responseDTO := &VlanGroupResponseDTO{}
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
	if goNetboxModel.VidRanges != nil {
		responseDTO.VidRanges = []*VlanGroupResponseDTOVidRanges{}
		for _, pair := range goNetboxModel.VidRanges {
			if len(pair) == 2 {
				s, e := pair[0], pair[1]
				responseDTO.VidRanges = append(responseDTO.VidRanges, &VlanGroupResponseDTOVidRanges{Start: &s, End: &e})
			}
		}
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
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
	{
		v := goNetboxModel.VlanCount
		responseDTO.VlanCount = &v
	}
	if goNetboxModel.Utilization != "" {
		v := goNetboxModel.Utilization
		responseDTO.Utilization = &v
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
