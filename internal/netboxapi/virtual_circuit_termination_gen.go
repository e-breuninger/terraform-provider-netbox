// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualCircuitTerminationRequestDTO is the request DTO of the virtual_circuit_termination resource; Payload writes it into
// the JSON request body.
type VirtualCircuitTerminationRequestDTO struct {
	VirtualCircuit *int64            `json:"virtual_circuit,omitempty"`
	Interface      *int64            `json:"interface,omitempty"`
	Role           *string           `json:"role,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_circuit_termination resource.
func (requestDTO *VirtualCircuitTerminationRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.VirtualCircuit != nil {
		payload["virtual_circuit"] = *requestDTO.VirtualCircuit
	}
	if requestDTO.Interface != nil {
		payload["interface"] = *requestDTO.Interface
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// VirtualCircuitTerminationResponseDTO is the response DTO of the virtual_circuit_termination resource, built from
// the go-netbox VirtualCircuitTermination by VirtualCircuitTerminationResponseDTOFromGoNetbox.
type VirtualCircuitTerminationResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	VirtualCircuit *int64            `json:"virtual_circuit,omitempty"`
	Interface      *int64            `json:"interface,omitempty"`
	Role           *string           `json:"role,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Created        *string           `json:"created,omitempty"`
	LastUpdated    *string           `json:"last_updated,omitempty"`
	URL            *string           `json:"url,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	TagsAll        []string          `json:"tags_all,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// VirtualCircuitTerminationResponseDTOFromGoNetbox converts a *models.VirtualCircuitTermination to the response DTO.
func VirtualCircuitTerminationResponseDTOFromGoNetbox(goNetboxModel *models.VirtualCircuitTermination) *VirtualCircuitTerminationResponseDTO {
	responseDTO := &VirtualCircuitTerminationResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.VirtualCircuit != nil {
		v := goNetboxModel.VirtualCircuit.ID
		responseDTO.VirtualCircuit = &v
	}
	if goNetboxModel.Interface != nil {
		v := goNetboxModel.Interface.ID
		responseDTO.Interface = &v
	}
	if goNetboxModel.Role != nil {
		responseDTO.Role = choiceValue[string](goNetboxModel.Role.Value)
	}
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
