// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CustomLinkRequestDTO is the request DTO of the custom_link resource; Payload writes it into
// the JSON request body.
type CustomLinkRequestDTO struct {
	Name        *string  `json:"name,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	LinkText    *string  `json:"link_text,omitempty"`
	LinkURL     *string  `json:"link_url,omitempty"`
	Weight      *int64   `json:"weight,omitempty"`
	GroupName   *string  `json:"group_name,omitempty"`
	ButtonClass *string  `json:"button_class,omitempty"`
	NewWindow   *bool    `json:"new_window,omitempty"`
	Owner       *int64   `json:"owner,omitempty"`
}

// Payload returns the JSON request body of the custom_link resource.
func (requestDTO *CustomLinkRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.ObjectTypes != nil {
		payload["object_types"] = requestDTO.ObjectTypes
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.LinkText != nil {
		payload["link_text"] = *requestDTO.LinkText
	}
	if requestDTO.LinkURL != nil {
		payload["link_url"] = *requestDTO.LinkURL
	}
	if requestDTO.Weight != nil {
		payload["weight"] = *requestDTO.Weight
	}
	if requestDTO.GroupName != nil {
		payload["group_name"] = *requestDTO.GroupName
	} else {
		payload["group_name"] = ""
	}
	if requestDTO.ButtonClass != nil {
		payload["button_class"] = *requestDTO.ButtonClass
	}
	if requestDTO.NewWindow != nil {
		payload["new_window"] = *requestDTO.NewWindow
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	return payload
}

// CustomLinkResponseDTO is the response DTO of the custom_link resource, built from
// the go-netbox CustomLink by CustomLinkResponseDTOFromGoNetbox.
type CustomLinkResponseDTO struct {
	ID          *int64   `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	LinkText    *string  `json:"link_text,omitempty"`
	LinkURL     *string  `json:"link_url,omitempty"`
	Weight      *int64   `json:"weight,omitempty"`
	GroupName   *string  `json:"group_name,omitempty"`
	ButtonClass *string  `json:"button_class,omitempty"`
	NewWindow   *bool    `json:"new_window,omitempty"`
	Owner       *int64   `json:"owner,omitempty"`
	Created     *string  `json:"created,omitempty"`
	LastUpdated *string  `json:"last_updated,omitempty"`
	URL         *string  `json:"url,omitempty"`
}

// CustomLinkResponseDTOFromGoNetbox converts a *models.CustomLink to the response DTO.
func CustomLinkResponseDTOFromGoNetbox(goNetboxModel *models.CustomLink) *CustomLinkResponseDTO {
	responseDTO := &CustomLinkResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.ObjectTypes = goNetboxModel.ObjectTypes
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	responseDTO.LinkText = goNetboxModel.LinkText
	responseDTO.LinkURL = goNetboxModel.LinkURL
	responseDTO.Weight = goNetboxModel.Weight
	if goNetboxModel.GroupName != "" {
		v := goNetboxModel.GroupName
		responseDTO.GroupName = &v
	}
	if goNetboxModel.ButtonClass != "" {
		v := goNetboxModel.ButtonClass
		responseDTO.ButtonClass = &v
	}
	{
		v := goNetboxModel.NewWindow
		responseDTO.NewWindow = &v
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
