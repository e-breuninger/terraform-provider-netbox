// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// L2vpnRequestDTO is the request DTO of the l2vpn resource; Payload writes it into
// the JSON request body.
type L2vpnRequestDTO struct {
	Name          *string           `json:"name,omitempty"`
	Slug          *string           `json:"slug,omitempty"`
	Type          *string           `json:"type,omitempty"`
	Identifier    *int64            `json:"identifier,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	ImportTargets []int64           `json:"import_targets,omitempty"`
	ExportTargets []int64           `json:"export_targets,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the l2vpn resource.
func (requestDTO *L2vpnRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Slug != nil {
		payload["slug"] = *requestDTO.Slug
	} else if requestDTO.Name != nil {
		payload["slug"] = Slugify(*requestDTO.Name)
	}
	if requestDTO.Type != nil {
		payload["type"] = *requestDTO.Type
	}
	if requestDTO.Identifier != nil {
		payload["identifier"] = *requestDTO.Identifier
	} else {
		payload["identifier"] = nil
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.ImportTargets != nil {
		payload["import_targets"] = requestDTO.ImportTargets
	} else {
		payload["import_targets"] = []any{}
	}
	if requestDTO.ExportTargets != nil {
		payload["export_targets"] = requestDTO.ExportTargets
	} else {
		payload["export_targets"] = []any{}
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

// L2vpnResponseDTO is the response DTO of the l2vpn resource, built from
// the go-netbox L2VPN by L2vpnResponseDTOFromGoNetbox.
type L2vpnResponseDTO struct {
	ID            *int64            `json:"id,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Slug          *string           `json:"slug,omitempty"`
	Type          *string           `json:"type,omitempty"`
	Identifier    *int64            `json:"identifier,omitempty"`
	Tenant        *int64            `json:"tenant,omitempty"`
	ImportTargets []int64           `json:"import_targets,omitempty"`
	ExportTargets []int64           `json:"export_targets,omitempty"`
	Description   *string           `json:"description,omitempty"`
	Comments      *string           `json:"comments,omitempty"`
	Owner         *int64            `json:"owner,omitempty"`
	Created       *string           `json:"created,omitempty"`
	LastUpdated   *string           `json:"last_updated,omitempty"`
	URL           *string           `json:"url,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	TagsAll       []string          `json:"tags_all,omitempty"`
	CustomFields  map[string]string `json:"custom_fields,omitempty"`
}

// L2vpnResponseDTOFromGoNetbox converts a *models.L2VPN to the response DTO.
func L2vpnResponseDTOFromGoNetbox(goNetboxModel *models.L2VPN) *L2vpnResponseDTO {
	responseDTO := &L2vpnResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Slug = goNetboxModel.Slug
	if goNetboxModel.Type != nil {
		responseDTO.Type = choiceValue[string](goNetboxModel.Type.Value)
	}
	responseDTO.Identifier = goNetboxModel.Identifier
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.ImportTargets != nil {
		responseDTO.ImportTargets = []int64{}
		for _, ref := range goNetboxModel.ImportTargets {
			if ref != nil {
				responseDTO.ImportTargets = append(responseDTO.ImportTargets, ref.ID)
			}
		}
	}
	if goNetboxModel.ExportTargets != nil {
		responseDTO.ExportTargets = []int64{}
		for _, ref := range goNetboxModel.ExportTargets {
			if ref != nil {
				responseDTO.ExportTargets = append(responseDTO.ExportTargets, ref.ID)
			}
		}
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
