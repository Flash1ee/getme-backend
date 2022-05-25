package dto

import (
	"database/sql"
	"time"

	"getme-backend/internal/app/user/entities"
)

type UserUsecase struct {
	ID           int64     `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Nickname     string    `db:"nickname"`
	About        string    `db:"about"`
	Avatar       string    `db:"avatar"`
	Email        string    `db:"email"`
	IsSearchable bool      `db:"is_searchable"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
type UserWithSkillsUsecase struct {
	UserUsecase
	Skill []string
}

func (m *UserUsecase) ToUserEntity() *entities.User {
	return &entities.User{
		ID: m.ID,
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
	}
}

func ToUserUsecase(data *entities.User) *UserUsecase {
	return &UserUsecase{
		ID:           data.ID,
		FirstName:    data.FirstName.String,
		LastName:     data.LastName.String,
		Nickname:     data.Nickname,
		About:        data.About.String,
		Avatar:       data.Avatar.String,
		Email:        data.Email.String,
		IsSearchable: data.IsSearchable,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToUsersWithSkillUsecase(data []entities.UserWithSkills) []UserWithSkillsUsecase {
	res := make([]UserWithSkillsUsecase, 0, len(data))
	for _, val := range data {
		res = append(res, *ToUserWithSkillUsecase(&val))
	}

	return res
}
func ToUserWithSkillsUsecase(data []entities.UserWithSkills) []UserWithSkillsUsecase {
	res := make([]UserWithSkillsUsecase, 0, len(data))
	for _, val := range data {
		res = append(res, *ToUserWithSkillUsecase(&val))
	}

	return res
}
func ToUserWithSkillUsecase(data *entities.UserWithSkills) *UserWithSkillsUsecase {
	return &UserWithSkillsUsecase{
		UserUsecase: UserUsecase{
			ID:           data.ID,
			FirstName:    data.FirstName.String,
			LastName:     data.LastName.String,
			Nickname:     data.Nickname,
			About:        data.About.String,
			Avatar:       data.Avatar.String,
			Email:        data.Email.String,
			IsSearchable: data.IsSearchable,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
		},
		Skill: data.Skills,
	}
}
