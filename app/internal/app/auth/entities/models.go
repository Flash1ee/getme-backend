package entities

import (
	"time"
)

type SimpleAuth struct {
	ID                int64  `db:"id"`
	Login             string `db:"login"`
	EncryptedPassword string `db:"encrypted_password,omitempty"`
	UserID            int64  `db:"user_id"`
}

type TelegramAuth struct {
	TelegramID int64     `db:"tg_id"`
	CreatedAt  time.Time `db:"created_at"`
	LastAuth   time.Time `db:"last_auth"`
	UserID     int64     `db:"user_id"`
}
