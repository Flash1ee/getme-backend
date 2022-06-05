package auth_repository

import "github.com/pkg/errors"

var (
	EmailAlreadyExist    = errors.New("email already exist")
	NicknameAlreadyExist = errors.New("nickname already exist")
)
