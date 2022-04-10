package postgresql

import (
	"getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/utilits/postgresql"
	"github.com/lib/pq"
)

const (
	codeDuplicateVal   = "23505"
	emailConstraint    = "users_email_key"
	nicknameConstraint = "users_pkey"
)

func parsePQError(err *pq.Error) error {
	switch {
	case err.Code == codeDuplicateVal && err.Constraint == emailConstraint:
		return repository.EmailAlreadyExist
	case err.Code == codeDuplicateVal && err.Constraint == nicknameConstraint:
		return repository.NicknameAlreadyExist
	default:
		return postgresql_utilits.NewDBError(err)
	}
}
