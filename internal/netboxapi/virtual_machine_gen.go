// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualMachineRequestDTO is the request DTO of the virtual_machine resource; Payload writes it into
// the JSON request body.
type VirtualMachineRequestDTO struct {
	Name               *string           `json:"name,omitempty"`
	Status             *string           `json:"status,omitempty"`
	StartOnBoot        *string           `json:"start_on_boot,omitempty"`
	Cluster            *int64            `json:"cluster,omitempty"`
	Site               *int64            `json:"site,omitempty"`
	Role               *int64            `json:"role,omitempty"`
	VirtualMachineType *int64            `json:"virtual_machine_type,omitempty"`
	Tenant             *int64            `json:"tenant,omitempty"`
	Platform           *int64            `json:"platform,omitempty"`
	Device             *int64            `json:"device,omitempty"`
	ConfigTemplate     *int64            `json:"config_template,omitempty"`
	Serial             *string           `json:"serial,omitempty"`
	Vcpus              *float64          `json:"vcpus,omitempty"`
	Memory             *int64            `json:"memory,omitempty"`
	Disk               *int64            `json:"disk,omitempty"`
	LocalContextData   *string           `json:"local_context_data,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Comments           *string           `json:"comments,omitempty"`
	Owner              *int64            `json:"owner,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	CustomFields       map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the virtual_machine resource.
func (requestDTO *VirtualMachineRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.StartOnBoot != nil {
		payload["start_on_boot"] = *requestDTO.StartOnBoot
	}
	if requestDTO.Cluster != nil {
		payload["cluster"] = *requestDTO.Cluster
	} else {
		payload["cluster"] = nil
	}
	if requestDTO.Site != nil {
		payload["site"] = *requestDTO.Site
	} else {
		payload["site"] = nil
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	} else {
		payload["role"] = nil
	}
	if requestDTO.VirtualMachineType != nil {
		payload["virtual_machine_type"] = *requestDTO.VirtualMachineType
	} else {
		payload["virtual_machine_type"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Platform != nil {
		payload["platform"] = *requestDTO.Platform
	} else {
		payload["platform"] = nil
	}
	if requestDTO.Device != nil {
		payload["device"] = *requestDTO.Device
	} else {
		payload["device"] = nil
	}
	if requestDTO.ConfigTemplate != nil {
		payload["config_template"] = *requestDTO.ConfigTemplate
	} else {
		payload["config_template"] = nil
	}
	if requestDTO.Serial != nil {
		payload["serial"] = *requestDTO.Serial
	} else {
		payload["serial"] = ""
	}
	if requestDTO.Vcpus != nil {
		payload["vcpus"] = *requestDTO.Vcpus
	} else {
		payload["vcpus"] = nil
	}
	if requestDTO.Memory != nil {
		payload["memory"] = *requestDTO.Memory
	} else {
		payload["memory"] = nil
	}
	if requestDTO.Disk != nil {
		payload["disk"] = *requestDTO.Disk
	}
	if requestDTO.LocalContextData != nil {
		payload["local_context_data"] = parseJSONText(*requestDTO.LocalContextData)
	} else {
		payload["local_context_data"] = nil
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

// VirtualMachineResponseDTO is the response DTO of the virtual_machine resource, built from
// the go-netbox VirtualMachineWithConfigContext by VirtualMachineResponseDTOFromGoNetbox.
type VirtualMachineResponseDTO struct {
	ID                 *int64            `json:"id,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Status             *string           `json:"status,omitempty"`
	StartOnBoot        *string           `json:"start_on_boot,omitempty"`
	Cluster            *int64            `json:"cluster,omitempty"`
	Site               *int64            `json:"site,omitempty"`
	Role               *int64            `json:"role,omitempty"`
	VirtualMachineType *int64            `json:"virtual_machine_type,omitempty"`
	Tenant             *int64            `json:"tenant,omitempty"`
	Platform           *int64            `json:"platform,omitempty"`
	Device             *int64            `json:"device,omitempty"`
	ConfigTemplate     *int64            `json:"config_template,omitempty"`
	Serial             *string           `json:"serial,omitempty"`
	PrimaryIp4         *int64            `json:"primary_ip4,omitempty"`
	PrimaryIp6         *int64            `json:"primary_ip6,omitempty"`
	Vcpus              *float64          `json:"vcpus,omitempty"`
	Memory             *int64            `json:"memory,omitempty"`
	Disk               *int64            `json:"disk,omitempty"`
	LocalContextData   *string           `json:"local_context_data,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Comments           *string           `json:"comments,omitempty"`
	Owner              *int64            `json:"owner,omitempty"`
	Created            *string           `json:"created,omitempty"`
	LastUpdated        *string           `json:"last_updated,omitempty"`
	URL                *string           `json:"url,omitempty"`
	InterfaceCount     *int64            `json:"interface_count,omitempty"`
	VirtualDiskCount   *int64            `json:"virtual_disk_count,omitempty"`
	Tags               []string          `json:"tags,omitempty"`
	TagsAll            []string          `json:"tags_all,omitempty"`
	CustomFields       map[string]string `json:"custom_fields,omitempty"`
}

// VirtualMachineResponseDTOFromGoNetbox converts a *models.VirtualMachineWithConfigContext to the response DTO.
func VirtualMachineResponseDTOFromGoNetbox(goNetboxModel *models.VirtualMachineWithConfigContext) *VirtualMachineResponseDTO {
	responseDTO := &VirtualMachineResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.StartOnBoot != nil {
		responseDTO.StartOnBoot = choiceValue[string](goNetboxModel.StartOnBoot.Value)
	}
	if goNetboxModel.Cluster != nil {
		v := goNetboxModel.Cluster.ID
		responseDTO.Cluster = &v
	}
	if goNetboxModel.Site != nil {
		v := goNetboxModel.Site.ID
		responseDTO.Site = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.VirtualMachineType != nil {
		v := goNetboxModel.VirtualMachineType.ID
		responseDTO.VirtualMachineType = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Platform != nil {
		v := goNetboxModel.Platform.ID
		responseDTO.Platform = &v
	}
	if goNetboxModel.Device != nil {
		v := goNetboxModel.Device.ID
		responseDTO.Device = &v
	}
	if goNetboxModel.ConfigTemplate != nil {
		v := goNetboxModel.ConfigTemplate.ID
		responseDTO.ConfigTemplate = &v
	}
	if goNetboxModel.Serial != "" {
		v := goNetboxModel.Serial
		responseDTO.Serial = &v
	}
	if goNetboxModel.PrimaryIp4 != nil {
		v := goNetboxModel.PrimaryIp4.ID
		responseDTO.PrimaryIp4 = &v
	}
	if goNetboxModel.PrimaryIp6 != nil {
		v := goNetboxModel.PrimaryIp6.ID
		responseDTO.PrimaryIp6 = &v
	}
	responseDTO.Vcpus = goNetboxModel.Vcpus
	responseDTO.Memory = goNetboxModel.Memory
	responseDTO.Disk = goNetboxModel.Disk
	responseDTO.LocalContextData = jsonText(goNetboxModel.LocalContextData)
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
	{
		v := goNetboxModel.InterfaceCount
		responseDTO.InterfaceCount = &v
	}
	{
		v := goNetboxModel.VirtualDiskCount
		responseDTO.VirtualDiskCount = &v
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
