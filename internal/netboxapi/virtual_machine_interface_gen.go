// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualMachineInterfaceRequestDTO is the request DTO of the virtual_machine_interface resource; Payload writes it into
// the JSON request body.
type VirtualMachineInterfaceRequestDTO struct {
	VirtualMachine        *int64            `json:"virtual_machine,omitempty"`
	Name                  *string           `json:"name,omitempty"`
	Enabled               *bool             `json:"enabled,omitempty"`
	Mtu                   *int64            `json:"mtu,omitempty"`
	Mode                  *string           `json:"mode,omitempty"`
	UntaggedVlan          *int64            `json:"untagged_vlan,omitempty"`
	TaggedVlans           []int64           `json:"tagged_vlans,omitempty"`
	QinqSvlan             *int64            `json:"qinq_svlan,omitempty"`
	VlanTranslationPolicy *int64            `json:"vlan_translation_policy,omitempty"`
	Vrf                   *int64            `json:"vrf,omitempty"`
	Parent                *int64            `json:"parent,omitempty"`
	Bridge                *int64            `json:"bridge,omitempty"`
	Description           *string           `json:"description,omitempty"`
	Owner                 *int64            `json:"owner,omitempty"`
	Tags                  []string          `json:"tags,omitempty"`
	CustomFields          map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_machine_interface resource.
func (requestDTO *VirtualMachineInterfaceRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.VirtualMachine != nil {
		payload["virtual_machine"] = *requestDTO.VirtualMachine
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.Mtu != nil {
		payload["mtu"] = *requestDTO.Mtu
	} else {
		payload["mtu"] = nil
	}
	if requestDTO.Mode != nil {
		payload["mode"] = *requestDTO.Mode
	} else {
		payload["mode"] = nil
	}
	if requestDTO.UntaggedVlan != nil {
		payload["untagged_vlan"] = *requestDTO.UntaggedVlan
	} else {
		payload["untagged_vlan"] = nil
	}
	if requestDTO.TaggedVlans != nil {
		payload["tagged_vlans"] = requestDTO.TaggedVlans
	} else {
		payload["tagged_vlans"] = []any{}
	}
	if requestDTO.QinqSvlan != nil {
		payload["qinq_svlan"] = *requestDTO.QinqSvlan
	} else {
		payload["qinq_svlan"] = nil
	}
	if requestDTO.VlanTranslationPolicy != nil {
		payload["vlan_translation_policy"] = *requestDTO.VlanTranslationPolicy
	} else {
		payload["vlan_translation_policy"] = nil
	}
	if requestDTO.Vrf != nil {
		payload["vrf"] = *requestDTO.Vrf
	} else {
		payload["vrf"] = nil
	}
	if requestDTO.Parent != nil {
		payload["parent"] = *requestDTO.Parent
	} else {
		payload["parent"] = nil
	}
	if requestDTO.Bridge != nil {
		payload["bridge"] = *requestDTO.Bridge
	} else {
		payload["bridge"] = nil
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

// VirtualMachineInterfaceResponseDTO is the response DTO of the virtual_machine_interface resource, built from
// the go-netbox VMInterface by VirtualMachineInterfaceResponseDTOFromGoNetbox.
type VirtualMachineInterfaceResponseDTO struct {
	ID                    *int64            `json:"id,omitempty"`
	VirtualMachine        *int64            `json:"virtual_machine,omitempty"`
	Name                  *string           `json:"name,omitempty"`
	Enabled               *bool             `json:"enabled,omitempty"`
	Mtu                   *int64            `json:"mtu,omitempty"`
	Mode                  *string           `json:"mode,omitempty"`
	UntaggedVlan          *int64            `json:"untagged_vlan,omitempty"`
	TaggedVlans           []int64           `json:"tagged_vlans,omitempty"`
	QinqSvlan             *int64            `json:"qinq_svlan,omitempty"`
	VlanTranslationPolicy *int64            `json:"vlan_translation_policy,omitempty"`
	Vrf                   *int64            `json:"vrf,omitempty"`
	Parent                *int64            `json:"parent,omitempty"`
	Bridge                *int64            `json:"bridge,omitempty"`
	Description           *string           `json:"description,omitempty"`
	Owner                 *int64            `json:"owner,omitempty"`
	Created               *string           `json:"created,omitempty"`
	LastUpdated           *string           `json:"last_updated,omitempty"`
	URL                   *string           `json:"url,omitempty"`
	Tags                  []string          `json:"tags,omitempty"`
	TagsAll               []string          `json:"tags_all,omitempty"`
	CustomFields          map[string]string `json:"custom_fields,omitempty"`
}

// VirtualMachineInterfaceResponseDTOFromGoNetbox converts a *models.VMInterface to the response DTO.
func VirtualMachineInterfaceResponseDTOFromGoNetbox(goNetboxModel *models.VMInterface) *VirtualMachineInterfaceResponseDTO {
	responseDTO := &VirtualMachineInterfaceResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	if goNetboxModel.VirtualMachine != nil {
		v := goNetboxModel.VirtualMachine.ID
		responseDTO.VirtualMachine = &v
	}
	responseDTO.Name = goNetboxModel.Name
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	responseDTO.Mtu = goNetboxModel.Mtu
	if goNetboxModel.Mode != nil {
		responseDTO.Mode = choiceValue[string](goNetboxModel.Mode.Value)
	}
	if goNetboxModel.UntaggedVlan != nil {
		v := goNetboxModel.UntaggedVlan.ID
		responseDTO.UntaggedVlan = &v
	}
	if goNetboxModel.TaggedVlans != nil {
		responseDTO.TaggedVlans = []int64{}
		for _, ref := range goNetboxModel.TaggedVlans {
			if ref != nil {
				responseDTO.TaggedVlans = append(responseDTO.TaggedVlans, ref.ID)
			}
		}
	}
	if goNetboxModel.QinqSvlan != nil {
		v := goNetboxModel.QinqSvlan.ID
		responseDTO.QinqSvlan = &v
	}
	if goNetboxModel.VlanTranslationPolicy != nil {
		v := goNetboxModel.VlanTranslationPolicy.ID
		responseDTO.VlanTranslationPolicy = &v
	}
	if goNetboxModel.Vrf != nil {
		v := goNetboxModel.Vrf.ID
		responseDTO.Vrf = &v
	}
	if goNetboxModel.Parent != nil {
		v := goNetboxModel.Parent.ID
		responseDTO.Parent = &v
	}
	if goNetboxModel.Bridge != nil {
		v := goNetboxModel.Bridge.ID
		responseDTO.Bridge = &v
	}
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
