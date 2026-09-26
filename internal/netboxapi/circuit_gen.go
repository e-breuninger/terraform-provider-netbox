// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitRequestDTO is the request DTO of the circuit resource; Payload writes it into
// the JSON request body.
type CircuitRequestDTO struct {
	Cid             *string           `json:"cid,omitempty"`
	Provider        *int64            `json:"provider,omitempty"`
	Type            *int64            `json:"type,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Tenant          *int64            `json:"tenant,omitempty"`
	InstallDate     *string           `json:"install_date,omitempty"`
	TerminationDate *string           `json:"termination_date,omitempty"`
	CommitRate      *int64            `json:"commit_rate,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Comments        *string           `json:"comments,omitempty"`
	Owner           *int64            `json:"owner,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	CustomFields    map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the circuit resource.
func (requestDTO *CircuitRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Cid != nil {
		payload["cid"] = *requestDTO.Cid
	}
	if requestDTO.Provider != nil {
		payload["provider"] = *requestDTO.Provider
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.InstallDate != nil {
		payload["install_date"] = *requestDTO.InstallDate
	}
	if requestDTO.TerminationDate != nil {
		payload["termination_date"] = *requestDTO.TerminationDate
	}
	if requestDTO.CommitRate != nil {
		payload["commit_rate"] = *requestDTO.CommitRate
	} else {
		payload["commit_rate"] = nil
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

// CircuitResponseDTO is the response DTO of the circuit resource, built from
// the go-netbox Circuit by CircuitResponseDTOFromGoNetbox.
type CircuitResponseDTO struct {
	ID              *int64            `json:"id,omitempty"`
	Cid             *string           `json:"cid,omitempty"`
	Provider        *int64            `json:"provider,omitempty"`
	Type            *int64            `json:"type,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Tenant          *int64            `json:"tenant,omitempty"`
	InstallDate     *string           `json:"install_date,omitempty"`
	TerminationDate *string           `json:"termination_date,omitempty"`
	CommitRate      *int64            `json:"commit_rate,omitempty"`
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

// CircuitResponseDTOFromGoNetbox converts a *models.Circuit to the response DTO.
func CircuitResponseDTOFromGoNetbox(goNetboxModel *models.Circuit) *CircuitResponseDTO {
	responseDTO := &CircuitResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Cid = goNetboxModel.Cid
	if goNetboxModel.Provider != nil {
		v := goNetboxModel.Provider.ID
		responseDTO.Provider = &v
	}
	if goNetboxModel.Type != nil {
		v := goNetboxModel.Type.ID
		responseDTO.Type = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.InstallDate != nil {
		v := goNetboxModel.InstallDate.String()
		responseDTO.InstallDate = &v
	}
	if goNetboxModel.TerminationDate != nil {
		v := goNetboxModel.TerminationDate.String()
		responseDTO.TerminationDate = &v
	}
	responseDTO.CommitRate = goNetboxModel.CommitRate
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
