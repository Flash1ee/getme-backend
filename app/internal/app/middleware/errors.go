package middleware

import "github.com/pkg/errors"

var (
	InvalidParameters = errors.New("invalid parameters")
	BDError           = errors.New("can not do bd operation")
	InternalError     = errors.New("server error")
)
