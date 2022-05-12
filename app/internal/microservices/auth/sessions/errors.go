package sessions

import "errors"

var (
	StatusNotOK = errors.New("can not add values to redis\nstatus not OK")
)
