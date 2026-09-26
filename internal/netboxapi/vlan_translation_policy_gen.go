// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VlanTranslationPolicyRequestDTO is the request DTO of the vlan_translation_policy resource; Payload writes it into
// the JSON request body.
type VlanTranslationPolicyRequestDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments    *string `json:"comments,omitempty"`
	Owner       *int64  `json:"owner,omitempty"`
}

// Payload returns the JSON request body of the vlan_translation_policy resource.
func (requestDTO *VlanTranslationPolicyRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
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
	return payload
}

// VlanTranslationPolicyResponseDTO is the response DTO of the vlan_translation_policy resource, built from
// the go-netbox VLANTranslationPolicy by VlanTranslationPolicyResponseDTOFromGoNetbox.
type VlanTranslationPolicyResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments    *string `json:"comments,omitempty"`
	Owner       *int64  `json:"owner,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// VlanTranslationPolicyResponseDTOFromGoNetbox converts a *models.VLANTranslationPolicy to the response DTO.
func VlanTranslationPolicyResponseDTOFromGoNetbox(goNetboxModel *models.VLANTranslationPolicy) *VlanTranslationPolicyResponseDTO {
	responseDTO := &VlanTranslationPolicyResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
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
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
