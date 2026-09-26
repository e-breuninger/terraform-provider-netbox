// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// CustomFieldChoiceSetRequestDTO is the request DTO of the custom_field_choice_set resource; Payload writes it into
// the JSON request body.
type CustomFieldChoiceSetRequestDTO struct {
	Name                *string                                       `json:"name,omitempty"`
	Description         *string                                       `json:"description,omitempty"`
	BaseChoices         *string                                       `json:"base_choices,omitempty"`
	ExtraChoices        []*CustomFieldChoiceSetRequestDTOExtraChoices `json:"extra_choices,omitempty"`
	OrderAlphabetically *bool                                         `json:"order_alphabetically,omitempty"`
	Owner               *int64                                        `json:"owner,omitempty"`
}

// CustomFieldChoiceSetRequestDTOExtraChoices is the request DTO of the extra_choices object; it is written into the
// request body as JSON.
type CustomFieldChoiceSetRequestDTOExtraChoices struct {
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// Payload returns the JSON request body of the custom_field_choice_set resource.
func (requestDTO *CustomFieldChoiceSetRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.BaseChoices != nil {
		payload["base_choices"] = *requestDTO.BaseChoices
	} else {
		payload["base_choices"] = nil
	}
	{
		pairs := make([][]string, 0, len(requestDTO.ExtraChoices))
		for _, pair := range requestDTO.ExtraChoices {
			v, l := "", ""
			if pair.Value != nil {
				v = *pair.Value
			}
			if pair.Label != nil {
				l = *pair.Label
			}
			pairs = append(pairs, []string{v, l})
		}
		payload["extra_choices"] = pairs
	}
	if requestDTO.OrderAlphabetically != nil {
		payload["order_alphabetically"] = *requestDTO.OrderAlphabetically
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	return payload
}

// CustomFieldChoiceSetResponseDTO is the response DTO of the custom_field_choice_set resource, built from
// the go-netbox CustomFieldChoiceSet by CustomFieldChoiceSetResponseDTOFromGoNetbox.
type CustomFieldChoiceSetResponseDTO struct {
	ID                  *int64                                         `json:"id,omitempty"`
	Name                *string                                        `json:"name,omitempty"`
	Description         *string                                        `json:"description,omitempty"`
	BaseChoices         *string                                        `json:"base_choices,omitempty"`
	ExtraChoices        []*CustomFieldChoiceSetResponseDTOExtraChoices `json:"extra_choices,omitempty"`
	OrderAlphabetically *bool                                          `json:"order_alphabetically,omitempty"`
	ChoicesCount        *int64                                         `json:"choices_count,omitempty"`
	Owner               *int64                                         `json:"owner,omitempty"`
	Created             *string                                        `json:"created,omitempty"`
	LastUpdated         *string                                        `json:"last_updated,omitempty"`
	URL                 *string                                        `json:"url,omitempty"`
}

// CustomFieldChoiceSetResponseDTOExtraChoices is the response DTO of the extra_choices object, decoded from the
// go-netbox value as JSON.
type CustomFieldChoiceSetResponseDTOExtraChoices struct {
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// CustomFieldChoiceSetResponseDTOFromGoNetbox converts a *models.CustomFieldChoiceSet to the response DTO.
func CustomFieldChoiceSetResponseDTOFromGoNetbox(goNetboxModel *models.CustomFieldChoiceSet) *CustomFieldChoiceSetResponseDTO {
	responseDTO := &CustomFieldChoiceSetResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.BaseChoices != nil {
		responseDTO.BaseChoices = choiceValue[string](goNetboxModel.BaseChoices.Value)
	}
	if goNetboxModel.ExtraChoices != nil {
		responseDTO.ExtraChoices = []*CustomFieldChoiceSetResponseDTOExtraChoices{}
		for _, pair := range goNetboxModel.ExtraChoices {
			if len(pair) == 2 {
				v, l := pair[0], pair[1]
				responseDTO.ExtraChoices = append(responseDTO.ExtraChoices, &CustomFieldChoiceSetResponseDTOExtraChoices{Value: &v, Label: &l})
			}
		}
	}
	{
		v := goNetboxModel.OrderAlphabetically
		responseDTO.OrderAlphabetically = &v
	}
	{
		v := goNetboxModel.ChoicesCount
		responseDTO.ChoicesCount = &v
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
