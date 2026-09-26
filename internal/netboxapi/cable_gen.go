// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CableRequestDTO is the request DTO of the cable resource; Payload writes it into
// the JSON request body.
type CableRequestDTO struct {
	ATerminations *CableRequestDTOATerminations `json:"a_terminations,omitempty"`
	BTerminations *CableRequestDTOBTerminations `json:"b_terminations,omitempty"`
	Type          *string                       `json:"type,omitempty"`
	Status        *string                       `json:"status,omitempty"`
	Profile       *string                       `json:"profile,omitempty"`
	Tenant        *int64                        `json:"tenant,omitempty"`
	Bundle        *int64                        `json:"bundle,omitempty"`
	Label         *string                       `json:"label,omitempty"`
	Color         *string                       `json:"color,omitempty"`
	Length        *float64                      `json:"length,omitempty"`
	LengthUnit    *string                       `json:"length_unit,omitempty"`
	Description   *string                       `json:"description,omitempty"`
	Comments      *string                       `json:"comments,omitempty"`
	Owner         *int64                        `json:"owner,omitempty"`
	Tags          []string                      `json:"tags,omitempty"`
	CustomFields  map[string]string             `json:"custom_fields,omitempty"`
}

// CableRequestDTOATerminations is the request DTO of the a_side object; it is written into the
// request body as JSON.
type CableRequestDTOATerminations struct {
	ObjectType            *string `json:"object_type,omitempty"`
	Ids                   []int64 `json:"ids,omitempty"`
	DeviceInterfaceIds    []int64 `json:"device_interface_ids,omitempty"`
	FrontPortIds          []int64 `json:"front_port_ids,omitempty"`
	RearPortIds           []int64 `json:"rear_port_ids,omitempty"`
	ConsolePortIds        []int64 `json:"console_port_ids,omitempty"`
	ConsoleServerPortIds  []int64 `json:"console_server_port_ids,omitempty"`
	PowerPortIds          []int64 `json:"power_port_ids,omitempty"`
	PowerOutletIds        []int64 `json:"power_outlet_ids,omitempty"`
	PowerFeedIds          []int64 `json:"power_feed_ids,omitempty"`
	CircuitTerminationIds []int64 `json:"circuit_termination_ids,omitempty"`
}

// CableRequestDTOBTerminations is the request DTO of the b_side object; it is written into the
// request body as JSON.
type CableRequestDTOBTerminations struct {
	ObjectType            *string `json:"object_type,omitempty"`
	Ids                   []int64 `json:"ids,omitempty"`
	DeviceInterfaceIds    []int64 `json:"device_interface_ids,omitempty"`
	FrontPortIds          []int64 `json:"front_port_ids,omitempty"`
	RearPortIds           []int64 `json:"rear_port_ids,omitempty"`
	ConsolePortIds        []int64 `json:"console_port_ids,omitempty"`
	ConsoleServerPortIds  []int64 `json:"console_server_port_ids,omitempty"`
	PowerPortIds          []int64 `json:"power_port_ids,omitempty"`
	PowerOutletIds        []int64 `json:"power_outlet_ids,omitempty"`
	PowerFeedIds          []int64 `json:"power_feed_ids,omitempty"`
	CircuitTerminationIds []int64 `json:"circuit_termination_ids,omitempty"`
}

