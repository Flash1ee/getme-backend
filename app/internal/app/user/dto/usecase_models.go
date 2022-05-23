package dto

import (
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
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Nickname:  m.About,
		Avatar:    m.Avatar,
	}
}
