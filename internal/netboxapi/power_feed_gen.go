// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// PowerFeedRequestDTO is the request DTO of the power_feed resource; Payload writes it into
// the JSON request body.
type PowerFeedRequestDTO struct {
	Name           *string           `json:"name,omitempty"`
	PowerPanel     *int64            `json:"power_panel,omitempty"`
	Rack           *int64            `json:"rack,omitempty"`
	Status         *string           `json:"status,omitempty"`
	Type           *string           `json:"type,omitempty"`
	Supply         *string           `json:"supply,omitempty"`
	Phase          *string           `json:"phase,omitempty"`
	Voltage        *int64            `json:"voltage,omitempty"`
	Amperage       *int64            `json:"amperage,omitempty"`
	MaxUtilization *int64            `json:"max_utilization,omitempty"`
	MarkConnected  *bool             `json:"mark_connected,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Comments       *string           `json:"comments,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the power_feed resource.
func (requestDTO *PowerFeedRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.PowerPanel != nil {
		payload["power_panel"] = *requestDTO.PowerPanel
	}
	if requestDTO.Rack != nil {
		payload["rack"] = *requestDTO.Rack
	} else {
		payload["rack"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.Supply != nil {
		payload["supply"] = *requestDTO.Supply
	}
	if requestDTO.Phase != nil {
		payload["phase"] = *requestDTO.Phase
	}
	if requestDTO.Voltage != nil {
		payload["voltage"] = *requestDTO.Voltage
	}
	if requestDTO.Amperage != nil {
		payload["amperage"] = *requestDTO.Amperage
	}
	if requestDTO.MaxUtilization != nil {
		payload["max_utilization"] = *requestDTO.MaxUtilization
	}
	if requestDTO.MarkConnected != nil {
		payload["mark_connected"] = *requestDTO.MarkConnected
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

// PowerFeedResponseDTO is the response DTO of the power_feed resource, built from
// the go-netbox PowerFeed by PowerFeedResponseDTOFromGoNetbox.
type PowerFeedResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	Name           *string           `json:"name,omitempty"`
	PowerPanel     *int64            `json:"power_panel,omitempty"`
	Rack           *int64            `json:"rack,omitempty"`
	Status         *string           `json:"status,omitempty"`
	Type           *string           `json:"type,omitempty"`
	Supply         *string           `json:"supply,omitempty"`
	Phase          *string           `json:"phase,omitempty"`
	Voltage        *int64            `json:"voltage,omitempty"`
	Amperage       *int64            `json:"amperage,omitempty"`
	MaxUtilization *int64            `json:"max_utilization,omitempty"`
	MarkConnected  *bool             `json:"mark_connected,omitempty"`
	Description    *string           `json:"description,omitempty"`
	Comments       *string           `json:"comments,omitempty"`
	Owner          *int64            `json:"owner,omitempty"`
	Created        *string           `json:"created,omitempty"`
	LastUpdated    *string           `json:"last_updated,omitempty"`
	URL            *string           `json:"url,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	TagsAll        []string          `json:"tags_all,omitempty"`
	CustomFields   map[string]string `json:"custom_fields,omitempty"`
}

// PowerFeedResponseDTOFromGoNetbox converts a *models.PowerFeed to the response DTO.
func PowerFeedResponseDTOFromGoNetbox(goNetboxModel *models.PowerFeed) *PowerFeedResponseDTO {
	responseDTO := &PowerFeedResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.PowerPanel != nil {
		v := goNetboxModel.PowerPanel.ID
		responseDTO.PowerPanel = &v
	}
	if goNetboxModel.Rack != nil {
		v := goNetboxModel.Rack.ID
		responseDTO.Rack = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
	if goNetboxModel.Supply != nil {
		responseDTO.Supply = choiceValue[string](goNetboxModel.Supply.Value)
	}
	if goNetboxModel.Phase != nil {
		responseDTO.Phase = choiceValue[string](goNetboxModel.Phase.Value)
	}
	responseDTO.Voltage = goNetboxModel.Voltage
	{
		v := goNetboxModel.Amperage
		responseDTO.Amperage = &v
	}
	{
		v := goNetboxModel.MaxUtilization
		responseDTO.MaxUtilization = &v
	}
	{
		v := goNetboxModel.MarkConnected
		responseDTO.MarkConnected = &v
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
