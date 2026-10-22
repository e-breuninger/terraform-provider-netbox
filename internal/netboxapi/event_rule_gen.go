// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// EventRuleRequestDTO is the request DTO of the event_rule resource; Payload writes it into
// the JSON request body.
type EventRuleRequestDTO struct {
	Name             *string           `json:"name,omitempty"`
	ObjectTypes      []string          `json:"object_types,omitempty"`
	EventTypes       []string          `json:"event_types,omitempty"`
	Enabled          *bool             `json:"enabled,omitempty"`
	Conditions       *string           `json:"conditions,omitempty"`
	ActionType       *string           `json:"action_type,omitempty"`
	ActionObjectType *string           `json:"action_object_type,omitempty"`
	ActionObjectID   *int64            `json:"action_object_id,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Owner            *int64            `json:"owner,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	CustomFields     map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the event_rule resource.
func (requestDTO *EventRuleRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.ObjectTypes != nil {
		payload["object_types"] = requestDTO.ObjectTypes
	}
	if requestDTO.EventTypes != nil {
		payload["event_types"] = requestDTO.EventTypes
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.Conditions != nil {
		payload["conditions"] = parseJSONText(*requestDTO.Conditions)
	} else {
		payload["conditions"] = nil
	}
	if requestDTO.ActionType != nil {
		payload["action_type"] = *requestDTO.ActionType
	}
	if requestDTO.ActionObjectType != nil {
		payload["action_object_type"] = *requestDTO.ActionObjectType
	}
	if requestDTO.ActionObjectID != nil {
		payload["action_object_id"] = *requestDTO.ActionObjectID
	} else {
		payload["action_object_id"] = nil
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

// EventRuleResponseDTO is the response DTO of the event_rule resource, built from
// the go-netbox EventRule by EventRuleResponseDTOFromGoNetbox.
type EventRuleResponseDTO struct {
	ID               *int64            `json:"id,omitempty"`
	Name             *string           `json:"name,omitempty"`
	ObjectTypes      []string          `json:"object_types,omitempty"`
	EventTypes       []string          `json:"event_types,omitempty"`
	Enabled          *bool             `json:"enabled,omitempty"`
	Conditions       *string           `json:"conditions,omitempty"`
	ActionType       *string           `json:"action_type,omitempty"`
	ActionObjectType *string           `json:"action_object_type,omitempty"`
	ActionObjectID   *int64            `json:"action_object_id,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Owner            *int64            `json:"owner,omitempty"`
	Created          *string           `json:"created,omitempty"`
	LastUpdated      *string           `json:"last_updated,omitempty"`
	URL              *string           `json:"url,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	TagsAll          []string          `json:"tags_all,omitempty"`
	CustomFields     map[string]string `json:"custom_fields,omitempty"`
}

// EventRuleResponseDTOFromGoNetbox converts a *models.EventRule to the response DTO.
func EventRuleResponseDTOFromGoNetbox(goNetboxModel *models.EventRule) *EventRuleResponseDTO {
	responseDTO := &EventRuleResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.ObjectTypes = goNetboxModel.ObjectTypes
	responseDTO.EventTypes = goNetboxModel.EventTypes
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	responseDTO.Conditions = jsonText(goNetboxModel.Conditions)
	if goNetboxModel.ActionType != nil {
		responseDTO.ActionType = choiceValue[string](goNetboxModel.ActionType.Value)
	}
	responseDTO.ActionObjectType = goNetboxModel.ActionObjectType
	responseDTO.ActionObjectID = goNetboxModel.ActionObjectID
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
