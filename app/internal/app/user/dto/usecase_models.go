package dto

import (
	"database/sql"
	"time"

	"getme-backend/internal/app/user/entities"
)

//go:generate easyjson -all -disallow_unknown_fields usecase_models.go

type UserUsecase struct {
	ID           int64     `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Nickname     string    `db:"nickname"`
	About        string    `db:"about"`
	Avatar       string    `db:"avatar"`
	TgTag        string    `db:"tg_tag"`
	Email        string    `db:"email"`
	IsSearchable bool      `db:"is_searchable"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

//easyjson:json
type UserWithSkillsUsecaseSlice []UserWithSkillsUsecase

type UserWithSkillsUsecase struct {
	UserUsecase
	Skills []string
}

type UserWithOfferIDUsecase struct {
	UserUsecase
	OfferID int64
}
type UserStatusUsecase struct {
	UserID   int64
	IsMentor bool
}

func (m *UserStatusUsecase) ToResponseStatus() UserStatusResponse {
	return UserStatusResponse{
		IsMentor: m.IsMentor,
	}
}

func (m *UserWithSkillsUsecase) ToUserWithSkillEntity() *entities_user.UserWithSkills {
	return &entities_user.UserWithSkills{
		User:   *m.UserUsecase.ToUserEntity(),
		Skills: m.Skills,
	}
}

func (m *UserUsecase) ToUserEntity() *entities_user.User {
	return &entities_user.User{
		ID: sql.NullInt64{
			Int64: m.ID,
		},
		FirstName: sql.NullString{
			String: m.FirstName,
		},
		LastName: sql.NullString{
			String: m.LastName,
		},
		Nickname: m.Nickname,
		Avatar: sql.NullString{
			String: m.Avatar,
		},
		TgTag: sql.NullString{
			String: m.TgTag,
		},
	}
}

func ToUserWithOfferIDUsecases(data []entities_user.UserWithOfferID) []UserWithOfferIDUsecase {
	res := make([]UserWithOfferIDUsecase, 0, len(data))
	for _, val := range data {
		res = append(res, *ToUserWithOfferIDUsecase(&val))
	}
	return res
}

func ToUserUsecase(data *entities_user.User) *UserUsecase {
	return &UserUsecase{
		ID:           data.ID.Int64,
		FirstName:    data.FirstName.String,
		LastName:     data.LastName.String,
		Nickname:     data.Nickname,
		About:        data.About.String,
		TgTag:        data.TgTag.String,
		Avatar:       data.Avatar.String,
		Email:        data.Email.String,
		IsSearchable: data.IsSearchable,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToUserWithSkillsUsecase(data *entities_user.UserWithSkills) *UserWithSkillsUsecase {
	return &UserWithSkillsUsecase{
		UserUsecase: *ToUserUsecase(&data.User),
		Skills:      data.Skills,
	}
}
func ToUserWithOfferIDUsecase(data *entities_user.UserWithOfferID) *UserWithOfferIDUsecase {
	u := ToUserUsecase(&data.User)
	return &UserWithOfferIDUsecase{
		UserUsecase: *u,
		OfferID:     data.OfferID,
	}
}

func ToUsersWithSkillUsecase(data []entities_user.UserWithSkills) []UserWithSkillsUsecase {
	res := make([]UserWithSkillsUsecase, 0, len(data))
	for _, val := range data {
		res = append(res, *ToUserWithSkillUsecase(&val))
	}

	return res
}

func ToUsersWithSkillsUsecase(data []entities_user.UserWithSkills) []UserWithSkillsUsecase {
	res := make([]UserWithSkillsUsecase, 0, len(data))
	for _, val := range data {
		res = append(res, *ToUserWithSkillUsecase(&val))
	}

	return res
}
func ToUserWithSkillUsecase(data *entities_user.UserWithSkills) *UserWithSkillsUsecase {
	return &UserWithSkillsUsecase{
		UserUsecase: *ToUserUsecase(&data.User),
		Skills:      data.Skills,
	}
}
