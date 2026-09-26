// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// InventoryItemRequestDTO is the request DTO of the inventory_item resource; Payload writes it into
// the JSON request body.
type InventoryItemRequestDTO struct {
	Device        *int64            `json:"device,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Parent        *int64            `json:"parent,omitempty"`
	Role          *int64            `json:"role,omitempty"`
	Manufacturer  *int64            `json:"manufacturer,omitempty"`
	PartID        *string           `json:"part_id,omitempty"`
	Serial        *string           `json:"serial,omitempty"`
	AssetTag      *string           `json:"asset_tag,omitempty"`
	Discovered    *bool             `json:"discovered,omitempty"`
	ComponentType *string           `json:"component_type,omitempty"`
	ComponentID   *int64            `json:"component_id,omitempty"`
	Label         *string           `json:"label,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the inventory_item resource.
func (requestDTO *InventoryItemRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Parent != nil {
		payload["parent"] = *requestDTO.Parent
	} else {
		payload["parent"] = nil
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	} else {
		payload["role"] = nil
	}
	if requestDTO.Manufacturer != nil {
		payload["manufacturer"] = *requestDTO.Manufacturer
	} else {
		payload["manufacturer"] = nil
	}
	if requestDTO.PartID != nil {
		payload["part_id"] = *requestDTO.PartID
	} else {
		payload["part_id"] = ""
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
	if requestDTO.Discovered != nil {
		payload["discovered"] = *requestDTO.Discovered
	}
	if requestDTO.ComponentType != nil {
		payload["component_type"] = *requestDTO.ComponentType
	} else {
		payload["component_type"] = nil
	}
	if requestDTO.ComponentID != nil {
		payload["component_id"] = *requestDTO.ComponentID
	} else {
		payload["component_id"] = nil
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
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

// InventoryItemResponseDTO is the response DTO of the inventory_item resource, built from
// the go-netbox InventoryItem by InventoryItemResponseDTOFromGoNetbox.
type InventoryItemResponseDTO struct {
	ID            *int64            `json:"id,omitempty"`
	Device        *int64            `json:"device,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Parent        *int64            `json:"parent,omitempty"`
	Role          *int64            `json:"role,omitempty"`
	Manufacturer  *int64            `json:"manufacturer,omitempty"`
	PartID        *string           `json:"part_id,omitempty"`
	Serial        *string           `json:"serial,omitempty"`
	AssetTag      *string           `json:"asset_tag,omitempty"`
	Discovered    *bool             `json:"discovered,omitempty"`
	ComponentType *string           `json:"component_type,omitempty"`
	ComponentID   *int64            `json:"component_id,omitempty"`
	Label         *string           `json:"label,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Created       *string           `json:"created,omitempty"`
	LastUpdated   *string           `json:"last_updated,omitempty"`
	URL           *string           `json:"url,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	TagsAll       []string          `json:"tags_all,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// InventoryItemResponseDTOFromGoNetbox converts a *models.InventoryItem to the response DTO.
func InventoryItemResponseDTOFromGoNetbox(goNetboxModel *models.InventoryItem) *InventoryItemResponseDTO {
	responseDTO := &InventoryItemResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Device != nil {
		v := goNetboxModel.Device.ID
		responseDTO.Device = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Parent = goNetboxModel.Parent
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.Manufacturer != nil {
		v := goNetboxModel.Manufacturer.ID
		responseDTO.Manufacturer = &v
	}
	if goNetboxModel.PartID != "" {
		v := goNetboxModel.PartID
		responseDTO.PartID = &v
	}
	if goNetboxModel.Serial != "" {
		v := goNetboxModel.Serial
		responseDTO.Serial = &v
	}
	responseDTO.AssetTag = goNetboxModel.AssetTag
	{
		v := goNetboxModel.Discovered
		responseDTO.Discovered = &v
	}
	responseDTO.ComponentType = goNetboxModel.ComponentType
	responseDTO.ComponentID = goNetboxModel.ComponentID
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
