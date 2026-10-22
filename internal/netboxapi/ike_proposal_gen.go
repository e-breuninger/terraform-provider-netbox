// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// IkeProposalRequestDTO is the request DTO of the ike_proposal resource; Payload writes it into
// the JSON request body.
type IkeProposalRequestDTO struct {
	Name                    *string           `json:"name,omitempty"`
	AuthenticationMethod    *string           `json:"authentication_method,omitempty"`
	EncryptionAlgorithm     *string           `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *string           `json:"authentication_algorithm,omitempty"`
	Group                   *int64            `json:"group,omitempty"`
	SaLifetime              *int64            `json:"sa_lifetime,omitempty"`
	Description             *string           `json:"description,omitempty"`
	Comments                *string           `json:"comments,omitempty"`
	Owner                   *int64            `json:"owner,omitempty"`
	Tags                    []string          `json:"tags,omitempty"`
	CustomFields            map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the ike_proposal resource.
func (requestDTO *IkeProposalRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.AuthenticationMethod != nil {
		payload["authentication_method"] = *requestDTO.AuthenticationMethod
	}
	if requestDTO.EncryptionAlgorithm != nil {
		payload["encryption_algorithm"] = *requestDTO.EncryptionAlgorithm
	}
	if requestDTO.AuthenticationAlgorithm != nil {
		payload["authentication_algorithm"] = *requestDTO.AuthenticationAlgorithm
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	}
	if requestDTO.SaLifetime != nil {
		payload["sa_lifetime"] = *requestDTO.SaLifetime
	} else {
		payload["sa_lifetime"] = nil
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

// IkeProposalResponseDTO is the response DTO of the ike_proposal resource, built from
// the go-netbox IKEProposal by IkeProposalResponseDTOFromGoNetbox.
type IkeProposalResponseDTO struct {
	ID                      *int64            `json:"id,omitempty"`
	Name                    *string           `json:"name,omitempty"`
	AuthenticationMethod    *string           `json:"authentication_method,omitempty"`
	EncryptionAlgorithm     *string           `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *string           `json:"authentication_algorithm,omitempty"`
	Group                   *int64            `json:"group,omitempty"`
	SaLifetime              *int64            `json:"sa_lifetime,omitempty"`
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

// IkeProposalResponseDTOFromGoNetbox converts a *models.IKEProposal to the response DTO.
func IkeProposalResponseDTOFromGoNetbox(goNetboxModel *models.IKEProposal) *IkeProposalResponseDTO {
	responseDTO := &IkeProposalResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.AuthenticationMethod != nil {
		responseDTO.AuthenticationMethod = choiceValue[string](goNetboxModel.AuthenticationMethod.Value)
	}
	if goNetboxModel.EncryptionAlgorithm != nil {
		responseDTO.EncryptionAlgorithm = choiceValue[string](goNetboxModel.EncryptionAlgorithm.Value)
	}
	if goNetboxModel.AuthenticationAlgorithm != nil {
		responseDTO.AuthenticationAlgorithm = choiceValue[string](goNetboxModel.AuthenticationAlgorithm.Value)
	}
	if goNetboxModel.Group != nil {
		responseDTO.Group = choiceValue[int64](goNetboxModel.Group.Value)
	}
	responseDTO.SaLifetime = goNetboxModel.SaLifetime
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
