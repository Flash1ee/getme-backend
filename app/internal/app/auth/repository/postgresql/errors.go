package repository_postgresql

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"getme-backend/internal/app/auth/repository"
	"getme-backend/internal/pkg/utilits/postgresql"
)

const (
	codeDuplicateVal   = "23505"
	emailConstraint    = "users_email_key"
	nicknameConstraint = "users_pkey"
)

func parsePQError(err *pq.Error) error {
	switch {
	case err.Code == codeDuplicateVal && err.Constraint == emailConstraint:
		return auth_repository.EmailAlreadyExist
	case err.Code == codeDuplicateVal && err.Constraint == nicknameConstraint:
		return auth_repository.NicknameAlreadyExist
	default:
		return postgresql_utilits.NewDBError(err)
	}
}

var (
	CreateError = errors.New("can not create user, internal error")
	GetError    = errors.New("can not get user from db, internal error")
)
