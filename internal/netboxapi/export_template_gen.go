// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExportTemplateRequestDTO is the request DTO of the export_template resource; Payload writes it into
// the JSON request body.
type ExportTemplateRequestDTO struct {
	Name              *string  `json:"name,omitempty"`
	ObjectTypes       []string `json:"object_types,omitempty"`
	TemplateCode      *string  `json:"template_code,omitempty"`
	MimeType          *string  `json:"mime_type,omitempty"`
	FileName          *string  `json:"file_name,omitempty"`
	FileExtension     *string  `json:"file_extension,omitempty"`
	AsAttachment      *bool    `json:"as_attachment,omitempty"`
	EnvironmentParams *string  `json:"environment_params,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Owner             *int64   `json:"owner,omitempty"`
}

// Payload returns the JSON request body of the export_template resource.
func (requestDTO *ExportTemplateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.ObjectTypes != nil {
		payload["object_types"] = requestDTO.ObjectTypes
	}
	if requestDTO.TemplateCode != nil {
		payload["template_code"] = *requestDTO.TemplateCode
	}
	if requestDTO.MimeType != nil {
		payload["mime_type"] = *requestDTO.MimeType
	} else {
		payload["mime_type"] = ""
	}
	if requestDTO.FileName != nil {
		payload["file_name"] = *requestDTO.FileName
	} else {
		payload["file_name"] = ""
	}
	if requestDTO.FileExtension != nil {
		payload["file_extension"] = *requestDTO.FileExtension
	} else {
		payload["file_extension"] = ""
	}
	if requestDTO.AsAttachment != nil {
		payload["as_attachment"] = *requestDTO.AsAttachment
	}
	if requestDTO.EnvironmentParams != nil {
		payload["environment_params"] = parseJSONText(*requestDTO.EnvironmentParams)
	} else {
		payload["environment_params"] = nil
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	return payload
}

// ExportTemplateResponseDTO is the response DTO of the export_template resource, built from
// the go-netbox ExportTemplate by ExportTemplateResponseDTOFromGoNetbox.
type ExportTemplateResponseDTO struct {
	ID                *int64   `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	ObjectTypes       []string `json:"object_types,omitempty"`
	TemplateCode      *string  `json:"template_code,omitempty"`
	MimeType          *string  `json:"mime_type,omitempty"`
	FileName          *string  `json:"file_name,omitempty"`
	FileExtension     *string  `json:"file_extension,omitempty"`
	AsAttachment      *bool    `json:"as_attachment,omitempty"`
	EnvironmentParams *string  `json:"environment_params,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Owner             *int64   `json:"owner,omitempty"`
	Created           *string  `json:"created,omitempty"`
	LastUpdated       *string  `json:"last_updated,omitempty"`
	URL               *string  `json:"url,omitempty"`
}

// ExportTemplateResponseDTOFromGoNetbox converts a *models.ExportTemplate to the response DTO.
func ExportTemplateResponseDTOFromGoNetbox(goNetboxModel *models.ExportTemplate) *ExportTemplateResponseDTO {
	responseDTO := &ExportTemplateResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.ObjectTypes = goNetboxModel.ObjectTypes
	responseDTO.TemplateCode = goNetboxModel.TemplateCode
	if goNetboxModel.MimeType != "" {
		v := goNetboxModel.MimeType
		responseDTO.MimeType = &v
	}
	if goNetboxModel.FileName != "" {
		v := goNetboxModel.FileName
		responseDTO.FileName = &v
	}
	if goNetboxModel.FileExtension != "" {
		v := goNetboxModel.FileExtension
		responseDTO.FileExtension = &v
	}
	{
		v := goNetboxModel.AsAttachment
		responseDTO.AsAttachment = &v
	}
	responseDTO.EnvironmentParams = jsonText(goNetboxModel.EnvironmentParams)
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
	return responseDTO
}
