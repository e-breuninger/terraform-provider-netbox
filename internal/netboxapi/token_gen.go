// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// TokenRequestDTO is the request DTO of the token resource; Payload writes it into
// the JSON request body.
type TokenRequestDTO struct {
	User         *int64  `json:"user,omitempty"`
	Token        *string `json:"token,omitempty"`
	Enabled      *bool   `json:"enabled,omitempty"`
	Version      *int64  `json:"version,omitempty"`
	Description  *string `json:"description,omitempty"`
	WriteEnabled *bool   `json:"write_enabled,omitempty"`
	Expires      *string `json:"expires,omitempty"`
}

// Payload returns the JSON request body of the token resource.
func (requestDTO *TokenRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.User != nil {
		payload["user"] = *requestDTO.User
	}
	if requestDTO.Token != nil {
		payload["token"] = *requestDTO.Token
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.Version != nil {
		payload["version"] = *requestDTO.Version
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.WriteEnabled != nil {
		payload["write_enabled"] = *requestDTO.WriteEnabled
	}
	if requestDTO.Expires != nil {
		payload["expires"] = *requestDTO.Expires
	}
	return payload
}

// TokenResponseDTO is the response DTO of the token resource, built from
// the go-netbox Token by TokenResponseDTOFromGoNetbox.
type TokenResponseDTO struct {
	ID           *int64  `json:"id,omitempty"`
	User         *int64  `json:"user,omitempty"`
	Key          *string `json:"key,omitempty"`
	Token        *string `json:"token,omitempty"`
	Enabled      *bool   `json:"enabled,omitempty"`
	Version      *int64  `json:"version,omitempty"`
	PepperID     *int64  `json:"pepper_id,omitempty"`
	Description  *string `json:"description,omitempty"`
	WriteEnabled *bool   `json:"write_enabled,omitempty"`
	Expires      *string `json:"expires,omitempty"`
	LastUsed     *string `json:"last_used,omitempty"`
	Created      *string `json:"created,omitempty"`
	URL          *string `json:"url,omitempty"`
}

// TokenResponseDTOFromGoNetbox converts a *models.Token to the response DTO.
func TokenResponseDTOFromGoNetbox(goNetboxModel *models.Token) *TokenResponseDTO {
	responseDTO := &TokenResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.User != nil {
		v := goNetboxModel.User.ID
		responseDTO.User = &v
	}
	responseDTO.Key = goNetboxModel.Key
	if goNetboxModel.Token != "" {
		v := goNetboxModel.Token
		responseDTO.Token = &v
	}
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	responseDTO.Version = goNetboxModel.Version
	responseDTO.PepperID = goNetboxModel.PepperID
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	{
		v := goNetboxModel.WriteEnabled
		responseDTO.WriteEnabled = &v
	}
	if goNetboxModel.Expires != nil {
		v := goNetboxModel.Expires.String()
		responseDTO.Expires = &v
	}
	if goNetboxModel.LastUsed != nil {
		v := goNetboxModel.LastUsed.String()
		responseDTO.LastUsed = &v
	}
	if !goNetboxModel.Created.IsZero() {
		v := goNetboxModel.Created.String()
		responseDTO.Created = &v
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
