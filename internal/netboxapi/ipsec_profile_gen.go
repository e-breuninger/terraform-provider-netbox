// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpsecProfileRequestDTO is the request DTO of the ipsec_profile resource; Payload writes it into
// the JSON request body.
type IpsecProfileRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Mode         *string           `json:"mode,omitempty"`
	IkePolicy    *int64            `json:"ike_policy,omitempty"`
	IpsecPolicy  *int64            `json:"ipsec_policy,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ipsec_profile resource.
func (requestDTO *IpsecProfileRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Mode != nil {
		payload["mode"] = *requestDTO.Mode
	}
	if requestDTO.IkePolicy != nil {
		payload["ike_policy"] = *requestDTO.IkePolicy
	}
	if requestDTO.IpsecPolicy != nil {
		payload["ipsec_policy"] = *requestDTO.IpsecPolicy
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

// IpsecProfileResponseDTO is the response DTO of the ipsec_profile resource, built from
// the go-netbox IPSecProfile by IpsecProfileResponseDTOFromGoNetbox.
type IpsecProfileResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Mode         *string           `json:"mode,omitempty"`
	IkePolicy    *int64            `json:"ike_policy,omitempty"`
	IpsecPolicy  *int64            `json:"ipsec_policy,omitempty"`
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

// IpsecProfileResponseDTOFromGoNetbox converts a *models.IPSecProfile to the response DTO.
func IpsecProfileResponseDTOFromGoNetbox(goNetboxModel *models.IPSecProfile) *IpsecProfileResponseDTO {
	responseDTO := &IpsecProfileResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Mode != nil {
		responseDTO.Mode = choiceValue[string](goNetboxModel.Mode.Value)
	}
	if goNetboxModel.IkePolicy != nil {
		v := goNetboxModel.IkePolicy.ID
		responseDTO.IkePolicy = &v
	}
	if goNetboxModel.IpsecPolicy != nil {
		v := goNetboxModel.IpsecPolicy.ID
		responseDTO.IpsecPolicy = &v
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
