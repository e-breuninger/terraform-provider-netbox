// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// L2vpnTerminationRequestDTO is the request DTO of the l2vpn_termination resource; Payload writes it into
// the JSON request body.
type L2vpnTerminationRequestDTO struct {
	L2vpn                     *int64            `json:"l2vpn,omitempty"`
	AssignedObjectType        *string           `json:"assigned_object_type,omitempty"`
	AssignedObjectID          *int64            `json:"assigned_object_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	VlanID                    *int64            `json:"vlan_id,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the l2vpn_termination resource.
func (requestDTO *L2vpnTerminationRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.L2vpn != nil {
		payload["l2vpn"] = *requestDTO.L2vpn
	}
	if requestDTO.AssignedObjectType != nil {
		payload["assigned_object_type"] = *requestDTO.AssignedObjectType
	}
	if requestDTO.AssignedObjectID != nil {
		payload["assigned_object_id"] = *requestDTO.AssignedObjectID
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// L2vpnTerminationResponseDTO is the response DTO of the l2vpn_termination resource, built from
// the go-netbox L2VPNTermination by L2vpnTerminationResponseDTOFromGoNetbox.
type L2vpnTerminationResponseDTO struct {
	ID                        *int64            `json:"id,omitempty"`
	L2vpn                     *int64            `json:"l2vpn,omitempty"`
	AssignedObjectType        *string           `json:"assigned_object_type,omitempty"`
	AssignedObjectID          *int64            `json:"assigned_object_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	VlanID                    *int64            `json:"vlan_id,omitempty"`
	Created                   *string           `json:"created,omitempty"`
	LastUpdated               *string           `json:"last_updated,omitempty"`
	URL                       *string           `json:"url,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	TagsAll                   []string          `json:"tags_all,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// L2vpnTerminationResponseDTOFromGoNetbox converts a *models.L2VPNTermination to the response DTO.
func L2vpnTerminationResponseDTOFromGoNetbox(goNetboxModel *models.L2VPNTermination) *L2vpnTerminationResponseDTO {
	responseDTO := &L2vpnTerminationResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.L2vpn != nil {
		v := goNetboxModel.L2vpn.ID
		responseDTO.L2vpn = &v
	}
	responseDTO.AssignedObjectType = goNetboxModel.AssignedObjectType
	responseDTO.AssignedObjectID = goNetboxModel.AssignedObjectID
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
