// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DataSourceRequestDTO is the request DTO of the data_source resource; Payload writes it into
// the JSON request body.
type DataSourceRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Type         *string           `json:"type,omitempty"`
	SourceURL    *string           `json:"source_url,omitempty"`
	Enabled      *bool             `json:"enabled,omitempty"`
	IgnoreRules  *string           `json:"ignore_rules,omitempty"`
	Parameters   *string           `json:"parameters,omitempty"`
	SyncInterval *int64            `json:"sync_interval,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the data_source resource.
func (requestDTO *DataSourceRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.SourceURL != nil {
		payload["source_url"] = *requestDTO.SourceURL
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.IgnoreRules != nil {
		payload["ignore_rules"] = *requestDTO.IgnoreRules
	} else {
		payload["ignore_rules"] = ""
	}
	if requestDTO.Parameters != nil {
		payload["parameters"] = parseJSONText(*requestDTO.Parameters)
	} else {
		payload["parameters"] = nil
	}
	if requestDTO.SyncInterval != nil {
		payload["sync_interval"] = *requestDTO.SyncInterval
	} else {
		payload["sync_interval"] = nil
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
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// DataSourceResponseDTO is the response DTO of the data_source resource, built from
// the go-netbox DataSource by DataSourceResponseDTOFromGoNetbox.
type DataSourceResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Type         *string           `json:"type,omitempty"`
	SourceURL    *string           `json:"source_url,omitempty"`
	Enabled      *bool             `json:"enabled,omitempty"`
	IgnoreRules  *string           `json:"ignore_rules,omitempty"`
	Parameters   *string           `json:"parameters,omitempty"`
	SyncInterval *int64            `json:"sync_interval,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Status       *string           `json:"status,omitempty"`
	LastSynced   *string           `json:"last_synced,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	FileCount    *int64            `json:"file_count,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// DataSourceResponseDTOFromGoNetbox converts a *models.DataSource to the response DTO.
func DataSourceResponseDTOFromGoNetbox(goNetboxModel *models.DataSource) *DataSourceResponseDTO {
	responseDTO := &DataSourceResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
	responseDTO.SourceURL = goNetboxModel.SourceURL
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	if goNetboxModel.IgnoreRules != "" {
		v := goNetboxModel.IgnoreRules
		responseDTO.IgnoreRules = &v
	}
	responseDTO.Parameters = jsonText(goNetboxModel.Parameters)
	responseDTO.SyncInterval = goNetboxModel.SyncInterval
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Comments != "" {
		v := goNetboxModel.Comments
		responseDTO.Comments = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.LastSynced != nil {
		v := goNetboxModel.LastSynced.String()
		responseDTO.LastSynced = &v
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
	{
		v := goNetboxModel.FileCount
		responseDTO.FileCount = &v
	}
	responseDTO.CustomFields = customFieldValues(goNetboxModel.CustomFields)
	return responseDTO
}
