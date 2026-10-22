// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// PowerPanelRequestDTO is the request DTO of the power_panel resource; Payload writes it into
// the JSON request body.
type PowerPanelRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Site         *int64            `json:"site,omitempty"`
	Location     *int64            `json:"location,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the power_panel resource.
func (requestDTO *PowerPanelRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Site != nil {
		payload["site"] = *requestDTO.Site
	}
	if requestDTO.Location != nil {
		payload["location"] = *requestDTO.Location
	} else {
		payload["location"] = nil
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

// PowerPanelResponseDTO is the response DTO of the power_panel resource, built from
// the go-netbox PowerPanel by PowerPanelResponseDTOFromGoNetbox.
type PowerPanelResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Site           *int64            `json:"site,omitempty"`
	Location       *int64            `json:"location,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Comments       *string           `json:"comments,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Created        *string           `json:"created,omitempty"`
	LastUpdated    *string           `json:"last_updated,omitempty"`
	URL            *string           `json:"url,omitempty"`
	PowerfeedCount *int64            `json:"powerfeed_count,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	TagsAll        []string          `json:"tags_all,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// PowerPanelResponseDTOFromGoNetbox converts a *models.PowerPanel to the response DTO.
func PowerPanelResponseDTOFromGoNetbox(goNetboxModel *models.PowerPanel) *PowerPanelResponseDTO {
	responseDTO := &PowerPanelResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Site != nil {
		v := goNetboxModel.Site.ID
		responseDTO.Site = &v
	}
	if goNetboxModel.Location != nil {
		v := goNetboxModel.Location.ID
		responseDTO.Location = &v
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
	{
		v := goNetboxModel.PowerfeedCount
		responseDTO.PowerfeedCount = &v
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
