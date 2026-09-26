// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// RackRequestDTO is the request DTO of the rack resource; Payload writes it into
// the JSON request body.
type RackRequestDTO struct {
	Name          *string           `json:"name,omitempty"`
	Site          *int64            `json:"site,omitempty"`
	Location      *int64            `json:"location,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	Role          *int64            `json:"role,omitempty"`
	RackType      *int64            `json:"rack_type,omitempty"`
	Status        *string           `json:"status,omitempty"`
	FormFactor    *string           `json:"form_factor,omitempty"`
	Width         *int64            `json:"width,omitempty"`
	UHeight       *int64            `json:"u_height,omitempty"`
	DescUnits     *bool             `json:"desc_units,omitempty"`
	Serial        *string           `json:"serial,omitempty"`
	AssetTag      *string           `json:"asset_tag,omitempty"`
	FacilityID    *string           `json:"facility_id,omitempty"`
	OuterWidth    *int64            `json:"outer_width,omitempty"`
	OuterDepth    *int64            `json:"outer_depth,omitempty"`
	OuterUnit     *string           `json:"outer_unit,omitempty"`
	MountingDepth *int64            `json:"mounting_depth,omitempty"`
	Weight        *float64          `json:"weight,omitempty"`
	MaxWeight     *int64            `json:"max_weight,omitempty"`
	WeightUnit    *string           `json:"weight_unit,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the rack resource.
func (requestDTO *RackRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Site != nil {
		payload["site"] = *requestDTO.Site
	}
	if requestDTO.Location != nil {
		payload["location"] = *requestDTO.Location
	} else {
		payload["location"] = nil
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
	if requestDTO.RackType != nil {
		payload["rack_type"] = *requestDTO.RackType
	} else {
		payload["rack_type"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.FormFactor != nil {
		payload["form_factor"] = *requestDTO.FormFactor
	} else {
		payload["form_factor"] = nil
	}
	if requestDTO.Width != nil {
		payload["width"] = *requestDTO.Width
	}
	if requestDTO.UHeight != nil {
		payload["u_height"] = *requestDTO.UHeight
	}
	if requestDTO.DescUnits != nil {
		payload["desc_units"] = *requestDTO.DescUnits
	}
	if requestDTO.Serial != nil {
		payload["serial"] = *requestDTO.Serial
	} else {
		payload["serial"] = ""
	}
	if requestDTO.AssetTag != nil {
		payload["asset_tag"] = *requestDTO.AssetTag
	} else {
		payload["asset_tag"] = nil
	}
	if requestDTO.FacilityID != nil {
		payload["facility_id"] = *requestDTO.FacilityID
	} else {
		payload["facility_id"] = nil
	}
	if requestDTO.OuterWidth != nil {
		payload["outer_width"] = *requestDTO.OuterWidth
	} else {
		payload["outer_width"] = nil
	}
	if requestDTO.OuterDepth != nil {
		payload["outer_depth"] = *requestDTO.OuterDepth
	} else {
		payload["outer_depth"] = nil
	}
	if requestDTO.OuterUnit != nil {
		payload["outer_unit"] = *requestDTO.OuterUnit
	} else {
		payload["outer_unit"] = nil
	}
	if requestDTO.MountingDepth != nil {
		payload["mounting_depth"] = *requestDTO.MountingDepth
	} else {
		payload["mounting_depth"] = nil
	}
	if requestDTO.Weight != nil {
		payload["weight"] = *requestDTO.Weight
	} else {
		payload["weight"] = nil
	}
	if requestDTO.MaxWeight != nil {
		payload["max_weight"] = *requestDTO.MaxWeight
	} else {
		payload["max_weight"] = nil
	}
	if requestDTO.WeightUnit != nil {
		payload["weight_unit"] = *requestDTO.WeightUnit
	} else {
		payload["weight_unit"] = nil
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

// RackResponseDTO is the response DTO of the rack resource, built from
// the go-netbox Rack by RackResponseDTOFromGoNetbox.
type RackResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Site           *int64            `json:"site,omitempty"`
	Location       *int64            `json:"location,omitempty"`
	Tenant         *int64            `json:"tenant,omitempty"`
	Role           *int64            `json:"role,omitempty"`
	RackType       *int64            `json:"rack_type,omitempty"`
	Status         *string           `json:"status,omitempty"`
	FormFactor     *string           `json:"form_factor,omitempty"`
	Width          *int64            `json:"width,omitempty"`
	UHeight        *int64            `json:"u_height,omitempty"`
	DescUnits      *bool             `json:"desc_units,omitempty"`
	Serial         *string           `json:"serial,omitempty"`
	AssetTag       *string           `json:"asset_tag,omitempty"`
	FacilityID     *string           `json:"facility_id,omitempty"`
	OuterWidth     *int64            `json:"outer_width,omitempty"`
	OuterDepth     *int64            `json:"outer_depth,omitempty"`
	OuterUnit      *string           `json:"outer_unit,omitempty"`
	MountingDepth  *int64            `json:"mounting_depth,omitempty"`
	Weight         *float64          `json:"weight,omitempty"`
	MaxWeight      *int64            `json:"max_weight,omitempty"`
	WeightUnit     *string           `json:"weight_unit,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Comments       *string           `json:"comments,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Created        *string           `json:"created,omitempty"`
	LastUpdated    *string           `json:"last_updated,omitempty"`
	URL            *string           `json:"url,omitempty"`
	DeviceCount    *int64            `json:"device_count,omitempty"`
	PowerfeedCount *int64            `json:"powerfeed_count,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	TagsAll        []string          `json:"tags_all,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// RackResponseDTOFromGoNetbox converts a *models.Rack to the response DTO.
func RackResponseDTOFromGoNetbox(goNetboxModel *models.Rack) *RackResponseDTO {
	responseDTO := &RackResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Site != nil {
		v := goNetboxModel.Site.ID
		responseDTO.Site = &v
	}
	if goNetboxModel.Location != nil {
		v := goNetboxModel.Location.ID
		responseDTO.Location = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.RackType != nil {
		v := goNetboxModel.RackType.ID
		responseDTO.RackType = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.FormFactor != nil {
		responseDTO.FormFactor = choiceValue[string](goNetboxModel.FormFactor.Value)
	}
	if goNetboxModel.Width != nil {
		responseDTO.Width = choiceValue[int64](goNetboxModel.Width.Value)
	}
	{
		v := goNetboxModel.UHeight
		responseDTO.UHeight = &v
	}
	{
		v := goNetboxModel.DescUnits
		responseDTO.DescUnits = &v
	}
	if goNetboxModel.Serial != "" {
		v := goNetboxModel.Serial
		responseDTO.Serial = &v
	}
	responseDTO.AssetTag = goNetboxModel.AssetTag
	responseDTO.FacilityID = goNetboxModel.FacilityID
	responseDTO.OuterWidth = goNetboxModel.OuterWidth
	responseDTO.OuterDepth = goNetboxModel.OuterDepth
	if goNetboxModel.OuterUnit != nil {
		responseDTO.OuterUnit = choiceValue[string](goNetboxModel.OuterUnit.Value)
	}
	responseDTO.MountingDepth = goNetboxModel.MountingDepth
	responseDTO.Weight = goNetboxModel.Weight
	responseDTO.MaxWeight = goNetboxModel.MaxWeight
	if goNetboxModel.WeightUnit != nil {
		responseDTO.WeightUnit = choiceValue[string](goNetboxModel.WeightUnit.Value)
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
		v := goNetboxModel.DeviceCount
		responseDTO.DeviceCount = &v
	}
	{
		v := goNetboxModel.PowerfeedCount
		responseDTO.PowerfeedCount = &v
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
