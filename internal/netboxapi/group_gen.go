// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// GroupRequestDTO is the request DTO of the group resource; Payload writes it into
// the JSON request body.
type GroupRequestDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the group resource.
func (requestDTO *GroupRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	return payload
}

// GroupResponseDTO is the response DTO of the group resource, built from
// the go-netbox Group by GroupResponseDTOFromGoNetbox.
type GroupResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	UserCount   *int64  `json:"user_count,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// GroupResponseDTOFromGoNetbox converts a *models.Group to the response DTO.
func GroupResponseDTOFromGoNetbox(goNetboxModel *models.Group) *GroupResponseDTO {
	responseDTO := &GroupResponseDTO{}
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
		v := goNetboxModel.UserCount
		responseDTO.UserCount = &v
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
