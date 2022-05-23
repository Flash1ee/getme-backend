package entities

import (
	"time"
)

type User struct {
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
