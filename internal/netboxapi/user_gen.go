// Code generated automatically. DO NOT EDIT.

package netboxapi

import (
	"github.com/fbreckle/go-netbox/netbox/models"
)

// UserRequestDTO is the request DTO of the user resource; Payload writes it into
// the JSON request body.
type UserRequestDTO struct {
	Username  *string `json:"username,omitempty"`
	Password  *string `json:"password,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
	Groups    []int64 `json:"groups,omitempty"`
}

// Payload returns the JSON request body of the user resource.
func (requestDTO *UserRequestDTO) Payload() map[string]any {
	payload := map[string]any{}
	if requestDTO.Username != nil {
		payload["username"] = *requestDTO.Username
	}
	if requestDTO.Password != nil {
		payload["password"] = *requestDTO.Password
	}
	if requestDTO.FirstName != nil {
		payload["first_name"] = *requestDTO.FirstName
	} else {
		payload["first_name"] = ""
	}
	if requestDTO.LastName != nil {
		payload["last_name"] = *requestDTO.LastName
	} else {
		payload["last_name"] = ""
	}
	if requestDTO.Email != nil {
		payload["email"] = *requestDTO.Email
	} else {
		payload["email"] = ""
	}
	if requestDTO.IsActive != nil {
		payload["is_active"] = *requestDTO.IsActive
	}
	if requestDTO.Groups != nil {
		payload["groups"] = requestDTO.Groups
	} else {
		payload["groups"] = []any{}
	}
	return payload
}

// UserResponseDTO is the response DTO of the user resource, built from
// the go-netbox User by UserResponseDTOFromGoNetbox.
type UserResponseDTO struct {
	ID         *int64  `json:"id,omitempty"`
	Username   *string `json:"username,omitempty"`
	FirstName  *string `json:"first_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	Email      *string `json:"email,omitempty"`
	IsActive   *bool   `json:"is_active,omitempty"`
	Groups     []int64 `json:"groups,omitempty"`
	DateJoined *string `json:"date_joined,omitempty"`
	LastLogin  *string `json:"last_login,omitempty"`
	URL        *string `json:"url,omitempty"`
}

// UserResponseDTOFromGoNetbox converts a *models.User to the response DTO.
func UserResponseDTOFromGoNetbox(goNetboxModel *models.User) *UserResponseDTO {
	responseDTO := &UserResponseDTO{}
	if goNetboxModel == nil {
		return responseDTO
	}
	{
		v := goNetboxModel.ID
		responseDTO.ID = &v
	}
	responseDTO.Username = goNetboxModel.Username
	if goNetboxModel.FirstName != "" {
		v := goNetboxModel.FirstName
		responseDTO.FirstName = &v
	}
	if goNetboxModel.LastName != "" {
		v := goNetboxModel.LastName
		responseDTO.LastName = &v
	}
	if goNetboxModel.Email != "" {
		v := string(goNetboxModel.Email)
		responseDTO.Email = &v
	}
	{
		v := goNetboxModel.IsActive
		responseDTO.IsActive = &v
	}
	if goNetboxModel.Groups != nil {
		responseDTO.Groups = []int64{}
		for _, ref := range goNetboxModel.Groups {
			if ref != nil {
				responseDTO.Groups = append(responseDTO.Groups, ref.ID)
			}
		}
	}
	if !goNetboxModel.DateJoined.IsZero() {
		v := goNetboxModel.DateJoined.String()
		responseDTO.DateJoined = &v
	}
	if goNetboxModel.LastLogin != nil {
		v := goNetboxModel.LastLogin.String()
		responseDTO.LastLogin = &v
	}
	if goNetboxModel.URL != "" {
		v := string(goNetboxModel.URL)
		responseDTO.URL = &v
	}
	return responseDTO
}
