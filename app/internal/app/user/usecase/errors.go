package user_usecase

import "github.com/pkg/errors"

var (
	ArgError = errors.New("invalid argument, expected not nil")
	BadAuth  = errors.New("authorization error - not valid data")
)
