package token_usecase

import "github.com/pkg/errors"

var (
	InvalidUserToken = errors.New("this user was not given this token")
)
