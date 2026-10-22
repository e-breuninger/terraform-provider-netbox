// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitProviderAccountRequestDTO is the request DTO of the circuit_provider_account resource; Payload writes it into
// the JSON request body.
type CircuitProviderAccountRequestDTO struct {
	Account      *string           `json:"account,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Provider     *int64            `json:"provider,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the circuit_provider_account resource.
func (requestDTO *CircuitProviderAccountRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Account != nil {
		payload["account"] = *requestDTO.Account
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	} else {
		payload["name"] = ""
	}
	if requestDTO.Provider != nil {
		payload["provider"] = *requestDTO.Provider
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

// CircuitProviderAccountResponseDTO is the response DTO of the circuit_provider_account resource, built from
// the go-netbox ProviderAccount by CircuitProviderAccountResponseDTOFromGoNetbox.
type CircuitProviderAccountResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Account      *string           `json:"account,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Provider     *int64            `json:"provider,omitempty"`
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

// CircuitProviderAccountResponseDTOFromGoNetbox converts a *models.ProviderAccount to the response DTO.
func CircuitProviderAccountResponseDTOFromGoNetbox(goNetboxModel *models.ProviderAccount) *CircuitProviderAccountResponseDTO {
	responseDTO := &CircuitProviderAccountResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Account = goNetboxModel.Account
	if goNetboxModel.Name != "" {
		v := goNetboxModel.Name
		responseDTO.Name = &v
	}
	if goNetboxModel.Provider != nil {
		v := goNetboxModel.Provider.ID
		responseDTO.Provider = &v
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
