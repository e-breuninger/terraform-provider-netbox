// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitGroupAssignmentRequestDTO is the request DTO of the circuit_group_assignment resource; Payload writes it into
// the JSON request body.
type CircuitGroupAssignmentRequestDTO struct {
	Group      *int64   `json:"group,omitempty"`
	MemberType *string  `json:"member_type,omitempty"`
	MemberID   *int64   `json:"member_id,omitempty"`
	Priority   *string  `json:"priority,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

// Payload returns the JSON request body of the circuit_group_assignment resource.
func (requestDTO *CircuitGroupAssignmentRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	}
	if requestDTO.MemberType != nil {
		payload["member_type"] = *requestDTO.MemberType
	}
	if requestDTO.MemberID != nil {
		payload["member_id"] = *requestDTO.MemberID
	}
	if requestDTO.Priority != nil {
		payload["priority"] = *requestDTO.Priority
	} else {
		payload["priority"] = ""
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	return payload
}

// CircuitGroupAssignmentResponseDTO is the response DTO of the circuit_group_assignment resource, built from
// the go-netbox CircuitGroupAssignment by CircuitGroupAssignmentResponseDTOFromGoNetbox.
type CircuitGroupAssignmentResponseDTO struct {
	ID          *int64   `json:"id,omitempty"`
	Group       *int64   `json:"group,omitempty"`
	MemberType  *string  `json:"member_type,omitempty"`
	MemberID    *int64   `json:"member_id,omitempty"`
	Priority    *string  `json:"priority,omitempty"`
	Created     *string  `json:"created,omitempty"`
	LastUpdated *string  `json:"last_updated,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TagsAll     []string `json:"tags_all,omitempty"`
}

// CircuitGroupAssignmentResponseDTOFromGoNetbox converts a *models.CircuitGroupAssignment to the response DTO.
func CircuitGroupAssignmentResponseDTOFromGoNetbox(goNetboxModel *models.CircuitGroupAssignment) *CircuitGroupAssignmentResponseDTO {
	responseDTO := &CircuitGroupAssignmentResponseDTO{}
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
	responseDTO.MemberType = goNetboxModel.MemberType
	responseDTO.MemberID = goNetboxModel.MemberID
	if goNetboxModel.Priority != nil {
		responseDTO.Priority = choiceValue[string](goNetboxModel.Priority.Value)
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
	return responseDTO
}
