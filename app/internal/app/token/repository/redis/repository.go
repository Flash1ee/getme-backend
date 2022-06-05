package token_redis_repository

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	"getme-backend/internal/app"
)

type TokenRepository struct {
	redisPool *redis.Pool
}

func NewTokenRepository(pool *redis.Pool) *TokenRepository {
	return &TokenRepository{
		redisPool: pool,
	}
}

// Set Errors:
// 		app.GeneralError with Errors
// 			SetError
func (repo *TokenRepository) Set(key string, value string, timeExp int) error {
	con := repo.redisPool.Get()
	defer func(con redis.Conn) {
		err := con.Close()
		if err != nil {
			panic(fmt.Sprintf("Unsuccessful close connection to redis with error: %s, with key: %s value: %s", err.Error(), key, value))
		}
	}(con)
	res, err := redis.String(con.Do("SET", key, value, "EX", timeExp))
	if res != "OK" {
		return app.GeneralError{
			Err: errors.Wrapf(SetError,
				"error when try set with key: %s value: %s", key, value),
			ExternalErr: err,
		}
	}
	return nil
}

// Get Errors:
//		NotFound
// 		app.GeneralError with Errors
// 			InvalidStorageData
func (repo *TokenRepository) Get(key string) (string, error) {
	con := repo.redisPool.Get()
	defer func(con redis.Conn) {
		err := con.Close()
		if err != nil {
			panic(errors.Errorf("Unsuccessful close connection to redis with error: %s, with key: %s",
				err.Error(), key))
		}
	}(con)
	res, err := redis.String(con.Do("GET", key))
	if err == redis.ErrNil {
		return "", NotFound
	}
	if err != nil {
		return "", app.GeneralError{
			Err: errors.Wrapf(InvalidStorageData,
				"error when try get from TokenRepository with key: %s", key),
			ExternalErr: err,
		}
	}

	return res, nil
}
