// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// InventoryItemTemplateRequestDTO is the request DTO of the inventory_item_template resource; Payload writes it into
// the JSON request body.
type InventoryItemTemplateRequestDTO struct {
	DeviceType    *int64  `json:"device_type,omitempty"`
	Name          *string `json:"name,omitempty"`
	Parent        *int64  `json:"parent,omitempty"`
	Role          *int64  `json:"role,omitempty"`
	Manufacturer  *int64  `json:"manufacturer,omitempty"`
	PartID        *string `json:"part_id,omitempty"`
	Label         *string `json:"label,omitempty"`
	ComponentType *string `json:"component_type,omitempty"`
	ComponentID   *int64  `json:"component_id,omitempty"`
	Description   *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the inventory_item_template resource.
func (requestDTO *InventoryItemTemplateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.DeviceType != nil {
		payload["device_type"] = *requestDTO.DeviceType
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
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
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
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	return payload
}

// InventoryItemTemplateResponseDTO is the response DTO of the inventory_item_template resource, built from
// the go-netbox InventoryItemTemplate by InventoryItemTemplateResponseDTOFromGoNetbox.
type InventoryItemTemplateResponseDTO struct {
	ID            *int64  `json:"id,omitempty"`
	DeviceType    *int64  `json:"device_type,omitempty"`
	Name          *string `json:"name,omitempty"`
	Parent        *int64  `json:"parent,omitempty"`
	Role          *int64  `json:"role,omitempty"`
	Manufacturer  *int64  `json:"manufacturer,omitempty"`
	PartID        *string `json:"part_id,omitempty"`
	Label         *string `json:"label,omitempty"`
	ComponentType *string `json:"component_type,omitempty"`
	ComponentID   *int64  `json:"component_id,omitempty"`
	Description   *string `json:"description,omitempty"`
	Created       *string `json:"created,omitempty"`
	LastUpdated   *string `json:"last_updated,omitempty"`
	URL           *string `json:"url,omitempty"`
}

// InventoryItemTemplateResponseDTOFromGoNetbox converts a *models.InventoryItemTemplate to the response DTO.
func InventoryItemTemplateResponseDTOFromGoNetbox(goNetboxModel *models.InventoryItemTemplate) *InventoryItemTemplateResponseDTO {
	responseDTO := &InventoryItemTemplateResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.DeviceType != nil {
		v := goNetboxModel.DeviceType.ID
		responseDTO.DeviceType = &v
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
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	responseDTO.ComponentType = goNetboxModel.ComponentType
	responseDTO.ComponentID = goNetboxModel.ComponentID
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
	return responseDTO
}
