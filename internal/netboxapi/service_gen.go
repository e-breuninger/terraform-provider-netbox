// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ServiceRequestDTO is the request DTO of the service resource; Payload writes it into
// the JSON request body.
type ServiceRequestDTO struct {
	Name             *string           `json:"name,omitempty"`
	Protocol         *string           `json:"protocol,omitempty"`
	Ports            []int64           `json:"ports,omitempty"`
	ParentObjectType *string           `json:"parent_object_type,omitempty"`
	ParentObjectID   *int64            `json:"parent_object_id,omitempty"`
	DeviceID         *int64            `json:"device_id,omitempty"`
	VirtualMachineID *int64            `json:"virtual_machine_id,omitempty"`
	Ipaddresses      []int64           `json:"ipaddresses,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Comments         *string           `json:"comments,omitempty"`
	Owner            *int64            `json:"owner,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	CustomFields     map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the service resource.
func (requestDTO *ServiceRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Protocol != nil {
		payload["protocol"] = *requestDTO.Protocol
	}
	if requestDTO.Ports != nil {
		payload["ports"] = requestDTO.Ports
	}
	if requestDTO.ParentObjectType != nil {
		payload["parent_object_type"] = *requestDTO.ParentObjectType
	}
	if requestDTO.ParentObjectID != nil {
		payload["parent_object_id"] = *requestDTO.ParentObjectID
	}
	if requestDTO.Ipaddresses != nil {
		payload["ipaddresses"] = requestDTO.Ipaddresses
	} else {
		payload["ipaddresses"] = []any{}
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Comments != nil {
		payload["comments"] = *requestDTO.Comments
	} else {
		payload["comments"] = ""
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

// ServiceResponseDTO is the response DTO of the service resource, built from
// the go-netbox Service by ServiceResponseDTOFromGoNetbox.
type ServiceResponseDTO struct {
	ID               *int64            `json:"id,omitempty"`
	Name             *string           `json:"name,omitempty"`
	Protocol         *string           `json:"protocol,omitempty"`
	Ports            []int64           `json:"ports,omitempty"`
	ParentObjectType *string           `json:"parent_object_type,omitempty"`
	ParentObjectID   *int64            `json:"parent_object_id,omitempty"`
	DeviceID         *int64            `json:"device_id,omitempty"`
	VirtualMachineID *int64            `json:"virtual_machine_id,omitempty"`
	Ipaddresses      []int64           `json:"ipaddresses,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Comments         *string           `json:"comments,omitempty"`
	Owner            *int64            `json:"owner,omitempty"`
	Created          *string           `json:"created,omitempty"`
	LastUpdated      *string           `json:"last_updated,omitempty"`
	URL              *string           `json:"url,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	TagsAll          []string          `json:"tags_all,omitempty"`
	CustomFields     map[string]string `json:"custom_fields,omitempty"`
}

// ServiceResponseDTOFromGoNetbox converts a *models.Service to the response DTO.
func ServiceResponseDTOFromGoNetbox(goNetboxModel *models.Service) *ServiceResponseDTO {
	responseDTO := &ServiceResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Protocol != nil {
		responseDTO.Protocol = choiceValue[string](goNetboxModel.Protocol.Value)
	}
	responseDTO.Ports = goNetboxModel.Ports
	responseDTO.ParentObjectType = goNetboxModel.ParentObjectType
	responseDTO.ParentObjectID = goNetboxModel.ParentObjectID
	if goNetboxModel.Ipaddresses != nil {
		responseDTO.Ipaddresses = []int64{}
		for _, ref := range goNetboxModel.Ipaddresses {
			if ref != nil {
				responseDTO.Ipaddresses = append(responseDTO.Ipaddresses, ref.ID)
			}
		}
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Comments != "" {
		v := goNetboxModel.Comments
		responseDTO.Comments = &v
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
