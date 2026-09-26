// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// InventoryItemRoleRequestDTO is the request DTO of the inventory_item_role resource; Payload writes it into
// the JSON request body.
type InventoryItemRoleRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Color        *string           `json:"color,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the inventory_item_role resource.
func (requestDTO *InventoryItemRoleRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Name != nil {
		payload["slug"] = Slugify(*requestDTO.Name)
	}
	if requestDTO.Color != nil {
		payload["color"] = *requestDTO.Color
	} else {
		payload["color"] = ""
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

// InventoryItemRoleResponseDTO is the response DTO of the inventory_item_role resource, built from
// the go-netbox InventoryItemRole by InventoryItemRoleResponseDTOFromGoNetbox.
type InventoryItemRoleResponseDTO struct {
	ID                 *int64            `json:"id,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Slug               *string           `json:"slug,omitempty"`
	Color              *string           `json:"color,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Owner              *int64            `json:"owner,omitempty"`
	Created            *string           `json:"created,omitempty"`
	LastUpdated        *string           `json:"last_updated,omitempty"`
	URL                *string           `json:"url,omitempty"`
	InventoryitemCount *int64            `json:"inventoryitem_count,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	TagsAll            []string          `json:"tags_all,omitempty"`
	CustomFields       map[string]string `json:"custom_fields,omitempty"`
}

// InventoryItemRoleResponseDTOFromGoNetbox converts a *models.InventoryItemRole to the response DTO.
func InventoryItemRoleResponseDTOFromGoNetbox(goNetboxModel *models.InventoryItemRole) *InventoryItemRoleResponseDTO {
	responseDTO := &InventoryItemRoleResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.Color != "" {
		v := goNetboxModel.Color
		responseDTO.Color = &v
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
	{
		v := goNetboxModel.InventoryitemCount
		responseDTO.InventoryitemCount = &v
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
