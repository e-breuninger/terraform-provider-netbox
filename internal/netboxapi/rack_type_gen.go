// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// RackTypeRequestDTO is the request DTO of the rack_type resource; Payload writes it into
// the JSON request body.
type RackTypeRequestDTO struct {
	Manufacturer  *int64            `json:"manufacturer,omitempty"`
	Model         *string           `json:"model,omitempty"`
	Slug          *string           `json:"slug,omitempty"`
	Description   *string           `json:"description,omitempty"`
	FormFactor    *string           `json:"form_factor,omitempty"`
	Width         *int64            `json:"width,omitempty"`
	UHeight       *int64            `json:"u_height,omitempty"`
	StartingUnit  *int64            `json:"starting_unit,omitempty"`
	DescUnits     *bool             `json:"desc_units,omitempty"`
	OuterWidth    *int64            `json:"outer_width,omitempty"`
	OuterHeight   *int64            `json:"outer_height,omitempty"`
	OuterDepth    *int64            `json:"outer_depth,omitempty"`
	OuterUnit     *string           `json:"outer_unit,omitempty"`
	MountingDepth *int64            `json:"mounting_depth,omitempty"`
	Weight        *float64          `json:"weight,omitempty"`
	MaxWeight     *int64            `json:"max_weight,omitempty"`
	WeightUnit    *string           `json:"weight_unit,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the rack_type resource.
func (requestDTO *RackTypeRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Manufacturer != nil {
		payload["manufacturer"] = *requestDTO.Manufacturer
	}
	if requestDTO.Model != nil {
		payload["model"] = *requestDTO.Model
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Model != nil {
		payload["slug"] = Slugify(*requestDTO.Model)
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.FormFactor != nil {
		payload["form_factor"] = *requestDTO.FormFactor
	}
	if requestDTO.Width != nil {
		payload["width"] = *requestDTO.Width
	}
	if requestDTO.UHeight != nil {
		payload["u_height"] = *requestDTO.UHeight
	}
	if requestDTO.StartingUnit != nil {
		payload["starting_unit"] = *requestDTO.StartingUnit
	}
	if requestDTO.DescUnits != nil {
		payload["desc_units"] = *requestDTO.DescUnits
	}
	if requestDTO.OuterWidth != nil {
		payload["outer_width"] = *requestDTO.OuterWidth
	} else {
		payload["outer_width"] = nil
	}
	if requestDTO.OuterHeight != nil {
		payload["outer_height"] = *requestDTO.OuterHeight
	} else {
		payload["outer_height"] = nil
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

// RackTypeResponseDTO is the response DTO of the rack_type resource, built from
// the go-netbox RackType by RackTypeResponseDTOFromGoNetbox.
type RackTypeResponseDTO struct {
	ID            *int64            `json:"id,omitempty"`
	Manufacturer  *int64            `json:"manufacturer,omitempty"`
	Model         *string           `json:"model,omitempty"`
	Slug          *string           `json:"slug,omitempty"`
	Description   *string           `json:"description,omitempty"`
	FormFactor    *string           `json:"form_factor,omitempty"`
	Width         *int64            `json:"width,omitempty"`
	UHeight       *int64            `json:"u_height,omitempty"`
	StartingUnit  *int64            `json:"starting_unit,omitempty"`
	DescUnits     *bool             `json:"desc_units,omitempty"`
	OuterWidth    *int64            `json:"outer_width,omitempty"`
	OuterHeight   *int64            `json:"outer_height,omitempty"`
	OuterDepth    *int64            `json:"outer_depth,omitempty"`
	OuterUnit     *string           `json:"outer_unit,omitempty"`
	MountingDepth *int64            `json:"mounting_depth,omitempty"`
	Weight        *float64          `json:"weight,omitempty"`
	MaxWeight     *int64            `json:"max_weight,omitempty"`
	WeightUnit    *string           `json:"weight_unit,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Created       *string           `json:"created,omitempty"`
	LastUpdated   *string           `json:"last_updated,omitempty"`
	URL           *string           `json:"url,omitempty"`
	RackCount     *int64            `json:"rack_count,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	TagsAll       []string          `json:"tags_all,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// RackTypeResponseDTOFromGoNetbox converts a *models.RackType to the response DTO.
func RackTypeResponseDTOFromGoNetbox(goNetboxModel *models.RackType) *RackTypeResponseDTO {
	responseDTO := &RackTypeResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Manufacturer != nil {
		v := goNetboxModel.Manufacturer.ID
		responseDTO.Manufacturer = &v
	}
	responseDTO.Model = goNetboxModel.Model
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
		v := goNetboxModel.StartingUnit
		responseDTO.StartingUnit = &v
	}
	{
		v := goNetboxModel.DescUnits
		responseDTO.DescUnits = &v
	}
	responseDTO.OuterWidth = goNetboxModel.OuterWidth
	responseDTO.OuterHeight = goNetboxModel.OuterHeight
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
		v := goNetboxModel.RackCount
		responseDTO.RackCount = &v
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
