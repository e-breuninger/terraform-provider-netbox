// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// AvailableVlanRequestDTO is the request DTO of the available_vlan resource; Payload writes it into
// the JSON request body.
type AvailableVlanRequestDTO struct {
	Name         *string           `json:"name,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Site         *int64            `json:"site,omitempty"`
	Group        *int64            `json:"group,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Role         *int64            `json:"role,omitempty"`
	QinqRole     *string           `json:"qinq_role,omitempty"`
	QinqSvlan    *int64            `json:"qinq_svlan,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// Payload returns the JSON request body of the available_vlan resource.
func (requestDTO *AvailableVlanRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Status != nil {
		payload["status"] = *requestDTO.Status
	}
	if requestDTO.Site != nil {
		payload["site"] = *requestDTO.Site
	} else {
		payload["site"] = nil
	}
	if requestDTO.Group != nil {
		payload["group"] = *requestDTO.Group
	}
	if requestDTO.Tenant != nil {
		payload["tenant"] = *requestDTO.Tenant
	} else {
		payload["tenant"] = nil
	}
	if requestDTO.Role != nil {
		payload["role"] = *requestDTO.Role
	} else {
		payload["role"] = nil
	}
	if requestDTO.QinqRole != nil {
		payload["qinq_role"] = *requestDTO.QinqRole
	} else {
		payload["qinq_role"] = nil
	}
	if requestDTO.QinqSvlan != nil {
		payload["qinq_svlan"] = *requestDTO.QinqSvlan
	} else {
		payload["qinq_svlan"] = nil
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

// AvailableVlanResponseDTO is the response DTO of the available_vlan resource, built from
// the go-netbox VLAN by AvailableVlanResponseDTOFromGoNetbox.
type AvailableVlanResponseDTO struct {
	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Vid          *int64            `json:"vid,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Site         *int64            `json:"site,omitempty"`
	Group        *int64            `json:"group,omitempty"`
	Tenant       *int64            `json:"tenant,omitempty"`
	Role         *int64            `json:"role,omitempty"`
	QinqRole     *string           `json:"qinq_role,omitempty"`
	QinqSvlan    *int64            `json:"qinq_svlan,omitempty"`
	Description  *string           `json:"description,omitempty"`
	Comments     *string           `json:"comments,omitempty"`
	Owner        *int64            `json:"owner,omitempty"`
	Created      *string           `json:"created,omitempty"`
	LastUpdated  *string           `json:"last_updated,omitempty"`
	URL          *string           `json:"url,omitempty"`
	PrefixCount  *int64            `json:"prefix_count,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	TagsAll      []string          `json:"tags_all,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

// AvailableVlanResponseDTOFromGoNetbox converts a *models.VLAN to the response DTO.
func AvailableVlanResponseDTOFromGoNetbox(goNetboxModel *models.VLAN) *AvailableVlanResponseDTO {
	responseDTO := &AvailableVlanResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	responseDTO.Vid = goNetboxModel.Vid
	if goNetboxModel.Status != nil {
		responseDTO.Status = choiceValue[string](goNetboxModel.Status.Value)
	}
	if goNetboxModel.Site != nil {
		v := goNetboxModel.Site.ID
		responseDTO.Site = &v
	}
	if goNetboxModel.Group != nil {
		v := goNetboxModel.Group.ID
		responseDTO.Group = &v
	}
	if goNetboxModel.Tenant != nil {
		v := goNetboxModel.Tenant.ID
		responseDTO.Tenant = &v
	}
	if goNetboxModel.Role != nil {
		v := goNetboxModel.Role.ID
		responseDTO.Role = &v
	}
	if goNetboxModel.QinqRole != nil {
		responseDTO.QinqRole = choiceValue[string](goNetboxModel.QinqRole.Value)
	}
	if goNetboxModel.QinqSvlan != nil {
		v := goNetboxModel.QinqSvlan.ID
		responseDTO.QinqSvlan = &v
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
	{
		v := goNetboxModel.PrefixCount
		responseDTO.PrefixCount = &v
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
