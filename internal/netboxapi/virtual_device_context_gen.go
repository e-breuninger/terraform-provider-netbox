// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualDeviceContextRequestDTO is the request DTO of the virtual_device_context resource; Payload writes it into
// the JSON request body.
type VirtualDeviceContextRequestDTO struct {
	Device       *int64            `json:"device,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Identifier   *int64            `json:"identifier,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	PrimaryIp4   *int64            `json:"primary_ip4,omitempty"`
	PrimaryIp6   *int64            `json:"primary_ip6,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_device_context resource.
func (requestDTO *VirtualDeviceContextRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Identifier != nil {
		payload["identifier"] = *requestDTO.Identifier
	} else {
		payload["identifier"] = nil
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.PrimaryIp4 != nil {
		payload["primary_ip4"] = *requestDTO.PrimaryIp4
	} else {
		payload["primary_ip4"] = nil
	}
	if requestDTO.PrimaryIp6 != nil {
		payload["primary_ip6"] = *requestDTO.PrimaryIp6
	} else {
		payload["primary_ip6"] = nil
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

// VirtualDeviceContextResponseDTO is the response DTO of the virtual_device_context resource, built from
// the go-netbox VirtualDeviceContext by VirtualDeviceContextResponseDTOFromGoNetbox.
type VirtualDeviceContextResponseDTO struct {
	ID             *int64            `json:"id,omitempty"`
	Device         *int64            `json:"device,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Identifier     *int64            `json:"identifier,omitempty"`
	Status         *string           `json:"status,omitempty"`
	Tenant         *int64            `json:"tenant,omitempty"`
	PrimaryIp4     *int64            `json:"primary_ip4,omitempty"`
	PrimaryIp6     *int64            `json:"primary_ip6,omitempty"`
	InterfaceCount *int64            `json:"interface_count,omitempty"`
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

// VirtualDeviceContextResponseDTOFromGoNetbox converts a *models.VirtualDeviceContext to the response DTO.
func VirtualDeviceContextResponseDTOFromGoNetbox(goNetboxModel *models.VirtualDeviceContext) *VirtualDeviceContextResponseDTO {
	responseDTO := &VirtualDeviceContextResponseDTO{}
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
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Identifier = goNetboxModel.Identifier
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.PrimaryIp4 != nil {
		v := goNetboxModel.PrimaryIp4.ID
		responseDTO.PrimaryIp4 = &v
	}
	if goNetboxModel.PrimaryIp6 != nil {
		v := goNetboxModel.PrimaryIp6.ID
		responseDTO.PrimaryIp6 = &v
	}
	{
		v := goNetboxModel.InterfaceCount
		responseDTO.InterfaceCount = &v
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
