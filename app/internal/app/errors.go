package app

import (
	"fmt"
	"github.com/pkg/errors"
)

const InvalidInt = -1
const InvalidFloat = -1.0

var UnknownError = errors.New("gotten unspecified error")

type GeneralError struct {
	Err         error
	ExternalErr error
}

func (e GeneralError) Error() string {
	return fmt.Sprintf("%v: %s", e.Err, e.ExternalErr.Error())
}
