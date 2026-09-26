// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CustomFieldRequestDTO is the request DTO of the custom_field resource; Payload writes it into
// the JSON request body.
type CustomFieldRequestDTO struct {
	Name              *string  `json:"name,omitempty"`
	ObjectTypes       []string `json:"object_types,omitempty"`
	Type              *string  `json:"type,omitempty"`
	RelatedObjectType *string  `json:"related_object_type,omitempty"`
	Label             *string  `json:"label,omitempty"`
	Description       *string  `json:"description,omitempty"`
	GroupName         *string  `json:"group_name,omitempty"`
	Required          *bool    `json:"required,omitempty"`
	FilterLogic       *string  `json:"filter_logic,omitempty"`
	Weight            *int64   `json:"weight,omitempty"`
	ValidationMinimum *float64 `json:"validation_minimum,omitempty"`
	ValidationMaximum *float64 `json:"validation_maximum,omitempty"`
	ValidationRegex   *string  `json:"validation_regex,omitempty"`
	Default           *string  `json:"default,omitempty"`
	ChoiceSet         *int64   `json:"choice_set,omitempty"`
	Owner             *int64   `json:"owner,omitempty"`
}

// Payload returns the JSON request body of the custom_field resource.
func (requestDTO *CustomFieldRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.ObjectTypes != nil {
		payload["object_types"] = requestDTO.ObjectTypes
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.RelatedObjectType != nil {
		payload["related_object_type"] = *requestDTO.RelatedObjectType
	} else {
		payload["related_object_type"] = nil
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.GroupName != nil {
		payload["group_name"] = *requestDTO.GroupName
	} else {
		payload["group_name"] = ""
	}
	if requestDTO.Required != nil {
		payload["required"] = *requestDTO.Required
	}
	if requestDTO.FilterLogic != nil {
		payload["filter_logic"] = *requestDTO.FilterLogic
	}
	if requestDTO.Weight != nil {
		payload["weight"] = *requestDTO.Weight
	}
	if requestDTO.ValidationMinimum != nil {
		payload["validation_minimum"] = *requestDTO.ValidationMinimum
	} else {
		payload["validation_minimum"] = nil
	}
	if requestDTO.ValidationMaximum != nil {
		payload["validation_maximum"] = *requestDTO.ValidationMaximum
	} else {
		payload["validation_maximum"] = nil
	}
	if requestDTO.ValidationRegex != nil {
		payload["validation_regex"] = *requestDTO.ValidationRegex
	} else {
		payload["validation_regex"] = ""
	}
	if requestDTO.Default != nil {
		payload["default"] = parseJSONText(*requestDTO.Default)
	} else {
		payload["default"] = nil
	}
	if requestDTO.ChoiceSet != nil {
		payload["choice_set"] = *requestDTO.ChoiceSet
	} else {
		payload["choice_set"] = nil
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	return payload
}

// CustomFieldResponseDTO is the response DTO of the custom_field resource, built from
// the go-netbox CustomField by CustomFieldResponseDTOFromGoNetbox.
type CustomFieldResponseDTO struct {
	ID                *int64   `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	ObjectTypes       []string `json:"object_types,omitempty"`
	Type              *string  `json:"type,omitempty"`
	RelatedObjectType *string  `json:"related_object_type,omitempty"`
	Label             *string  `json:"label,omitempty"`
	Description       *string  `json:"description,omitempty"`
	GroupName         *string  `json:"group_name,omitempty"`
	Required          *bool    `json:"required,omitempty"`
	FilterLogic       *string  `json:"filter_logic,omitempty"`
	Weight            *int64   `json:"weight,omitempty"`
	ValidationMinimum *float64 `json:"validation_minimum,omitempty"`
	ValidationMaximum *float64 `json:"validation_maximum,omitempty"`
	ValidationRegex   *string  `json:"validation_regex,omitempty"`
	Default           *string  `json:"default,omitempty"`
	ChoiceSet         *int64   `json:"choice_set,omitempty"`
	Owner             *int64   `json:"owner,omitempty"`
	Created           *string  `json:"created,omitempty"`
	LastUpdated       *string  `json:"last_updated,omitempty"`
	URL               *string  `json:"url,omitempty"`
}

// CustomFieldResponseDTOFromGoNetbox converts a *models.CustomField to the response DTO.
func CustomFieldResponseDTOFromGoNetbox(goNetboxModel *models.CustomField) *CustomFieldResponseDTO {
	responseDTO := &CustomFieldResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.ObjectTypes = goNetboxModel.ObjectTypes
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
	responseDTO.RelatedObjectType = goNetboxModel.RelatedObjectType
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.GroupName != "" {
		v := goNetboxModel.GroupName
		responseDTO.GroupName = &v
	}
	{
		v := goNetboxModel.Required
		responseDTO.Required = &v
	}
	if goNetboxModel.FilterLogic != nil {
		responseDTO.FilterLogic = choiceValue[string](goNetboxModel.FilterLogic.Value)
	}
	responseDTO.Weight = goNetboxModel.Weight
	responseDTO.ValidationMinimum = goNetboxModel.ValidationMinimum
	responseDTO.ValidationMaximum = goNetboxModel.ValidationMaximum
	if goNetboxModel.ValidationRegex != "" {
		v := goNetboxModel.ValidationRegex
		responseDTO.ValidationRegex = &v
	}
	responseDTO.Default = jsonText(goNetboxModel.Default)
	if goNetboxModel.ChoiceSet != nil {
		v := goNetboxModel.ChoiceSet.ID
		responseDTO.ChoiceSet = &v
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
