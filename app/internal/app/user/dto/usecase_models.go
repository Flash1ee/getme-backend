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

func (m *UserUsecase) ToUserEntity() *entities.User {
	return &entities.User{
		ID: m.ID,
		FirstName: sql.NullString{
			String: m.FirstName,
		},
		LastName: sql.NullString{
			String: m.LastName,
		},
		Nickname: m.About,
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
