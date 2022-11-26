package models

import (
	"time"
)

type Token string

type TokenSources struct {
	UserId int64

	ExpiredTime time.Time
}
