// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// PermissionRequestDTO is the request DTO of the permission resource; Payload writes it into
// the JSON request body.
type PermissionRequestDTO struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	Actions     []string `json:"actions,omitempty"`
	Groups      []int64  `json:"groups,omitempty"`
	Users       []int64  `json:"users,omitempty"`
	Constraints *string  `json:"constraints,omitempty"`
}

// Payload returns the JSON request body of the permission resource.
func (requestDTO *PermissionRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.ObjectTypes != nil {
		payload["object_types"] = requestDTO.ObjectTypes
	}
	if requestDTO.Actions != nil {
		payload["actions"] = requestDTO.Actions
	}
	if requestDTO.Groups != nil {
		payload["groups"] = requestDTO.Groups
	} else {
		payload["groups"] = []any{}
	}
	if requestDTO.Users != nil {
		payload["users"] = requestDTO.Users
	} else {
		payload["users"] = []any{}
	}
	if requestDTO.Constraints != nil {
		payload["constraints"] = parseJSONText(*requestDTO.Constraints)
	} else {
		payload["constraints"] = nil
	}
	return payload
}

// PermissionResponseDTO is the response DTO of the permission resource, built from
// the go-netbox ObjectPermission by PermissionResponseDTOFromGoNetbox.
type PermissionResponseDTO struct {
	ID          *int64   `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	Actions     []string `json:"actions,omitempty"`
	Groups      []int64  `json:"groups,omitempty"`
	Users       []int64  `json:"users,omitempty"`
	Constraints *string  `json:"constraints,omitempty"`
	URL         *string  `json:"url,omitempty"`
}

// PermissionResponseDTOFromGoNetbox converts a *models.ObjectPermission to the response DTO.
func PermissionResponseDTOFromGoNetbox(goNetboxModel *models.ObjectPermission) *PermissionResponseDTO {
	responseDTO := &PermissionResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	responseDTO.ObjectTypes = goNetboxModel.ObjectTypes
	responseDTO.Actions = goNetboxModel.Actions
	if goNetboxModel.Groups != nil {
		responseDTO.Groups = []int64{}
		for _, ref := range goNetboxModel.Groups {
			if ref != nil {
				responseDTO.Groups = append(responseDTO.Groups, ref.ID)
			}
		}
	}
	if goNetboxModel.Users != nil {
		responseDTO.Users = []int64{}
		for _, ref := range goNetboxModel.Users {
			if ref != nil {
				responseDTO.Users = append(responseDTO.Users, ref.ID)
			}
		}
	}
	responseDTO.Constraints = jsonText(goNetboxModel.Constraints)
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
