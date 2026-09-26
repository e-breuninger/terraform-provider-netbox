// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// NotificationGroupRequestDTO is the request DTO of the notification_group resource; Payload writes it into
// the JSON request body.
type NotificationGroupRequestDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Groups      []int64 `json:"groups,omitempty"`
	Users       []int64 `json:"users,omitempty"`
}

// Payload returns the JSON request body of the notification_group resource.
func (requestDTO *NotificationGroupRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Name != nil {
		payload["name"] = *requestDTO.Name
	}
	if requestDTO.Description != nil {
		payload["description"] = *requestDTO.Description
	} else {
		payload["description"] = ""
	}
	if requestDTO.Groups != nil {
		payload["groups"] = requestDTO.Groups
	} else {
		payload["groups"] = []any{}
	}
	if requestDTO.Users != nil {
		payload["users"] = requestDTO.Users
	} else {
		payload["users"] = []any{}
	}
	return payload
}

// NotificationGroupResponseDTO is the response DTO of the notification_group resource, built from
// the go-netbox NotificationGroup by NotificationGroupResponseDTOFromGoNetbox.
type NotificationGroupResponseDTO struct {
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Groups      []int64 `json:"groups,omitempty"`
	Users       []int64 `json:"users,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// NotificationGroupResponseDTOFromGoNetbox converts a *models.NotificationGroup to the response DTO.
func NotificationGroupResponseDTOFromGoNetbox(goNetboxModel *models.NotificationGroup) *NotificationGroupResponseDTO {
	responseDTO := &NotificationGroupResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Name = goNetboxModel.Name
	if goNetboxModel.Description != "" {
		v := goNetboxModel.Description
		responseDTO.Description = &v
	}
	if goNetboxModel.Groups != nil {
		responseDTO.Groups = []int64{}
		for _, ref := range goNetboxModel.Groups {
			if ref != nil {
				responseDTO.Groups = append(responseDTO.Groups, ref.ID)
			}
		}
	}
	if goNetboxModel.Users != nil {
		responseDTO.Users = []int64{}
		for _, ref := range goNetboxModel.Users {
			if ref != nil {
				responseDTO.Users = append(responseDTO.Users, ref.ID)
			}
		}
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
