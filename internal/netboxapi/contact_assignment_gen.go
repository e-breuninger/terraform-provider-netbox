// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ContactAssignmentRequestDTO is the request DTO of the contact_assignment resource; Payload writes it into
// the JSON request body.
type ContactAssignmentRequestDTO struct {
	ObjectType   *string           `json:"object_type,omitempty"`
	ObjectID     *int64            `json:"object_id,omitempty"`
	Contact      *int64            `json:"contact,omitempty"`
	Role         *int64            `json:"role,omitempty"`
	Priority     *string           `json:"priority,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the contact_assignment resource.
func (requestDTO *ContactAssignmentRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.ObjectType != nil {
		payload["object_type"] = *requestDTO.ObjectType
	}
	if requestDTO.ObjectID != nil {
		payload["object_id"] = *requestDTO.ObjectID
	}
	if requestDTO.Contact != nil {
		payload["contact"] = *requestDTO.Contact
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	}
	if requestDTO.Priority != nil {
		payload["priority"] = *requestDTO.Priority
	} else {
		payload["priority"] = ""
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// ContactAssignmentResponseDTO is the response DTO of the contact_assignment resource, built from
// the go-netbox ContactAssignment by ContactAssignmentResponseDTOFromGoNetbox.
type ContactAssignmentResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	ObjectType   *string           `json:"object_type,omitempty"`
	ObjectID     *int64            `json:"object_id,omitempty"`
	Contact      *int64            `json:"contact,omitempty"`
	Role         *int64            `json:"role,omitempty"`
	Priority     *string           `json:"priority,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// ContactAssignmentResponseDTOFromGoNetbox converts a *models.ContactAssignment to the response DTO.
func ContactAssignmentResponseDTOFromGoNetbox(goNetboxModel *models.ContactAssignment) *ContactAssignmentResponseDTO {
	responseDTO := &ContactAssignmentResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.ObjectType = goNetboxModel.ObjectType
	responseDTO.ObjectID = goNetboxModel.ObjectID
	if goNetboxModel.Contact != nil {
		v := goNetboxModel.Contact.ID
		responseDTO.Contact = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
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
	responseDTO.CustomFields = customFieldValues(goNetboxModel.CustomFields)
	return responseDTO
}
