// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VpnTunnelTerminationRequestDTO is the request DTO of the vpn_tunnel_termination resource; Payload writes it into
// the JSON request body.
type VpnTunnelTerminationRequestDTO struct {
	Tunnel                    *int64            `json:"tunnel,omitempty"`
	Role                      *string           `json:"role,omitempty"`
	TerminationType           *string           `json:"termination_type,omitempty"`
	TerminationID             *int64            `json:"termination_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	OutsideIP                 *int64            `json:"outside_ip,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the vpn_tunnel_termination resource.
func (requestDTO *VpnTunnelTerminationRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Tunnel != nil {
		payload["tunnel"] = *requestDTO.Tunnel
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	}
	if requestDTO.TerminationType != nil {
		payload["termination_type"] = *requestDTO.TerminationType
	}
	if requestDTO.TerminationID != nil {
		payload["termination_id"] = *requestDTO.TerminationID
	}
	if requestDTO.OutsideIP != nil {
		payload["outside_ip"] = *requestDTO.OutsideIP
	} else {
		payload["outside_ip"] = nil
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// VpnTunnelTerminationResponseDTO is the response DTO of the vpn_tunnel_termination resource, built from
// the go-netbox TunnelTermination by VpnTunnelTerminationResponseDTOFromGoNetbox.
type VpnTunnelTerminationResponseDTO struct {
	ID                        *int64            `json:"id,omitempty"`
	Tunnel                    *int64            `json:"tunnel,omitempty"`
	Role                      *string           `json:"role,omitempty"`
	TerminationType           *string           `json:"termination_type,omitempty"`
	TerminationID             *int64            `json:"termination_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	OutsideIP                 *int64            `json:"outside_ip,omitempty"`
	Created                   *string           `json:"created,omitempty"`
	LastUpdated               *string           `json:"last_updated,omitempty"`
	URL                       *string           `json:"url,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	TagsAll                   []string          `json:"tags_all,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// VpnTunnelTerminationResponseDTOFromGoNetbox converts a *models.TunnelTermination to the response DTO.
func VpnTunnelTerminationResponseDTOFromGoNetbox(goNetboxModel *models.TunnelTermination) *VpnTunnelTerminationResponseDTO {
	responseDTO := &VpnTunnelTerminationResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Tunnel != nil {
		v := goNetboxModel.Tunnel.ID
		responseDTO.Tunnel = &v
	}
	if goNetboxModel.Role != nil {
		responseDTO.Role = choiceValue[string](goNetboxModel.Role.Value)
	}
	responseDTO.TerminationType = goNetboxModel.TerminationType
	responseDTO.TerminationID = goNetboxModel.TerminationID
	if goNetboxModel.OutsideIP != nil {
		v := goNetboxModel.OutsideIP.ID
		responseDTO.OutsideIP = &v
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
