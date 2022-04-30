package entities

import "time"

type User struct {
	ID           int64     `json:"id"`
	TelegramID   int64     `json:"tg_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Nickname     string    `json:"nickname"`
	About        string    `json:"about"`
	Avatar       string    `json:"avatar"`
	Email        string    `json:"email"`
	IsSearchable bool      `json:"is_searchable"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
