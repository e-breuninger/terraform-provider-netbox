// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// MacAddressRequestDTO is the request DTO of the mac_address resource; Payload writes it into
// the JSON request body.
type MacAddressRequestDTO struct {
	MacAddress                *string           `json:"mac_address,omitempty"`
	AssignedObjectType        *string           `json:"assigned_object_type,omitempty"`
	AssignedObjectID          *int64            `json:"assigned_object_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	Description               *string           `json:"description,omitempty"`
	Comments                  *string           `json:"comments,omitempty"`
	Owner                     *int64            `json:"owner,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the mac_address resource.
func (requestDTO *MacAddressRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.MacAddress != nil {
		payload["mac_address"] = *requestDTO.MacAddress
	}
	if requestDTO.AssignedObjectType != nil {
		payload["assigned_object_type"] = *requestDTO.AssignedObjectType
	} else {
		payload["assigned_object_type"] = nil
	}
	if requestDTO.AssignedObjectID != nil {
		payload["assigned_object_id"] = *requestDTO.AssignedObjectID
	} else {
		payload["assigned_object_id"] = nil
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

// MacAddressResponseDTO is the response DTO of the mac_address resource, built from
// the go-netbox MACAddress by MacAddressResponseDTOFromGoNetbox.
type MacAddressResponseDTO struct {
	ID                        *int64            `json:"id,omitempty"`
	MacAddress                *string           `json:"mac_address,omitempty"`
	AssignedObjectType        *string           `json:"assigned_object_type,omitempty"`
	AssignedObjectID          *int64            `json:"assigned_object_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	Description               *string           `json:"description,omitempty"`
	Comments                  *string           `json:"comments,omitempty"`
	Owner                     *int64            `json:"owner,omitempty"`
	LastUpdated               *string           `json:"last_updated,omitempty"`
	URL                       *string           `json:"url,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	TagsAll                   []string          `json:"tags_all,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// MacAddressResponseDTOFromGoNetbox converts a *models.MACAddress to the response DTO.
func MacAddressResponseDTOFromGoNetbox(goNetboxModel *models.MACAddress) *MacAddressResponseDTO {
	responseDTO := &MacAddressResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.MacAddress = goNetboxModel.MacAddress
	responseDTO.AssignedObjectType = goNetboxModel.AssignedObjectType
	responseDTO.AssignedObjectID = goNetboxModel.AssignedObjectID
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
