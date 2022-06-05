package postgresql_utilits

import (
	"errors"

	"getme-backend/internal/app"
)

var (
	DefaultErrDB = errors.New("something wrong DB")
	NotFound     = errors.New("user not found")
	Conflict     = errors.New("conflict in db")
	Exists       = errors.New("data exists")
)

func NewDBError(externalErr error) *app.GeneralError {
	return &app.GeneralError{
		Err:         DefaultErrDB,
		ExternalErr: externalErr,
	}

}
