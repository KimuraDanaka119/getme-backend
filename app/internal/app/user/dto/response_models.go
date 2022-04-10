package dto

import (
	"getme-backend/internal/app/user/entities"
)

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type UserResponse struct {
	Nickname string `json:"nickname"`
	Fullname string `json:"fullname"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
}

func ToUserResponse(usr *entities.User) *UserResponse {
	return &UserResponse{
		Nickname: usr.Nickname,
		Fullname: usr.Fullname,
		About:    usr.About,
		Email:    usr.Email,
	}
}

//easyjson:json
type UsersResponse []UserResponse

func ToUsersResponse(usrs []entities.User) *UsersResponse {
	res := UsersResponse{}
	for _, usr := range usrs {
		res = append(res, *ToUserResponse(&usr))
	}
	return &res
}
