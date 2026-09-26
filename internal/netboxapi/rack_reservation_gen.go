// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// RackReservationRequestDTO is the request DTO of the rack_reservation resource; Payload writes it into
// the JSON request body.
type RackReservationRequestDTO struct {
	Rack         *int64            `json:"rack,omitempty"`
	Units        []int64           `json:"units,omitempty"`
	User         *int64            `json:"user,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the rack_reservation resource.
func (requestDTO *RackReservationRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Rack != nil {
		payload["rack"] = *requestDTO.Rack
	}
	if requestDTO.Units != nil {
		payload["units"] = requestDTO.Units
	}
	if requestDTO.User != nil {
		payload["user"] = *requestDTO.User
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
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

// RackReservationResponseDTO is the response DTO of the rack_reservation resource, built from
// the go-netbox RackReservation by RackReservationResponseDTOFromGoNetbox.
type RackReservationResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Rack         *int64            `json:"rack,omitempty"`
	Units        []int64           `json:"units,omitempty"`
	User         *int64            `json:"user,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	UnitCount    *int64            `json:"unit_count,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// RackReservationResponseDTOFromGoNetbox converts a *models.RackReservation to the response DTO.
func RackReservationResponseDTOFromGoNetbox(goNetboxModel *models.RackReservation) *RackReservationResponseDTO {
	responseDTO := &RackReservationResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Rack != nil {
		v := goNetboxModel.Rack.ID
		responseDTO.Rack = &v
	}
	if goNetboxModel.Units != nil {
		responseDTO.Units = []int64{}
		for _, value := range goNetboxModel.Units {
			if value != nil {
				responseDTO.Units = append(responseDTO.Units, *value)
			}
		}
	}
	if goNetboxModel.User != nil {
		v := goNetboxModel.User.ID
		responseDTO.User = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	responseDTO.Description = goNetboxModel.Description
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
		v := int64(goNetboxModel.UnitCount)
		responseDTO.UnitCount = &v
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
