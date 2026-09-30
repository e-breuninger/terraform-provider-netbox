// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpsecPolicyRequestDTO is the request DTO of the ipsec_policy resource; Payload writes it into
// the JSON request body.
type IpsecPolicyRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Proposals    []int64           `json:"proposals,omitempty"`
	PfsGroup     *int64            `json:"pfs_group,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ipsec_policy resource.
func (requestDTO *IpsecPolicyRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Proposals != nil {
		payload["proposals"] = requestDTO.Proposals
	}
	if requestDTO.PfsGroup != nil {
		payload["pfs_group"] = *requestDTO.PfsGroup
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

// IpsecPolicyResponseDTO is the response DTO of the ipsec_policy resource, built from
// the go-netbox IPSecPolicy by IpsecPolicyResponseDTOFromGoNetbox.
type IpsecPolicyResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Proposals    []int64           `json:"proposals,omitempty"`
	PfsGroup     *int64            `json:"pfs_group,omitempty"`
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

// IpsecPolicyResponseDTOFromGoNetbox converts a *models.IPSecPolicy to the response DTO.
func IpsecPolicyResponseDTOFromGoNetbox(goNetboxModel *models.IPSecPolicy) *IpsecPolicyResponseDTO {
	responseDTO := &IpsecPolicyResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Proposals != nil {
		responseDTO.Proposals = []int64{}
		for _, ref := range goNetboxModel.Proposals {
			if ref != nil {
				responseDTO.Proposals = append(responseDTO.Proposals, ref.ID)
			}
		}
	}
	if goNetboxModel.PfsGroup != nil {
		responseDTO.PfsGroup = choiceValue[int64](goNetboxModel.PfsGroup.Value)
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
