// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// WirelessLanRequestDTO is the request DTO of the wireless_lan resource; Payload writes it into
// the JSON request body.
type WirelessLanRequestDTO struct {
	Ssid         *string           `json:"ssid,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Group        *int64            `json:"group,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Vlan         *int64            `json:"vlan,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	AuthType     *string           `json:"auth_type,omitempty"`
	AuthCipher   *string           `json:"auth_cipher,omitempty"`
	AuthPsk      *string           `json:"auth_psk,omitempty"`
	ScopeType    *string           `json:"scope_type,omitempty"`
	ScopeID      *int64            `json:"scope_id,omitempty"`
	SiteID       *int64            `json:"site_id,omitempty"`
	LocationID   *int64            `json:"location_id,omitempty"`
	RegionID     *int64            `json:"region_id,omitempty"`
	SiteGroupID  *int64            `json:"site_group_id,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the wireless_lan resource.
func (requestDTO *WirelessLanRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Ssid != nil {
		payload["ssid"] = *requestDTO.Ssid
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	} else {
		payload["group"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Vlan != nil {
		payload["vlan"] = *requestDTO.Vlan
	} else {
		payload["vlan"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.AuthType != nil {
		payload["auth_type"] = *requestDTO.AuthType
	} else {
		payload["auth_type"] = nil
	}
	if requestDTO.AuthCipher != nil {
		payload["auth_cipher"] = *requestDTO.AuthCipher
	} else {
		payload["auth_cipher"] = nil
	}
	if requestDTO.AuthPsk != nil {
		payload["auth_psk"] = *requestDTO.AuthPsk
	} else {
		payload["auth_psk"] = ""
	}
	if requestDTO.ScopeType != nil {
		payload["scope_type"] = *requestDTO.ScopeType
	} else {
		payload["scope_type"] = nil
	}
	if requestDTO.ScopeID != nil {
		payload["scope_id"] = *requestDTO.ScopeID
	} else {
		payload["scope_id"] = nil
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

// WirelessLanResponseDTO is the response DTO of the wireless_lan resource, built from
// the go-netbox WirelessLAN by WirelessLanResponseDTOFromGoNetbox.
type WirelessLanResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Ssid         *string           `json:"ssid,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Group        *int64            `json:"group,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Vlan         *int64            `json:"vlan,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	AuthType     *string           `json:"auth_type,omitempty"`
	AuthCipher   *string           `json:"auth_cipher,omitempty"`
	AuthPsk      *string           `json:"auth_psk,omitempty"`
	ScopeType    *string           `json:"scope_type,omitempty"`
	ScopeID      *int64            `json:"scope_id,omitempty"`
	SiteID       *int64            `json:"site_id,omitempty"`
	LocationID   *int64            `json:"location_id,omitempty"`
	RegionID     *int64            `json:"region_id,omitempty"`
	SiteGroupID  *int64            `json:"site_group_id,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// WirelessLanResponseDTOFromGoNetbox converts a *models.WirelessLAN to the response DTO.
func WirelessLanResponseDTOFromGoNetbox(goNetboxModel *models.WirelessLAN) *WirelessLanResponseDTO {
	responseDTO := &WirelessLanResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Ssid = goNetboxModel.Ssid
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Vlan != nil {
		v := goNetboxModel.Vlan.ID
		responseDTO.Vlan = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.AuthType != nil {
		responseDTO.AuthType = choiceValue[string](goNetboxModel.AuthType.Value)
	}
	if goNetboxModel.AuthCipher != nil {
		responseDTO.AuthCipher = choiceValue[string](goNetboxModel.AuthCipher.Value)
	}
	if goNetboxModel.AuthPsk != "" {
		v := goNetboxModel.AuthPsk
		responseDTO.AuthPsk = &v
	}
	responseDTO.ScopeType = goNetboxModel.ScopeType
	responseDTO.ScopeID = goNetboxModel.ScopeID
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
