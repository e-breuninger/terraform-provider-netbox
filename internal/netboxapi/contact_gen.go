// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ContactRequestDTO is the request DTO of the contact resource; Payload writes it into
// the JSON request body.
type ContactRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Title        *string           `json:"title,omitempty"`
	Phone        *string           `json:"phone,omitempty"`
	Email        *string           `json:"email,omitempty"`
	Address      *string           `json:"address,omitempty"`
	Link         *string           `json:"link,omitempty"`
	Groups       []int64           `json:"groups,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the contact resource.
func (requestDTO *ContactRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Title != nil {
		payload["title"] = *requestDTO.Title
	} else {
		payload["title"] = ""
	}
	if requestDTO.Phone != nil {
		payload["phone"] = *requestDTO.Phone
	} else {
		payload["phone"] = ""
	}
	if requestDTO.Email != nil {
		payload["email"] = *requestDTO.Email
	} else {
		payload["email"] = ""
	}
	if requestDTO.Address != nil {
		payload["address"] = *requestDTO.Address
	} else {
		payload["address"] = ""
	}
	if requestDTO.Link != nil {
		payload["link"] = *requestDTO.Link
	} else {
		payload["link"] = ""
	}
	if requestDTO.Groups != nil {
		payload["groups"] = requestDTO.Groups
	} else {
		payload["groups"] = []any{}
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

// ContactResponseDTO is the response DTO of the contact resource, built from
// the go-netbox Contact by ContactResponseDTOFromGoNetbox.
type ContactResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Title        *string           `json:"title,omitempty"`
	Phone        *string           `json:"phone,omitempty"`
	Email        *string           `json:"email,omitempty"`
	Address      *string           `json:"address,omitempty"`
	Link         *string           `json:"link,omitempty"`
	Groups       []int64           `json:"groups,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// ContactResponseDTOFromGoNetbox converts a *models.Contact to the response DTO.
func ContactResponseDTOFromGoNetbox(goNetboxModel *models.Contact) *ContactResponseDTO {
	responseDTO := &ContactResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Title != "" {
		v := goNetboxModel.Title
		responseDTO.Title = &v
	}
	if goNetboxModel.Phone != "" {
		v := goNetboxModel.Phone
		responseDTO.Phone = &v
	}
	if goNetboxModel.Email != "" {
		v := string(goNetboxModel.Email)
		responseDTO.Email = &v
	}
	if goNetboxModel.Address != "" {
		v := goNetboxModel.Address
		responseDTO.Address = &v
	}
	if goNetboxModel.Link != "" {
		v := string(goNetboxModel.Link)
		responseDTO.Link = &v
	}
	if goNetboxModel.Groups != nil {
		responseDTO.Groups = []int64{}
		for _, ref := range goNetboxModel.Groups {
			if ref != nil {
				responseDTO.Groups = append(responseDTO.Groups, ref.ID)
			}
		}
	}
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
