// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// OwnerRequestDTO is the request DTO of the owner resource; Payload writes it into
// the JSON request body.
type OwnerRequestDTO struct {
	Name        *string `json:"name,omitempty"`
	Group       *int64  `json:"group,omitempty"`
	Users       []int64 `json:"users,omitempty"`
	UserGroups  []int64 `json:"user_groups,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the owner resource.
func (requestDTO *OwnerRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	}
	if requestDTO.Users != nil {
		payload["users"] = requestDTO.Users
	} else {
		payload["users"] = []any{}
	}
	if requestDTO.UserGroups != nil {
		payload["user_groups"] = requestDTO.UserGroups
	} else {
		payload["user_groups"] = []any{}
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	return payload
}

// OwnerResponseDTO is the response DTO of the owner resource, built from
// the go-netbox Owner by OwnerResponseDTOFromGoNetbox.
type OwnerResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Group       *int64  `json:"group,omitempty"`
	Users       []int64 `json:"users,omitempty"`
	UserGroups  []int64 `json:"user_groups,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// OwnerResponseDTOFromGoNetbox converts a *models.Owner to the response DTO.
func OwnerResponseDTOFromGoNetbox(goNetboxModel *models.Owner) *OwnerResponseDTO {
	responseDTO := &OwnerResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	if goNetboxModel.Users != nil {
		responseDTO.Users = []int64{}
		for _, ref := range goNetboxModel.Users {
			if ref != nil {
				responseDTO.Users = append(responseDTO.Users, ref.ID)
			}
		}
	}
	if goNetboxModel.UserGroups != nil {
		responseDTO.UserGroups = []int64{}
		for _, ref := range goNetboxModel.UserGroups {
			if ref != nil {
				responseDTO.UserGroups = append(responseDTO.UserGroups, ref.ID)
			}
		}
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
