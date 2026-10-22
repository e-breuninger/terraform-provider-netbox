// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// OwnerGroupRequestDTO is the request DTO of the owner_group resource; Payload writes it into
// the JSON request body.
type OwnerGroupRequestDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the owner_group resource.
func (requestDTO *OwnerGroupRequestDTO) Payload() map[string]any {
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

// OwnerGroupResponseDTO is the response DTO of the owner_group resource, built from
// the go-netbox OwnerGroup by OwnerGroupResponseDTOFromGoNetbox.
type OwnerGroupResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
	MemberCount *int64  `json:"member_count,omitempty"`
}

// OwnerGroupResponseDTOFromGoNetbox converts a *models.OwnerGroup to the response DTO.
func OwnerGroupResponseDTOFromGoNetbox(goNetboxModel *models.OwnerGroup) *OwnerGroupResponseDTO {
	responseDTO := &OwnerGroupResponseDTO{}
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
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	{
		v := goNetboxModel.MemberCount
		responseDTO.MemberCount = &v
	}
	return responseDTO
}
