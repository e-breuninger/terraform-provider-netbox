// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualCircuitRequestDTO is the request DTO of the virtual_circuit resource; Payload writes it into
// the JSON request body.
type VirtualCircuitRequestDTO struct {
	Cid             *string           `json:"cid,omitempty"`
	ProviderNetwork *int64            `json:"provider_network,omitempty"`
	Type            *int64            `json:"type,omitempty"`
	ProviderAccount *int64            `json:"provider_account,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Tenant          *int64            `json:"tenant,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Comments        *string           `json:"comments,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_circuit resource.
func (requestDTO *VirtualCircuitRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Cid != nil {
		payload["cid"] = *requestDTO.Cid
	}
	if requestDTO.ProviderNetwork != nil {
		payload["provider_network"] = *requestDTO.ProviderNetwork
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.ProviderAccount != nil {
		payload["provider_account"] = *requestDTO.ProviderAccount
	} else {
		payload["provider_account"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
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

// VirtualCircuitResponseDTO is the response DTO of the virtual_circuit resource, built from
// the go-netbox VirtualCircuit by VirtualCircuitResponseDTOFromGoNetbox.
type VirtualCircuitResponseDTO struct {
	ID              *int64            `json:"id,omitempty"`
	Cid             *string           `json:"cid,omitempty"`
	ProviderNetwork *int64            `json:"provider_network,omitempty"`
	Type            *int64            `json:"type,omitempty"`
	ProviderAccount *int64            `json:"provider_account,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Tenant          *int64            `json:"tenant,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Comments        *string           `json:"comments,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Created         *string           `json:"created,omitempty"`
	LastUpdated     *string           `json:"last_updated,omitempty"`
	URL             *string           `json:"url,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	TagsAll         []string          `json:"tags_all,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// VirtualCircuitResponseDTOFromGoNetbox converts a *models.VirtualCircuit to the response DTO.
func VirtualCircuitResponseDTOFromGoNetbox(goNetboxModel *models.VirtualCircuit) *VirtualCircuitResponseDTO {
	responseDTO := &VirtualCircuitResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Cid = goNetboxModel.Cid
	if goNetboxModel.ProviderNetwork != nil {
		v := goNetboxModel.ProviderNetwork.ID
		responseDTO.ProviderNetwork = &v
	}
	if goNetboxModel.Type != nil {
		v := goNetboxModel.Type.ID
		responseDTO.Type = &v
	}
	if goNetboxModel.ProviderAccount != nil {
		v := goNetboxModel.ProviderAccount.ID
		responseDTO.ProviderAccount = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
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
