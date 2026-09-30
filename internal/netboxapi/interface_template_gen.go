// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// InterfaceTemplateRequestDTO is the request DTO of the interface_template resource; Payload writes it into
// the JSON request body.
type InterfaceTemplateRequestDTO struct {
	DeviceType  *int64  `json:"device_type,omitempty"`
	ModuleType  *int64  `json:"module_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	MgmtOnly    *bool   `json:"mgmt_only,omitempty"`
	Bridge      *int64  `json:"bridge,omitempty"`
	PoeMode     *string `json:"poe_mode,omitempty"`
	PoeType     *string `json:"poe_type,omitempty"`
	RfRole      *string `json:"rf_role,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Payload returns the JSON request body of the interface_template resource.
func (requestDTO *InterfaceTemplateRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.DeviceType != nil {
		payload["device_type"] = *requestDTO.DeviceType
	} else {
		payload["device_type"] = nil
	}
	if requestDTO.ModuleType != nil {
		payload["module_type"] = *requestDTO.ModuleType
	} else {
		payload["module_type"] = nil
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.MgmtOnly != nil {
		payload["mgmt_only"] = *requestDTO.MgmtOnly
	}
	if requestDTO.Bridge != nil {
		payload["bridge"] = *requestDTO.Bridge
	} else {
		payload["bridge"] = nil
	}
	if requestDTO.PoeMode != nil {
		payload["poe_mode"] = *requestDTO.PoeMode
	} else {
		payload["poe_mode"] = nil
	}
	if requestDTO.PoeType != nil {
		payload["poe_type"] = *requestDTO.PoeType
	} else {
		payload["poe_type"] = nil
	}
	if requestDTO.RfRole != nil {
		payload["rf_role"] = *requestDTO.RfRole
	} else {
		payload["rf_role"] = nil
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
	return payload
}

// InterfaceTemplateResponseDTO is the response DTO of the interface_template resource, built from
// the go-netbox InterfaceTemplate by InterfaceTemplateResponseDTOFromGoNetbox.
type InterfaceTemplateResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	DeviceType  *int64  `json:"device_type,omitempty"`
	ModuleType  *int64  `json:"module_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	MgmtOnly    *bool   `json:"mgmt_only,omitempty"`
	Bridge      *int64  `json:"bridge,omitempty"`
	PoeMode     *string `json:"poe_mode,omitempty"`
	PoeType     *string `json:"poe_type,omitempty"`
	RfRole      *string `json:"rf_role,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Created     *string `json:"created,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// InterfaceTemplateResponseDTOFromGoNetbox converts a *models.InterfaceTemplate to the response DTO.
func InterfaceTemplateResponseDTOFromGoNetbox(goNetboxModel *models.InterfaceTemplate) *InterfaceTemplateResponseDTO {
	responseDTO := &InterfaceTemplateResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.DeviceType != nil {
		v := goNetboxModel.DeviceType.ID
		responseDTO.DeviceType = &v
	}
	if goNetboxModel.ModuleType != nil {
		v := goNetboxModel.ModuleType.ID
		responseDTO.ModuleType = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	{
		v := goNetboxModel.MgmtOnly
		responseDTO.MgmtOnly = &v
	}
	if goNetboxModel.Bridge != nil {
		v := goNetboxModel.Bridge.ID
		responseDTO.Bridge = &v
	}
	if goNetboxModel.PoeMode != nil {
		responseDTO.PoeMode = choiceValue[string](goNetboxModel.PoeMode.Value)
	}
	if goNetboxModel.PoeType != nil {
		responseDTO.PoeType = choiceValue[string](goNetboxModel.PoeType.Value)
	}
	if goNetboxModel.RfRole != nil {
		responseDTO.RfRole = choiceValue[string](goNetboxModel.RfRole.Value)
	}
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
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
	return responseDTO
}
