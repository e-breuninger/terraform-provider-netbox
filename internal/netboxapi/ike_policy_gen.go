// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IkePolicyRequestDTO is the request DTO of the ike_policy resource; Payload writes it into
// the JSON request body.
type IkePolicyRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Version      *int64            `json:"version,omitempty"`
	Mode         *string           `json:"mode,omitempty"`
	Proposals    []int64           `json:"proposals,omitempty"`
	PresharedKey *string           `json:"preshared_key,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ike_policy resource.
func (requestDTO *IkePolicyRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Version != nil {
		payload["version"] = *requestDTO.Version
	}
	if requestDTO.Mode != nil {
		payload["mode"] = *requestDTO.Mode
	}
	if requestDTO.Proposals != nil {
		payload["proposals"] = requestDTO.Proposals
	}
	if requestDTO.PresharedKey != nil {
		payload["preshared_key"] = *requestDTO.PresharedKey
	} else {
		payload["preshared_key"] = ""
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

// IkePolicyResponseDTO is the response DTO of the ike_policy resource, built from
// the go-netbox IKEPolicy by IkePolicyResponseDTOFromGoNetbox.
type IkePolicyResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Version      *int64            `json:"version,omitempty"`
	Mode         *string           `json:"mode,omitempty"`
	Proposals    []int64           `json:"proposals,omitempty"`
	PresharedKey *string           `json:"preshared_key,omitempty"`
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

// IkePolicyResponseDTOFromGoNetbox converts a *models.IKEPolicy to the response DTO.
func IkePolicyResponseDTOFromGoNetbox(goNetboxModel *models.IKEPolicy) *IkePolicyResponseDTO {
	responseDTO := &IkePolicyResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Version != nil {
		responseDTO.Version = choiceValue[int64](goNetboxModel.Version.Value)
	}
	if goNetboxModel.Mode != nil {
		responseDTO.Mode = choiceValue[string](goNetboxModel.Mode.Value)
	}
	if goNetboxModel.Proposals != nil {
		responseDTO.Proposals = []int64{}
		for _, ref := range goNetboxModel.Proposals {
			if ref != nil {
				responseDTO.Proposals = append(responseDTO.Proposals, ref.ID)
			}
		}
	}
	if goNetboxModel.PresharedKey != "" {
		v := goNetboxModel.PresharedKey
		responseDTO.PresharedKey = &v
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
