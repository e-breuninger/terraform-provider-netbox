// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VpnTunnelRequestDTO is the request DTO of the vpn_tunnel resource; Payload writes it into
// the JSON request body.
type VpnTunnelRequestDTO struct {
	Name          *string           `json:"name,omitempty"`
	Status        *string           `json:"status,omitempty"`
	Encapsulation *string           `json:"encapsulation,omitempty"`
	Group         *int64            `json:"group,omitempty"`
	IpsecProfile  *int64            `json:"ipsec_profile,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	TunnelID      *int64            `json:"tunnel_id,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the vpn_tunnel resource.
func (requestDTO *VpnTunnelRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Encapsulation != nil {
		payload["encapsulation"] = *requestDTO.Encapsulation
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	} else {
		payload["group"] = nil
	}
	if requestDTO.IpsecProfile != nil {
		payload["ipsec_profile"] = *requestDTO.IpsecProfile
	} else {
		payload["ipsec_profile"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.TunnelID != nil {
		payload["tunnel_id"] = *requestDTO.TunnelID
	} else {
		payload["tunnel_id"] = nil
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

// VpnTunnelResponseDTO is the response DTO of the vpn_tunnel resource, built from
// the go-netbox Tunnel by VpnTunnelResponseDTOFromGoNetbox.
type VpnTunnelResponseDTO struct {
	ID                *int64            `json:"id,omitempty"`
	Name              *string           `json:"name,omitempty"`
	Status            *string           `json:"status,omitempty"`
	Encapsulation     *string           `json:"encapsulation,omitempty"`
	Group             *int64            `json:"group,omitempty"`
	IpsecProfile      *int64            `json:"ipsec_profile,omitempty"`
	Tenant            *int64            `json:"tenant,omitempty"`
	TunnelID          *int64            `json:"tunnel_id,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Comments          *string           `json:"comments,omitempty"`
	Owner             *int64            `json:"owner,omitempty"`
	Created           *string           `json:"created,omitempty"`
	LastUpdated       *string           `json:"last_updated,omitempty"`
	URL               *string           `json:"url,omitempty"`
	TerminationsCount *int64            `json:"terminations_count,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	TagsAll           []string          `json:"tags_all,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
}

// VpnTunnelResponseDTOFromGoNetbox converts a *models.Tunnel to the response DTO.
func VpnTunnelResponseDTOFromGoNetbox(goNetboxModel *models.Tunnel) *VpnTunnelResponseDTO {
	responseDTO := &VpnTunnelResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Encapsulation != nil {
		responseDTO.Encapsulation = choiceValue[string](goNetboxModel.Encapsulation.Value)
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	if goNetboxModel.IpsecProfile != nil {
		v := goNetboxModel.IpsecProfile.ID
		responseDTO.IpsecProfile = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	responseDTO.TunnelID = goNetboxModel.TunnelID
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
	{
		v := goNetboxModel.TerminationsCount
		responseDTO.TerminationsCount = &v
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
