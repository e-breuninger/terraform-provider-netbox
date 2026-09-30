// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// DeviceInterfaceRequestDTO is the request DTO of the device_interface resource; Payload writes it into
// the JSON request body.
type DeviceInterfaceRequestDTO struct {
	Device                *int64            `json:"device,omitempty"`
	Name                  *string           `json:"name,omitempty"`
	Type                  *string           `json:"type,omitempty"`
	Label                 *string           `json:"label,omitempty"`
	Enabled               *bool             `json:"enabled,omitempty"`
	MgmtOnly              *bool             `json:"mgmt_only,omitempty"`
	MarkConnected         *bool             `json:"mark_connected,omitempty"`
	Mtu                   *int64            `json:"mtu,omitempty"`
	Speed                 *int64            `json:"speed,omitempty"`
	Duplex                *string           `json:"duplex,omitempty"`
	PoeMode               *string           `json:"poe_mode,omitempty"`
	PoeType               *string           `json:"poe_type,omitempty"`
	RfRole                *string           `json:"rf_role,omitempty"`
	RfChannel             *string           `json:"rf_channel,omitempty"`
	RfChannelFrequency    *float64          `json:"rf_channel_frequency,omitempty"`
	RfChannelWidth        *float64          `json:"rf_channel_width,omitempty"`
	TxPower               *int64            `json:"tx_power,omitempty"`
	Wwn                   *string           `json:"wwn,omitempty"`
	Mode                  *string           `json:"mode,omitempty"`
	UntaggedVlan          *int64            `json:"untagged_vlan,omitempty"`
	TaggedVlans           []int64           `json:"tagged_vlans,omitempty"`
	QinqSvlan             *int64            `json:"qinq_svlan,omitempty"`
	VlanTranslationPolicy *int64            `json:"vlan_translation_policy,omitempty"`
	Module                *int64            `json:"module,omitempty"`
	Lag                   *int64            `json:"lag,omitempty"`
	Parent                *int64            `json:"parent,omitempty"`
	Bridge                *int64            `json:"bridge,omitempty"`
	Vrf                   *int64            `json:"vrf,omitempty"`
	Vdcs                  []int64           `json:"vdcs,omitempty"`
	WirelessLans          []int64           `json:"wireless_lans,omitempty"`
	Description           *string           `json:"description,omitempty"`
	Owner                 *int64            `json:"owner,omitempty"`
	Tags                  []string          `json:"tags,omitempty"`
	CustomFields          map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the device_interface resource.
func (requestDTO *DeviceInterfaceRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.Label != nil {
		payload["label"] = *requestDTO.Label
	} else {
		payload["label"] = ""
	}
	if requestDTO.Enabled != nil {
		payload["enabled"] = *requestDTO.Enabled
	}
	if requestDTO.MgmtOnly != nil {
		payload["mgmt_only"] = *requestDTO.MgmtOnly
	}
	if requestDTO.MarkConnected != nil {
		payload["mark_connected"] = *requestDTO.MarkConnected
	}
	if requestDTO.Mtu != nil {
		payload["mtu"] = *requestDTO.Mtu
	} else {
		payload["mtu"] = nil
	}
	if requestDTO.Speed != nil {
		payload["speed"] = *requestDTO.Speed
	} else {
		payload["speed"] = nil
	}
	if requestDTO.Duplex != nil {
		payload["duplex"] = *requestDTO.Duplex
	} else {
		payload["duplex"] = nil
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
	if requestDTO.RfChannel != nil {
		payload["rf_channel"] = *requestDTO.RfChannel
	} else {
		payload["rf_channel"] = nil
	}
	if requestDTO.RfChannelFrequency != nil {
		payload["rf_channel_frequency"] = *requestDTO.RfChannelFrequency
	} else {
		payload["rf_channel_frequency"] = nil
	}
	if requestDTO.RfChannelWidth != nil {
		payload["rf_channel_width"] = *requestDTO.RfChannelWidth
	} else {
		payload["rf_channel_width"] = nil
	}
	if requestDTO.TxPower != nil {
		payload["tx_power"] = *requestDTO.TxPower
	} else {
		payload["tx_power"] = nil
	}
	if requestDTO.Wwn != nil {
		payload["wwn"] = *requestDTO.Wwn
	} else {
		payload["wwn"] = nil
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
	if requestDTO.Module != nil {
		payload["module"] = *requestDTO.Module
	} else {
		payload["module"] = nil
	}
	if requestDTO.Lag != nil {
		payload["lag"] = *requestDTO.Lag
	} else {
		payload["lag"] = nil
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
	if requestDTO.Vrf != nil {
		payload["vrf"] = *requestDTO.Vrf
	} else {
		payload["vrf"] = nil
	}
	if requestDTO.Vdcs != nil {
		payload["vdcs"] = requestDTO.Vdcs
	} else {
		payload["vdcs"] = []any{}
	}
	if requestDTO.WirelessLans != nil {
		payload["wireless_lans"] = requestDTO.WirelessLans
	} else {
		payload["wireless_lans"] = []any{}
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

// DeviceInterfaceResponseDTO is the response DTO of the device_interface resource, built from
// the go-netbox Interface by DeviceInterfaceResponseDTOFromGoNetbox.
type DeviceInterfaceResponseDTO struct {
	ID                    *int64            `json:"id,omitempty"`
	Device                *int64            `json:"device,omitempty"`
	Name                  *string           `json:"name,omitempty"`
	Type                  *string           `json:"type,omitempty"`
	Label                 *string           `json:"label,omitempty"`
	Enabled               *bool             `json:"enabled,omitempty"`
	MgmtOnly              *bool             `json:"mgmt_only,omitempty"`
	MarkConnected         *bool             `json:"mark_connected,omitempty"`
	Mtu                   *int64            `json:"mtu,omitempty"`
	Speed                 *int64            `json:"speed,omitempty"`
	Duplex                *string           `json:"duplex,omitempty"`
	PoeMode               *string           `json:"poe_mode,omitempty"`
	PoeType               *string           `json:"poe_type,omitempty"`
	RfRole                *string           `json:"rf_role,omitempty"`
	RfChannel             *string           `json:"rf_channel,omitempty"`
	RfChannelFrequency    *float64          `json:"rf_channel_frequency,omitempty"`
	RfChannelWidth        *float64          `json:"rf_channel_width,omitempty"`
	TxPower               *int64            `json:"tx_power,omitempty"`
	Wwn                   *string           `json:"wwn,omitempty"`
	Mode                  *string           `json:"mode,omitempty"`
	UntaggedVlan          *int64            `json:"untagged_vlan,omitempty"`
	TaggedVlans           []int64           `json:"tagged_vlans,omitempty"`
	QinqSvlan             *int64            `json:"qinq_svlan,omitempty"`
	VlanTranslationPolicy *int64            `json:"vlan_translation_policy,omitempty"`
	Module                *int64            `json:"module,omitempty"`
	Lag                   *int64            `json:"lag,omitempty"`
	Parent                *int64            `json:"parent,omitempty"`
	Bridge                *int64            `json:"bridge,omitempty"`
	Vrf                   *int64            `json:"vrf,omitempty"`
	Vdcs                  []int64           `json:"vdcs,omitempty"`
	WirelessLans          []int64           `json:"wireless_lans,omitempty"`
	Description           *string           `json:"description,omitempty"`
	Owner                 *int64            `json:"owner,omitempty"`
	Created               *string           `json:"created,omitempty"`
	LastUpdated           *string           `json:"last_updated,omitempty"`
	URL                   *string           `json:"url,omitempty"`
	Tags                  []string          `json:"tags,omitempty"`
	TagsAll               []string          `json:"tags_all,omitempty"`
	CustomFields          map[string]string `json:"custom_fields,omitempty"`
}

// DeviceInterfaceResponseDTOFromGoNetbox converts a *models.Interface to the response DTO.
func DeviceInterfaceResponseDTOFromGoNetbox(goNetboxModel *models.Interface) *DeviceInterfaceResponseDTO {
	responseDTO := &DeviceInterfaceResponseDTO{}
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
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
	if goNetboxModel.Label != "" {
		v := goNetboxModel.Label
		responseDTO.Label = &v
	}
	{
		v := goNetboxModel.Enabled
		responseDTO.Enabled = &v
	}
	{
		v := goNetboxModel.MgmtOnly
		responseDTO.MgmtOnly = &v
	}
	{
		v := goNetboxModel.MarkConnected
		responseDTO.MarkConnected = &v
	}
	responseDTO.Mtu = goNetboxModel.Mtu
	responseDTO.Speed = goNetboxModel.Speed
	if goNetboxModel.Duplex != nil {
		responseDTO.Duplex = choiceValue[string](goNetboxModel.Duplex.Value)
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
	if goNetboxModel.RfChannel != nil {
		responseDTO.RfChannel = choiceValue[string](goNetboxModel.RfChannel.Value)
	}
	responseDTO.RfChannelFrequency = goNetboxModel.RfChannelFrequency
	responseDTO.RfChannelWidth = goNetboxModel.RfChannelWidth
	responseDTO.TxPower = goNetboxModel.TxPower
	responseDTO.Wwn = goNetboxModel.Wwn
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
	if goNetboxModel.Module != nil {
		v := goNetboxModel.Module.ID
		responseDTO.Module = &v
	}
	if goNetboxModel.Lag != nil {
		v := goNetboxModel.Lag.ID
		responseDTO.Lag = &v
	}
	if goNetboxModel.Parent != nil {
		v := goNetboxModel.Parent.ID
		responseDTO.Parent = &v
	}
	if goNetboxModel.Bridge != nil {
		v := goNetboxModel.Bridge.ID
		responseDTO.Bridge = &v
	}
	if goNetboxModel.Vrf != nil {
		v := goNetboxModel.Vrf.ID
		responseDTO.Vrf = &v
	}
	if goNetboxModel.Vdcs != nil {
		responseDTO.Vdcs = []int64{}
		for _, ref := range goNetboxModel.Vdcs {
			if ref != nil {
				responseDTO.Vdcs = append(responseDTO.Vdcs, ref.ID)
			}
		}
	}
	if goNetboxModel.WirelessLans != nil {
		responseDTO.WirelessLans = []int64{}
		for _, ref := range goNetboxModel.WirelessLans {
			if ref != nil {
				responseDTO.WirelessLans = append(responseDTO.WirelessLans, ref.ID)
			}
		}
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
