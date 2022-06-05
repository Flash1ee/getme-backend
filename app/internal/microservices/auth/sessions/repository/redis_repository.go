package repository

import (
	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/microservices/auth/sessions/models"
)

type RedisRepository struct {
	redisPool *redis.Pool
	log       *logrus.Logger
}

func NewRedisRepository(pool *redis.Pool, log *logrus.Logger) *RedisRepository {
	return &RedisRepository{
		redisPool: pool,
		log:       log,
	}
}

func (repo *RedisRepository) Set(session *models.Session) error {
	con := repo.redisPool.Get()
	defer func(con redis.Conn) {
		err := con.Close()
		if err != nil {
			repo.log.Errorf("Unsuccessful close connection to redis with error: %s, with session: %s",
				err.Error(), session)
		}
	}(con)

	res, err := redis.String(con.Do("SET", session.UniqID, session.UserID,
		"PX", session.Expiration))
	if res != "OK" {
		return errors.Wrapf(err,
			"error when try create session with uniqId: %s, and userId: %s", session.UniqID, session.UserID)
	}
	return nil
}

func (repo *RedisRepository) GetUserId(uniqID string, updExpiration int) (string, error) {
	con := repo.redisPool.Get()
	defer func() {
		err := con.Close()
		if err != nil {
			repo.log.Errorf("Unsuccessful close connection to redis with error: %s, with session id: %s",
				err.Error(), uniqID)
		}
	}()

	res, err := redis.String(con.Do("GET", uniqID))
	if err != nil {
		return "", errors.Wrapf(err,
			"error when try get session with uniqId: %s", uniqID)
	}

	_, err = redis.Int64(con.Do("EXPIRE", uniqID, updExpiration/100))
	if err != nil {
		return "", errors.Wrapf(err,
			"error when try update expire session with uniqId: %s", uniqID)
	}
	return res, nil
}

func (repo *RedisRepository) Del(session *models.Session) error {
	con := repo.redisPool.Get()
	defer func() {
		err := con.Close()
		if err != nil {
			repo.log.Errorf("Unsuccessful close connection to redis with error: %s, with session: %s",
				err.Error(), session)
		}
	}()

	_, err := redis.Int(con.Do("DEL", session.UniqID))
	return errors.Wrapf(err,
		"error when try delete session with uniqId: %s, and userId: %s", session.UniqID, session.UserID)
}
