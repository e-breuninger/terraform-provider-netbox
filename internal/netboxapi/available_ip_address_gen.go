// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// AvailableIPAddressRequestDTO is the request DTO of the available_ip_address resource; Payload writes it into
// the JSON request body.
type AvailableIPAddressRequestDTO struct {
	PrefixID                  *int64            `json:"prefix_id,omitempty"`
	IPRangeID                 *int64            `json:"ip_range_id,omitempty"`
	Status                    *string           `json:"status,omitempty"`
	Role                      *string           `json:"role,omitempty"`
	DNSName                   *string           `json:"dns_name,omitempty"`
	Description               *string           `json:"description,omitempty"`
	Comments                  *string           `json:"comments,omitempty"`
	Tenant                    *int64            `json:"tenant,omitempty"`
	Vrf                       *int64            `json:"vrf,omitempty"`
	NatInside                 *int64            `json:"nat_inside,omitempty"`
	AssignedObjectType        *string           `json:"assigned_object_type,omitempty"`
	AssignedObjectID          *int64            `json:"assigned_object_id,omitempty"`
	DeviceInterfaceID         *int64            `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64            `json:"virtual_machine_interface_id,omitempty"`
	Owner                     *int64            `json:"owner,omitempty"`
	Tags                      []string          `json:"tags,omitempty"`
	CustomFields              map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the available_ip_address resource.
func (requestDTO *AvailableIPAddressRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	} else {
		payload["role"] = nil
	}
	if requestDTO.DNSName != nil {
		payload["dns_name"] = *requestDTO.DNSName
	} else {
		payload["dns_name"] = ""
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
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Vrf != nil {
		payload["vrf"] = *requestDTO.Vrf
	}
	if requestDTO.NatInside != nil {
		payload["nat_inside"] = *requestDTO.NatInside
	} else {
		payload["nat_inside"] = nil
	}
	if requestDTO.AssignedObjectType != nil {
		payload["assigned_object_type"] = *requestDTO.AssignedObjectType
	} else {
		payload["assigned_object_type"] = nil
	}
	if requestDTO.AssignedObjectID != nil {
		payload["assigned_object_id"] = *requestDTO.AssignedObjectID
	} else {
		payload["assigned_object_id"] = nil
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

// AvailableIPAddressResponseDTO is the response DTO of the available_ip_address resource, built from
// the go-netbox IPAddress by AvailableIPAddressResponseDTOFromGoNetbox.
type AvailableIPAddressResponseDTO struct {
	ID                        *int64                                       `json:"id,omitempty"`
	Address                   *string                                      `json:"address,omitempty"`
	Status                    *string                                      `json:"status,omitempty"`
	Role                      *string                                      `json:"role,omitempty"`
	DNSName                   *string                                      `json:"dns_name,omitempty"`
	Description               *string                                      `json:"description,omitempty"`
	Comments                  *string                                      `json:"comments,omitempty"`
	Tenant                    *int64                                       `json:"tenant,omitempty"`
	Vrf                       *int64                                       `json:"vrf,omitempty"`
	NatInside                 *int64                                       `json:"nat_inside,omitempty"`
	AssignedObjectType        *string                                      `json:"assigned_object_type,omitempty"`
	AssignedObjectID          *int64                                       `json:"assigned_object_id,omitempty"`
	DeviceInterfaceID         *int64                                       `json:"device_interface_id,omitempty"`
	VirtualMachineInterfaceID *int64                                       `json:"virtual_machine_interface_id,omitempty"`
	AssignedObject            *AvailableIPAddressResponseDTOAssignedObject `json:"assigned_object,omitempty"`
	Family                    *int64                                       `json:"family,omitempty"`
	Owner                     *int64                                       `json:"owner,omitempty"`
	Created                   *string                                      `json:"created,omitempty"`
	LastUpdated               *string                                      `json:"last_updated,omitempty"`
	URL                       *string                                      `json:"url,omitempty"`
	Tags                      []string                                     `json:"tags,omitempty"`
	TagsAll                   []string                                     `json:"tags_all,omitempty"`
	CustomFields              map[string]string                            `json:"custom_fields,omitempty"`
}

// AvailableIPAddressResponseDTOAssignedObject is the response DTO of the assigned_object object, decoded from the
// go-netbox value as JSON.
type AvailableIPAddressResponseDTOAssignedObject struct {
	ID     *int64                                             `json:"id,omitempty"`
	Name   *string                                            `json:"name,omitempty"`
	Device *AvailableIPAddressResponseDTOAssignedObjectDevice `json:"device,omitempty"`
}

// AvailableIPAddressResponseDTOAssignedObjectDevice is the response DTO of the device object, decoded from the
// go-netbox value as JSON.
type AvailableIPAddressResponseDTOAssignedObjectDevice struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// AvailableIPAddressResponseDTOFromGoNetbox converts a *models.IPAddress to the response DTO.
func AvailableIPAddressResponseDTOFromGoNetbox(goNetboxModel *models.IPAddress) *AvailableIPAddressResponseDTO {
	responseDTO := &AvailableIPAddressResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Address = goNetboxModel.Address
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Role != nil {
		responseDTO.Role = choiceValue[string](goNetboxModel.Role.Value)
	}
	if goNetboxModel.DNSName != "" {
		v := goNetboxModel.DNSName
		responseDTO.DNSName = &v
	}
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Comments != "" {
		v := goNetboxModel.Comments
		responseDTO.Comments = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Vrf != nil {
		v := goNetboxModel.Vrf.ID
		responseDTO.Vrf = &v
	}
	if goNetboxModel.NatInside != nil {
		v := goNetboxModel.NatInside.ID
		responseDTO.NatInside = &v
	}
	responseDTO.AssignedObjectType = goNetboxModel.AssignedObjectType
	responseDTO.AssignedObjectID = goNetboxModel.AssignedObjectID
	if goNetboxModel.AssignedObject != nil {
		var v *AvailableIPAddressResponseDTOAssignedObject
		if err := reencode(goNetboxModel.AssignedObject, &v); err == nil {
			responseDTO.AssignedObject = v
		}
	}
	if goNetboxModel.Family != nil {
		responseDTO.Family = choiceValue[int64](goNetboxModel.Family.Value)
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
