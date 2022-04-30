package server

import "github.com/pkg/errors"

var (
	ArgError = errors.New("invalid initialization BaseServer values, expected not nil")
)
