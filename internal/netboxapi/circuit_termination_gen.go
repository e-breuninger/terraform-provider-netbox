// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitTerminationRequestDTO is the request DTO of the circuit_termination resource; Payload writes it into
// the JSON request body.
type CircuitTerminationRequestDTO struct {
	Circuit           *int64            `json:"circuit,omitempty"`
	TermSide          *string           `json:"term_side,omitempty"`
	TerminationType   *string           `json:"termination_type,omitempty"`
	TerminationID     *int64            `json:"termination_id,omitempty"`
	SiteID            *int64            `json:"site_id,omitempty"`
	LocationID        *int64            `json:"location_id,omitempty"`
	RegionID          *int64            `json:"region_id,omitempty"`
	SiteGroupID       *int64            `json:"site_group_id,omitempty"`
	ProviderNetworkID *int64            `json:"provider_network_id,omitempty"`
	PortSpeed         *int64            `json:"port_speed,omitempty"`
	UpstreamSpeed     *int64            `json:"upstream_speed,omitempty"`
	XconnectID        *string           `json:"xconnect_id,omitempty"`
	PpInfo            *string           `json:"pp_info,omitempty"`
	MarkConnected     *bool             `json:"mark_connected,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the circuit_termination resource.
func (requestDTO *CircuitTerminationRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Circuit != nil {
		payload["circuit"] = *requestDTO.Circuit
	}
	if requestDTO.TermSide != nil {
		payload["term_side"] = *requestDTO.TermSide
	}
	if requestDTO.TerminationType != nil {
		payload["termination_type"] = *requestDTO.TerminationType
	}
	if requestDTO.TerminationID != nil {
		payload["termination_id"] = *requestDTO.TerminationID
	}
	if requestDTO.PortSpeed != nil {
		payload["port_speed"] = *requestDTO.PortSpeed
	} else {
		payload["port_speed"] = nil
	}
	if requestDTO.UpstreamSpeed != nil {
		payload["upstream_speed"] = *requestDTO.UpstreamSpeed
	} else {
		payload["upstream_speed"] = nil
	}
	if requestDTO.XconnectID != nil {
		payload["xconnect_id"] = *requestDTO.XconnectID
	} else {
		payload["xconnect_id"] = ""
	}
	if requestDTO.PpInfo != nil {
		payload["pp_info"] = *requestDTO.PpInfo
	} else {
		payload["pp_info"] = ""
	}
	if requestDTO.MarkConnected != nil {
		payload["mark_connected"] = *requestDTO.MarkConnected
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	payload["tags"] = tagRefs(requestDTO.Tags)
	if requestDTO.CustomFields != nil {
		payload["custom_fields"] = requestDTO.CustomFields
	}
	return payload
}

// CircuitTerminationResponseDTO is the response DTO of the circuit_termination resource, built from
// the go-netbox CircuitTermination by CircuitTerminationResponseDTOFromGoNetbox.
type CircuitTerminationResponseDTO struct {
	ID                *int64            `json:"id,omitempty"`
	Circuit           *int64            `json:"circuit,omitempty"`
	TermSide          *string           `json:"term_side,omitempty"`
	TerminationType   *string           `json:"termination_type,omitempty"`
	TerminationID     *int64            `json:"termination_id,omitempty"`
	SiteID            *int64            `json:"site_id,omitempty"`
	LocationID        *int64            `json:"location_id,omitempty"`
	RegionID          *int64            `json:"region_id,omitempty"`
	SiteGroupID       *int64            `json:"site_group_id,omitempty"`
	ProviderNetworkID *int64            `json:"provider_network_id,omitempty"`
	PortSpeed         *int64            `json:"port_speed,omitempty"`
	UpstreamSpeed     *int64            `json:"upstream_speed,omitempty"`
	XconnectID        *string           `json:"xconnect_id,omitempty"`
	PpInfo            *string           `json:"pp_info,omitempty"`
	MarkConnected     *bool             `json:"mark_connected,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Created           *string           `json:"created,omitempty"`
	LastUpdated       *string           `json:"last_updated,omitempty"`
	URL               *string           `json:"url,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	TagsAll           []string          `json:"tags_all,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
}

// CircuitTerminationResponseDTOFromGoNetbox converts a *models.CircuitTermination to the response DTO.
func CircuitTerminationResponseDTOFromGoNetbox(goNetboxModel *models.CircuitTermination) *CircuitTerminationResponseDTO {
	responseDTO := &CircuitTerminationResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Circuit != nil {
		v := goNetboxModel.Circuit.ID
		responseDTO.Circuit = &v
	}
	responseDTO.TermSide = goNetboxModel.TermSide
	responseDTO.TerminationType = goNetboxModel.TerminationType
	responseDTO.TerminationID = goNetboxModel.TerminationID
	responseDTO.PortSpeed = goNetboxModel.PortSpeed
	responseDTO.UpstreamSpeed = goNetboxModel.UpstreamSpeed
	if goNetboxModel.XconnectID != "" {
		v := goNetboxModel.XconnectID
		responseDTO.XconnectID = &v
	}
	if goNetboxModel.PpInfo != "" {
		v := goNetboxModel.PpInfo
		responseDTO.PpInfo = &v
	}
	{
		v := goNetboxModel.MarkConnected
		responseDTO.MarkConnected = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
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
