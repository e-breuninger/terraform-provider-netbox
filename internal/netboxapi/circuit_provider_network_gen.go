// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitProviderNetworkRequestDTO is the request DTO of the circuit_provider_network resource; Payload writes it into
// the JSON request body.
type CircuitProviderNetworkRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Provider     *int64            `json:"provider,omitempty"`
	ServiceID    *string           `json:"service_id,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the circuit_provider_network resource.
func (requestDTO *CircuitProviderNetworkRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Provider != nil {
		payload["provider"] = *requestDTO.Provider
	}
	if requestDTO.ServiceID != nil {
		payload["service_id"] = *requestDTO.ServiceID
	} else {
		payload["service_id"] = ""
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

// CircuitProviderNetworkResponseDTO is the response DTO of the circuit_provider_network resource, built from
// the go-netbox ProviderNetwork by CircuitProviderNetworkResponseDTOFromGoNetbox.
type CircuitProviderNetworkResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Provider     *int64            `json:"provider,omitempty"`
	ServiceID    *string           `json:"service_id,omitempty"`
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

// CircuitProviderNetworkResponseDTOFromGoNetbox converts a *models.ProviderNetwork to the response DTO.
func CircuitProviderNetworkResponseDTOFromGoNetbox(goNetboxModel *models.ProviderNetwork) *CircuitProviderNetworkResponseDTO {
	responseDTO := &CircuitProviderNetworkResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Provider != nil {
		v := goNetboxModel.Provider.ID
		responseDTO.Provider = &v
	}
	if goNetboxModel.ServiceID != "" {
		v := goNetboxModel.ServiceID
		responseDTO.ServiceID = &v
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
