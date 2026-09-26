// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ModuleRequestDTO is the request DTO of the module resource; Payload writes it into
// the JSON request body.
type ModuleRequestDTO struct {
	Device       *int64            `json:"device,omitempty"`
	ModuleBay    *int64            `json:"module_bay,omitempty"`
	ModuleType   *int64            `json:"module_type,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Serial       *string           `json:"serial,omitempty"`
	AssetTag     *string           `json:"asset_tag,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the module resource.
func (requestDTO *ModuleRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	}
	if requestDTO.ModuleBay != nil {
		payload["module_bay"] = *requestDTO.ModuleBay
	}
	if requestDTO.ModuleType != nil {
		payload["module_type"] = *requestDTO.ModuleType
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Serial != nil {
		payload["serial"] = *requestDTO.Serial
	} else {
		payload["serial"] = ""
	}
	if requestDTO.AssetTag != nil {
		payload["asset_tag"] = *requestDTO.AssetTag
	} else {
		payload["asset_tag"] = nil
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

// ModuleResponseDTO is the response DTO of the module resource, built from
// the go-netbox Module by ModuleResponseDTOFromGoNetbox.
type ModuleResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Device       *int64            `json:"device,omitempty"`
	ModuleBay    *int64            `json:"module_bay,omitempty"`
	ModuleType   *int64            `json:"module_type,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Serial       *string           `json:"serial,omitempty"`
	AssetTag     *string           `json:"asset_tag,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// ModuleResponseDTOFromGoNetbox converts a *models.Module to the response DTO.
func ModuleResponseDTOFromGoNetbox(goNetboxModel *models.Module) *ModuleResponseDTO {
	responseDTO := &ModuleResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.Device != nil {
		v := goNetboxModel.Device.ID
		responseDTO.Device = &v
	}
	if goNetboxModel.ModuleBay != nil {
		v := goNetboxModel.ModuleBay.ID
		responseDTO.ModuleBay = &v
	}
	if goNetboxModel.ModuleType != nil {
		v := goNetboxModel.ModuleType.ID
		responseDTO.ModuleType = &v
	}
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Serial != "" {
		v := goNetboxModel.Serial
		responseDTO.Serial = &v
	}
	responseDTO.AssetTag = goNetboxModel.AssetTag
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