// Payload returns the JSON request body of the cable resource.
func (requestDTO *CableRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.ATerminations != nil {
		payload["a_terminations"] = genericObjects(requestDTO.ATerminations.ObjectType, requestDTO.ATerminations.Ids)
	}
	if requestDTO.BTerminations != nil {
		payload["b_terminations"] = genericObjects(requestDTO.BTerminations.ObjectType, requestDTO.BTerminations.Ids)
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	} else {
		payload["type"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Profile != nil {
		payload["profile"] = *requestDTO.Profile
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Bundle != nil {
		payload["bundle"] = *requestDTO.Bundle
	} else {
		payload["bundle"] = nil
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.Color != nil {
		payload["color"] = *requestDTO.Color
	} else {
		payload["color"] = ""
	}
	if requestDTO.Length != nil {
		payload["length"] = *requestDTO.Length
	} else {
		payload["length"] = nil
	}
	if requestDTO.LengthUnit != nil {
		payload["length_unit"] = *requestDTO.LengthUnit
	} else {
		payload["length_unit"] = nil
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

// CableResponseDTO is the response DTO of the cable resource, built from
// the go-netbox Cable by CableResponseDTOFromGoNetbox.
type CableResponseDTO struct {
	ID            *int64                         `json:"id,omitempty"`
	ATerminations *CableResponseDTOATerminations `json:"a_terminations,omitempty"`
	BTerminations *CableResponseDTOBTerminations `json:"b_terminations,omitempty"`
	Type          *string                        `json:"type,omitempty"`
	Status        *string                        `json:"status,omitempty"`
	Profile       *string                        `json:"profile,omitempty"`
	Tenant        *int64                         `json:"tenant,omitempty"`
	Bundle        *int64                         `json:"bundle,omitempty"`
	Label         *string                        `json:"label,omitempty"`
	Color         *string                        `json:"color,omitempty"`
	Length        *float64                       `json:"length,omitempty"`
	LengthUnit    *string                        `json:"length_unit,omitempty"`
	Description   *string                        `json:"description,omitempty"`
	Comments      *string                        `json:"comments,omitempty"`
	Owner         *int64                         `json:"owner,omitempty"`
	Created       *string                        `json:"created,omitempty"`
	LastUpdated   *string                        `json:"last_updated,omitempty"`
	URL           *string                        `json:"url,omitempty"`
	Tags          []string                       `json:"tags,omitempty"`
	TagsAll       []string                       `json:"tags_all,omitempty"`
	CustomFields  map[string]string              `json:"custom_fields,omitempty"`
}

// CableResponseDTOATerminations is the response DTO of the a_side object, decoded from the
// go-netbox value as JSON.
type CableResponseDTOATerminations struct {
	ObjectType            *string `json:"object_type,omitempty"`
	Ids                   []int64 `json:"ids,omitempty"`
	DeviceInterfaceIds    []int64 `json:"device_interface_ids,omitempty"`
	FrontPortIds          []int64 `json:"front_port_ids,omitempty"`
	RearPortIds           []int64 `json:"rear_port_ids,omitempty"`
	ConsolePortIds        []int64 `json:"console_port_ids,omitempty"`
	ConsoleServerPortIds  []int64 `json:"console_server_port_ids,omitempty"`
	PowerPortIds          []int64 `json:"power_port_ids,omitempty"`
	PowerOutletIds        []int64 `json:"power_outlet_ids,omitempty"`
	PowerFeedIds          []int64 `json:"power_feed_ids,omitempty"`
	CircuitTerminationIds []int64 `json:"circuit_termination_ids,omitempty"`
}

// CableResponseDTOBTerminations is the response DTO of the b_side object, decoded from the
// go-netbox value as JSON.
type CableResponseDTOBTerminations struct {
	ObjectType            *string `json:"object_type,omitempty"`
	Ids                   []int64 `json:"ids,omitempty"`
	DeviceInterfaceIds    []int64 `json:"device_interface_ids,omitempty"`
	FrontPortIds          []int64 `json:"front_port_ids,omitempty"`
	RearPortIds           []int64 `json:"rear_port_ids,omitempty"`
	ConsolePortIds        []int64 `json:"console_port_ids,omitempty"`
	ConsoleServerPortIds  []int64 `json:"console_server_port_ids,omitempty"`
	PowerPortIds          []int64 `json:"power_port_ids,omitempty"`
	PowerOutletIds        []int64 `json:"power_outlet_ids,omitempty"`
	PowerFeedIds          []int64 `json:"power_feed_ids,omitempty"`
	CircuitTerminationIds []int64 `json:"circuit_termination_ids,omitempty"`
}

// CableResponseDTOFromGoNetbox converts a *models.Cable to the response DTO.
func CableResponseDTOFromGoNetbox(goNetboxModel *models.Cable) *CableResponseDTO {
	responseDTO := &CableResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.ATerminations != nil {
		v := &CableResponseDTOATerminations{}
		v.ObjectType, v.Ids = terminationSideOf(goNetboxModel.ATerminations)
		responseDTO.ATerminations = v
	}
	if goNetboxModel.BTerminations != nil {
		v := &CableResponseDTOBTerminations{}
		v.ObjectType, v.Ids = terminationSideOf(goNetboxModel.BTerminations)
		responseDTO.BTerminations = v
	}
	responseDTO.Type = goNetboxModel.Type
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Profile != nil {
		responseDTO.Profile = choiceValue[string](goNetboxModel.Profile.Value)
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Bundle != nil {
		v := goNetboxModel.Bundle.ID
		responseDTO.Bundle = &v
	}
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	if goNetboxModel.Color != "" {
		v := goNetboxModel.Color
		responseDTO.Color = &v
	}
	responseDTO.Length = goNetboxModel.Length
	if goNetboxModel.LengthUnit != nil {
		responseDTO.LengthUnit = choiceValue[string](goNetboxModel.LengthUnit.Value)
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
