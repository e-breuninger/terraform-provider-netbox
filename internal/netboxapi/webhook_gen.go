// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// WebhookRequestDTO is the request DTO of the webhook resource; Payload writes it into
// the JSON request body.
type WebhookRequestDTO struct {
	Name              *string           `json:"name,omitempty"`
	PayloadURL        *string           `json:"payload_url,omitempty"`
	HTTPMethod        *string           `json:"http_method,omitempty"`
	HTTPContentType   *string           `json:"http_content_type,omitempty"`
	AdditionalHeaders *string           `json:"additional_headers,omitempty"`
	BodyTemplate      *string           `json:"body_template,omitempty"`
	Secret            *string           `json:"secret,omitempty"`
	SslVerification   *bool             `json:"ssl_verification,omitempty"`
	CaFilePath        *string           `json:"ca_file_path,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Owner             *int64            `json:"owner,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the webhook resource.
func (requestDTO *WebhookRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.PayloadURL != nil {
		payload["payload_url"] = *requestDTO.PayloadURL
	}
	if requestDTO.HTTPMethod != nil {
		payload["http_method"] = *requestDTO.HTTPMethod
	}
	if requestDTO.HTTPContentType != nil {
		payload["http_content_type"] = *requestDTO.HTTPContentType
	}
	if requestDTO.AdditionalHeaders != nil {
		payload["additional_headers"] = *requestDTO.AdditionalHeaders
	} else {
		payload["additional_headers"] = ""
	}
	if requestDTO.BodyTemplate != nil {
		payload["body_template"] = *requestDTO.BodyTemplate
	} else {
		payload["body_template"] = ""
	}
	if requestDTO.Secret != nil {
		payload["secret"] = *requestDTO.Secret
	} else {
		payload["secret"] = ""
	}
	if requestDTO.SslVerification != nil {
		payload["ssl_verification"] = *requestDTO.SslVerification
	}
	if requestDTO.CaFilePath != nil {
		payload["ca_file_path"] = *requestDTO.CaFilePath
	} else {
		payload["ca_file_path"] = nil
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
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// WebhookResponseDTO is the response DTO of the webhook resource, built from
// the go-netbox Webhook by WebhookResponseDTOFromGoNetbox.
type WebhookResponseDTO struct {
	ID                *int64            `json:"id,omitempty"`
	Name              *string           `json:"name,omitempty"`
	PayloadURL        *string           `json:"payload_url,omitempty"`
	HTTPMethod        *string           `json:"http_method,omitempty"`
	HTTPContentType   *string           `json:"http_content_type,omitempty"`
	AdditionalHeaders *string           `json:"additional_headers,omitempty"`
	BodyTemplate      *string           `json:"body_template,omitempty"`
	Secret            *string           `json:"secret,omitempty"`
	SslVerification   *bool             `json:"ssl_verification,omitempty"`
	CaFilePath        *string           `json:"ca_file_path,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Owner             *int64            `json:"owner,omitempty"`
	Created           *string           `json:"created,omitempty"`
	LastUpdated       *string           `json:"last_updated,omitempty"`
	URL               *string           `json:"url,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	TagsAll           []string          `json:"tags_all,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
}

// WebhookResponseDTOFromGoNetbox converts a *models.Webhook to the response DTO.
func WebhookResponseDTOFromGoNetbox(goNetboxModel *models.Webhook) *WebhookResponseDTO {
	responseDTO := &WebhookResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.PayloadURL = goNetboxModel.PayloadURL
	if goNetboxModel.HTTPMethod != "" {
		v := goNetboxModel.HTTPMethod
		responseDTO.HTTPMethod = &v
	}
	if goNetboxModel.HTTPContentType != "" {
		v := goNetboxModel.HTTPContentType
		responseDTO.HTTPContentType = &v
	}
	if goNetboxModel.AdditionalHeaders != "" {
		v := goNetboxModel.AdditionalHeaders
		responseDTO.AdditionalHeaders = &v
	}
	if goNetboxModel.BodyTemplate != "" {
		v := goNetboxModel.BodyTemplate
		responseDTO.BodyTemplate = &v
	}
	if goNetboxModel.Secret != "" {
		v := goNetboxModel.Secret
		responseDTO.Secret = &v
	}
	{
		v := goNetboxModel.SslVerification
		responseDTO.SslVerification = &v
	}
	responseDTO.CaFilePath = goNetboxModel.CaFilePath
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
