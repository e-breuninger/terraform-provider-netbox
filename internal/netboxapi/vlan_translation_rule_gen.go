// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VlanTranslationRuleRequestDTO is the request DTO of the vlan_translation_rule resource; Payload writes it into
// the JSON request body.
type VlanTranslationRuleRequestDTO struct {
	Policy      *int64  `json:"policy,omitempty"`
	LocalVid    *int64  `json:"local_vid,omitempty"`
	RemoteVid   *int64  `json:"remote_vid,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the vlan_translation_rule resource.
func (requestDTO *VlanTranslationRuleRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Policy != nil {
		payload["policy"] = *requestDTO.Policy
	}
	if requestDTO.LocalVid != nil {
		payload["local_vid"] = *requestDTO.LocalVid
	}
	if requestDTO.RemoteVid != nil {
		payload["remote_vid"] = *requestDTO.RemoteVid
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	return payload
}

// VlanTranslationRuleResponseDTO is the response DTO of the vlan_translation_rule resource, built from
// the go-netbox VLANTranslationRule by VlanTranslationRuleResponseDTOFromGoNetbox.
type VlanTranslationRuleResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	Policy      *int64  `json:"policy,omitempty"`
	LocalVid    *int64  `json:"local_vid,omitempty"`
	RemoteVid   *int64  `json:"remote_vid,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// VlanTranslationRuleResponseDTOFromGoNetbox converts a *models.VLANTranslationRule to the response DTO.
func VlanTranslationRuleResponseDTOFromGoNetbox(goNetboxModel *models.VLANTranslationRule) *VlanTranslationRuleResponseDTO {
	responseDTO := &VlanTranslationRuleResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Policy = goNetboxModel.Policy
	responseDTO.LocalVid = goNetboxModel.LocalVid
	responseDTO.RemoteVid = goNetboxModel.RemoteVid
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
