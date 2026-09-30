// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// FhrpGroupAssignmentRequestDTO is the request DTO of the fhrp_group_assignment resource; Payload writes it into
// the JSON request body.
type FhrpGroupAssignmentRequestDTO struct {
	Group         *int64  `json:"group,omitempty"`
	InterfaceType *string `json:"interface_type,omitempty"`
	InterfaceID   *int64  `json:"interface_id,omitempty"`
	Priority      *int64  `json:"priority,omitempty"`
}

// Payload returns the JSON request body of the fhrp_group_assignment resource.
func (requestDTO *FhrpGroupAssignmentRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	}
	if requestDTO.InterfaceType != nil {
		payload["interface_type"] = *requestDTO.InterfaceType
	}
	if requestDTO.InterfaceID != nil {
		payload["interface_id"] = *requestDTO.InterfaceID
	}
	if requestDTO.Priority != nil {
		payload["priority"] = *requestDTO.Priority
	}
	return payload
}

// FhrpGroupAssignmentResponseDTO is the response DTO of the fhrp_group_assignment resource, built from
// the go-netbox FHRPGroupAssignment by FhrpGroupAssignmentResponseDTOFromGoNetbox.
type FhrpGroupAssignmentResponseDTO struct {
	ID            *int64  `json:"id,omitempty"`
	Group         *int64  `json:"group,omitempty"`
	InterfaceType *string `json:"interface_type,omitempty"`
	InterfaceID   *int64  `json:"interface_id,omitempty"`
	Priority      *int64  `json:"priority,omitempty"`
	Created       *string `json:"created,omitempty"`
	LastUpdated   *string `json:"last_updated,omitempty"`
	URL           *string `json:"url,omitempty"`
}

// FhrpGroupAssignmentResponseDTOFromGoNetbox converts a *models.FHRPGroupAssignment to the response DTO.
func FhrpGroupAssignmentResponseDTOFromGoNetbox(goNetboxModel *models.FHRPGroupAssignment) *FhrpGroupAssignmentResponseDTO {
	responseDTO := &FhrpGroupAssignmentResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	responseDTO.InterfaceType = goNetboxModel.InterfaceType
	responseDTO.InterfaceID = goNetboxModel.InterfaceID
	responseDTO.Priority = goNetboxModel.Priority
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
