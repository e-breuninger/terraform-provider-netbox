// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ConfigTemplateRequestDTO is the request DTO of the config_template resource; Payload writes it into
// the JSON request body.
type ConfigTemplateRequestDTO struct {
	Name              *string  `json:"name,omitempty"`
	Description       *string  `json:"description,omitempty"`
	TemplateCode      *string  `json:"template_code,omitempty"`
	EnvironmentParams *string  `json:"environment_params,omitempty"`
	MimeType          *string  `json:"mime_type,omitempty"`
	FileName          *string  `json:"file_name,omitempty"`
	FileExtension     *string  `json:"file_extension,omitempty"`
	AsAttachment      *bool    `json:"as_attachment,omitempty"`
	Debug             *bool    `json:"debug,omitempty"`
	Owner             *int64   `json:"owner,omitempty"`
	Tags              []string `json:"tags,omitempty"`
}

// Payload returns the JSON request body of the config_template resource.
func (requestDTO *ConfigTemplateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.TemplateCode != nil {
		payload["template_code"] = *requestDTO.TemplateCode
	}
	if requestDTO.EnvironmentParams != nil {
		payload["environment_params"] = parseJSONText(*requestDTO.EnvironmentParams)
	} else {
		payload["environment_params"] = nil
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
	if requestDTO.Debug != nil {
		payload["debug"] = *requestDTO.Debug
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	return payload
}

// ConfigTemplateResponseDTO is the response DTO of the config_template resource, built from
// the go-netbox ConfigTemplate by ConfigTemplateResponseDTOFromGoNetbox.
type ConfigTemplateResponseDTO struct {
	ID                *int64   `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	Description       *string  `json:"description,omitempty"`
	TemplateCode      *string  `json:"template_code,omitempty"`
	EnvironmentParams *string  `json:"environment_params,omitempty"`
	MimeType          *string  `json:"mime_type,omitempty"`
	FileName          *string  `json:"file_name,omitempty"`
	FileExtension     *string  `json:"file_extension,omitempty"`
	AsAttachment      *bool    `json:"as_attachment,omitempty"`
	Debug             *bool    `json:"debug,omitempty"`
	Owner             *int64   `json:"owner,omitempty"`
	Created           *string  `json:"created,omitempty"`
	LastUpdated       *string  `json:"last_updated,omitempty"`
	URL               *string  `json:"url,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	TagsAll           []string `json:"tags_all,omitempty"`
}

// ConfigTemplateResponseDTOFromGoNetbox converts a *models.ConfigTemplate to the response DTO.
func ConfigTemplateResponseDTOFromGoNetbox(goNetboxModel *models.ConfigTemplate) *ConfigTemplateResponseDTO {
	responseDTO := &ConfigTemplateResponseDTO{}
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
	responseDTO.TemplateCode = goNetboxModel.TemplateCode
	responseDTO.EnvironmentParams = jsonText(goNetboxModel.EnvironmentParams)
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
	{
		v := goNetboxModel.Debug
		responseDTO.Debug = &v
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
	return responseDTO
}
