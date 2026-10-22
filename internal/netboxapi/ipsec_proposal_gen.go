// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpsecProposalRequestDTO is the request DTO of the ipsec_proposal resource; Payload writes it into
// the JSON request body.
type IpsecProposalRequestDTO struct {
	Name                    *string           `json:"name,omitempty"`
	EncryptionAlgorithm     *string           `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *string           `json:"authentication_algorithm,omitempty"`
	SaLifetimeSeconds       *int64            `json:"sa_lifetime_seconds,omitempty"`
	SaLifetimeData          *int64            `json:"sa_lifetime_data,omitempty"`
	Description             *string           `json:"description,omitempty"`
	Comments                *string           `json:"comments,omitempty"`
	Owner                   *int64            `json:"owner,omitempty"`
	Tags                    []string          `json:"tags,omitempty"`
	CustomFields            map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ipsec_proposal resource.
func (requestDTO *IpsecProposalRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.EncryptionAlgorithm != nil {
		payload["encryption_algorithm"] = *requestDTO.EncryptionAlgorithm
	}
	if requestDTO.AuthenticationAlgorithm != nil {
		payload["authentication_algorithm"] = *requestDTO.AuthenticationAlgorithm
	}
	if requestDTO.SaLifetimeSeconds != nil {
		payload["sa_lifetime_seconds"] = *requestDTO.SaLifetimeSeconds
	} else {
		payload["sa_lifetime_seconds"] = nil
	}
	if requestDTO.SaLifetimeData != nil {
		payload["sa_lifetime_data"] = *requestDTO.SaLifetimeData
	} else {
		payload["sa_lifetime_data"] = nil
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

// IpsecProposalResponseDTO is the response DTO of the ipsec_proposal resource, built from
// the go-netbox IPSecProposal by IpsecProposalResponseDTOFromGoNetbox.
type IpsecProposalResponseDTO struct {
	ID                      *int64            `json:"id,omitempty"`
	Name                    *string           `json:"name,omitempty"`
	EncryptionAlgorithm     *string           `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *string           `json:"authentication_algorithm,omitempty"`
	SaLifetimeSeconds       *int64            `json:"sa_lifetime_seconds,omitempty"`
	SaLifetimeData          *int64            `json:"sa_lifetime_data,omitempty"`
	Description             *string           `json:"description,omitempty"`
	Comments                *string           `json:"comments,omitempty"`
	Owner                   *int64            `json:"owner,omitempty"`
	Created                 *string           `json:"created,omitempty"`
	LastUpdated             *string           `json:"last_updated,omitempty"`
	URL                     *string           `json:"url,omitempty"`
	Tags                    []string          `json:"tags,omitempty"`
	TagsAll                 []string          `json:"tags_all,omitempty"`
	CustomFields            map[string]string `json:"custom_fields,omitempty"`
}

// IpsecProposalResponseDTOFromGoNetbox converts a *models.IPSecProposal to the response DTO.
func IpsecProposalResponseDTOFromGoNetbox(goNetboxModel *models.IPSecProposal) *IpsecProposalResponseDTO {
	responseDTO := &IpsecProposalResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.EncryptionAlgorithm != nil {
		responseDTO.EncryptionAlgorithm = choiceValue[string](goNetboxModel.EncryptionAlgorithm.Value)
	}
	if goNetboxModel.AuthenticationAlgorithm != nil {
		responseDTO.AuthenticationAlgorithm = choiceValue[string](goNetboxModel.AuthenticationAlgorithm.Value)
	}
	responseDTO.SaLifetimeSeconds = goNetboxModel.SaLifetimeSeconds
	responseDTO.SaLifetimeData = goNetboxModel.SaLifetimeData
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
