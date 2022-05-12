package token_redis_repository

import "errors"

var (
	InvalidStorageData = errors.New("can not parse data from storage")
	SetError           = errors.New("can not set value to storage")
	NotFound           = errors.New("data in storage not found")
)
