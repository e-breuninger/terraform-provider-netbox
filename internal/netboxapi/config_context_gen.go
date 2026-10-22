// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// ConfigContextRequestDTO is the request DTO of the config_context resource; Payload writes it into
// the JSON request body.
type ConfigContextRequestDTO struct {
	Name          *string  `json:"name,omitempty"`
	Weight        *int64   `json:"weight,omitempty"`
	Description   *string  `json:"description,omitempty"`
	IsActive      *bool    `json:"is_active,omitempty"`
	Data          *string  `json:"data,omitempty"`
	Profile       *int64   `json:"profile,omitempty"`
	Regions       []int64  `json:"regions,omitempty"`
	SiteGroups    []int64  `json:"site_groups,omitempty"`
	Sites         []int64  `json:"sites,omitempty"`
	Locations     []int64  `json:"locations,omitempty"`
	DeviceTypes   []int64  `json:"device_types,omitempty"`
	Roles         []int64  `json:"roles,omitempty"`
	Platforms     []int64  `json:"platforms,omitempty"`
	ClusterTypes  []int64  `json:"cluster_types,omitempty"`
	ClusterGroups []int64  `json:"cluster_groups,omitempty"`
	Clusters      []int64  `json:"clusters,omitempty"`
	TenantGroups  []int64  `json:"tenant_groups,omitempty"`
	Tenants       []int64  `json:"tenants,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Owner         *int64   `json:"owner,omitempty"`
}

// Payload returns the JSON request body of the config_context resource.
func (requestDTO *ConfigContextRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Weight != nil {
		payload["weight"] = *requestDTO.Weight
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.IsActive != nil {
		payload["is_active"] = *requestDTO.IsActive
	}
	if requestDTO.Data != nil {
		payload["data"] = parseJSONText(*requestDTO.Data)
	}
	if requestDTO.Profile != nil {
		payload["profile"] = *requestDTO.Profile
	} else {
		payload["profile"] = nil
	}
	if requestDTO.Regions != nil {
		payload["regions"] = requestDTO.Regions
	} else {
		payload["regions"] = []any{}
	}
	if requestDTO.SiteGroups != nil {
		payload["site_groups"] = requestDTO.SiteGroups
	} else {
		payload["site_groups"] = []any{}
	}
	if requestDTO.Sites != nil {
		payload["sites"] = requestDTO.Sites
	} else {
		payload["sites"] = []any{}
	}
	if requestDTO.Locations != nil {
		payload["locations"] = requestDTO.Locations
	} else {
		payload["locations"] = []any{}
	}
	if requestDTO.DeviceTypes != nil {
		payload["device_types"] = requestDTO.DeviceTypes
	} else {
		payload["device_types"] = []any{}
	}
	if requestDTO.Roles != nil {
		payload["roles"] = requestDTO.Roles
	} else {
		payload["roles"] = []any{}
	}
	if requestDTO.Platforms != nil {
		payload["platforms"] = requestDTO.Platforms
	} else {
		payload["platforms"] = []any{}
	}
	if requestDTO.ClusterTypes != nil {
		payload["cluster_types"] = requestDTO.ClusterTypes
	} else {
		payload["cluster_types"] = []any{}
	}
	if requestDTO.ClusterGroups != nil {
		payload["cluster_groups"] = requestDTO.ClusterGroups
	} else {
		payload["cluster_groups"] = []any{}
	}
	if requestDTO.Clusters != nil {
		payload["clusters"] = requestDTO.Clusters
	} else {
		payload["clusters"] = []any{}
	}
	if requestDTO.TenantGroups != nil {
		payload["tenant_groups"] = requestDTO.TenantGroups
	} else {
		payload["tenant_groups"] = []any{}
	}
	if requestDTO.Tenants != nil {
		payload["tenants"] = requestDTO.Tenants
	} else {
		payload["tenants"] = []any{}
	}
	if requestDTO.Tags != nil {
		payload["tags"] = requestDTO.Tags
	} else {
		payload["tags"] = []any{}
	}
	if requestDTO.Owner != nil {
		payload["owner"] = *requestDTO.Owner
	} else {
		payload["owner"] = nil
	}
	return payload
}

// ConfigContextResponseDTO is the response DTO of the config_context resource, built from
// the go-netbox ConfigContext by ConfigContextResponseDTOFromGoNetbox.
type ConfigContextResponseDTO struct {
	ID            *int64   `json:"id,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Weight        *int64   `json:"weight,omitempty"`
	Description   *string  `json:"description,omitempty"`
	IsActive      *bool    `json:"is_active,omitempty"`
	Data          *string  `json:"data,omitempty"`
	Profile       *int64   `json:"profile,omitempty"`
	Regions       []int64  `json:"regions,omitempty"`
	SiteGroups    []int64  `json:"site_groups,omitempty"`
	Sites         []int64  `json:"sites,omitempty"`
	Locations     []int64  `json:"locations,omitempty"`
	DeviceTypes   []int64  `json:"device_types,omitempty"`
	Roles         []int64  `json:"roles,omitempty"`
	Platforms     []int64  `json:"platforms,omitempty"`
	ClusterTypes  []int64  `json:"cluster_types,omitempty"`
	ClusterGroups []int64  `json:"cluster_groups,omitempty"`
	Clusters      []int64  `json:"clusters,omitempty"`
	TenantGroups  []int64  `json:"tenant_groups,omitempty"`
	Tenants       []int64  `json:"tenants,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Owner         *int64   `json:"owner,omitempty"`
	Created       *string  `json:"created,omitempty"`
	LastUpdated   *string  `json:"last_updated,omitempty"`
	URL           *string  `json:"url,omitempty"`
}

// ConfigContextResponseDTOFromGoNetbox converts a *models.ConfigContext to the response DTO.
func ConfigContextResponseDTOFromGoNetbox(goNetboxModel *models.ConfigContext) *ConfigContextResponseDTO {
	responseDTO := &ConfigContextResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Weight = goNetboxModel.Weight
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	{
		v := goNetboxModel.IsActive
		responseDTO.IsActive = &v
	}
	responseDTO.Data = jsonText(goNetboxModel.Data)
	if goNetboxModel.Profile != nil {
		v := goNetboxModel.Profile.ID
		responseDTO.Profile = &v
	}
	if goNetboxModel.Regions != nil {
		responseDTO.Regions = []int64{}
		for _, ref := range goNetboxModel.Regions {
			if ref != nil {
				responseDTO.Regions = append(responseDTO.Regions, ref.ID)
			}
		}
	}
	if goNetboxModel.SiteGroups != nil {
		responseDTO.SiteGroups = []int64{}
		for _, ref := range goNetboxModel.SiteGroups {
			if ref != nil {
				responseDTO.SiteGroups = append(responseDTO.SiteGroups, ref.ID)
			}
		}
	}
	if goNetboxModel.Sites != nil {
		responseDTO.Sites = []int64{}
		for _, ref := range goNetboxModel.Sites {
			if ref != nil {
				responseDTO.Sites = append(responseDTO.Sites, ref.ID)
			}
		}
	}
	if goNetboxModel.Locations != nil {
		responseDTO.Locations = []int64{}
		for _, ref := range goNetboxModel.Locations {
			if ref != nil {
				responseDTO.Locations = append(responseDTO.Locations, ref.ID)
			}
		}
	}
	if goNetboxModel.DeviceTypes != nil {
		responseDTO.DeviceTypes = []int64{}
		for _, ref := range goNetboxModel.DeviceTypes {
			if ref != nil {
				responseDTO.DeviceTypes = append(responseDTO.DeviceTypes, ref.ID)
			}
		}
	}
	if goNetboxModel.Roles != nil {
		responseDTO.Roles = []int64{}
		for _, ref := range goNetboxModel.Roles {
			if ref != nil {
				responseDTO.Roles = append(responseDTO.Roles, ref.ID)
			}
		}
	}
	if goNetboxModel.Platforms != nil {
		responseDTO.Platforms = []int64{}
		for _, ref := range goNetboxModel.Platforms {
			if ref != nil {
				responseDTO.Platforms = append(responseDTO.Platforms, ref.ID)
			}
		}
	}
	if goNetboxModel.ClusterTypes != nil {
		responseDTO.ClusterTypes = []int64{}
		for _, ref := range goNetboxModel.ClusterTypes {
			if ref != nil {
				responseDTO.ClusterTypes = append(responseDTO.ClusterTypes, ref.ID)
			}
		}
	}
	if goNetboxModel.ClusterGroups != nil {
		responseDTO.ClusterGroups = []int64{}
		for _, ref := range goNetboxModel.ClusterGroups {
			if ref != nil {
				responseDTO.ClusterGroups = append(responseDTO.ClusterGroups, ref.ID)
			}
		}
	}
	if goNetboxModel.Clusters != nil {
		responseDTO.Clusters = []int64{}
		for _, ref := range goNetboxModel.Clusters {
			if ref != nil {
				responseDTO.Clusters = append(responseDTO.Clusters, ref.ID)
			}
		}
	}
	if goNetboxModel.TenantGroups != nil {
		responseDTO.TenantGroups = []int64{}
		for _, ref := range goNetboxModel.TenantGroups {
			if ref != nil {
				responseDTO.TenantGroups = append(responseDTO.TenantGroups, ref.ID)
			}
		}
	}
	if goNetboxModel.Tenants != nil {
		responseDTO.Tenants = []int64{}
		for _, ref := range goNetboxModel.Tenants {
			if ref != nil {
				responseDTO.Tenants = append(responseDTO.Tenants, ref.ID)
			}
		}
	}
	responseDTO.Tags = goNetboxModel.Tags
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
