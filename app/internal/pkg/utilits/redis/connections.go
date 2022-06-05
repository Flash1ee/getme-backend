package utilits

import "github.com/gomodule/redigo/redis"

func NewRedisPool(redisUrl string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisUrl)
		},
	}
}
