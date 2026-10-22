// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// WirelessLinkRequestDTO is the request DTO of the wireless_link resource; Payload writes it into
// the JSON request body.
type WirelessLinkRequestDTO struct {
	Interfacea   *int64            `json:"interface_a,omitempty"`
	Interfaceb   *int64            `json:"interface_b,omitempty"`
	Ssid         *string           `json:"ssid,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	AuthType     *string           `json:"auth_type,omitempty"`
	AuthCipher   *string           `json:"auth_cipher,omitempty"`
	AuthPsk      *string           `json:"auth_psk,omitempty"`
	Distance     *float64          `json:"distance,omitempty"`
	DistanceUnit *string           `json:"distance_unit,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the wireless_link resource.
func (requestDTO *WirelessLinkRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Interfacea != nil {
		payload["interface_a"] = *requestDTO.Interfacea
	}
	if requestDTO.Interfaceb != nil {
		payload["interface_b"] = *requestDTO.Interfaceb
	}
	if requestDTO.Ssid != nil {
		payload["ssid"] = *requestDTO.Ssid
	} else {
		payload["ssid"] = ""
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
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
	if requestDTO.Distance != nil {
		payload["distance"] = *requestDTO.Distance
	} else {
		payload["distance"] = nil
	}
	if requestDTO.DistanceUnit != nil {
		payload["distance_unit"] = *requestDTO.DistanceUnit
	} else {
		payload["distance_unit"] = nil
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

// WirelessLinkResponseDTO is the response DTO of the wireless_link resource, built from
// the go-netbox WirelessLink by WirelessLinkResponseDTOFromGoNetbox.
type WirelessLinkResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Interfacea   *int64            `json:"interface_a,omitempty"`
	Interfaceb   *int64            `json:"interface_b,omitempty"`
	Ssid         *string           `json:"ssid,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	AuthType     *string           `json:"auth_type,omitempty"`
	AuthCipher   *string           `json:"auth_cipher,omitempty"`
	AuthPsk      *string           `json:"auth_psk,omitempty"`
	Distance     *float64          `json:"distance,omitempty"`
	DistanceUnit *string           `json:"distance_unit,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// WirelessLinkResponseDTOFromGoNetbox converts a *models.WirelessLink to the response DTO.
func WirelessLinkResponseDTOFromGoNetbox(goNetboxModel *models.WirelessLink) *WirelessLinkResponseDTO {
	responseDTO := &WirelessLinkResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Interfacea != nil {
		v := goNetboxModel.Interfacea.ID
		responseDTO.Interfacea = &v
	}
	if goNetboxModel.Interfaceb != nil {
		v := goNetboxModel.Interfaceb.ID
		responseDTO.Interfaceb = &v
	}
	if goNetboxModel.Ssid != "" {
		v := goNetboxModel.Ssid
		responseDTO.Ssid = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
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
	responseDTO.Distance = goNetboxModel.Distance
	if goNetboxModel.DistanceUnit != nil {
		responseDTO.DistanceUnit = choiceValue[string](goNetboxModel.DistanceUnit.Value)
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
