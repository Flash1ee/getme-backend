package entities

import "time"

type Token struct {
	Token string `db:"token"`
}

type TokenSources struct {
	IdentifierData string
	ExpiredTime    time.Time
}
